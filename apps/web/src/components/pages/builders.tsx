import { Code2, Network, ShieldCheck, Users } from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

const content: InformationPageContent = {
  audience: "BUILD WITH EDU247",
  activeAudience: "For builders",
  title: "Build tools that",
  highlight: "help learning grow.",
  description: "Explore the planned developer ecosystem for extending EDU247 through APIs, SDKs, documentation, and responsible integrations. The developer experience is being shaped to help builders understand platform capabilities, test ideas safely, and contribute tools that support education.",
  dashboardHref: "/Dashboard/developer-dashboard",
  dashboardLabel: "Build tools",
  roleHref: "/developer",
  roleLabel: "Explore the developer experience",
  illustration: {
    center: Code2,
    nodes: [Network, ShieldCheck, Users, Code2],
  },
  features: [
    { title: "APIs and SDKs", description: "Versioned APIs and developer libraries are planned to help teams connect useful services and learning tools.", icon: Code2 },
    { title: "Developer workspace", description: "A future portal for registered applications, API usage, technical documentation, and service health.", icon: Network },
    { title: "Hackathons and contributions", description: "Collaborate on open education ideas, community projects, and build events when they are announced.", icon: Users },
    { title: "Security by design", description: "Plan integrations around scoped access, careful data handling, and clear security practices.", icon: ShieldCheck },
  ],
  steps: [
    { title: "Explore the platform", description: "Review the project structure and identify a useful integration or contribution." },
    { title: "Build and test", description: "Use published API contracts and sandbox tools when they become available." },
    { title: "Share your work", description: "Contribute improvements and join community events." },
  ],
  related: [{ label: "Community", href: "/community" }, { label: "Institutions", href: "/institutions" }, { label: "Developer role page", href: "/developer" }],
};

export function BuildersPage() {
  return <PageLayout content={content} />;
}