import { Building2, Network, ShieldCheck, Users } from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

const content: InformationPageContent = {
  audience: "EDU247 FOR INSTITUTIONS",
  activeAudience: "For institutions",
  title: "A connected foundation",
  highlight: "for education teams.",
  description: "EDU247 is a planned education platform for schools, colleges, and training centers. It is designed to connect learners, educators, content, and progress with role-based administration, helping education teams plan a responsible digital learning experience around their institutional needs.",
  dashboardHref: "/Dashboard/manager-dashboard",
  dashboardLabel: "Open manager dashboard",
  roleHref: "/Dashboard/superadmin-dashboard",
  roleLabel: "Platform administration view",
  illustration: {
    center: Building2,
    nodes: [ShieldCheck, Network, Users, Building2],
  },
  trustSection: true,
  features: [
    { title: "A shared learning space", description: "Bring learners, teachers, course content, progress, and learning resources into a connected environment.", icon: Building2 },
    { title: "Clear roles and oversight", description: "Plan role-based access, content review, audit trails, and operational reporting with defined responsibilities.", icon: ShieldCheck },
    { title: "Institution dashboards", description: "Provide managers with scoped views of team activity, content review, and reports relevant to their work.", icon: Network },
    { title: "Support for education teams", description: "Coordinate learners, educators, administrators, and institution leaders around shared goals.", icon: Users },
  ],
  steps: [
    { title: "Discuss your needs", description: "Identify learner groups, existing systems, accessibility requirements, and institutional goals." },
    { title: "Plan onboarding", description: "Define responsibilities, roles, data access, integrations, and a staged rollout before implementation." },
    { title: "Support your teams", description: "Use institution and manager views to review relevant activity and respond to operational needs." },
  ],
  related: [{ label: "Teaching", href: "/teach" }, { label: "Learning", href: "/learn" }, { label: "Trust and safety", href: "/institutions#trust" }],
};

export function InstitutionPage() {
  return <PageLayout content={content} />;
}
