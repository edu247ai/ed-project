# EDU247 Platform Services

## Overview

`services/` is the planned home for EDU247's Go platform microservices. Each
service is intended to own a distinct domain, API, data lifecycle, and
deployment lifecycle. Services should communicate through versioned APIs or
events rather than importing one another's internal code or accessing one
another's databases.

> **Current status:** The service directories listed below are empty
> placeholders. No Go modules, endpoints, domain logic, database integrations,
> tests, or service deployments are implemented here yet. The descriptions
> below document intended responsibilities, not available features.

The Go API gateway is separately scaffolded under [`../gateway/`](../gateway/)
and is also not yet runnable.

## Service inventory

| Service | Intended responsibility |
| --- | --- |
| `admin-service/` | Administrative operations, platform configuration, and access management |
| `analytics-service/` | Product, learning, and marketplace reporting |
| `assessment-service/` | Assessment definitions, rubrics, and grading workflows |
| `assignment-service/` | Assignment creation, submissions, deadlines, and feedback |
| `audit-service/` | Audit events and compliance records |
| `catalog-service/` | Course, content, and marketplace catalogs |
| `content-service/` | Educational content metadata, versions, and lifecycle |
| `contribution-service/` | User-contributed content, projects, and uploads |
| `course-service/` | Course structure, creation, and enrollment rules |
| `delivery-service/` | Digital product delivery, access, and licensing |
| `developer-service/` | Developer-facing platform tools and integration management |
| `hackathon-service/` | Hackathon events, teams, and submissions |
| `identity-service/` | Authentication, identity, and federation integrations |
| `integration-service/` | Connections to external providers and systems |
| `inventory-service/` | Marketplace inventory and availability |
| `learning-service/` | Learning activities, paths, and curriculum coordination |
| `marketplace-service/` | Listings, pricing, vendor marketplace rules, and discovery |
| `media-service/` | Media metadata and processing coordination |
| `moderation-service/` | Content review, moderation policy, and review workflows |
| `notification-service/` | Email, push, and in-app notification coordination |
| `order-service/` | Order lifecycle, tracking, and refunds |
| `payment-service/` | Payment, subscription, and vendor payout workflows |
| `progress-service/` | Learner progress and progress reporting |
| `recommendation-orchestrator/` | Coordination of recommendation requests across approved capabilities |
| `search-service/` | Search indexing and query APIs for platform content |
| `user-service/` | User profiles and account preferences |
| `vendor-service/` | Vendor onboarding, profile, and product management |

These boundaries are an initial domain map. Before implementing a service,
confirm its ownership and contracts to avoid overlapping responsibilities.

## Standard service layout

Each implemented Go service should be independently buildable and follow the
repository's Go service convention:

```text
<service-name>/
├── cmd/<service-name>/main.go
├── internal/
│   ├── domain/          # Service-owned entities and rules
│   ├── application/     # Service use cases
│   ├── repository/      # Persistence interfaces and implementations
│   ├── transport/http/  # HTTP handlers and routing
│   ├── transport/grpc/  # gRPC handlers, when needed
│   ├── messaging/       # Event consumers and publishers, when needed
│   ├── config/          # Service configuration
│   └── observability/   # Service instrumentation
├── tests/
├── config/
├── go.mod
├── Dockerfile
└── README.md
```

Only add components the service needs. Keep implementation details under
`internal/` unless a deliberately designed public Go library is required.
Contracts shared between services belong in the versioned contract locations,
not in another service's source tree.

## Service boundaries

- A service owns its data, domain rules, API, and operational lifecycle.
- Use documented, versioned OpenAPI, gRPC, or event contracts for integration.
- Do not read or write another service's database directly.
- Keep external systems behind adapters owned by the service using them.
- Authenticate and authorize requests at the owning service; a gateway does
  not replace service-level authorization.
- Make errors, timeouts, retries, and idempotency behavior explicit.
- Store secrets outside source control and load them from approved runtime
  configuration or secret-management systems.
- Add structured logs, health checks, metrics, and traces as part of service
  implementation; do not log credentials or sensitive payloads.

## Developing a service

Before implementation:

1. Define the service's responsibilities, API consumers, and data ownership.
2. Identify required contracts and dependencies on other services.
3. Create a service README with setup, run, test, configuration, and deployment
   instructions.
4. Implement a small vertical slice with unit and integration tests.
5. Add the Go module and deployment configuration required for independent
   builds and releases.
6. Document operational requirements, including migrations, health checks, and
   rollback behavior.

Until those files and implementations are added, the service folders cannot be
built, run, or tested.

## Related areas

- [`../gateway/`](../gateway/) — planned Go API gateway scaffold
- [`../contracts/`](../contracts/) — planned API, event, and schema contracts
- [`../packages/`](../packages/) — planned reusable Go libraries; see its README
  for the current empty-scaffold status
- [`../deployments/`](../deployments/) — planned environment deployment
  configuration

## Current platform status

The platform services under `services/` are planned and not yet implemented.
Keep this README and each future service README aligned with verified
implementation, test, and deployment status as the platform develops.
