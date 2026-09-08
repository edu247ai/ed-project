from .base import Skill

class ForecastSkill(Skill):
    name = "forecast"

    def run(self, series, steps=3):
        model = self.supervisor.registry.get("forecast")
        if not model:
            return {"error": "Forecast model not trained"}
        return {
            "forecast": model.forecast(series, steps)
        }
