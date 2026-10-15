import { BookOpen, Code2, MessageCircle, Sparkles, Users } from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

const content: InformationPageContent = {
  audience: "EDU247 COMMUNITY",
  title: "Learning grows",
  highlight: "when we connect.",
  description: "Find a welcoming place to connect around learning, teaching, and educational creativity. The EDU247 community is being designed to help learners, educators, and creators exchange ideas, ask questions, share useful resources, and collaborate on projects that support learning.",
  dashboardHref: "/auth/register",
  dashboardLabel: "Join EDU247",
  roleHref: "/student",
  roleLabel: "Explore learner participation",
  illustration: {
    center: Users,
    nodes: [MessageCircle, BookOpen, Sparkles, Code2],
  },
  features: [
    { title: "Share and learn", description: "Connect in topic-based groups around learning goals, teaching ideas, and useful educational resources.", icon: Users },
    { title: "Build together", description: "Use discussions to ask questions, exchange feedback, and collaborate on community learning projects.", icon: Code2 },
    { title: "Stay informed", description: "Follow project updates, community events, and new learning opportunities as they become available.", icon: Sparkles },
  ],
  steps: [
    { title: "Choose how to participate", description: "Join conversations as a learner, educator, creator, vendor, or institutional partner." },
    { title: "Bring a question or idea", description: "Start a discussion, share feedback, or propose a practical improvement for the community." },
    { title: "Keep connected", description: "Return to group conversations, follow useful updates, and contribute as the platform grows." },
  ],
  related: [{ label: "Learning", href: "/learn" }, { label: "Teaching", href: "/teach" }, { label: "Builders", href: "/developers" }],
};

export function CommunityPage() {
  return <PageLayout content={content} />;
}
