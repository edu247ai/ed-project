from edu247model.model.engine import build_engine

engine = build_engine()

def anomaly_tool(args):
    value = args.get("value")
    return engine.run_skill("anomaly", value=value)
