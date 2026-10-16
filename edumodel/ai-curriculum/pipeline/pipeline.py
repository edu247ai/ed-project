from typing import Protocol


class Predictor(Protocol):
    def predict(self, payload: dict[str, object]) -> dict[str, object]: ...


class Pipeline:
    """Connect service-specific preprocessing and postprocessing to a model."""

    def __init__(self, model: Predictor) -> None:
        self.model = model

    def run(self, payload: dict[str, object]) -> dict[str, object]:
        """Pass input to the model; add preprocessing and postprocessing here."""
        return self.model.predict(payload)
