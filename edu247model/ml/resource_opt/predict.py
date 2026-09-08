from .model import ResourceOptModel

def optimize_resources(model: ResourceOptModel, load, memory_mb):
    return model.optimize(load, memory_mb)
