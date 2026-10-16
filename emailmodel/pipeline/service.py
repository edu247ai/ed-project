from dataclasses import dataclass
from collections.abc import Callable, Sequence
from typing import Generic, TypeVar

from model import EmailClassification, EmailClassifier, EmailMessage

ContextT = TypeVar("ContextT")


@dataclass(frozen=True)
class WorkflowResult(Generic[ContextT]):
    context: ContextT
    completed_steps: tuple[str, ...]


class WorkflowStepError(RuntimeError):
    def __init__(self, failed_step: str, completed_steps: tuple[str, ...]) -> None:
        self.failed_step = failed_step
        self.completed_steps = completed_steps
        super().__init__(
            f"Email workflow failed at step {failed_step!r} after "
            f"{len(completed_steps)} completed step(s)"
        )


@dataclass(frozen=True)
class _WorkflowStep(Generic[ContextT]):
    name: str
    action: Callable[[ContextT], None]


class _Workflow(Generic[ContextT]):
    def __init__(self, steps: Sequence[_WorkflowStep[ContextT]]) -> None:
        self._steps = tuple(steps)

    def run(self, context: ContextT) -> WorkflowResult[ContextT]:
        completed_steps: list[str] = []
        for step in self._steps:
            try:
                step.action(context)
            except Exception as error:
                raise WorkflowStepError(step.name, tuple(completed_steps)) from error
            completed_steps.append(step.name)
        return WorkflowResult(context, tuple(completed_steps))


@dataclass
class EmailAnalysisContext:
    email: EmailMessage
    suggest_reply: bool
    classification: EmailClassification | None = None
    reply_suggestion: str | None = None


class EmailAnalysisService:
    def __init__(self, classifier: EmailClassifier | None = None) -> None:
        self._classifier = classifier or EmailClassifier()

    def analyze(
        self,
        email: EmailMessage,
        *,
        suggest_reply: bool = False,
    ) -> WorkflowResult[EmailAnalysisContext]:
        context = EmailAnalysisContext(email=email, suggest_reply=suggest_reply)
        workflow = _Workflow(
            [
                _WorkflowStep("validate_email", self._validate_email),
                _WorkflowStep("classify_email", self._classify_email),
                _WorkflowStep("suggest_reply", self._suggest_reply),
            ]
        )
        return workflow.run(context)

    @staticmethod
    def _validate_email(context: EmailAnalysisContext) -> None:
        if not context.email.subject.strip() and not context.email.body.strip():
            raise ValueError("email subject and body cannot both be empty")

    def _classify_email(self, context: EmailAnalysisContext) -> None:
        context.classification = self._classifier.classify(context.email)

    @staticmethod
    def _suggest_reply(context: EmailAnalysisContext) -> None:
        if context.suggest_reply:
            context.reply_suggestion = (
                "Thanks for your email. I’ll review it and follow up shortly."
            )
