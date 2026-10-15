"use client";

import Image from "next/image";
import { useMemo, useState, type FormEvent } from "react";
import {
  ArrowRight,
  ArrowUpRight,
  BookOpen,
  BookOpenCheck,
  BrainCircuit,
  CalendarDays,
  Check,
  ChevronDown,
  ChevronRight,
  Cloud,
  Code2,
  GraduationCap,
  Lightbulb,
  Menu,
  Network,
  Search,
  Settings,
  ShieldCheck,
  ShoppingBag,
  Sparkles,
  Store,
  Users,
  X,
  type LucideIcon,
} from "lucide-react";
import { SocialLinks } from "../components/SocialLinks";

const navigation = [
  { label: "Courses", href: "#learning" },
  { label: "For educators", href: "#teachers" },
  { label: "AI learning", href: "#ai-tools" },
  { label: "Marketplace", href: "#marketplace" },
];

const categories = [
  { name: "Artificial intelligence", icon: BrainCircuit, href: "#learning" },
  { name: "Technology", icon: Code2, href: "#learning" },
  { name: "Teaching & learning", icon: BookOpenCheck, href: "#teachers" },
  { name: "Creative skills", icon: Lightbulb, href: "#ai-tools" },
  { name: "Career development", icon: GraduationCap, href: "#roadmap" },
];

const roles: {
  name: string;
  description: string;
  icon: LucideIcon;
  color: string;
  href: string;
}[] = [
  {
    name: "Students",
    description: "Learn new skills and make progress at your own pace.",
    icon: GraduationCap,
    color: "mint",
    href: "#learning",
  },
  {
    name: "Teachers",
    description: "Create lessons, inspire learners and track progress.",
    icon: BookOpenCheck,
    color: "peach",
    href: "#teachers",
  },
  {
    name: "Developers",
    description: "Build useful tools and help shape open education.",
    icon: Code2,
    color: "blue",
    href: "#developers",
  },
  {
    name: "Vendors",
    description: "Bring trusted learning resources to more people.",
    icon: Store,
    color: "lavender",
    href: "#marketplace",
  },
  {
    name: "Security reviewers",
    description: "Help keep contributions and releases safe and trustworthy.",
    icon: ShieldCheck,
    color: "blue",
    href: "#trust",
  },
  {
    name: "Platform admins",
    description: "Oversee users, content, operations and platform health.",
    icon: Settings,
    color: "peach",
    href: "#platform",
  },
];

const courses = [
  {
    title: "Introduction to Artificial Intelligence",
    category: "Technology",
    level: "Beginner · 4 weeks",
    provider: "EDU247 Learning",
    credential: "Guided learning path",
    color: "course-lilac",
    icon: BrainCircuit,
  },
  {
    title: "Web Development Fundamentals",
    category: "Computer science",
    level: "Beginner · 6 weeks",
    provider: "EDU247 Learning",
    credential: "Course",
    color: "course-blue",
    icon: Code2,
  },
  {
    title: "Creative Problem Solving",
    category: "Personal development",
    level: "All levels · 3 weeks",
    provider: "EDU247 Learning",
    credential: "Course",
    color: "course-peach",
    icon: Lightbulb,
  },
];

const aiTools: {
  name: string;
  description: string;
  icon: LucideIcon;
  color: string;
}[] = [
  {
    name: "AI tutor",
    description: "Get explanations, practice and support that meet you where you are.",
    icon: BrainCircuit,
    color: "tool-mint",
  },
  {
    name: "Lesson builder",
    description: "Turn your ideas into structured, engaging lessons.",
    icon: BookOpenCheck,
    color: "tool-peach",
  },
  {
    name: "Comic creator",
    description: "Bring learning topics to life as memorable visual stories.",
    icon: Sparkles,
    color: "tool-lilac",
  },
  {
    name: "Infographic creator",
    description: "Make clear visual summaries from lessons and notes.",
    icon: Network,
    color: "tool-blue",
  },
  {
    name: "Mind map creator",
    description: "Organize complex ideas into connected, useful maps.",
    icon: Lightbulb,
    color: "tool-yellow",
  },
  {
    name: "Smart recommendations",
    description: "Find learning content and opportunities chosen for you.",
    icon: ArrowUpRight,
    color: "tool-rose",
  },
  {
    name: "AI code reviewer",
    description: "Help developers inspect changes and improve code quality.",
    icon: Code2,
    color: "tool-blue",
  },
  {
    name: "AI security auditor",
    description: "Support security reviews with focused analysis and findings.",
    icon: ShieldCheck,
    color: "tool-mint",
  },
  {
    name: "Documentation engine",
    description: "Turn technical work into clearer, easier-to-follow guides.",
    icon: BookOpen,
    color: "tool-peach",
  },
];

