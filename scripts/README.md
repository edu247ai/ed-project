# EDU247 Scripts

## Overview

`scripts/` is reserved for focused command-line utilities that support
development and operations across EDU247. Scripts should automate repeatable
tasks without replacing service-owned behavior or becoming an undocumented
deployment system.

> **Current status:** The `bootstrap/`, `build/`, `database/`, `deploy/`, and
> `security/` directories are empty placeholders. No scripts are currently
> available in this folder.

The repository does have a web build and deployment workflow, but it is
implemented separately in
[`.github/workflows/deploy-web.yml`](../.github/workflows/deploy-web.yml), not
under `scripts/`.

## Directory layout

```text
scripts/
├── bootstrap/  # Planned developer setup and prerequisite checks
├── build/      # Planned build helpers
├── database/   # Planned safe database utilities
├── deploy/     # Planned deployment helpers
├── security/   # Planned security and dependency checks
└── README.md
```

## Intended categories

| Directory | Intended use |
| --- | --- |
| `bootstrap/` | Verify prerequisites and help developers prepare local environments |
| `build/` | Run focused build or artifact-generation tasks |
| `database/` | Run documented database maintenance or migration tasks |
| `deploy/` | Support deployments where no existing CI/CD workflow is appropriate |
| `security/` | Run security checks and validate configuration |

These categories are planned; they do not imply that setup, database,
deployment, or security automation is implemented.

## Script guidelines

- Give each script one clearly defined task and a descriptive name.
- Document prerequisites, arguments, environment variables, and examples.
- Fail with a non-zero exit status and a useful error message when a task fails.
- Make potentially destructive actions explicit; require confirmation or a
  deliberate flag for deleting data, modifying remote resources, or deploying.
- Never print or commit passwords, access tokens, private keys, or other
  secrets. Read secret values from a protected runtime environment or approved
  secret manager.
- Avoid hard-coded environment-specific endpoints and credentials.
- Prefer existing package-manager and service tooling over duplicating their
  functionality in shell scripts.
- Test scripts with safe local fixtures before using them in CI or against
  shared environments.

## Adding a script

When adding a script:

1. Place it in the category that matches its purpose.
2. Add usage instructions next to the script or in the relevant component
   README.
3. Document required inputs, side effects, and cleanup or rollback steps.
4. Add automated tests or a safe validation mode when practical.
5. Update this README if the directory responsibilities or supported workflows
   change.

## Existing automation

The GitHub Actions workflow at
[`../.github/workflows/deploy-web.yml`](../.github/workflows/deploy-web.yml)
installs web dependencies, runs lint and build checks on pull requests, and
deploys the static web export to GitHub Pages on pushes to `main` or manual
workflow runs. See the [deployment guide](../deployments/README.md) and
[web app guide](../apps/web/README.md) for details.
