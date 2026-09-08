from fastapi import FastAPI
from edu247model.api.middleware import request_logger
from edu247model.api.router import router

app = FastAPI(title="EDU247 Intelligence API", version="1.0")

# Middleware
app.middleware("http")(request_logger)

# Routes
app.include_router(router)
