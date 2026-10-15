import Link from "next/link";
import {
  ArrowRight,
  ArrowUpRight,
  type LucideIcon,
} from "lucide-react";
import { SiteFooter } from "../SiteFooter";
import { SiteHeader } from "../SiteHeader";
import { DesignCard } from "../DesignCard";

export type InformationPageContent = {
  audience: string;
  title: string;
  highlight: string;
  description: string;
  dashboardHref: string;
  dashboardLabel: string;
  roleHref: string;
  roleLabel: string;
  features: { title: string; description: string; icon: LucideIcon }[];
  illustration: {
    center: LucideIcon;
    nodes: [LucideIcon, LucideIcon, LucideIcon, LucideIcon];
  };
  steps: { title: string; description: string }[];
  related: { label: string; href: string }[];
  activeAudience?: string;
  trustSection?: boolean;
};

export function PageLayout({ content }: { content: InformationPageContent }) {
  const [TopNode, RightNode, BottomNode, LeftNode] = content.illustration.nodes;
  const CenterIcon = content.illustration.center;

  return (
    <>
      <SiteHeader activeAudience={content.activeAudience} className="site-header platform-site-header" />
      <main className="page-shell platform-page">
        <section className="platform-hero">
          <div className="platform-hero-copy">
            <span className="section-kicker">{content.audience}</span>
            <h1>{content.title}<br /><em>{content.highlight}</em></h1>
            <p>{content.description}</p>
            <div className="platform-actions">
              <Link className="button button-dark" href={content.dashboardHref}>{content.dashboardLabel}<ArrowRight size={16} /></Link>
              <Link className="text-link" href={content.roleHref}>{content.roleLabel}<ArrowUpRight size={15} /></Link>
            </div>
          </div>
          <div
            className="platform-hero-art"
            aria-hidden="true"
            data-page-illustration
          >
            <span className="platform-orbit platform-orbit-one" />
            <span className="platform-orbit platform-orbit-two" />
            <span className="platform-art-center"><CenterIcon size={36} strokeWidth={1.8} /></span>
            <span className="platform-art-node platform-art-node-one"><TopNode size={20} strokeWidth={1.8} /></span>
            <span className="platform-art-node platform-art-node-two"><RightNode size={20} strokeWidth={1.8} /></span>
            <span className="platform-art-node platform-art-node-three"><BottomNode size={20} strokeWidth={1.8} /></span>
            <span className="platform-art-node platform-art-node-four"><LeftNode size={20} strokeWidth={1.8} /></span>
          </div>
        </section>

        <section className="platform-audience">
          <DesignCard as="div"><span>WHO THIS IS FOR</span><p>{content.audience}</p></DesignCard>
          <DesignCard as="div"><span>WHAT YOU CAN DO</span><p>{content.description}</p></DesignCard>
        </section>

        <section className="platform-content-section" id={content.trustSection ? "trust" : undefined} aria-labelledby="platform-features-title">
          <div className="platform-section-heading">
            <span className="section-kicker">A CLEAR NEXT STEP</span>
            <h2 id="platform-features-title">What you can explore</h2>
          </div>
          <div className="platform-feature-grid">
            {content.features.map(({ title, description, icon: Icon }, index) => (
              <DesignCard className="platform-feature-card" key={title}>
                <span className="platform-feature-icon"><Icon size={21} /></span>
                <span className="platform-feature-number">0{index + 1}</span>
                <h3>{title}</h3>
                <p>{description}</p>
              </DesignCard>
            ))}
          </div>
        </section>

        <section className="platform-content-section platform-steps-section" aria-labelledby="platform-steps-title">
          <div className="platform-section-heading">
            <span className="section-kicker">HOW IT WORKS</span>
            <h2 id="platform-steps-title">A simple path to get started</h2>
          </div>
          <ol className="platform-steps">
            {content.steps.map((step, index) => (
              <li key={step.title}>
                <span className="platform-step-number">0{index + 1}</span>
                <div><h3>{step.title}</h3><p>{step.description}</p></div>
              </li>
            ))}
          </ol>
        </section>

        <section className="platform-cta">
          <div>
            <span>YOUR NEXT STEP</span>
            <h2>Ready to move forward?</h2>
            <p>Explore your role workspace or learn more about the connected EDU247 platform.</p>
          </div>
          <div className="platform-actions">
            <Link className="button button-light" href={content.dashboardHref}>{content.dashboardLabel}<ArrowRight size={16} /></Link>
            <Link className="platform-cta-link" href="/marketplace"><ArrowUpRight size={15} /> Visit marketplace</Link>
          </div>
        </section>

        <nav className="platform-related" aria-label="Related EDU247 pages">
          <strong>Explore more</strong>
          {content.related.map(({ label, href }) => <Link href={href} key={`${label}-${href}`}>{label}<ArrowUpRight size={13} /></Link>)}
        </nav>
        <p className="platform-preview-note">EDU247 is in development. Dashboard data and backend features depend on the platform APIs becoming available.</p>
        <SiteFooter className="site-footer platform-site-footer" />
      </main>
    </>
  );
}
