EDU247 Full Project Structure — Go Platform with Python AI Services
EDU247 should be implemented as a modular monorepo. All core platform services, APIs, gateways, background workers, integrations, and command-line utilities will be written in Go. Only the dedicated artificial-intelligence services will be written in Python. Go and Python components will communicate through versioned REST or gRPC contracts and asynchronous events, ensuring that AI workloads remain isolated without introducing Python into the core platform.
Repository Structure
edu247/
├── README.md
├── LICENSE
├── Makefile
├── go.work
├── .env.example
├── .editorconfig
├── .gitignore
├── .github/
│   ├── workflows/
│   ├── ISSUE_TEMPLATE/
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── CODEOWNERS
├── apps/
│   ├── web/
│   │   ├── public/
│   │   ├── src/
│   │   └── README.md
│   ├── student-mobile/
│   ├── teacher-mobile/
│   ├── marketplace-mobile/
│   └── admin-console/
├── gateway/
│   ├── cmd/api-gateway/main.go
│   ├── internal/auth/
│   ├── internal/middleware/
│   ├── internal/routing/
│   ├── internal/ratelimit/
│   ├── internal/observability/
│   ├── config/
│   ├── go.mod
│   └── Dockerfile
├── services/
│   ├── identity-service/
│   ├── user-service/
│   ├── learning-service/
│   ├── course-service/
│   ├── assessment-service/
│   ├── assignment-service/
│   ├── progress-service/
│   ├── notification-service/
│   ├── content-service/
│   ├── media-service/
│   ├── search-service/
│   ├── recommendation-orchestrator/
│   ├── developer-service/
│   ├── contribution-service/
│   ├── hackathon-service/
│   ├── marketplace-service/
│   ├── catalog-service/
│   ├── inventory-service/
│   ├── order-service/
│   ├── payment-service/
│   ├── delivery-service/
│   ├── vendor-service/
│   ├── moderation-service/
│   ├── analytics-service/
│   ├── audit-service/
│   ├── admin-service/
│   └── integration-service/
├── ai/
│   ├── ai-gateway/
│   ├── tutor-service/
│   ├── lesson-builder-service/
│   ├── comic-creator-service/
│   ├── infographic-creator-service/
│   ├── mindmap-creator-service/
│   ├── recommendation-service/
│   ├── code-review-service/
│   ├── security-auditor-service/
│   ├── documentation-service/
│   └── shared/
├── packages/
│   ├── go/auth/
│   ├── go/config/
│   ├── go/database/
│   ├── go/events/
│   ├── go/httpx/
│   ├── go/logging/
│   ├── go/observability/
│   ├── go/security/
│   └── go/validation/
├── contracts/
│   ├── openapi/
│   ├── protobuf/
│   ├── events/
│   └── schemas/
├── database/
│   ├── migrations/
│   ├── seeds/
│   ├── queries/
│   └── README.md
├── infrastructure/
│   ├── docker/
│   ├── kubernetes/
│   ├── helm/
│   ├── terraform/
│   ├── monitoring/
│   ├── logging/
│   ├── secrets/
│   └── environments/
├── deployments/
│   ├── local/
│   ├── development/
│   ├── staging/
│   └── production/
├── scripts/
│   ├── bootstrap/
│   ├── build/
│   ├── database/
│   ├── deploy/
│   └── security/
├── tests/
│   ├── contract/
│   ├── integration/
│   ├── end-to-end/
│   ├── performance/
│   ├── security/
│   └── fixtures/
└── docs/
    ├── architecture/
    ├── api/
    ├── ai/
    ├── security/
    ├── operations/
    ├── product/
    ├── decisions/
    └── onboarding/
Standard Go Service Layout
service-name/
├── cmd/service-name/main.go
├── internal/
│   ├── domain/
│   ├── application/
│   ├── repository/
│   ├── transport/http/
│   ├── transport/grpc/
│   ├── messaging/
│   ├── config/
│   └── observability/
├── migrations/
├── tests/
├── config/
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
Standard Python AI Service Layout
ai-service-name/
├── app/
│   ├── main.py
│   ├── api/
│   ├── domain/
│   ├── pipelines/
│   ├── models/
│   ├── prompts/
│   ├── providers/
│   ├── safety/
│   ├── evaluation/
│   ├── storage/
│   └── observability/
├── tests/
├── model-config/
├── datasets/
├── pyproject.toml
├── requirements.lock
├── Dockerfile
└── README.md
Language and Integration Rules
•	Go-only platform boundary: authentication, users, courses, assignments, progress, notifications, marketplace, payments, delivery, administration, analytics, moderation, integrations, and all non-AI workers remain in Go.
•	Python-only AI boundary: AI Tutor, Lesson Builder, Comic Creator, Infographic Creator, Mind Map Creator, recommendation models, code review, security analysis, and documentation generation remain under the AI directory.
•	No direct cross-service database access: each service owns its data and exposes approved APIs or events.
•	Contract-first communication: OpenAPI specifications define public REST interfaces, Protocol Buffers define internal gRPC calls, and versioned schemas define asynchronous events.
•	AI orchestration: Go services submit authorized AI jobs through the AI Gateway; Python services process them and return status, output references, safety results, and model metadata.
•	Asynchronous workloads: comic, infographic, mind-map, document, and media generation use queues so requests can be retried, monitored, cancelled, and scaled independently.
•	Security: every request carries identity, tenant, role, consent, and trace context. Secrets are stored outside source control, and generated content passes moderation and audit checks.
•	Observability: Go and Python services publish consistent logs, metrics, traces, health checks, and audit events.
•	Deployment: every service is containerized independently and deployed through the same development, staging, and production pipeline.
•	Quality gates: unit, contract, integration, end-to-end, performance, accessibility, security, and AI evaluation tests are required before release.
