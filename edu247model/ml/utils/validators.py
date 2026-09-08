def ensure_non_empty(data, name="data"):
    if not data:
        raise ValueError(f"{name} must not be empty")