const deliveryPhases = [
  {
    number: "01",
    name: "Foundation",
    timing: "Weeks 1–6",
    description: "Validate requirements, architecture and design; establish delivery and developer workflows.",
    icon: CalendarDays,
    color: "phase-mint",
  },
  {
    number: "02",
    name: "Core platform",
    timing: "Weeks 7–16",
    description: "Build identity, core services, learning experiences and the first administration tools.",
    icon: Network,
    color: "phase-blue",
  },
  {
    number: "03",
    name: "AI capabilities",
    timing: "Weeks 17–28",
    description: "Introduce the tutor, lesson builder and creative tools with evaluation and human oversight.",
    icon: BrainCircuit,
    color: "phase-lilac",
  },
  {
    number: "04",
    name: "Marketplace",
    timing: "Weeks 29–36",
    description: "Develop vendor onboarding, product discovery, checkout and delivery integrations.",
    icon: ShoppingBag,
    color: "phase-peach",
  },
  {
    number: "05",
    name: "Community",
    timing: "Weeks 37–44",
    description: "Grow developer onboarding, contribution practices, hackathons and security review.",
    icon: Users,
    color: "phase-yellow",
  },
  {
    number: "06",
    name: "Pilot & readiness",
    timing: "Weeks 45–52",
    description: "Test with representative users, resolve priority issues and assess launch readiness.",
    icon: Check,
    color: "phase-green",
  },
];

const products = [
  {
    name: "STEM Activity Kit",
    kind: "Digital resource",
    price: "$19.99",
    icon: Network,
    color: "product-mint",
  },
  {
    name: "Interactive Science Workbook",
    kind: "eBook",
    price: "$9.99",
    icon: BookOpen,
    color: "product-peach",
  },
  {
    name: "Educator Template Pack",
    kind: "Resource bundle",
    price: "$14.99",
    icon: Sparkles,
    color: "product-lilac",
  },
];

function Brand({ footer = false }: { footer?: boolean }) {
  return (
    <a className={`brand${footer ? " brand-footer" : ""}`} href="#top" aria-label="EDU247 home">
      <Image
        src="/logoimg/edu247-logo.png"
        alt=""
        width={footer ? 42 : 38}
        height={footer ? 42 : 38}
        priority={!footer}
      />
      <span>
        <strong>EDU<span>247</span></strong>
        {footer && <small>Learn without limits.</small>}
      </span>
    </a>
  );
}

