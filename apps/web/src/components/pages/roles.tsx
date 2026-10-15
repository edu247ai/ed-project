import {
  Award,
  BookOpenCheck,
  BrainCircuit,
  Check,
  Code2,
  CreditCard,
  GraduationCap,
  Network,
  Package,
  ShieldCheck,
  Sparkles,
  Store,
  Truck,
  Users,
} from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

type RoleKey = "student" | "teacher" | "vendor" | "developer";

const roleContent: Record<RoleKey, InformationPageContent> = {
  student: {
    audience: "STUDENT ROLE",
    activeAudience: "For learners",
    title: "Your learning,",
    highlight: "all in one place.",
    description: "Discover courses, keep moving through guided learning paths, and follow your progress in a planned student workspace. Explore subjects at your pace, return to active study, and keep your next learning goal in view.",
    dashboardHref: "/Dashboard/student-dashboard",
    dashboardLabel: "Start learning",
    roleHref: "/learn",
    roleLabel: "Explore learning",
    illustration: {
      center: BookOpenCheck,
      nodes: [BrainCircuit, GraduationCap, Award, Sparkles],
    },
    features: [
      { title: "My courses", description: "Return to active learning, review course materials, and follow progress through each learning journey.", icon: BookOpenCheck },
      { title: "Assignments and activities", description: "Keep upcoming work, practice activities, and learning tasks organized in one place.", icon: Check },
      { title: "Tutor and achievements", description: "Access planned AI learning support and recognize milestones as you build knowledge and skills.", icon: BrainCircuit },
    ],
    steps: [
      { title: "Create your learner profile", description: "Choose interests and goals as account features become available." },
      { title: "Explore learning", description: "Find a course or path that fits your next step." },
      { title: "Return to your dashboard", description: "Review progress and continue where you left off." },
    ],
    related: [{ label: "Learning page", href: "/learn" }, { label: "Community", href: "/community" }, { label: "Marketplace", href: "/marketplace" }],
  },
  teacher: {
    audience: "TEACHER ROLE",
    activeAudience: "For educators",
    title: "Your classroom,",
    highlight: "with room to create.",
    description: "Create and review lessons, manage classes, and support learners through a planned educator workspace. Keep your expertise at the center while using connected tools to organize teaching and prepare resources for sharing.",
    dashboardHref: "/Dashboard/teacher-dashboard",
    dashboardLabel: "Teach now",
    roleHref: "/teach",
    roleLabel: "Explore teaching",
    illustration: {
      center: GraduationCap,
      nodes: [BookOpenCheck, Users, Sparkles, BrainCircuit],
    },
    features: [
      { title: "Create and share", description: "Build courses and lessons, then prepare eligible teaching resources for the planned marketplace.", icon: BookOpenCheck },
      { title: "Manage classes", description: "Organize learners, submissions, assignments, and progress to help you plan classroom support.", icon: Users },
      { title: "Create with AI support", description: "Use planned AI lesson-building tools to develop drafts while retaining educator review and control.", icon: BrainCircuit },
    ],
    steps: [
      { title: "Set up your educator profile", description: "Share your teaching areas and the learners you support." },
      { title: "Create a learning experience", description: "Draft and review courses, lessons, and assignments." },
      { title: "Manage and improve", description: "Use your dashboard to support classes and understand learner activity." },
    ],
    related: [{ label: "Teaching page", href: "/teach" }, { label: "Marketplace", href: "/marketplace" }, { label: "Learning", href: "/learn" }],
  },
  vendor: {
    audience: "VENDOR ROLE",
    title: "Bring your resources",
    highlight: "to more learners.",
    description: "Introduce educational products to a wider learning community and prepare for marketplace operations through the planned EDU247 vendor workspace. The future workflow is designed to bring product setup, pricing, orders, and delivery information together.",
    dashboardHref: "/Dashboard/vendor-dashboard",
    dashboardLabel: "Sell on the marketplace",
    roleHref: "/marketplace",
    roleLabel: "Explore the marketplace",
    illustration: {
      center: Store,
      nodes: [Package, Truck, CreditCard, Users],
    },
    features: [
      { title: "Onboard your store", description: "Create a vendor profile and provide clear information about your products for marketplace review.", icon: Package },
      { title: "Upload and price products", description: "Prepare digital or physical listings with useful descriptions, accurate pricing, and delivery details.", icon: Package },
      { title: "Fulfill orders", description: "Use planned tools to coordinate inventory, payments, delivery, returns, and customer feedback.", icon: Users },
    ],
    steps: [
      { title: "Set up your vendor profile", description: "Introduce your store and the learning products you provide." },
      { title: "Prepare product listings", description: "Add product information, pricing, files, stock, and delivery details." },
      { title: "Manage the marketplace", description: "Use the vendor dashboard for orders, fulfillment, and sales." },
    ],
    related: [{ label: "Marketplace page", href: "/marketplace" }, { label: "Teaching", href: "/teach" }, { label: "Community", href: "/community" }],
  },
  developer: {
    audience: "DEVELOPER ROLE",
    activeAudience: "For builders",
    title: "Build tools for",
    highlight: "open education.",
    description: "Explore APIs, SDK resources, integration guidance, and contribution opportunities in the planned developer workspace. The goal is to help builders understand platform capabilities and create useful tools while respecting security, privacy, and learner needs.",
    dashboardHref: "/Dashboard/developer-dashboard",
    dashboardLabel: "Build tools",
    roleHref: "/developers",
    roleLabel: "Explore developer resources",
    illustration: {
      center: Code2,
      nodes: [Network, ShieldCheck, Users, Sparkles],
    },
    features: [
      { title: "APIs and SDKs", description: "Build applications with documented, versioned platform interfaces as they become available.", icon: Code2 },
      { title: "Applications and keys", description: "Manage registered apps, API usage, and access credentials in a developer workspace.", icon: ShieldCheck },
      { title: "Contribute and collaborate", description: "Join issue discussions, contributions, and community hackathons.", icon: Users },
    ],
    steps: [
      { title: "Explore docs and contracts", description: "Understand available endpoints, integration patterns, and platform boundaries." },
      { title: "Build in a safe environment", description: "Use sandbox tools and scoped credentials when provided by the platform." },
      { title: "Share and maintain", description: "Monitor application health and contribute improvements." },
    ],
    related: [{ label: "Developers page", href: "/developers" }, { label: "Community", href: "/community" }, { label: "Institutions", href: "/institutions" }],
  },
};

export function RolePage({ role }: { role: RoleKey }) {
  return <PageLayout content={roleContent[role]} />;
}
