from edu247model.model.engine import build_engine

engine = build_engine()

def autoscale_tool(args):
    cpu = args.get("cpu")
    return engine.run_skill("autoscale", cpu_usage=cpu)
