# EDU247 Test Strategy

## Overview

`tests/` is reserved for repository-level tests that validate contracts,
integration boundaries, and end-to-end behavior across independently
developed services.

> **Current status:** The `contract/`, `end-to-end/`, `fixtures/`,
> `integration/`, `performance/`, and `security/` directories are empty
> placeholders. There are no repository-level test suites or shared fixtures
> here yet.

Tests that belong to one component should live with that component. For
example, the Python AI service starter tests are under each service's `tests/`
directory, while the web app currently has lint and build checks but no
dedicated test script.

## Directory layout

```text
tests/
├── contract/      # Planned API, event, and schema compatibility tests
├── end-to-end/    # Planned multi-service user workflow tests
├── fixtures/      # Planned sanitized, reusable cross-service test data
├── integration/   # Planned service-to-service and external adapter tests
├── performance/   # Planned load, latency, and throughput tests
├── security/      # Planned security-focused test suites
└── README.md
```

## Test categories

| Directory | Scope |
| --- | --- |
| `contract/` | Validate versioned OpenAPI, Protobuf, event, and schema contracts |
| `end-to-end/` | Exercise user-visible workflows across deployed or test instances |
| `fixtures/` | Store small, synthetic or sanitized data that can be safely shared by repository-level tests |
| `integration/` | Verify service boundaries, adapters, persistence, and event flows |
| `performance/` | Measure agreed latency, throughput, and resource-use targets |
| `security/` | Test authentication, authorization, validation, and security controls |

Keep unit tests next to the service or package whose behavior they verify.
Place tests here only when they need to validate behavior across component
boundaries or use repository-level tooling.

## Running available checks

Run component checks from that component's own directory and follow its README.
For example, the AI model service scaffolds document:

```sh
python -m pytest
```

The web application's available checks are:

```sh
cd apps/web
npm ci
npm run lint
npm run build
```

The root `tests/` directory does not currently provide a single test runner or
repository-wide test command. Do not assume that a check has run merely because
it is listed in a plan or workflow.

## Writing reliable tests

- Test observable behavior and service contracts rather than private
  implementation details.
- Keep tests deterministic, independent, and safe to run repeatedly.
- Use synthetic or sanitized fixtures; never commit credentials, private user
  data, or production dumps.
- Isolate external services and network dependencies behind mocks or
  disposable test environments.
- State required services, environment variables, and cleanup steps in the
  relevant README.
- Avoid destructive tests against shared or production environments.
- For performance tests, record workload, environment, and measurable
  acceptance thresholds.
- For security tests, document the control being checked and avoid including
  exploitable payloads or real secrets in committed test data.

## Adding repository-level tests

Before adding a test suite here:

1. Confirm that the behavior crosses service or package boundaries.
2. Identify its contract, test environment, and dependencies.
3. Keep test data minimal and sanitized under `fixtures/` when reuse is needed.
4. Document how to run the suite and what failures mean.
5. Add it to CI only when it is deterministic and its required infrastructure
   is provisioned explicitly.
6. Update this README to describe implemented suites and supported commands.

## Current quality status

The directories in this folder are structural placeholders, not evidence of
test coverage. Each service and application README is the source of truth for
the checks currently available for that component.
