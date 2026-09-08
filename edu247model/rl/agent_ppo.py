import random
from .env_autoscale import ACTIONS

class PPOAgent:
    def __init__(self):
        self.policy = {
            "high": "scale_up",
            "low": "scale_down",
            "mid": "stable",
        }

    def act(self, cpu: float) -> str:
        if cpu > 70:
            return self.policy["high"]
        if cpu < 30:
            return self.policy["low"]
        return self.policy["mid"]

    def update(self, experiences):
        # Placeholder for PPO update – you can extend later
        pass
