import Image from "next/image";
import Link from "next/link";
import { ArrowUpRight } from "lucide-react";

const basePath = process.env.NEXT_PUBLIC_BASE_PATH ?? "";

export default function FooterLogo() {
  return (
    <div className="footer-brand-col">
      <span className="footer-brand-ripple" aria-hidden="true" />
      <Link className="brand brand-footer" href="/" aria-label="EDU247 home">
        <Image src={`${basePath}/logoimg/edu247-logo.png`} alt="" width={68} height={68} />
        <span><strong>EDU<span>247</span></strong><small>Learn. Create. Thrive.</small></span>
      </Link>
      <p>
        EDU247 is a global learning ecosystem built to empower students, developers, educators, and creators.
        We combine intelligent tools, modern design, and community-driven innovation to make education accessible,
        inspiring, and limitless.
      </p>
      <a href="mailto:hello@edu247.com" className="footer-email">
        hello@edu247.com <ArrowUpRight size={15} />
      </a>
    </div>
  );
}
