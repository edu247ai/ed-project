from edu247model.model.engine import build_engine

engine = build_engine()

def rl_tool(args):
    return engine.run_skill("rl")
