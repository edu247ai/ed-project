# EDU247 Businessmodel Services

## Overview

`businessmodel/` contains independent Python microservices for AI-powered
creation, developer tooling, security analysis, recommendations, and service
routing. Each service owns its implementation, configuration, tests, and
documentation. Services communicate through explicit APIs or events; they do
not import one another's internal code.

> **Repository status:** The service directories listed below currently exist
> as placeholders. Their described capabilities are intended responsibilities,
> not claims that working APIs or models are already implemented.

## Services

| Service | Intended responsibility |
| --- | --- |
| `ai-gateway/` | Entry point for routing authorized requests to businessmodel AI services |
| `code-review-service/` | Analyze code changes and return review findings |
| `comic-creator-service/` | Create educational comic scripts and related visual assets |
| `documentation-service/` | Generate and maintain technical and service documentation |
| `infographic-creator-service/` | Turn educational content into infographic specifications or assets |
| `lesson-builder-service/` | Assemble lesson content and learning activities |
| `mindmap-creator-service/` | Organize topics into structured mindmaps |
| `recommendation-service/` | Produce recommendations from approved content and user signals |
| `security-auditor-service/` | Analyze code and service configuration for security issues |
| `tutor-service/` | Provide explanations and guided learning assistance |

These ten services have separate deployment and development lifecycles. The
gateway may route requests, but it should not contain the individual services'
domain logic or become a shared code package.

## Required service layout

Every service should keep its implementation within its own directory and use
the same predictable structure:

```text
<service-name>/
├── api/          # HTTP or other service endpoints
├── model/        # Domain and AI model interfaces/definitions
├── pipeline/     # Input processing, model execution, and output handling
├── adapters/     # Database, storage, provider, and external API connections
├── config/       # Environment-based configuration; no committed secrets
├── tests/        # Unit and integration tests, including mocks
├── scripts/      # Training, evaluation, and maintenance utilities
└── README.md     # Purpose, setup, usage, testing, and deployment
```

The exact modules inside these folders depend on the service. Keep provider-
specific integrations in `adapters/`, model and domain definitions in
`model/`, and processing flow in `pipeline/`. Do not add a `shared/` runtime
dependency: a reusable capability should be exposed through a stable API or
event contract, or be duplicated as a small service-owned implementation when
appropriate.

## Service boundaries and communication

- Each service owns its API, model, pipeline, configuration, tests, and
  deployment lifecycle.
- Do not import another service's source code or access its private database.
- Integrate through documented, versioned APIs or events.
- Authenticate and authorize service-to-service requests.
- Treat external model/provider access as an adapter and handle failures
  explicitly.
- Store secrets in the deployment environment or an approved secret manager.
- Keep generated files and model artifacts out of source control unless a
  service has an explicit versioning policy.

## Developing a service

When implementing a service:

1. Read its service README and define its inputs, outputs, and ownership.
2. Add the required folders and implement a small, testable vertical slice.
3. Add unit tests for model and pipeline behavior and API/integration tests for
   service boundaries.
4. Document local setup, run, test, training (if applicable), and deployment
   commands in that service's README.
5. Update API/event contracts when changing a service boundary.
6. Run the checks documented by the service before submitting changes.

Each service must provide its own dependency manifest and setup instructions.
There is no central Python environment or shared runtime package for these
services.

## Scope note

The `businessmodel/` directory is a grouping for these independent services;
it is not itself a deployable service or a shared application runtime. The
service descriptions above document intended scope and should be updated as
each implementation becomes available.
