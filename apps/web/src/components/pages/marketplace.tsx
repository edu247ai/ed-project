import { CreditCard, Package, Store, Truck } from "lucide-react";
import { PageLayout, type InformationPageContent } from "./PageLayout";

const content: InformationPageContent = {
  audience: "EDU247 MARKETPLACE",
  title: "Useful resources,",
  highlight: "made for learning.",
  description: "Discover a planned marketplace for digital learning products and educator-created resources. Learners and institutions will be able to find materials for their needs, while vendors can prepare listings, set prices, and manage fulfillment through a dedicated workspace.",
  dashboardHref: "/Dashboard/vendor-dashboard",
  dashboardLabel: "Sell on the marketplace",
  roleHref: "/vendor",
  roleLabel: "Explore the vendor experience",
  illustration: {
    center: Store,
    nodes: [Package, Truck, CreditCard, Store],
  },
  features: [
    { title: "Digital products", description: "Discover the planned range of course materials, templates, workbooks, and other downloadable learning resources.", icon: Package },
    { title: "Vendor onboarding", description: "A planned review and setup process for vendor profiles, product information, and marketplace listings.", icon: Store },
    { title: "Pricing and delivery", description: "Prepare pricing, orders, fulfillment, and delivery information for management from a vendor workspace.", icon: Truck },
  ],
  steps: [
    { title: "Set up a vendor profile", description: "Introduce your organization and explain the learning resources or products you provide." },
    { title: "Create product listings", description: "Prepare clear descriptions, accessible product details, pricing, and delivery information." },
    { title: "Manage orders", description: "Use the vendor dashboard to review orders, monitor inventory, and coordinate fulfillment when marketplace services launch." },
  ],
  related: [{ label: "Teaching", href: "/teach" }, { label: "Vendor role page", href: "/vendor" }, { label: "Learning", href: "/learn" }],
};

export function MarketplacePage() {
  return <PageLayout content={content} />;
}
