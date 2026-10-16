# EDU247 Go Packages

## Overview

`packages/` is reserved for reusable Go libraries that may be adopted by EDU247
Go services. It is not a deployable service. Services should depend on a
package only when they need its capability; package use must not create
unnecessary coupling between independently deployable services.

> **Current status:** The package directories below are empty placeholders.
> There is no Go module or shared library implementation here yet. The listed
> responsibilities describe intended scope, not available APIs.

## Package layout

```text
packages/
├── auth/           # Planned authentication primitives
├── config/         # Planned configuration helpers
├── database/       # Planned database clients and utilities
├── events/         # Planned event types and helpers
├── httpx/          # Planned HTTP client and response utilities
├── logging/        # Planned structured logging helpers
├── observability/  # Planned metrics and tracing helpers
├── security/       # Planned cryptographic/security primitives
├── validation/     # Planned shared input validation
└── README.md
```

The repository's target architecture describes these packages under
`packages/go/`. The current empty placeholders are directly under `packages/`;
when implementation begins, align the directory layout and Go module path with
the chosen package publishing and versioning strategy.

## Intended package responsibilities

| Package | Intended scope |
| --- | --- |
| `auth/` | Reusable authentication token or credential helpers; authorization policies remain owned by each service |
| `config/` | Typed configuration loading and validation without embedding environment-specific secrets |
| `database/` | Optional database connection and lifecycle helpers; schemas and data ownership remain with services |
| `events/` | Shared event envelope types and serialization helpers aligned with versioned contracts |
| `httpx/` | HTTP client defaults, timeouts, request/response encoding, and transport-level error helpers |
| `logging/` | Structured logging setup and common fields without logging sensitive values |
| `observability/` | Instrumentation helpers for metrics and tracing, with explicit opt-in by consumers |
| `security/` | Carefully reviewed reusable cryptographic primitives; service authorization decisions remain local |
| `validation/` | Small reusable validation primitives for well-defined shared formats |

These packages should stay focused and dependency-light. Do not put
service-specific workflows, business rules, API handlers, or private data
models here.

## Dependency and compatibility rules

- Keep each service independently buildable and deployable.
- Services may import stable shared libraries, but must not import another
  service's internal source code.
- Avoid cyclic package dependencies and unnecessary third-party dependencies.
- Keep public APIs small and document exported types and functions.
- Treat changes to exported behavior as compatibility changes; document
  breaking changes and follow the versioning policy chosen for the Go module.
- Add unit tests for behavior and compatibility tests where consumers depend
  on a contract.
- Do not centralize service-owned authorization or security policy in a helper
  package; shared primitives do not replace service-level controls.

## Development

There is no `go.mod` in `packages/` yet, so the shared Go libraries cannot
currently be built or tested as a module. When the first package is
implemented, add a Go module, define the supported Go version, and document
module-specific commands here. At minimum, run:

```sh
go test ./...
go vet ./...
```

Do not add imports to these placeholders until package code and a stable
module path exist.

## Adding a package

Before adding a shared library:

1. Confirm that the capability is genuinely reusable and not service-owned.
2. Define a small public API and document its consumers.
3. Add tests, package documentation, and any required compatibility guidance.
4. Review its dependency and security implications.
5. Update this README to describe implemented behavior rather than planned
   scope.
