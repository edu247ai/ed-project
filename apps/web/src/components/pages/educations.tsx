import { Award, BookOpenCheck, BrainCircuit, GraduationCap, Sparkles } from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

const content: InformationPageContent = {
  audience: "LEARN WITH EDU247",
  activeAudience: "For learners",
  title: "Build skills for",
  highlight: "what comes next.",
  description: "Explore a planned learning experience where courses, guided paths, AI-supported study, and progress tools come together. Start with a subject that interests you, build practical skills at your own pace, and work toward goals that matter in your studies, career, and everyday life.",
  dashboardHref: "/Dashboard/student-dashboard",
  dashboardLabel: "Start learning",
  roleHref: "/student",
  roleLabel: "Explore the student experience",
  illustration: {
    center: BookOpenCheck,
    nodes: [BrainCircuit, GraduationCap, Award, Sparkles],
  },
  features: [
    { title: "Learning categories", description: "Explore STEM, technology, business, creative arts, languages, and professional skills.", icon: BookOpenCheck },
    { title: "AI tutor", description: "Get explanations, examples, and practice support while staying in control of your learning.", icon: BrainCircuit },
    { title: "Paths and certificates", description: "Follow a sequence of learning activities and work toward credentials as these features launch.", icon: GraduationCap },
    { title: "Creative study tools", description: "Use planned AI tools to organize ideas, make visual summaries, and deepen understanding.", icon: Sparkles },
  ],
  steps: [
    { title: "Choose a direction", description: "Start with a subject, practical skill, or personal goal that interests you." },
    { title: "Build a learning rhythm", description: "Move through lessons, practice, and assignments at a pace that works for you." },
    { title: "Track your progress", description: "Use your student dashboard to review activity, recognize milestones, and plan a next step." },
  ],
  related: [{ label: "Teaching", href: "/teach" }, { label: "Community", href: "/community" }, { label: "Marketplace", href: "/marketplace" }],
};

export function EducationsPage() {
  return <PageLayout content={content} />;
}