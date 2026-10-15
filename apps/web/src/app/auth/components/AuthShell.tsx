import Link from "next/link";
import { BookOpenCheck, BrainCircuit, GraduationCap, Sparkles } from "lucide-react";
import { SiteFooter } from "../../../components/SiteFooter";
import { SiteHeader } from "../../../components/SiteHeader";
import { DesignCard } from "../../../components/DesignCard";

type AuthShellProps = {
  eyebrow: string;
  title: string;
  description: string;
  children: React.ReactNode;
};

export function AuthShell({ eyebrow, title, description, children }: AuthShellProps) {
  return (
    <div className="auth-page">
      <SiteHeader />

      <DesignCard as="section" className="auth-layout">
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
      </DesignCard>

      <SiteFooter />
    </div>
  );
}
