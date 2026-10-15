# EDU247 Web

The EDU247 public homepage is built with Next.js App Router, React, and TypeScript. It presents the end-to-end platform vision: participant roles, learning and AI tools, developer and vendor ecosystems, trust and administration, connected user journeys, and an indicative delivery roadmap.

## Run locally

```sh
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000). The course search filters the featured learning paths on the homepage. The design adapts to desktop, tablet, and mobile screens. The supplied EDU247 logo is also used as the browser tab and Apple touch icon.

The six delivery phases and week ranges are an initial planning model, not a committed schedule; scope and estimates remain subject to discovery and validation.

## Account and onboarding preview

The frontend preview includes `/auth/login`, `/auth/register`, `/auth/forgot-password`, and `/auth/onboarding`. These screens are not connected to an authentication service: sign-in, account creation, and password recovery do not contact a server. Registration and onboarding keep only a temporary draft in the current browser tab's session storage; passwords are never stored.

The homepage newsletter form is also a frontend preview until a subscription service is connected. It validates the email address in the browser but does not send or save it.

## Build

```sh
npm run build
npm start
```

The supplied EDU247 logo and homepage wireframe are in `public/logoimg/` and `public/homepageimg/`. The homepage follows the wireframe's main sections while using a responsive visual system and the project's teal, navy, and gold brand colors.
