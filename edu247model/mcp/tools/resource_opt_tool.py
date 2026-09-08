from edu247model.model.engine import build_engine

engine = build_engine()

def resource_opt_tool(args):
    load = args.get("load")
    memory = args.get("memory")
    return engine.run_skill("resource_opt", load=load, memory_mb=memory)
