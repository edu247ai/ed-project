from model import EmailCategory, EmailClassifier, EmailMessage


def make_email(subject: str, body: str = "") -> EmailMessage:
    return EmailMessage(
        message_id="test-1",
        sender="sender@example.com",
        recipients=("recipient@example.com",),
        subject=subject,
        body=body,
    )


def test_classifier_recognizes_transactional_email() -> None:
    result = EmailClassifier().classify(make_email("Your order confirmation"))

    assert result.category is EmailCategory.TRANSACTIONAL
    assert result.matched_terms == ("order confirmation",)
    assert 0.0 < result.score <= 1.0


def test_classifier_returns_unknown_when_no_rule_matches() -> None:
    result = EmailClassifier().classify(make_email("Hello", "Are we still meeting?"))

    assert result.category is EmailCategory.UNKNOWN
    assert result.score == 0.0
    assert result.matched_terms == ()


def test_classifier_uses_spam_as_tie_break_priority() -> None:
    email = make_email("Act now, and unsubscribe")

    result = EmailClassifier().classify(email)

    assert result.category is EmailCategory.SPAM
