from edu247model.skills.base import Skill

class AnomalySkill(Skill):
    name = "anomaly"

    def run(self, value):
        model = self.supervisor.registry.get("anomaly")
        if not model:
            return {"error": "Anomaly model not trained"}
        return {
            "value": value,
            "is_anomaly": model.is_anomaly(value),
            "score": model.score(value)
        }
