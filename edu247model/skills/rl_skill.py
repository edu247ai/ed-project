from .base import Skill
from edu247model.rl.evaluate_rl import evaluate_agent

class RLSkill(Skill):
    name = "rl"

    def run(self):
        agent = self.supervisor.registry.get("rl_agent")
        if not agent:
            return {"error": "RL agent not trained"}
        return {
            "evaluation": evaluate_agent(agent)
        }
