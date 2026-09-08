from edu247model.model.engine import build_engine

engine = build_engine()

def forecast_tool(args):
    series = args.get("series", [])
    steps = args.get("steps", 3)
    return engine.run_skill("forecast", series=series, steps=steps)
