# EDU247 email analysis service

This Python service is a starter for analyzing incoming email. It provides an
HTTP health check and an analysis endpoint that validates an email, assigns a
heuristic category, and can optionally return a generic reply draft.

The classifier is a deterministic keyword baseline, not a trained ML model. Its
score is a simple rule score and is not a calibrated probability. Replace or
extend it with a reviewed model/provider adapter when suitable training data
and evaluation criteria are available.

The service does **not** send, forward, delete, or persist email. A reply
suggestion is only draft text returned to the caller.

## Set up and run

Requires Python 3.11 or newer. From this directory:

```bash
python -m venv .venv
source .venv/bin/activate  # Windows: .venv\Scripts\activate
python -m pip install -e ".[dev]"
python -m uvicorn api.main:app --reload
```

The health endpoint is `GET /health`. Analyze an email with `POST
/v1/emails/analyze`:

```bash
curl -X POST http://127.0.0.1:8000/v1/emails/analyze \
  -H 'Content-Type: application/json' \
  -d '{"message_id":"demo-1","sender":"person@example.com","recipients":["you@example.com"],"subject":"Your receipt","body":"Payment received"}'
```

Pass `?suggest_reply=true` to request a generic draft suggestion. The API never
sends that suggestion.

## Test

```bash
python -m pytest
```

## Model and orchestration

- `model/` defines the email domain object and initial keyword classifier.
- `pipeline/` composes validation, classification, and optional reply drafting.
- `pipeline/` owns this service's synchronous, in-process orchestration; it has
  no shared runtime dependency on other services.
- `config/` holds the initial classifier keyword rules.
- `adapters/` is reserved for isolated external integrations; no email-sending
  adapter is included.

The workflow is synchronous and in-process: a process restart loses in-flight
work. It does not retry failed steps or persist results. Use a durable queue
before using it for long-running or delivery-critical work.

## Deployment

This is a development scaffold, not a production mail server. Before deployment,
add authentication, rate limiting, privacy-aware logging, operational
monitoring, a production ASGI server configuration, and deployment manifests.
Do not send user email content to an external model/provider without an
approved privacy and security review.
