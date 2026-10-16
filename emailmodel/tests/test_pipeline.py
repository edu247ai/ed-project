import pytest

from model import EmailCategory, EmailMessage
from pipeline import EmailAnalysisService, WorkflowStepError


def test_email_analysis_runs_in_order_without_suggesting_reply_by_default() -> None:
    email = EmailMessage(
        message_id="message-1",
        sender="sender@example.com",
        recipients=("recipient@example.com",),
        subject="Your receipt",
        body="Payment processed",
    )

    result = EmailAnalysisService().analyze(email)

    assert result.completed_steps == (
        "validate_email",
        "classify_email",
        "suggest_reply",
    )
    assert result.context.classification is not None
    assert result.context.classification.category is EmailCategory.TRANSACTIONAL
    assert result.context.reply_suggestion is None


def test_email_analysis_can_return_a_draft_without_sending_it() -> None:
    email = EmailMessage(
        message_id="message-2",
        sender="sender@example.com",
        recipients=("recipient@example.com",),
        subject="Hello",
        body="Can you review this?",
    )

    result = EmailAnalysisService().analyze(email, suggest_reply=True)

    assert result.context.reply_suggestion is not None
    assert "follow up shortly" in result.context.reply_suggestion


def test_email_analysis_reports_empty_content_as_workflow_failure() -> None:
    email = EmailMessage(
        message_id="message-3",
        sender="sender@example.com",
        recipients=("recipient@example.com",),
        subject=" ",
        body="",
    )

    with pytest.raises(WorkflowStepError) as error:
        EmailAnalysisService().analyze(email)

    assert error.value.failed_step == "validate_email"
    assert error.value.completed_steps == ()
    assert isinstance(error.value.__cause__, ValueError)
