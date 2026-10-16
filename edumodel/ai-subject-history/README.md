# Subject History AI service

## Purpose

Starter scaffold for the EDU247 subject history AI service. The model behavior is not implemented yet; the only working endpoint is a health check.

## Set up and run

Requires Python 3.11 or newer. From this service directory:

```bash
python -m venv .venv
source .venv/bin/activate  # Windows: .venv\Scripts\activate
python -m pip install -e ".[dev]"
python -m uvicorn api.main:app --reload
```

The health endpoint is available at `http://127.0.0.1:8000/health`.

## Test

```bash
python -m pytest
```

## Develop and train the model

Implement this service's behavior in `model/`, then connect preprocessing, inference, and postprocessing in `pipeline/`. Add training and evaluation entry points under `scripts/`; training is not implemented in this scaffold, so there is no training command yet. Keep generated checkpoints and evaluation outputs in `artifacts/`.

## Deploy

This scaffold is not production-ready. Before deployment, implement and validate the model and pipeline, configure secrets through the deployment environment, and add a production container/deployment definition. The Uvicorn command above is for local development only.
