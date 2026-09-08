from .model import AutoscaleModel

def autoscale_decision(model: AutoscaleModel, cpu_usage):
    return model.decide(cpu_usage)
