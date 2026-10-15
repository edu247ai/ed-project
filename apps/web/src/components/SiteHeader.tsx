"use client";

import Image from "next/image";
import Link from "next/link";
import { useState } from "react";
import { ChevronDown, Menu, Search, X } from "lucide-react";
import { SocialLinks } from "./SocialLinks";

const basePath = process.env.NEXT_PUBLIC_BASE_PATH ?? "";

const audiences = [
  { label: "For learners", href: "/learn" },
  { label: "For educators", href: "/teach" },
  { label: "For institutions", href: "/institutions" },
  { label: "For builders", href: "/developers" },
];

const navigation = [
  { label: "Courses", href: "/learn" },
  { label: "For educators", href: "/teach" },
  { label: "AI learning", href: "/learn#ai-tools" },
  { label: "Marketplace", href: "/marketplace" },
];

type SiteHeaderProps = {
  activeAudience?: string;
  className?: string;
  onSearchChange?: (value: string) => void;
  searchValue?: string;
};

export function SiteHeader({ activeAudience, className = "site-header", onSearchChange, searchValue }: SiteHeaderProps) {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <header className={className}>
      <div className="audience-bar">
        <div className="audience-inner">
          {audiences.map(({ label, href }) => (
            <Link className={label === activeAudience ? "audience-active" : undefined} href={href} key={href}>{label}</Link>
          ))}
          <SocialLinks variant="header" />
        </div>
      </div>
      <div className="header-inner">
        <Link className="brand" href="/" aria-label="EDU247 home">
          <Image src={`${basePath}/logoimg/edu247-logo.png`} alt="" width={38} height={38} priority />
          <span><strong>EDU<span>247</span></strong></span>
        </Link>
        <button
          aria-label={menuOpen ? "Close navigation" : "Open navigation"}
          aria-expanded={menuOpen}
          className="menu-toggle"
          onClick={() => setMenuOpen((open) => !open)}
          type="button"
        >
          {menuOpen ? <X size={22} /> : <Menu size={22} />}
        </button>
        <Link className="explore-link" href="/learn#categories">Explore <ChevronDown size={15} /></Link>
        <nav className={`main-nav${menuOpen ? " nav-open" : ""}`} aria-label="Main navigation">
          {navigation.map(({ label, href }) => (
            <Link href={href} key={label} onClick={() => setMenuOpen(false)}>{label}</Link>
          ))}
        </nav>
        <form className="header-search" action={`${basePath}/learn`} method="get" role="search">
          <Search aria-hidden="true" size={17} />
          <input
            aria-label="Search courses"
            name="query"
            onChange={(event) => onSearchChange?.(event.target.value)}
            placeholder="What do you want to learn?"
            type="search"
            value={searchValue}
          />
          <button className="search-submit" type="submit" aria-label="Search learning content"><Search size={17} /></button>
        </form>
        <div className="header-actions">
          <Link className="sign-in-link" href="/auth/login">Sign in</Link>
          <Link className="button button-small button-outline" href="/auth/register">Join EDU247</Link>
        </div>
      </div>
    </header>
  );
}
