# EDU247 Deployments

## Overview

`deployments/` is reserved for environment-specific deployment configuration
and operational instructions for EDU247 applications and services.

> **Current status:** The `local/`, `development/`, `staging/`, and
> `production/` directories are empty scaffolds. They do not currently contain
> runnable deployment definitions, service manifests, environment profiles, or
> promotion/rollback automation.

The only configured deployment in this repository is the static web
application workflow at
[`.github/workflows/deploy-web.yml`](../.github/workflows/deploy-web.yml).
It builds and deploys `apps/web` to GitHub Pages. It does not deploy backend
services or provision infrastructure.

## Directory layout

```text
deployments/
├── local/        # Planned developer-machine setup
├── development/  # Planned shared development environment
├── staging/      # Planned pre-production environment
├── production/   # Planned production environment
└── README.md
```

## Environment intent

| Directory | Intended use | Current status |
| --- | --- | --- |
| `local/` | Run selected services locally for development and testing | Empty scaffold |
| `development/` | Integrate active development builds | Empty scaffold |
| `staging/` | Validate release candidates against production-like settings | Empty scaffold |
| `production/` | Operate approved, released services for real users | Empty scaffold |

These folders describe intended environments, not deployed or provisioned
environments. Add configuration only after identifying the deployment owner,
target platform, required services, and operational responsibilities.

## Existing web deployment

The GitHub Actions workflow:

- Runs on pull requests to `main`, pushes to `main`, and manual dispatch.
- Uses Node.js 22 and installs web dependencies with `npm ci`.
- Runs `npm run lint` and builds a static export of `apps/web`.
- Uploads and deploys the `apps/web/out` artifact for non-PR runs.

To enable the workflow, configure GitHub Pages to use **GitHub Actions** as its
deployment source. The workflow sets the repository base path automatically.

An optional repository Actions variable,
`NEXT_PUBLIC_EDU247_API_BASE_URL`, provides the origin for separately hosted
dashboard APIs. The API must implement the contract documented in the
[web app guide](../apps/web/README.md) and allow browser requests from the
published site origin. GitHub Pages serves static frontend assets only; it
does not host backend APIs.

## Deployment principles

- Keep environment configuration separate and avoid copying secrets into Git.
- Use approved secret managers or protected runtime variables for secret
  values.
- Promote the same reviewed build artifact between environments where the
  platform allows it; avoid rebuilding different code for production.
- Validate configuration and test relevant health checks before promotion.
- Document ownership, dependencies, rollback steps, and data migration effects.
- Do not claim a service or environment is deployed until its configuration and
  runtime are operational and verified.

## Adding deployment configuration

When adding a deployment:

1. Document the target platform, supported environment, and responsible owner.
2. Add a service-specific deployment guide or README next to the configuration.
3. Keep non-secret settings in version control and reference secrets by name.
4. Add CI validation for deployment manifests and scripts where practical.
5. Define rollout, health-check, failure, and rollback procedures.
6. Update this README and the service documentation to match what is actually
   deployed.

## Roadmap

The deployment scaffold can be populated incrementally, starting with a local
development setup for one implemented service. Shared development, staging,
and production deployment should follow only after service dependencies,
security requirements, monitoring, backups, and operational ownership are
defined.
