class Skill:
    name = "base"

    def __init__(self, supervisor):
        self.supervisor = supervisor

    def run(self, *args, **kwargs):
        raise NotImplementedError("Skill must implement run()")
