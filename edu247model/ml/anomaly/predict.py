from .model import AnomalyModel

def predict_anomaly(model: AnomalyModel, value: float) -> bool:
    return model.is_anomaly(value)
