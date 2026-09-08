from .base import Skill

class ResourceOptSkill(Skill):
    name = "resource_opt"

    def run(self, load, memory_mb):
        model = self.supervisor.registry.get("resource_opt")
        if not model:
            return {"error": "Resource optimization model not trained"}
        return model.optimize(load, memory_mb)
