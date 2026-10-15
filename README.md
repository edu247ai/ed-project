<p align="center">
  <img src="apps/web/public/logoimg/edu247-logo.png" alt="EDU247" width="180">
</p>

<h1 align="center">EDU247</h1>

<p align="center">
  A connected learning platform for learners, educators, creators, and institutions.
</p>

<p align="center">
  <a href="#what-is-edu247">About</a> ·
  <a href="#get-started">Get started</a> ·
  <a href="#platform-architecture">Architecture</a> ·
  <a href="#project-status">Project status</a>
</p>

---

## What is EDU247?

EDU247 is being built to bring learning, teaching, creator tools, and educational commerce together in one platform. The project is designed as a modular monorepo: a Go-powered platform core, with dedicated Python services for AI workloads.

This repository is actively evolving. The web experience is the most complete application today; platform services and AI integrations described below represent the intended architecture and are not all implemented yet.

## Explore the platform

The web app introduces EDU247 to its main audiences:

- **Learners** — discover learning paths and manage learning activity.
- **Teachers** — explore teaching tools and classroom workflows.
- **Creators and developers** — find ways to build, contribute, and share.
- **Vendors and institutions** — learn about marketplace and organizational experiences.

The current web experience includes public information pages, account and onboarding previews, and role-based dashboard interfaces. See the [web app guide](apps/web/README.md) for routes, design details, and known limitations.

## Get started

### Requirements

- Node.js 22 or later
- npm

### Run the web app locally

```bash
cd apps/web
npm ci
npm run dev
```

