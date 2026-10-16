class Model:
    """Implement the service-specific model behavior in this class."""

    def predict(self, payload: dict[str, object]) -> dict[str, object]:
        """Return a prediction for the input payload once the model is implemented."""
        raise NotImplementedError("Implement the service model's predict method")
