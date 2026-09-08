import numpy as np

class Metrics:
    def mse(self, y_true, y_pred):
        return float(np.mean((np.array(y_true) - np.array(y_pred)) ** 2))

    def drift(self, old_mean, new_mean):
        return abs(new_mean - old_mean)
