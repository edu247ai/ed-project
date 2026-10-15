# EDU247 Web

The EDU247 public homepage is built with Next.js App Router, React, and TypeScript. It presents the platform's global learning vision for learners, teachers, creators, and institutions.

## Design system

The platform uses IBM Plex Sans for interface text, IBM Plex Serif for editorial quotations, and IBM Plex Mono for code. Surfaces use pure white (`#FFFFFF`) with no gradient or tinted page backgrounds.

The shared palette is defined in `src/app/globals.css`:

| Token | Value | Use |
| --- | --- | --- |
| `--ed-navy` | `#173E59` | Primary buttons and strong interface text |
| `--ed-teal` | `#087F72` | Links, actions, and primary icon accents |
| `--ed-aqua` | `#27A69A` | Secondary AI and progress accent |
| `--ed-violet` | `#7256C8` | Secondary highlight |
| `--ed-gold` | `#9A6A12` | Caution and warm highlight |
| `--ed-border` | `#DCE5E9` | White card and control borders |

Use the shared `DesignCard` component from `src/components/DesignCard.tsx` for card surfaces. It applies the shared white surface, border, radius, spacing, and shadow tokens.

## Run locally

```sh
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000). The course search filters the featured learning paths on the homepage and submits to the learning page. The homepage header and footer are shared across the platform, role, authentication, and dashboard pages. The design adapts to desktop, tablet, and mobile screens. The supplied EDU247 logo is also used as the browser tab and Apple touch icon.

## Platform pages

The information pages are `/learn`, `/teach`, `/developers`, `/marketplace`, `/community`, and `/institutions`. Role information pages are `/student`, `/teacher`, `/vendor`, and `/developer`. Each page explains its audience, purpose, features, and a suggested workflow, and links to the matching dashboard. Features are presented as platform plans where they are not yet implemented.

The homepage header and footer are reused from `src/components/SiteHeader.tsx` and `src/components/SiteFooter.tsx`; information page components live under `src/components/pages/`. The dashboard application also uses those shared site components around its role-specific workspace.

The six delivery phases and week ranges are an initial planning model, not a committed schedule; scope and estimates remain subject to discovery and validation.

## Account and onboarding preview

The frontend preview includes `/auth/login`, `/auth/register`, `/auth/forgot-password`, and `/auth/onboarding`. These screens are not connected to an authentication service: email sign-in, LinkedIn/GitHub/Facebook/X authentication, account creation, and password recovery do not contact a server. Registration and onboarding keep only a temporary draft in the current browser tab's session storage; passwords are never stored.

The homepage newsletter is a frontend preview until a subscription service is connected: it validates the email address in the browser but does not send or save it.

## Dashboards

The role dashboards are available at `/Dashboard/student-dashboard`, `/Dashboard/teacher-dashboard`, `/Dashboard/vendor-dashboard`, `/Dashboard/developer-dashboard`, `/Dashboard/manager-dashboard`, and `/Dashboard/superadmin-dashboard`. Each requests its role-specific data from `GET /api/v1/dashboard/{role}`. Set `NEXT_PUBLIC_EDU247_API_BASE_URL` to the API origin when the API is hosted separately; if unset, the frontend requests the endpoint from the current origin. Requests include browser credentials for a backend-managed session.

The frontend expects a JSON response with the requested role, an optional user display name, metric cards, and role-specific sections. Sections support `list`, `table`, `progress`, `summary`, and `chart` kinds, with optional layout sizes `full`, `half`, `third`, or `quarter`. For sections linked from the sidebar, use kebab-case IDs matching the role's sidebar labels (for example, `my-courses`, `assignments`, or `service-health`). For example:

```json
{
  "role": "student",
  "user": { "name": "Learner" },
  "metrics": [
    { "id": "courses", "label": "Active courses", "value": 4, "change": "2 this term", "trend": "up" }
  ],
  "sections": [
    {
      "id": "my-courses",
      "title": "My Courses",
      "kind": "list",
      "size": "half",
      "items": [{ "title": "Course title", "subtitle": "Module 2", "status": "In progress" }]
    },
    {
      "id": "activity",
      "title": "Learning activity",
      "kind": "chart",
      "size": "half",
      "series": [{ "name": "Study time", "values": [{ "label": "Mon", "value": 2 }] }]
    }
  ]
}
```

Dashboards show loading, retryable API/contract errors, and empty-data states instead of placeholder metrics. This repository does not currently contain implemented dashboard API endpoints; until the backend provides the endpoints and response contract above, the dashboard error state is expected. The API must authenticate the session and enforce role-level authorization server-side, especially for manager and superadmin data; the frontend is not an access-control boundary.

## GitHub Pages deployment

The workflow at `../../.github/workflows/deploy-web.yml` runs the web lint and static build for pull requests to `main`, and deploys the site to GitHub Pages after pushes to `main` (or a manual workflow run). In the GitHub repository, open **Settings → Pages** and select **GitHub Actions** as the build and deployment source. After the first successful deployment, the site is available at `https://<owner>.github.io/<repository>/`.

The workflow sets the repository subpath automatically and exports the app as a static site. To connect dashboards to an API hosted separately, add the repository Actions variable `NEXT_PUBLIC_EDU247_API_BASE_URL` under **Settings → Secrets and variables → Actions → Variables**. The API must allow browser requests from the published site origin and implement the dashboard endpoints. GitHub Pages only serves static frontend files; it does not host the Next.js server or backend APIs. Without a separately deployed API, dashboards show the explicit API-unavailable state.

## Build

```sh
npm run build
npm start
```

The supplied EDU247 logo, homepage image, homepage wireframe, and dashboard wireframes are in `public/logoimg/`, `public/homepageimg/`, `public/homepage/`, and `public/dasboards/`. The homepage uses a responsive design and the shared global design tokens.
