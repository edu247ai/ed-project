class ResourceOptModel:
    def optimize(self, load, memory_mb):
        cpu_units = max(1, int(load / 10))
        mem_units = max(256, int(memory_mb * 1.2))
        return {
            "cpu": cpu_units,
            "memory_mb": mem_units,
        }
