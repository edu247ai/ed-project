from fastapi import APIRouter
from pydantic import BaseModel
from edu247model.model.engine import build_engine

router = APIRouter()
engine = build_engine()


# ------------------------------
# Request Models
# ------------------------------

class AnomalyRequest(BaseModel):
    value: float

class ForecastRequest(BaseModel):
    series: list[float]
    steps: int = 3

class AutoscaleRequest(BaseModel):
    cpu: float

class ResourceOptRequest(BaseModel):
    load: float
    memory: float

class RAGRequest(BaseModel):
    query: str


# ------------------------------
# Routes
# ------------------------------

@router.get("/")
def root():
    return {"status": "EDU247 Intelligence API running"}


@router.get("/skills")
def list_skills():
    return {"skills": engine.skills.list()}


@router.post("/ml/anomaly")
def anomaly_api(req: AnomalyRequest):
    return engine.run_skill("anomaly", value=req.value)


@router.post("/ml/forecast")
def forecast_api(req: ForecastRequest):
    return engine.run_skill("forecast", series=req.series, steps=req.steps)


@router.post("/ml/autoscale")
def autoscale_api(req: AutoscaleRequest):
    return engine.run_skill("autoscale", cpu_usage=req.cpu)


@router.post("/ml/resource_opt")
def resource_opt_api(req: ResourceOptRequest):
    return engine.run_skill("resource_opt", load=req.load, memory_mb=req.memory)


@router.post("/rag/query")
def rag_api(req: RAGRequest):
    return engine.run_skill("rag", query=req.query)


@router.get("/rl/evaluate")
def rl_api():
    return engine.run_skill("rl")
