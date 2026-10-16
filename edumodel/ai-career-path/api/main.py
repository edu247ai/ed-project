from fastapi import FastAPI

app = FastAPI(title="Career Path AI service", version="0.1.0")


@app.get("/health", tags=["health"])
def health() -> dict[str, str]:
    return {"status": "ok", "service": "ai-career-path"}
