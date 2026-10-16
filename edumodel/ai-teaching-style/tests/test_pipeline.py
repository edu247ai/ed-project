from pipeline import Pipeline


class FakeModel:
    def predict(self, payload: dict[str, object]) -> dict[str, object]:
        return {"received": payload}


def test_pipeline_passes_payload_to_model() -> None:
    payload = {"input": "example"}

    assert Pipeline(FakeModel()).run(payload) == {"received": payload}