export default function HomePage() {
  const [query, setQuery] = useState("");
  const [menuOpen, setMenuOpen] = useState(false);
  const [newsletterMessage, setNewsletterMessage] = useState("");

  const filteredCourses = useMemo(() => {
    const normalizedQuery = query.trim().toLowerCase();
    if (!normalizedQuery) return courses;

    return courses.filter((course) =>
      `${course.title} ${course.category} ${course.level}`
        .toLowerCase()
        .includes(normalizedQuery),
    );
  }, [query]);

  function handleNewsletterSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setNewsletterMessage("Newsletter signup is not connected yet. Your email was not sent or saved.");
  }

  return (
    <main>
      <header className="site-header">
        <div className="audience-bar">
          <div className="audience-inner">
            <a className="audience-active" href="#learning">For learners</a>
            <a href="#teachers">For educators</a>
            <a href="#platform">For institutions</a>
            <a href="#developers">For builders</a>
            <SocialLinks variant="header" />
          </div>
        </div>
        <div className="header-inner">
          <Brand />
          <button
            aria-label={menuOpen ? "Close navigation" : "Open navigation"}
            aria-expanded={menuOpen}
            className="menu-toggle"
            onClick={() => setMenuOpen((open) => !open)}
            type="button"
          >
            {menuOpen ? <X size={22} /> : <Menu size={22} />}
          </button>
          <a className="explore-link" href="#categories">Explore <ChevronDown size={15} /></a>
          <nav className={`main-nav${menuOpen ? " nav-open" : ""}`} aria-label="Main navigation">
            {navigation.map((item) => (
              <a key={item.label} href={item.href} onClick={() => setMenuOpen(false)}>
                {item.label}
              </a>
            ))}
          </nav>
          <label className="header-search">
            <Search aria-hidden="true" size={17} />
            <input
              aria-label="Search courses"
              onChange={(event) => setQuery(event.target.value)}
              placeholder="What do you want to learn?"
              type="search"
              value={query}
            />
            <span className="search-submit"><Search size={17} /></span>
          </label>
          <div className="header-actions">
            <a className="sign-in-link" href="/auth/login">Sign in</a>
            <a className="button button-small button-outline" href="/auth/register">
              Join EDU247
            </a>
          </div>
        </div>
      </header>

      <div className="page-shell">
        <section className="hero" id="top" aria-labelledby="hero-heading">
          <div className="hero-feature">
            <div className="hero-content">
              <div className="eyebrow"><span /> LEARNING THAT MOVES YOU FORWARD</div>
              <h1 id="hero-heading">Learn without limits.</h1>
              <p>
                Build real-world skills, explore what interests you, and make
                your next move with learning designed for everyone.
              </p>
              <div className="hero-actions">
                <a className="button button-hero" href="#learning">
                  Explore courses <ArrowRight size={17} />
                </a>
                <a className="hero-text-link" href="#about">
                  Discover EDU247 <ArrowUpRight size={16} />
                </a>
              </div>
              <div className="hero-proof">
                <span><GraduationCap size={17} /></span>
                Courses, creative tools and community—all in one place.
              </div>
            </div>
            <div className="hero-art" aria-hidden="true">
              <div className="hero-art-orbit hero-art-orbit-one" />
              <div className="hero-art-orbit hero-art-orbit-two" />
              <div className="hero-art-book"><BookOpen size={75} strokeWidth={1.15} /><span>YOUR NEXT CHAPTER</span></div>
              <div className="hero-art-badge hero-art-badge-top"><Sparkles size={16} /> A new skill</div>
              <div className="hero-art-badge hero-art-badge-bottom"><Check size={15} /> Keep growing</div>
              <div className="hero-art-stamp">247</div>
            </div>
          </div>
          <aside className="hero-side-card">
            <span className="hero-side-kicker">A WORLD OF LEARNING</span>
            <h2>Make your next move a learning one.</h2>
            <p>Whether you&apos;re starting out or moving forward, there&apos;s a path for you.</p>
            <div className="hero-side-list">
              <a href="#learning"><span className="side-list-icon side-icon-blue"><BrainCircuit size={17} /></span><span><strong>Explore a new subject</strong><small>Start with what makes you curious</small></span><ChevronRight size={17} /></a>
              <a href="#teachers"><span className="side-list-icon side-icon-green"><BookOpenCheck size={17} /></span><span><strong>Bring learning to your class</strong><small>Tools for teachers and educators</small></span><ChevronRight size={17} /></a>
              <a href="#developers"><span className="side-list-icon side-icon-gold"><Code2 size={17} /></span><span><strong>Build the future with us</strong><small>Meet the EDU247 community</small></span><ChevronRight size={17} /></a>
            </div>
            <a className="hero-side-link" href="#about">See what you can do <ArrowRight size={15} /></a>
          </aside>
        </section>

        <section className="category-section" id="categories" aria-label="Explore learning categories">
          <div className="category-heading">
            <div><span className="section-kicker">FIND YOUR DIRECTION</span><h2>What would you like to learn?</h2></div>
            <a className="text-link" href="#learning">Browse all learning <ArrowRight size={16} /></a>
          </div>
          <div className="category-list">
            {categories.map(({ name, icon: Icon, href }) => (
              <a className="category-chip" href={href} key={name}>
                <Icon size={18} />
                <span>{name}</span>
                <ChevronRight size={15} />
              </a>
            ))}
          </div>
        </section>

        <section className="section intro-section" id="about">
          <div className="section-heading">
            <div>
              <span className="section-kicker">LEARNING FOR EVERY KIND OF AMBITION</span>
              <h2>Find the learning path that fits you.</h2>
            </div>
            <p>Build knowledge for your next goal, create better learning experiences, or join the people building what comes next.</p>
          </div>
          <div className="role-grid">
            {roles.map(({ name, description, icon: Icon, color, href }) => (
              <a className="role-card" href={href} key={name}>
                <span className={`role-icon ${color}`}><Icon size={23} strokeWidth={1.8} /></span>
                <span className="role-copy"><strong>{name}</strong><small>{description}</small></span>
                <ChevronRight className="role-arrow" size={18} />
              </a>
            ))}
          </div>
        </section>

        <section className="section learning-section" id="learning">
          <div className="section-heading section-heading-compact">
            <div>
              <span className="section-kicker">POPULAR ON EDU247</span>
              <h2>Start, switch, or advance your skills.</h2>
              <p className="section-subtitle">Explore instructor-guided courses and learning paths for what matters to you.</p>
            </div>
            <a className="text-link" href="#get-started">View all courses <ArrowRight size={16} /></a>
          </div>
          {query.trim() && (
            <p className="search-feedback" aria-live="polite">
              Showing {filteredCourses.length} learning {filteredCourses.length === 1 ? "path" : "paths"} for “{query}”
            </p>
          )}
          {filteredCourses.length > 0 ? (
            <div className="course-grid">
              {filteredCourses.map(({ title, category, level, provider, credential, color, icon: Icon }) => (
                <article className="course-card" key={title}>
                  <div className={`course-art ${color}`}>
                    <Icon size={31} strokeWidth={1.45} />
                    <span>{credential}</span>
                    <div className="art-shape art-shape-one" />
                    <div className="art-shape art-shape-two" />
                  </div>
                  <div className="course-info">
                    <div className="course-provider">{provider}</div>
                    <h3>{title}</h3>
                    <div className="course-rating">
                      <span>{category}</span>
                      <span className="course-dot" />
                      <span>{level}</span>
                    </div>
                  </div>
                  <a href="#get-started" className="course-link" aria-label={`Continue ${title}`}><ArrowUpRight size={17} /></a>
                </article>
              ))}
            </div>
          ) : (
            <div className="empty-search">
              <Search size={23} />
              <strong>No learning paths found</strong>
              <span>Try another course title or topic.</span>
              <button onClick={() => setQuery("")} type="button">Clear search</button>
            </div>
          )}
        </section>

        <section className="section ai-section" id="ai-tools">
          <div className="ai-section-intro">
            <span className="section-kicker">A LITTLE MORE “WHAT IF?”</span>
            <h2>Big ideas deserve<br />brilliant tools.</h2>
            <p>Helpful AI that supports the people doing the learning and teaching. Your imagination stays in the driver&apos;s seat.</p>
            <a className="button button-light" href="#get-started">Meet your AI toolkit <ArrowRight size={16} /></a>
            <div className="responsible-note"><ShieldCheck size={17} /><span>Built with care.<br /><strong>People stay in control.</strong></span></div>
          </div>
          <div className="ai-tool-grid">
            {aiTools.map(({ name, description, icon: Icon, color }, index) => (
              <a className="ai-tool-card" href="#get-started" key={name}>
                <span className={`tool-icon ${color}`}><Icon size={21} strokeWidth={1.8} /></span>
                <span className="tool-number">0{index + 1}</span>
                <strong>{name}</strong>
                <small>{description}</small>
                <span className="tool-arrow"><ArrowUpRight size={15} /></span>
              </a>
            ))}
          </div>
        </section>

        <section className="section teacher-section" id="teachers">
          <div className="teacher-visual">
            <div className="teacher-visual-top"><span className="live-dot" /> YOUR TEACHING STUDIO</div>
            <div className="studio-card">
              <div className="studio-card-head"><div><span>COURSE BUILDER</span><strong>Make room for wonder.</strong></div><Sparkles size={20} /></div>
              <div className="lesson-preview">
                <div className="lesson-preview-art"><Network size={31} /><span>LESSON 01</span></div>
                <div><span className="preview-label">A NEW LESSON</span><strong>Explore the natural world</strong><small>Made for your classroom, with a little help from AI.</small></div>
              </div>
              <div className="studio-footer"><span><Check size={14} /> Ready for your learners</span><span>Draft saved</span></div>
            </div>
            <div className="teacher-floating"><span><Lightbulb size={17} /></span><div><strong>Your idea, made teachable.</strong><small>Start with what inspires you.</small></div></div>
          </div>
          <div className="teacher-copy">
            <span className="section-kicker">FOR THE PEOPLE WHO INSPIRE US</span>
            <h2>Make teaching feel a little more <em>limitless.</em></h2>
            <p>Spend less time wrestling with busywork and more time doing what matters: helping every learner find their spark.</p>
            <ul>
              <li><Check size={16} /> Build courses and lessons with a thoughtful AI partner</li>
              <li><Check size={16} /> Bring learning to life with comics, maps and visuals</li>
              <li><Check size={16} /> Keep assignments and learner progress in one place</li>
            </ul>
            <a className="text-link" href="#get-started">Discover tools for teachers <ArrowRight size={16} /></a>
          </div>
        </section>

        <section className="section marketplace-section" id="marketplace">
          <div className="marketplace-heading">
            <div>
              <span className="section-kicker">GOOD THINGS, MADE FOR LEARNING</span>
              <h2>Find something that sparks.</h2>
              <p>Discover resources from educators, creators and trusted vendors.</p>
            </div>
            <a className="text-link" href="#get-started">Visit the marketplace <ArrowRight size={16} /></a>
          </div>
          <div className="product-grid">
            {products.map(({ name, kind, price, icon: Icon, color }) => (
              <article className="product-card" key={name}>
                <div className={`product-art ${color}`}>
                  <Icon size={42} strokeWidth={1.25} />
                  <span className="product-tag">EDU247 PICK</span>
                </div>
                <div className="product-details">
                  <div><span>{kind}</span><h3>{name}</h3></div>
                  <strong>{price}</strong>
                </div>
                <a href="#get-started" aria-label={`Discover ${name}`} className="product-action"><ArrowUpRight size={17} /></a>
              </article>
            ))}
          </div>
        </section>

        <section className="section ecosystem-section" id="developers">
          <div className="ecosystem-heading">
            <span className="section-kicker">A PLATFORM BUILT TO GROW</span>
            <h2>Many parts. One shared purpose.</h2>
            <p>EDU247 is a connected ecosystem, thoughtfully developed in stages — with people, safety and useful learning at the center.</p>
          </div>
          <div className="ecosystem-grid">
            <article className="ecosystem-card ecosystem-card-learning">
              <span className="ecosystem-icon"><GraduationCap size={22} /></span>
              <span className="ecosystem-label">01 / LEARNING</span>
              <h3>Learning that meets you where you are.</h3>
              <p>Courses, lessons, assignments and progress tools for students and teachers.</p>
              <div className="ecosystem-pills"><span>Courses</span><span>Assignments</span><span>Progress</span></div>
            </article>
            <article className="ecosystem-card ecosystem-card-ai">
              <span className="ecosystem-icon"><BrainCircuit size={22} /></span>
              <span className="ecosystem-label">02 / AI STUDIO</span>
              <h3>Human creativity, with a helpful boost.</h3>
              <p>Tutor, lesson and visual-creation tools built to support—not replace—people.</p>
              <div className="ecosystem-pills"><span>AI tutor</span><span>Creation tools</span><span>Human review</span></div>
            </article>
            <article className="ecosystem-card ecosystem-card-dev">
              <span className="ecosystem-icon"><Code2 size={22} /></span>
              <span className="ecosystem-label">03 / BUILDERS</span>
              <h3>Open doors for people who build.</h3>
              <p>Developer resources, contribution workflows, collaboration and hackathons.</p>
              <div className="ecosystem-pills"><span>Developer portal</span><span>Community</span><span>Hackathons</span></div>
            </article>
            <article className="ecosystem-card ecosystem-card-market">
              <span className="ecosystem-icon"><ShoppingBag size={22} /></span>
              <span className="ecosystem-label">04 / MARKET</span>
              <h3>Useful resources, all in one place.</h3>
              <p>A marketplace concept for trusted digital and physical education products.</p>
              <div className="ecosystem-pills"><span>Digital</span><span>Physical</span><span>Vendors</span></div>
            </article>
            <article className="ecosystem-card ecosystem-card-trust" id="trust">
              <span className="ecosystem-icon"><ShieldCheck size={22} /></span>
              <span className="ecosystem-label">05 / TRUST &amp; SAFETY</span>
              <h3>Designed with care from the start.</h3>
              <p>Privacy, access controls, moderation and responsible AI are part of the plan.</p>
              <div className="ecosystem-pills"><span>Privacy</span><span>Safety</span><span>Oversight</span></div>
            </article>
            <article className="ecosystem-card ecosystem-card-admin" id="platform">
              <span className="ecosystem-icon"><Network size={22} /></span>
              <span className="ecosystem-label">06 / PLATFORM</span>
              <h3>One foundation. Room to evolve.</h3>
              <p>Connected services, integrations and administration built to scale in phases.</p>
              <div className="ecosystem-pills"><span>Services</span><span>Integrations</span><span>Analytics</span></div>
            </article>
          </div>
          <div className="architecture-note">
            <div className="architecture-mark"><Network size={18} /></div>
            <p><strong>A clear technical boundary.</strong> Core platform services are planned in Go. Dedicated AI services are planned in Python, connected through versioned APIs and events.</p>
            <a href="#get-started">How we build <ArrowUpRight size={15} /></a>
          </div>
        </section>

        <section className="section journey-section" id="journey">
          <div className="ecosystem-heading">
            <span className="section-kicker">FROM FIRST IDEA TO REAL-WORLD IMPACT</span>
            <h2>One journey, connected end to end.</h2>
            <p>Learning, creation, review and delivery are designed to work together—with clear roles and human oversight at every step.</p>
          </div>
          <div className="journey-steps">
            <article className="journey-step">
              <span className="journey-number">01</span>
              <span className="journey-icon"><BookOpenCheck size={20} /></span>
              <span className="journey-label">CREATE</span>
              <h3>Teachers shape the learning.</h3>
              <p>Build courses, lessons and activities, with AI tools to help turn ideas into useful materials.</p>
            </article>
            <article className="journey-step">
              <span className="journey-number">02</span>
              <span className="journey-icon"><GraduationCap size={20} /></span>
              <span className="journey-label">LEARN</span>
              <h3>Students find their next step.</h3>
              <p>Explore lessons, get tutor support, complete assignments and follow personal progress.</p>
            </article>
            <article className="journey-step">
              <span className="journey-number">03</span>
              <span className="journey-icon"><ShieldCheck size={20} /></span>
              <span className="journey-label">REVIEW</span>
              <h3>People keep quality in focus.</h3>
              <p>Moderation, security review and audit workflows help protect users and platform integrity.</p>
            </article>
            <article className="journey-step">
              <span className="journey-number">04</span>
              <span className="journey-icon"><Store size={20} /></span>
              <span className="journey-label">DISCOVER</span>
              <h3>Good resources find their people.</h3>
              <p>Vendors can list educational products for discovery, purchase and delivery through the marketplace.</p>
            </article>
            <article className="journey-step">
              <span className="journey-number">05</span>
              <span className="journey-icon"><Settings size={20} /></span>
              <span className="journey-label">IMPROVE</span>
              <h3>Feedback makes the platform better.</h3>
              <p>Administrators use oversight and analytics to guide operations, improvements and future releases.</p>
            </article>
          </div>
          <div className="journey-footnote"><ShieldCheck size={16} /><span>Privacy, accessibility and responsible AI are part of the journey—not a final step.</span></div>
        </section>

        <section className="section roadmap-section" id="roadmap">
          <div className="section-heading roadmap-heading">
            <div>
              <span className="section-kicker">A THOUGHTFUL PATH TO DELIVERY</span>
              <h2>Build it in meaningful stages.</h2>
              <p className="section-subtitle">A proposed 52-week sequence from discovery foundations to pilot readiness.</p>
            </div>
            <span className="roadmap-status">Initial planning model</span>
          </div>
          <div className="roadmap-grid">
            {deliveryPhases.map(({ number, name, timing, description, icon: Icon, color }) => (
              <article className="roadmap-card" key={number}>
                <div className="roadmap-card-top">
                  <span className={`roadmap-icon ${color}`}><Icon size={19} /></span>
                  <span className="roadmap-number">{number}</span>
                </div>
                <span className="roadmap-timing">{timing}</span>
                <h3>{name}</h3>
                <p>{description}</p>
              </article>
            ))}
          </div>
          <p className="roadmap-disclaimer">Indicative sequencing only. Detailed scope, dependencies, resources and dates will be validated during discovery and pilot planning.</p>
        </section>

        <section className="community-band" id="community">
          <div className="community-art" aria-hidden="true">
            <span className="community-avatar avatar-a"><GraduationCap size={18} /></span>
            <span className="community-avatar avatar-b"><Code2 size={17} /></span>
            <span className="community-avatar avatar-c"><Lightbulb size={17} /></span>
            <span className="community-avatar avatar-d"><BookOpen size={17} /></span>
            <span className="community-center"><Users size={27} /></span>
          </div>
          <div className="community-copy">
            <span className="section-kicker">BETTER, TOGETHER</span>
            <h2>Build together.<br />Learn together.</h2>
            <p>Meet the educators, learners and builders who believe education can open more doors.</p>
          </div>
          <div className="community-cta">
            <span>Curiosity looks good on everyone.</span>
            <a className="button button-dark" href="#get-started">Find your community <ArrowRight size={16} /></a>
          </div>
        </section>

        <section className="stats-strip" aria-label="EDU247 community goals">
          <div><span className="stat-icon stat-mint"><Users size={21} /></span><span><strong>Every learner</strong><small>Belongs here</small></span></div>
          <div><span className="stat-icon stat-peach"><BookOpen size={21} /></span><span><strong>More ways to learn</strong><small>One connected place</small></span></div>
          <div><span className="stat-icon stat-blue"><Code2 size={21} /></span><span><strong>More room to build</strong><small>Ideas welcome</small></span></div>
          <div><span className="stat-icon stat-lilac"><Store size={21} /></span><span><strong>Good things to discover</strong><small>Learning resources</small></span></div>
        </section>

        <section className="get-started newsletter-section" id="get-started" aria-labelledby="newsletter-title">
          <div className="newsletter-copy">
            <span className="newsletter-kicker"><Cloud size={15} /> EDU247 COMMUNITY</span>
            <h2 id="newsletter-title">Subscribe to news<br />about <em>EDU247.</em></h2>
            <p>Get EDU247 updates, community news, events and practical learning highlights from EDU247.</p>
          </div>
          <div className="newsletter-panel">
            <span className="newsletter-panel-title">Stay up to date</span>
            <p>One useful update at a time. Sign up with your email address.</p>
            <form className="newsletter-form" onSubmit={handleNewsletterSubmit}>
              <label className="newsletter-email-label" htmlFor="newsletter-email">Email address</label>
              <input
                autoComplete="email"
                id="newsletter-email"
                name="email"
                placeholder="Your email address"
                type="email"
                required
              />
              <button className="button button-dark" type="submit">Subscribe <ArrowRight size={16} /></button>
            </form>
            {newsletterMessage ? (
              <p className="newsletter-message" aria-live="polite" role="status">{newsletterMessage}</p>
            ) : (
              <span className="get-started-caption">Preview only: your email won&apos;t be sent or saved.</span>
            )}
          </div>
        </section>

        <footer className="site-footer">
          <div className="footer-main">
            <div className="footer-brand-col">
              <Brand footer />
              <p>Learning without limits.<br />Anytime. Anywhere.</p>
              <a href="mailto:hello@edu247.com" className="footer-email">hello@edu247.com <ArrowUpRight size={14} /></a>
              <SocialLinks variant="footer" />
            </div>
            <div className="footer-links">
              <div><strong>Explore</strong><a href="#learning">Learn</a><a href="#teachers">Teach</a><a href="#ai-tools">AI tools</a><a href="#marketplace">Marketplace</a></div>
              <div><strong>Build with us</strong><a href="#developers">Developers</a><a href="#community">Community</a><a href="#about">About EDU247</a><a href="#get-started">Get in touch</a></div>
              <div><strong>Our promise</strong><a href="#about">Privacy by design</a><a href="#about">Responsible AI</a><a href="#about">Accessibility</a><a href="#about">Trust &amp; safety</a></div>
            </div>
            <a href="#top" className="back-to-top" aria-label="Back to top"><ChevronDown size={18} /></a>
          </div>
          <div className="footer-bottom">
            <span>© 2026 EDU247. Learning without limits.</span>
            <span>Made for a more curious world <span className="footer-heart">✳</span></span>
            <a href="#top">Back to top ↑</a>
          </div>
        </footer>
      </div>
    </main>
  );
}
