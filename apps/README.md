# EDU247 Applications

The `apps/` directory contains EDU247's user-facing applications. Each app is
an independent frontend project with its own dependencies, run commands, tests,
and release process. Apps communicate with backend services through documented
APIs; they must not contain service-owned business or data-access logic.

## Applications

| App | Audience and intended use | Current status |
| --- | --- | --- |
| [`web/`](./web/) | Public EDU247 site and planned learner, teacher, vendor, and administrator experiences | Implemented as a Next.js application; backend-dependent features are previews |
| `student-mobile/` | Learner lessons, practice, progress, and tutoring | Directory placeholder; not implemented |
| `teacher-mobile/` | Teacher planning, classroom content, and learning tools | Directory placeholder; not implemented |
| `marketplace-mobile/` | Browse and access learning resources | Directory placeholder; not implemented |
| `admin-console/` | Administrative operations, reporting, and service management | Directory placeholder; not implemented |

The planned responsibilities describe product direction, not existing
functionality. Update this table as each app is implemented.

## Web application

The web app is built with Next.js App Router, React, and TypeScript. See the
[web app guide](./web/README.md) for its pages, design system, development
commands, backend contracts, and deployment details.

Run locally:

```sh
cd apps/web
npm ci
npm run dev
```

Open [http://localhost:3000](http://localhost:3000).

Run the web checks:

```sh
npm run lint
npm run build
```

Authentication, newsletter subscriptions, and dashboard data are not backed by
implemented APIs in this repository. The web app documents these preview
behaviors and the expected dashboard API contract. Configure
`NEXT_PUBLIC_EDU247_API_BASE_URL` only when connecting to a separately hosted
API that implements that contract. The frontend is not an authorization
boundary; backend services must enforce identity, roles, and permissions.

## App boundaries and integration

- Keep each app independently buildable and deployable.
- Apps own presentation and client-side interaction; backend services own
  authorization, persistence, and business rules.
- Integrate with backend services through versioned APIs or events exposed for
  client use. Never import another service's internal code or access its
  database directly.
- Keep endpoint configuration environment-based. Do not commit credentials,
  tokens, or private user data.
- Show explicit loading, empty, and retryable error states when APIs are
  unavailable; do not display fabricated success data.
- Request only the user data needed by the current experience and handle it
  according to EDU247's privacy and security requirements.

## User experience and accessibility

Use the web app's shared visual language as a reference when building other
apps: consistent colors, typography, spacing, controls, and navigation help
users move between experiences. Follow platform conventions for mobile
navigation and interaction rather than copying web layouts directly.

All apps should support keyboard and assistive-technology use where applicable,
maintain readable contrast and text sizing, provide visible focus states, and
communicate validation and errors in text—not color alone.

## Developing an app

Each implemented app should include its own README with:

- Purpose and supported platforms
- Prerequisites and dependency installation
- Local development and build commands
- Relevant test and lint commands
- Required environment variables, without secret values
- API integrations and expected backend availability
- Deployment instructions and known limitations

Before submitting changes, run the checks documented by the app, test the
affected user flows, and update app documentation when setup or behavior
changes. Mobile directories are placeholders and do not yet have runnable
projects or commands.
