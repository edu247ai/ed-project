import { BookOpenCheck, BrainCircuit, GraduationCap, Package, Sparkles, Users } from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

const content: InformationPageContent = {
  audience: "TEACH WITH EDU247",
  activeAudience: "For educators",
  title: "Make teaching",
  highlight: "feel limitless.",
  description: "Plan, create, and share learning experiences in one connected educator journey. EDU247 is being shaped to bring course creation, class management, responsible AI lesson support, and opportunities to share useful teaching resources together—while keeping educators in control of what reaches learners.",
  dashboardHref: "/Dashboard/teacher-dashboard",
  dashboardLabel: "Teach now",
  roleHref: "/teacher",
  roleLabel: "Explore the teacher experience",
  illustration: {
    center: GraduationCap,
    nodes: [BookOpenCheck, Users, Sparkles, BrainCircuit],
  },
  features: [
    { title: "Create course content", description: "Develop structured courses, clear lessons, assessments, and reusable classroom resources.", icon: BookOpenCheck },
    { title: "AI lesson builder", description: "Use planned AI assistance to shape lesson drafts, then review and adapt every resource yourself.", icon: BrainCircuit },
    { title: "Manage classes", description: "Organize learners, assignments, submissions, and learning progress in a role-based workspace.", icon: Users },
    { title: "Earn from resources", description: "Prepare eligible teaching materials to share through the marketplace as vendor features become available.", icon: Package },
  ],
  steps: [
    { title: "Bring an idea", description: "Start with a subject, learning goal, or real classroom need you want to address." },
    { title: "Create and review", description: "Build a lesson or course, check it for accuracy and accessibility, then prepare it to share." },
    { title: "Support your class", description: "Use the teacher dashboard to organize classes, follow learner activity, and plan support." },
  ],
  related: [{ label: "Learning", href: "/learn" }, { label: "Teacher role page", href: "/teacher" }, { label: "Marketplace", href: "/marketplace" }],
};

export function TeachingPage() {
  return <PageLayout content={content} />;
}
