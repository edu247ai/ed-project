from .model import AutoscaleModel

def build_autoscale_model(high=80.0, low=20.0):
    return AutoscaleModel(high, low)
