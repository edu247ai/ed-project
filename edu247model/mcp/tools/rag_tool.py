from edu247model.model.engine import build_engine

engine = build_engine()

def rag_tool(args):
    query = args.get("query")
    return engine.run_skill("rag", query=query)
