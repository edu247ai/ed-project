from dataclasses import dataclass
from enum import StrEnum


class EmailCategory(StrEnum):
    SPAM = "spam"
    PROMOTIONAL = "promotional"
    TRANSACTIONAL = "transactional"
    PERSONAL = "personal"
    UNKNOWN = "unknown"


@dataclass(frozen=True)
class EmailMessage:
    message_id: str
    sender: str
    recipients: tuple[str, ...]
    subject: str
    body: str

    def __post_init__(self) -> None:
        if not self.message_id.strip():
            raise ValueError("message_id must not be empty")
        if not self.sender.strip():
            raise ValueError("sender must not be empty")
        if not self.recipients or any(not recipient.strip() for recipient in self.recipients):
            raise ValueError("at least one non-empty recipient is required")


@dataclass(frozen=True)
class EmailClassification:
    category: EmailCategory
    score: float
    matched_terms: tuple[str, ...]

    def __post_init__(self) -> None:
        if not 0.0 <= self.score <= 1.0:
            raise ValueError("score must be between 0 and 1")
