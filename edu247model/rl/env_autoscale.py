import random

ACTIONS = ["scale_up", "scale_down", "stable"]

class AutoscaleEnv:
    def __init__(self):
        self.reset()

    def reset(self):
        self.cpu = random.randint(10, 90)
        return self.cpu

    def step(self, action: str):
        if action == "scale_up":
            self.cpu -= 10
        elif action == "scale_down":
            self.cpu += 10
        elif action == "stable":
            self.cpu += random.randint(-5, 5)

        reward = -abs(self.cpu - 50)
        done = self.cpu <= 0 or self.cpu >= 100

        return self.cpu, reward, done
