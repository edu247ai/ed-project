import importlib

def load_skill(module_path, class_name, supervisor):
    module = importlib.import_module(module_path)
    cls = getattr(module, class_name)
    return cls(supervisor)
