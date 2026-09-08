import numpy as np

class ForecastModel:
    def __init__(self, window: int = 5):
        self.window = window

    def forecast(self, series, steps=1):
        arr = np.array(series, dtype=float)
        if len(arr) == 0:
            return [0.0] * steps
        avg = float(arr[-self.window:].mean())
        return [avg] * steps
