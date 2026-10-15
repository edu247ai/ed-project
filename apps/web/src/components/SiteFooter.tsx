import Link from "next/link";
import Image from "next/image";
import { ArrowUpRight, ChevronDown } from "lucide-react";
import { SocialLinks } from "./SocialLinks";

const basePath = process.env.NEXT_PUBLIC_BASE_PATH ?? "";

export function SiteFooter({ className = "site-footer" }: { className?: string }) {
  return (
    <footer className={className}>
      <div className="footer-main">
        <div className="footer-brand-col">
          <Link className="brand brand-footer" href="/" aria-label="EDU247 home">
            <Image src={`${basePath}/logoimg/edu247-logo.png`} alt="" width={42} height={42} />
            <span><strong>EDU<span>247</span></strong><small>Learn without limits.</small></span>
          </Link>
          <p>Learning without limits.<br />Anytime. Anywhere.</p>
          <a href="mailto:hello@edu247.com" className="footer-email">hello@edu247.com <ArrowUpRight size={14} /></a>
          <SocialLinks variant="footer" />
        </div>
        <nav className="footer-links" aria-label="Footer">
          <div><strong>Explore</strong><Link href="/learn">Learn</Link><Link href="/teach">Teach</Link><Link href="/learn#ai-tools">AI tools</Link><Link href="/marketplace">Marketplace</Link></div>
          <div><strong>Build with us</strong><Link href="/developers">Developers</Link><Link href="/community">Community</Link><Link href="/#about">About EDU247</Link><Link href="/#get-started">Get in touch</Link></div>
          <div><strong>Our promise</strong><Link href="/institutions#trust">Privacy by design</Link><Link href="/institutions#trust">Responsible AI</Link><Link href="/institutions#trust">Accessibility</Link><Link href="/institutions#trust">Trust &amp; safety</Link></div>
        </nav>
        <Link href="/#top" className="back-to-top" aria-label="Back to top"><ChevronDown size={18} /></Link>
      </div>
      <div className="footer-bottom">
        <span>© 2026 EDU247. Learning without limits.</span>
        <span>Made for a more curious world <span className="footer-heart">✳</span></span>
        <Link href="/#top">Back to top ↑</Link>
      </div>
    </footer>
  );
}
