import Link from "next/link";
import { ChevronDown } from "lucide-react";
import FooterAbout from "./FooterAbout";
import FooterCommunity from "./FooterCommunity";
import FooterLogo from "./FooterLogo";
import FooterMore from "./FooterMore";

type FooterProps = {
  className?: string;
};

export default function Footer({ className = "site-footer" }: FooterProps) {
  return (
    <footer className={className}>
      <div className="footer-main">
        <FooterLogo />
        <nav className="footer-links footer-primary-links" aria-label="Footer information">
          <FooterAbout />
          <FooterCommunity />
          <FooterMore />
        </nav>
        <Link href="/#top" className="back-to-top" aria-label="Back to top">
          <ChevronDown size={18} />
        </Link>
      </div>
      <div className="footer-bottom">
        <span>© {new Date().getFullYear()} EDU247. All rights reserved.</span>
        <span>Learn. Create. Thrive.</span>
        <Link href="/#top">Back to top ↑</Link>
      </div>
    </footer>
  );
}
