from edu247model.model.core.events import EventBus
from edu247model.model.core.registry import ModelRegistry
from edu247model.model.core.metrics import Metrics
from edu247model.skills.registry import SkillRegistry


class Supervisor:
    def __init__(self):
        # Core systems
        self.bus = EventBus()
        self.registry = ModelRegistry()
        self.skills = SkillRegistry()
        self.metrics = Metrics()

        # Emit startup event
        self.bus.emit("supervisor_started", {"status": "ok"})

    # -----------------------------
    # Model Management
    # -----------------------------
    def register_model(self, name, model):
        self.registry.register(name, model)
        self.bus.emit("model_registered", {"name": name})

    def get_model(self, name):
        return self.registry.get(name)

    # -----------------------------
    # Skill Management
    # -----------------------------
    def register_skill(self, skill):
        self.skills.register(skill)
        self.bus.emit("skill_registered", {"name": skill.name})

    def run_skill(self, name, *args, **kwargs):
        skill = self.skills.get(name)
        if not skill:
            return {"error": f"Skill '{name}' not found"}
        return skill.run(*args, **kwargs)

    # -----------------------------
    # Evaluation
    # -----------------------------
    def evaluate_model(self, name, eval_fn, *args, **kwargs):
        model = self.registry.get(name)
        if not model:
            return {"error": f"Model '{name}' not found"}

        result = eval_fn(model, *args, **kwargs)
        self.bus.emit("model_evaluated", {"name": name, "result": result})
        return result
