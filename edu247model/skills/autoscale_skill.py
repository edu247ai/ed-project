from .base import Skill

class AutoscaleSkill(Skill):
    name = "autoscale"

    def run(self, cpu_usage):
        model = self.supervisor.registry.get("autoscale")
        if not model:
            return {"error": "Autoscale model not trained"}
        return {
            "cpu": cpu_usage,
            "decision": model.decide(cpu_usage)
        }
