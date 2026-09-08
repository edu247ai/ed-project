class AutoscaleModel:
    def __init__(self, high=80.0, low=20.0):
        self.high = high
        self.low = low

    def decide(self, cpu_usage):
        if cpu_usage > self.high:
            return "scale_up"
        if cpu_usage < self.low:
            return "scale_down"
        return "stable"
