from model.email import EmailCategory

DEFAULT_KEYWORDS: dict[EmailCategory, tuple[str, ...]] = {
    EmailCategory.SPAM: (
        "claim your prize",
        "act now",
        "wire transfer",
        "guaranteed income",
    ),
    EmailCategory.PROMOTIONAL: (
        "unsubscribe",
        "limited time",
        "sale",
        "discount",
        "special offer",
    ),
    EmailCategory.TRANSACTIONAL: (
        "receipt",
        "invoice",
        "password reset",
        "shipping update",
        "order confirmation",
    ),
}
