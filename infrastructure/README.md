# EDU247 Infrastructure

## Overview

`infrastructure/` is reserved for deployment, provisioning, environment,
security, logging, and monitoring configuration for EDU247. These resources
are intended to support the independently deployable applications and services
in the repository.

> **Current status:** This directory currently contains only this README and
> empty scaffold folders. There are no Docker images or Compose files,
> Kubernetes manifests, Helm charts, Terraform modules, environment profiles,
> logging/monitoring configurations, or secret-manager integrations here.
> Descriptions below document intended scope, not available infrastructure.

The current web deployment workflow is maintained separately at
[`.github/workflows/deploy-web.yml`](../.github/workflows/deploy-web.yml). It
builds and deploys the static web app to GitHub Pages; it does not deploy
backend services or provision cloud infrastructure.

## Directory layout

```text
infrastructure/
├── docker/       # Planned container build and local development definitions
├── environments/ # Planned non-secret environment-specific configuration
├── helm/         # Planned Helm charts and release values
├── kubernetes/   # Planned Kubernetes manifests and policies
├── logging/      # Planned logging conventions and collection configuration
├── monitoring/   # Planned metrics, dashboards, and alerts
├── secrets/      # Planned secret-manager references and policy templates
├── terraform/    # Planned infrastructure-as-code modules
└── README.md
```

## Intended responsibilities

- **`docker/`** — Build and run container images, including local development
  compositions when required.
- **`environments/`** — Hold non-secret environment settings, service
  endpoints, and deployment profiles.
- **`helm/`** — Define versioned Helm charts and environment-specific values.
- **`kubernetes/`** — Hold Kubernetes resources such as Deployments, Services,
  Ingresses, policies, and autoscaling configuration.
- **`logging/`** — Define structured log conventions, routing, and retention.
- **`monitoring/`** — Define health checks, metrics, dashboards, and alerts.
- **`secrets/`** — Hold templates and references for an approved secret manager;
  never store secret values, credentials, tokens, or private keys here.
- **`terraform/`** — Define reviewed, repeatable provisioning for cloud
  resources such as networks, storage, databases, and compute.

Choose one authoritative deployment mechanism for each resource. Avoid
duplicating or conflicting configuration between Terraform, Helm, raw
Kubernetes manifests, and environment files.

## Security and operations

- Keep development, staging, and production resources and credentials isolated.
- Store secret values only in an approved secret manager or protected runtime
  environment; commit no secrets to source control.
- Apply least privilege to service identities, operators, and deployment
  automation.
- Require review and validation before applying infrastructure changes.
- Document resource ownership, data retention, backup and recovery, and
  teardown behavior before introducing stateful production resources.
- Do not treat an empty scaffold directory as evidence that a control or
  deployment capability is active.

## Adding infrastructure

When implementing infrastructure here:

1. Add a README alongside each substantial component, including purpose,
   prerequisites, usage, and rollback or cleanup steps.
2. Keep environment-specific values separate from reusable templates.
3. Use placeholders or secret-manager references instead of secret values.
4. Add validation to CI where practical (for example, formatting and
   configuration validation).
5. Update this README and the relevant service deployment documentation to
   describe what is actually implemented.

## Roadmap

The scaffold may be populated incrementally, beginning with a local development
environment and a minimal deployment for one implemented service. Add
provisioning, observability, and production controls only when their ownership,
security requirements, and operational procedures are defined.
