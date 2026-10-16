from fastapi import FastAPI
from pydantic import BaseModel, Field

from model import EmailMessage
from pipeline import EmailAnalysisService

app = FastAPI(title="EDU247 Email Analysis", version="0.1.0")
analysis_service = EmailAnalysisService()


class EmailAnalysisRequest(BaseModel):
    message_id: str
    sender: str
    recipients: list[str] = Field(min_length=1)
    subject: str = ""
    body: str = ""


class ClassificationResponse(BaseModel):
    category: str
    score: float
    matched_terms: list[str]


class EmailAnalysisResponse(BaseModel):
    message_id: str
    classification: ClassificationResponse
    reply_suggestion: str | None
    completed_steps: list[str]


@app.get("/health", tags=["health"])
def health() -> dict[str, str]:
    return {"status": "ok", "service": "email-analysis"}


@app.post(
    "/v1/emails/analyze",
    response_model=EmailAnalysisResponse,
    tags=["email-analysis"],
)
def analyze_email(
    request: EmailAnalysisRequest,
    suggest_reply: bool = False,
) -> EmailAnalysisResponse:
    email = EmailMessage(
        message_id=request.message_id,
        sender=request.sender,
        recipients=tuple(request.recipients),
        subject=request.subject,
        body=request.body,
    )
    result = analysis_service.analyze(email, suggest_reply=suggest_reply)
    classification = result.context.classification
    if classification is None:
        raise RuntimeError("Email workflow completed without a classification")
    return EmailAnalysisResponse(
        message_id=email.message_id,
        classification=ClassificationResponse(
            category=classification.category.value,
            score=classification.score,
            matched_terms=list(classification.matched_terms),
        ),
        reply_suggestion=result.context.reply_suggestion,
        completed_steps=list(result.completed_steps),
    )
