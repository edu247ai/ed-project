import pytest

from model import Model


def test_unimplemented_model_fails_explicitly() -> None:
    with pytest.raises(NotImplementedError, match="Implement the service model"):
        Model().predict({"input": "example"})
