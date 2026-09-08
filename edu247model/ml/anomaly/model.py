import numpy as np

class AnomalyModel:
    def __init__(self, threshold: float = 3.0):
        self.threshold = threshold
        self.mean = None
        self.std = None

    def fit(self, values):
        arr = np.array(values, dtype=float)
        self.mean = float(arr.mean())
        self.std = float(arr.std() or 1e-8)

    def score(self, value: float) -> float:
        return abs(value - self.mean) / self.std

    def is_anomaly(self, value: float) -> bool:
        return self.score(value) > self.threshold
