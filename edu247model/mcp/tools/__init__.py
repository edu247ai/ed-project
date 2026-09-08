from .anomaly_tool import anomaly_tool
from .forecast_tool import forecast_tool
from .autoscale_tool import autoscale_tool
from .resource_opt_tool import resource_opt_tool
from .rag_tool import rag_tool
from .rl_tool import rl_tool

TOOLS = {
    "anomaly": anomaly_tool,
    "forecast": forecast_tool,
    "autoscale": autoscale_tool,
    "resource_opt": resource_opt_tool,
    "rag": rag_tool,
    "rl": rl_tool,
}

def list_tools():
    return list(TOOLS.keys())

def get_tool(name):
    return TOOLS.get(name)
