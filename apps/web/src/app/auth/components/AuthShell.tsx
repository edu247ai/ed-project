"use client";

import Image from "next/image";
import Link from "next/link";
import { useState } from "react";
import { BookOpenCheck, BrainCircuit, ChevronDown, GraduationCap, Menu, Sparkles, X } from "lucide-react";
import { SocialLinks } from "../../../components/SocialLinks";

type AuthShellProps = {
  eyebrow: string;
  title: string;
  description: string;
  children: React.ReactNode;
};

export function AuthShell({ eyebrow, title, description, children }: AuthShellProps) {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <div className="auth-page">
      <header className="auth-site-header">
        <div className="auth-audience-bar">
          <div className="auth-audience-inner">
            <nav className="auth-audience-links" aria-label="EDU247 audiences">
              <Link href="/#learning">For learners</Link>
              <Link href="/#teachers">For educators</Link>
              <Link href="/#platform">For institutions</Link>
              <Link href="/#developers">For builders</Link>
            </nav>
            <SocialLinks variant="header" />
          </div>
        </div>
        <div className="auth-header-inner">
          <Link className="auth-brand" href="/" aria-label="EDU247 home">
            <Image src="/logoimg/edu247-logo.png" alt="" width={40} height={40} priority />
            <span>EDU<span>247</span></span>
          </Link>
          <button
            aria-label={menuOpen ? "Close navigation" : "Open navigation"}
            aria-expanded={menuOpen}
            className="auth-menu-toggle"
            onClick={() => setMenuOpen((open) => !open)}
            type="button"
          >
            {menuOpen ? <X size={20} /> : <Menu size={20} />}
          </button>
          <Link className="auth-explore-link" href="/#categories">
            Explore <ChevronDown size={15} />
          </Link>
          <nav className={`auth-main-nav${menuOpen ? " auth-nav-open" : ""}`} aria-label="Main navigation">
            <Link href="/#learning" onClick={() => setMenuOpen(false)}>Courses</Link>
            <Link href="/#teachers" onClick={() => setMenuOpen(false)}>For educators</Link>
            <Link href="/#ai-tools" onClick={() => setMenuOpen(false)}>AI learning</Link>
            <Link href="/#marketplace" onClick={() => setMenuOpen(false)}>Marketplace</Link>
          </nav>
          <div className="auth-header-actions">
            <Link className="auth-header-sign-in" href="/auth/login">Sign in</Link>
            <Link className="auth-header-join" href="/auth/register">Join EDU247</Link>
          </div>
        </div>
      </header>

      <main className="auth-layout">
        <aside className="auth-aside">
          <div className="auth-aside-art" aria-hidden="true">
            <div className="auth-art-ring auth-art-ring-one" />
            <div className="auth-art-ring auth-art-ring-two" />
            <div className="auth-art-center"><BookOpenCheck size={40} /></div>
            <span className="auth-art-orb auth-orb-one"><BrainCircuit size={21} /></span>
            <span className="auth-art-orb auth-orb-two"><GraduationCap size={22} /></span>
            <span className="auth-art-orb auth-orb-three"><Sparkles size={19} /></span>
          </div>
          <div className="auth-aside-copy">
            <span className="auth-aside-kicker">ONE SPACE FOR YOUR NEXT STEP</span>
            <h1>Learning opens<br />new doors.</h1>
            <p>Find a course, build a new skill, make something useful, and grow alongside a community that cares about education.</p>
            <div className="auth-aside-points">
              <span><BookOpenCheck size={16} /> Learning for students and educators</span>
              <span><BrainCircuit size={16} /> Helpful AI, with people in control</span>
              <span><Sparkles size={16} /> Room to create, contribute and grow</span>
            </div>
          </div>
          <div className="auth-aside-footer">Learning without limits. <span>Anytime. Anywhere.</span></div>
        </aside>

        <section className="auth-content">
          <div className="auth-form-wrap">
            <span className="auth-eyebrow">{eyebrow}</span>
            <h2>{title}</h2>
            <p className="auth-description">{description}</p>
            {children}
          </div>
          <p className="auth-legal">
            By continuing, you agree to EDU247&apos;s <Link href="/#about">Terms</Link> and <Link href="/#about">Privacy Policy</Link>.
          </p>
        </section>
      </main>

      <footer className="auth-site-footer">
        <div className="auth-footer-main">
          <div className="auth-footer-brand">
            <Link className="auth-brand" href="/" aria-label="EDU247 home">
              <Image src="/logoimg/edu247-logo.png" alt="" width={40} height={40} />
              <span>EDU<span>247</span><small>Learn without limits.</small></span>
            </Link>
            <p>Learning without limits.<br />Anytime. Anywhere.</p>
            <a href="mailto:hello@edu247.com">hello@edu247.com</a>
            <SocialLinks variant="footer" />
          </div>
          <nav className="auth-footer-links" aria-label="Footer">
            <div><strong>Explore</strong><Link href="/#learning">Learn</Link><Link href="/#teachers">Teach</Link><Link href="/#ai-tools">AI tools</Link><Link href="/#marketplace">Marketplace</Link></div>
            <div><strong>Build with us</strong><Link href="/#developers">Developers</Link><Link href="/#community">Community</Link><Link href="/#about">About EDU247</Link><Link href="/#get-started">Get in touch</Link></div>
            <div><strong>Our promise</strong><Link href="/#about">Privacy by design</Link><Link href="/#about">Responsible AI</Link><Link href="/#about">Accessibility</Link><Link href="/#about">Trust &amp; safety</Link></div>
          </nav>
        </div>
        <div className="auth-footer-bottom">
          <span>© 2026 EDU247. Learning without limits.</span>
          <span>Made for a more curious world <span className="auth-footer-mark">✳</span></span>
          <Link href="/">Back to EDU247</Link>
        </div>
      </footer>
    </div>
  );
}
