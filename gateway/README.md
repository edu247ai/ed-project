# EDU247 API Gateway

## Overview

`gateway/` is the planned Go entry point for client requests to EDU247 backend
services. It is intended to route requests to approved service APIs and provide
a consistent place for cross-cutting HTTP concerns.

> **Current status:** This directory is an incomplete scaffold. The Go source
> file, `go.mod`, and `Dockerfile` are empty, and the `config/` and `internal/`
> directories do not contain implementations. There is no runnable gateway,
> configured route, authentication middleware, or deployed container yet.

The gateway must not own business-domain behavior or access service databases
directly. Each backend service remains responsible for its own domain,
authorization rules, and data.

## Intended responsibilities

When implemented, the gateway may provide:

- Routing to explicitly configured backend API endpoints
- Consistent request IDs, structured access logs, and HTTP error responses
- Authentication integration and forwarding of verified identity context
- Rate limiting and request size/time limits
- Health and readiness endpoints for operations

These are planned capabilities only. Do not rely on them being available or
secure until implemented, tested, and deployed.

## Directory layout

```text
gateway/
├── cmd/
│   └── api-gateway/
│       └── main.go       # Planned process entry point
├── config/               # Planned environment-based configuration
├── internal/
│   ├── auth/             # Planned authentication integration
│   ├── middleware/       # Planned HTTP middleware
│   ├── observability/    # Planned logs, metrics, and tracing
│   ├── ratelimit/        # Planned request limits
│   └── routing/          # Planned route and proxy handling
├── Dockerfile            # Empty placeholder
├── go.mod                # Empty placeholder
└── README.md
```

## Development

The project currently cannot be built or tested: `go.mod` does not yet contain
a Go module declaration, and `cmd/api-gateway/main.go` is empty. Once the Go
module and implementation are added, use the module's documented commands. The
expected Go checks will be:

```sh
go test ./...
go build ./cmd/api-gateway
```

These commands are not expected to pass until the scaffold is implemented.
Document the required Go version and exact local run command here when the
service becomes runnable.

## Configuration and security

Configuration should be read from the deployment environment or an approved
configuration source. Do not commit credentials, signing keys, or tokens.

Before exposing routes, define:

- Which clients and services may call each route
- Authentication and authorization responsibilities
- Request and response limits, timeouts, and error behavior
- How trusted identity context is conveyed to downstream services
- Logging rules that avoid credentials and sensitive request bodies
- Health-check behavior and operational metrics

The gateway must not be treated as the only authorization boundary. Downstream
services must validate access to their own resources.

## Integrating backend services

Add routes only for services with a documented, versioned API contract. Keep
service-specific logic inside the owning service; the gateway should perform
transport-level routing and cross-cutting request handling only. Do not route
to planned services as if they were already available.

## Deployment

The `Dockerfile` is empty, and no gateway deployment configuration exists.
Containerization and deployment instructions should be added after the service
has a working build, configuration, health checks, and automated tests. Do not
deploy this scaffold as a production gateway.

## Implementation checklist

1. Define the Go module path and supported Go version.
2. Implement a minimal HTTP server with explicit health and readiness behavior.
3. Add versioned route configuration and one tested downstream integration.
4. Implement and test authentication, authorization context, limits, and
   consistent errors before exposing the gateway externally.
5. Add structured observability without logging secrets or private payloads.
6. Add a non-empty, multi-stage container build and deployment documentation.