Open [http://localhost:3000](http://localhost:3000).

### Check the web app

From `apps/web`, run:

```bash
npm run lint
npm run build
```

The current Go gateway directory is an early scaffold. Platform APIs and Python AI services are not yet available as locally runnable services in this repository.

## Platform architecture

EDU247 separates core platform responsibilities from AI workloads so each can be developed, deployed, and scaled independently.

```mermaid
flowchart LR
    Clients["Web and mobile apps"] --> Gateway["Go API gateway"]
    Gateway --> Platform["Go platform services"]
    Platform --> Contracts["Versioned APIs and events"]
    Contracts --> AIGateway["AI gateway"]
    AIGateway --> AIServices["Python AI services"]
    Platform --> Events["Asynchronous event and job processing"]
    Events --> AIGateway
```

### Go platform boundary

Go is the target language for the API gateway, authentication and identity, learning and course management, assignments, progress, notifications, marketplace, payments, administration, analytics, integrations, and non-AI background workers.

### Python AI boundary

Python is reserved for dedicated AI services, including tutoring, lesson and media generation, recommendation models, code review, security analysis, and documentation generation. Core platform logic remains outside this boundary.

### Service communication

- Public REST APIs are described with versioned OpenAPI contracts.
- Internal gRPC APIs use Protocol Buffers where appropriate.
- Asynchronous messages use versioned event schemas.
- Each service owns its data; services access other service data through approved APIs or events, not direct database access.
- Go services submit authorized AI jobs through the AI gateway. AI services return job status, output references, safety results, and model metadata.
- Long-running generation jobs are designed to support retries, monitoring, cancellation, and independent scaling.

## Project status

| Area | Status |
| --- | --- |
| EDU247 web experience | Implemented as a Next.js, React, and TypeScript app in `apps/web` |
| API gateway | Early Go scaffold in `gateway` |
| Platform microservices | Planned; not yet implemented |
| Python AI services | Planned; not yet implemented |
| Service contracts and event schemas | Part of the target architecture; not yet established throughout the repository |
| GitHub Pages | Web app build and deployment workflow is configured |

**Frontend previews are not production services.** Sign-in, registration, password recovery, and onboarding are not connected to an authentication backend. Newsletter validation happens in the browser only. Dashboards expect role-specific API endpoints; without a separately deployed API, they show an API-unavailable state. GitHub Pages hosts the static web app only, not backend APIs.

## Project structure

The repository is intended to grow into the modular monorepo below. **This is the target structure, not a claim that every directory or service exists today.** The current implementation snapshot is listed separately afterward.

### Target monorepo

```text
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
│   ├── web/                         # Next.js web experience
│   ├── student-mobile/              # Planned student app
│   ├── teacher-mobile/              # Planned teacher app
│   ├── marketplace-mobile/          # Planned marketplace app
│   └── admin-console/               # Planned administration UI
├── gateway/                         # Go API gateway
│   ├── cmd/api-gateway/
│   ├── internal/
│   │   ├── auth/
│   │   ├── middleware/
│   │   ├── routing/
│   │   ├── ratelimit/
│   │   └── observability/
│   ├── config/
│   ├── go.mod
│   └── Dockerfile
├── services/                        # Go platform services
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
├── ai/                              # Python AI workloads only
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
│   └── go/                          # Shared Go libraries
│       ├── auth/
│       ├── config/
│       ├── database/
│       ├── events/
│       ├── httpx/
│       ├── logging/
│       ├── observability/
│       ├── security/
│       └── validation/
├── contracts/
│   ├── openapi/                     # Public REST APIs
│   ├── protobuf/                   # Internal gRPC APIs
│   ├── events/                     # Versioned event definitions
│   └── schemas/                    # Shared data schemas
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
```

### Standard service layouts

Each Go service should be independently buildable and own its domain, persistence, transport, and operational configuration:

```text
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
```

Each Python AI service should keep its API, model/provider integrations, generation pipeline, safety, evaluation, and observability concerns within the AI boundary:

```text
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
```

### Current repository snapshot

Only the web application and an early gateway scaffold are currently present as application code. The other top-level areas in the target tree are planned:

```text
.
├── apps/
│   └── web/                         # Implemented Next.js app
├── gateway/                         # Early Go scaffold
├── .github/
│   ├── CODEOWNERS
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── workflows/deploy-web.yml
├── .env.example
├── .editorconfig
├── LICENSE
├── Makefile
└── README.md
```

## Engineering principles

- **Clear boundaries:** Go owns core platform behavior; Python is limited to AI workloads.
- **Contract first:** APIs and events have explicit, versioned contracts.
- **Service ownership:** each service owns its data and its operational lifecycle.
- **Security and privacy:** enforce identity, tenant, role, and consent on the server; keep secrets out of source control; moderate and audit generated content.
- **Observable by default:** services should expose structured logs, metrics, traces, health checks, and audit events.
- **Quality before release:** build confidence with unit, contract, integration, end-to-end, accessibility, performance, security, and AI evaluation tests as each capability is introduced.

## Deployment

The web app is configured for static export and GitHub Pages deployment. Pull requests to `main` run lint and build checks; pushes to `main` build and deploy the site.

To enable dashboard requests to an API hosted separately, configure the repository Actions variable `NEXT_PUBLIC_EDU247_API_BASE_URL`. The API must allow browser requests from the published site origin and implement the dashboard contract documented in the [web app guide](apps/web/README.md). GitHub Pages itself does not host APIs.

## Contributing

Contributions are welcome as the project takes shape. Before opening a pull request:

1. Keep changes focused and follow the Go/Python service boundaries above.
2. Run the relevant checks for the area you changed.
3. Update the closest documentation when behavior or setup changes.
4. Never commit credentials, secrets, or private user data.

Please use the repository's [pull request template](.github/PULL_REQUEST_TEMPLATE.md) when proposing a change.

## Community

Questions, ideas, or want to meet other contributors? Join us:

<p>
  <a href="https://discord.gg/98HJ7bMfY"><img alt="Discord — Developers & AI community" src="https://img.shields.io/badge/Discord-Developers%20%26%20AI%20community-5865F2?logo=discord&logoColor=white"></a>
  <a href="https://edu247.slack.com/archives/C0C9GBJCVPT"><img alt="Slack — Team & contributors" src="https://img.shields.io/badge/Slack-Team%20%26%20contributors-4A154B?logo=slack&logoColor=white"></a>
  <a href="https://github.com/edu247ai"><img alt="GitHub — Explore EDU247 projects" src="https://img.shields.io/badge/GitHub-Explore%20EDU247%20projects-181717?logo=github&logoColor=white"></a>
</p>

### Follow & Contact

<p>
  <img alt="LinkedIn" src="https://img.shields.io/badge/LinkedIn-0A66C2?logo=linkedin&logoColor=white">
  <img alt="Facebook" src="https://img.shields.io/badge/Facebook-0866FF?logo=facebook&logoColor=white">
  <img alt="X" src="https://img.shields.io/badge/X-000000?logo=x&logoColor=white">
  <img alt="TikTok" src="https://img.shields.io/badge/TikTok-000000?logo=tiktok&logoColor=white">
  <a href="mailto:developer@edu247.co.za"><img alt="Email: developer@edu247.co.za" src="https://img.shields.io/badge/Email-developer%40edu247.co.za-EA4335?logo=gmail&logoColor=white"></a>
</p>

Social profile links will be added when the official EDU247 accounts are available.

## License

See [LICENSE](LICENSE).
