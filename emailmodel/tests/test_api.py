from fastapi.testclient import TestClient

from api.main import app

client = TestClient(app)


def test_health_endpoint() -> None:
    response = client.get("/health")

    assert response.status_code == 200
    assert response.json() == {"status": "ok", "service": "email-analysis"}


def test_email_analysis_endpoint_returns_classification() -> None:
    response = client.post(
        "/v1/emails/analyze",
        json={
            "message_id": "message-4",
            "sender": "sender@example.com",
            "recipients": ["recipient@example.com"],
            "subject": "Your invoice",
            "body": "Details attached",
        },
    )

    assert response.status_code == 200
    result = response.json()
    assert result["classification"]["category"] == "transactional"
    assert result["reply_suggestion"] is None
    assert result["completed_steps"] == [
        "validate_email",
        "classify_email",
        "suggest_reply",
    ]


def test_email_analysis_endpoint_can_return_draft() -> None:
    response = client.post(
        "/v1/emails/analyze?suggest_reply=true",
        json={
            "message_id": "message-5",
            "sender": "sender@example.com",
            "recipients": ["recipient@example.com"],
            "subject": "Hello",
            "body": "A message",
        },
    )

    assert response.status_code == 200
    assert response.json()["reply_suggestion"]


def test_email_analysis_endpoint_rejects_missing_recipients() -> None:
    response = client.post(
        "/v1/emails/analyze",
        json={
            "message_id": "message-6",
            "sender": "sender@example.com",
            "recipients": [],
            "subject": "Hello",
            "body": "A message",
        },
    )

    assert response.status_code == 422
