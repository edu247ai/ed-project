import importlib

class ModelLoader:
    def load(self, module_path, class_name):
        module = importlib.import_module(module_path)
        cls = getattr(module, class_name)
        return cls()
