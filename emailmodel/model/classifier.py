from collections.abc import Mapping

from config.classification import DEFAULT_KEYWORDS
from model.email import EmailCategory, EmailClassification, EmailMessage


class EmailClassifier:
    """Deterministic keyword baseline; scores are not calibrated probabilities."""

    def __init__(
        self,
        keywords: Mapping[EmailCategory, tuple[str, ...]] = DEFAULT_KEYWORDS,
    ) -> None:
        self._keywords = {
            category: tuple(term.casefold() for term in terms)
            for category, terms in keywords.items()
        }

    def classify(self, email: EmailMessage) -> EmailClassification:
        content = f"{email.subject}\n{email.body}".casefold()
        matches = {
            category: tuple(term for term in terms if term in content)
            for category, terms in self._keywords.items()
        }
        candidates = [
            (category, terms)
            for category, terms in matches.items()
            if terms
        ]
        if not candidates:
            return EmailClassification(
                category=EmailCategory.UNKNOWN,
                score=0.0,
                matched_terms=(),
            )

        # Stable tie-break order: spam, promotional, then transactional.
        priority = {
            EmailCategory.SPAM: 0,
            EmailCategory.PROMOTIONAL: 1,
            EmailCategory.TRANSACTIONAL: 2,
        }
        category, terms = min(
            candidates,
            key=lambda candidate: (
                -len(candidate[1]),
                priority.get(candidate[0], len(priority)),
            ),
        )
        score = min(len(terms) / 3.0, 1.0)
        return EmailClassification(
            category=category,
            score=score,
            matched_terms=terms,
        )
