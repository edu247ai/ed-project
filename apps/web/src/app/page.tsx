"use client";

import Link from "next/link";
import { useState, type FormEvent } from "react";
import {
  ArrowRight,
  ArrowUpRight,
  Atom,
  Award,
  BookOpenCheck,
  BriefcaseBusiness,
  BrainCircuit,
  Check,
  Code2,
  GraduationCap,
  HandHeart,
  Languages,
  Lightbulb,
  Network,
  Palette,
  Route,
  School,
  Sparkles,
  Store,
  UsersRound,
  Users,
  type LucideIcon,
} from "lucide-react";
import { SiteFooter } from "../components/SiteFooter";
import { SiteHeader } from "../components/SiteHeader";
import { DesignCard } from "../components/DesignCard";

const impactFeatures: { title: string; description: string; icon: LucideIcon }[] = [
  {
    title: "More access to learning",
    description: "Open learning pathways designed to help more people build useful knowledge and skills.",
    icon: BookOpenCheck,
  },
  {
    title: "AI that opens possibilities",
    description: "Thoughtful AI tools can give learners and educators new ways to explore and create.",
    icon: BrainCircuit,
  },
  {
    title: "Opportunity for young people",
    description: "Support curiosity, practical skills, confidence, and pathways to future opportunities.",
    icon: Sparkles,
  },
  {
    title: "Support for teachers",
    description: "Give educators more ways to prepare lessons, engage classes, and focus on learners.",
    icon: HandHeart,
  },
  {
    title: "A platform for creators",
    description: "Help educators and creators share learning resources with wider communities.",
    icon: Store,
  },
  {
    title: "Transformation for institutions",
    description: "Connect schools, colleges, and training centers with modern learning experiences.",
    icon: School,
  },
];

const roles = [
  {
    title: "Students",
    description: "Explore new subjects, grow your skills, and find a path that fits your ambition.",
    href: "/student",
    icon: GraduationCap,
    tone: "mint",
    action: "Start learning",
  },
  {
    title: "Teachers",
    description: "Create meaningful lessons and bring fresh ideas into your classroom.",
    href: "/teacher",
    icon: BookOpenCheck,
    tone: "gold",
    action: "Explore teaching",
  },
  {
    title: "Vendors & creators",
    description: "Share useful educational products and resources with learners and educators.",
    href: "/vendor",
    icon: Store,
    tone: "blue",
    action: "Explore the marketplace",
  },
];

const learningAreas: { title: string; description: string; icon: LucideIcon }[] = [
  {
    title: "Courses for every goal",
    description: "Explore subjects and skills for learning, work, and everyday life.",
    icon: BookOpenCheck,
  },
  {
    title: "Practical skills",
    description: "Build knowledge through guided activities and relevant real-world practice.",
    icon: Lightbulb,
  },
  {
    title: "Learning paths",
    description: "Follow a clear sequence of lessons as you work toward your next goal.",
    icon: Network,
  },
  {
    title: "AI tutor",
    description: "Get extra explanations and practice support when you need a hand.",
    icon: BrainCircuit,
  },
  {
    title: "Certifications",
    description: "Work toward recognizing the skills and knowledge you build.",
    icon: Check,
  },
];

const popularCourses: { title: string; description: string; icon: LucideIcon; label: string }[] = [
  {
    title: "AI foundations for everyone",
    description: "Understand core AI ideas and learn how to use emerging tools thoughtfully.",
    icon: BrainCircuit,
    label: "Technology",
  },
  {
    title: "Digital skills for work and life",
    description: "Build practical confidence with the digital tools used in study and work.",
    icon: Code2,
    label: "Career skills",
  },
  {
    title: "Creative problem solving",
    description: "Explore ways to turn questions into ideas, prototypes, and useful solutions.",
    icon: Lightbulb,
    label: "Creative skills",
  },
];

const categories: { title: string; description: string; icon: LucideIcon }[] = [
  {
    title: "STEM",
    description: "Explore science, mathematics, engineering, and discovery.",
    icon: Atom,
  },
  {
    title: "Business",
    description: "Develop practical knowledge for enterprise and work.",
    icon: BriefcaseBusiness,
  },
  {
    title: "Technology",
    description: "Build confidence with digital tools, computing, and AI.",
    icon: Code2,
  },
  {
    title: "Creative arts",
    description: "Learn through design, storytelling, media, and making.",
    icon: Palette,
  },
  {
    title: "Languages",
    description: "Practice communication and connect across languages.",
    icon: Languages,
  },
  {
    title: "Professional skills",
    description: "Grow the transferable skills that support lifelong careers.",
    icon: Award,
  },
];

const platformOffers: { title: string; description: string; icon: LucideIcon; href: string; action: string }[] = [
  {
    title: "Learning",
    description: "Explore courses, build skills, and follow a path toward your goals.",
    icon: BookOpenCheck,
    href: "/learn",
    action: "Explore learning",
  },
  {
    title: "Teaching",
    description: "Plan lessons, create resources, and support learners in your classes.",
    icon: GraduationCap,
    href: "/teach",
    action: "Explore teaching",
  },
  {
    title: "Marketplace",
    description: "Discover educational resources and support independent creators.",
    icon: Store,
    href: "/marketplace",
    action: "Visit marketplace",
  },
  {
    title: "AI tools",
    description: "Use thoughtful AI assistance to explore ideas and support learning.",
    icon: BrainCircuit,
    href: "/learn#ai-tools",
    action: "Explore AI tools",
  },
  {
    title: "Community",
    description: "Connect with learners, educators, and creators around shared goals.",
    icon: UsersRound,
    href: "/community",
    action: "Meet the community",
  },
  {
    title: "Institutional solutions",
    description: "Bring connected digital learning experiences to your institution.",
    icon: School,
    href: "/institutions",
    action: "Explore institutions",
  },
];

const creatorTools: { title: string; description: string; icon: LucideIcon }[] = [
  {
    title: "Lesson builder",
    description: "Shape an idea into a clear lesson plan and classroom activity.",
    icon: BookOpenCheck,
  },
  {
    title: "Infographic creator",
    description: "Turn important concepts into visual summaries learners can revisit.",
    icon: Network,
  },
  {
    title: "Comic creator",
    description: "Use visual storytelling to make lessons memorable and engaging.",
    icon: Palette,
  },
];

const impactMeasures = [
  { label: "Learners reached", icon: GraduationCap },
  { label: "Teachers supported", icon: HandHeart },
  { label: "Institutions onboarded", icon: School },
  { label: "Marketplace creators", icon: Store },
  { label: "Courses available", icon: BookOpenCheck },
];

export default function HomePage() {
  const [newsletterMessage, setNewsletterMessage] = useState("");

  function handleNewsletterSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setNewsletterMessage("Newsletter signup is not connected yet. Your email was not sent or saved.");
  }

  return (
    <main className="africa-home">
      <SiteHeader activeAudience="For learners" />
      <div className="page-shell africa-page-shell">
        <section className="africa-hero" id="top" aria-labelledby="hero-heading">
          <div className="africa-hero-copy">
            <span className="africa-eyebrow"><Sparkles size={15} /> LEARNING POWERED BY POSSIBILITY</span>
          <h1 id="hero-heading">Learn, create, and <em>grow without limits.</em></h1>
            <p>
            A modern education ecosystem connecting learners, teachers, creators, and
            institutions with AI-powered tools and opportunities around the world.
            </p>
            <div className="africa-hero-actions">
              <Link className="africa-button africa-button-primary" href="/Dashboard/student-dashboard">
                Start learning <ArrowRight size={18} />
              </Link>
              <a className="africa-button africa-button-quiet" href="#why-edu247">
                Explore EDU247 <ArrowUpRight size={17} />
              </a>
            </div>
            <div className="africa-hero-note">
              <span><HandHeart size={17} /></span>
              People. Knowledge. AI. A brighter tomorrow.
            </div>
          </div>
          <div className="africa-hero-visual" aria-hidden="true">
            <div className="africa-visual-orbit africa-visual-orbit-one" />
            <div className="africa-visual-orbit africa-visual-orbit-two" />
            <div className="africa-visual-hub"><BrainCircuit size={38} /><span>AI</span></div>
            <div className="africa-visual-node africa-visual-node-student"><GraduationCap size={22} /></div>
            <div className="africa-visual-node africa-visual-node-community"><Users size={22} /></div>
            <div className="africa-visual-node africa-visual-node-learning"><BookOpenCheck size={22} /></div>
            <div className="africa-visual-node africa-visual-node-progress"><Network size={22} /></div>
            <span className="africa-visual-caption">People. Knowledge. A brighter tomorrow.</span>
          </div>
          <div className="africa-hero-bottom" aria-label="EDU247 values">
            <span>LEARN</span><span>CREATE</span><span>GROW</span><span>THRIVE</span>
          </div>
        </section>

        <section className="africa-impact-intro" id="why-edu247" aria-labelledby="impact-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">WHY EDU247 MATTERS</span>
            <h2 id="impact-heading">Education can change what&apos;s possible.</h2>
            <p>
              Talent and ambition are everywhere. EDU247 is being shaped to connect
              that potential with learning, tools, and opportunity around the world.
            </p>
          </div>
          <div className="africa-impact-grid">
            {impactFeatures.map(({ title, description, icon: Icon }, index) => (
              <DesignCard className="africa-impact-card" key={title}>
                <span className={`africa-card-icon africa-tone-${index % 3}`}><Icon size={22} /></span>
                <h3>{title}</h3>
                <p>{description}</p>
              </DesignCard>
            ))}
          </div>
        </section>

        <section className="africa-roles" id="roles" aria-labelledby="roles-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">A PLACE FOR YOUR AMBITION</span>
            <h2 id="roles-heading">Choose how you&apos;ll grow.</h2>
            <p>Whether you are learning, teaching, or sharing resources, there is a place to begin.</p>
          </div>
          <div className="africa-role-grid">
            {roles.map(({ title, description, href, icon: Icon, tone, action }) => (
              <Link className="africa-role-link" href={href} key={title}>
                <DesignCard className="africa-role-card">
                  <span className={`africa-role-icon africa-role-${tone}`}><Icon size={25} /></span>
                  <h3>{title}</h3>
                  <p>{description}</p>
                  <span className="africa-card-action">{action}<ArrowRight size={16} /></span>
                </DesignCard>
              </Link>
            ))}
          </div>
        </section>

        <section className="africa-learning" id="learning" aria-labelledby="learning-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">LEARNING OPPORTUNITIES</span>
            <h2 id="learning-heading">Take your next step, your way.</h2>
            <p>Build practical skills and explore subjects that matter to your future.</p>
          </div>
          <div className="africa-learning-grid">
            {learningAreas.map(({ title, description, icon: Icon }) => (
              <DesignCard className="africa-learning-card" key={title}>
                <span><Icon size={21} /></span>
                <h3>{title}</h3>
                <p>{description}</p>
              </DesignCard>
            ))}
          </div>
          <Link className="africa-inline-link" href="/learn">Explore learning <ArrowRight size={16} /></Link>
        </section>

        <section className="africa-learning" id="courses" aria-labelledby="courses-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">COURSES &amp; SKILLS</span>
            <h2 id="courses-heading">Start with a skill that moves you forward.</h2>
            <p>
              Explore sample learning journeys across technology, work, and creativity.
              Course listings and certifications are being prepared for the platform.
            </p>
          </div>
          <div className="africa-impact-grid">
            {popularCourses.map(({ title, description, icon: Icon, label }) => (
              <DesignCard className="africa-impact-card" key={title}>
                <span className="africa-card-icon"><Icon size={22} /></span>
                <span className="africa-course-label">{label} · Learning journey preview</span>
                <h3>{title}</h3>
                <p>{description}</p>
                <Link className="africa-inline-link" href="/learn">
                  Explore learning <ArrowRight size={16} />
                </Link>
              </DesignCard>
            ))}
          </div>
          <div className="africa-course-paths" aria-label="Learning features">
            <span><BrainCircuit size={18} /> AI-powered learning tools</span>
            <span><Route size={18} /> Guided learning paths</span>
            <span><Award size={18} /> Certifications in development</span>
          </div>
        </section>

        <section className="africa-learning" id="categories" aria-labelledby="categories-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">EXPLORE BY CATEGORY</span>
            <h2 id="categories-heading">A place to begin, whatever your ambition.</h2>
            <p>Find a subject area that matches your interests, experience, and next goal.</p>
          </div>
          <div className="africa-learning-grid">
            {categories.map(({ title, description, icon: Icon }) => (
              <DesignCard className="africa-learning-card" key={title}>
                <span><Icon size={21} /></span>
                <h3>{title}</h3>
                <p>{description}</p>
                <Link className="africa-inline-link" href="/learn">
                  Explore {title} <ArrowRight size={16} />
                </Link>
              </DesignCard>
            ))}
          </div>
        </section>

        <section className="africa-teaching" id="teaching" aria-labelledby="teaching-heading">
          <div className="africa-teaching-copy">
            <span className="africa-kicker">FOR EDUCATORS</span>
            <h2 id="teaching-heading">More time for the moments that matter.</h2>
            <p>
              Support your teaching with tools that help you plan, create, and bring ideas
              to life—while keeping your expertise at the heart of every lesson.
            </p>
            <ul>
              {creatorTools.map(({ title }) => <li key={title}><Check size={17} />{title}</li>)}
              <li><Check size={17} />Tools for classes, learners, and progress</li>
            </ul>
            <Link className="africa-inline-link" href="/teach">Explore tools for teachers <ArrowRight size={16} /></Link>
          </div>
          <div className="africa-creator-grid">
            {creatorTools.map(({ title, description, icon: Icon }, index) => (
              <DesignCard className={`africa-creator-card africa-creator-${index + 1}`} key={title}>
                <span><Icon size={23} /></span>
                <h3>{title}</h3>
                <p>{description}</p>
              </DesignCard>
            ))}
            <div className="africa-creator-quote"><Sparkles size={19} /><span>Your ideas can inspire a whole classroom.</span></div>
          </div>
        </section>

        <section className="africa-marketplace" id="marketplace" aria-labelledby="marketplace-heading">
          <div className="africa-marketplace-art" aria-hidden="true">
            <span className="africa-marketplace-orbit" />
            <span className="africa-marketplace-symbol"><Store size={38} /></span>
            <span className="africa-marketplace-badge"><Palette size={17} /> Creator-led learning</span>
          </div>
          <div className="africa-marketplace-copy">
            <span className="africa-kicker">MADE FOR LEARNING. OPEN TO EVERYONE.</span>
            <h2 id="marketplace-heading">A marketplace where creators can thrive.</h2>
            <p>
              Discover digital resources, educational products, and affordable learning
              tools—while creating more ways for African educators and creators to share
              their work with learners around the world.
            </p>
            <div className="africa-marketplace-tags"><span>Digital resources</span><span>Learning products</span><span>Independent creators</span></div>
            <Link className="africa-button africa-button-primary" href="/marketplace">
              Explore the marketplace <ArrowRight size={17} />
            </Link>
          </div>
        </section>

        <section className="africa-community" id="community" aria-labelledby="community-heading">
          <div className="africa-community-art" aria-hidden="true">
            <span className="africa-community-node node-one"><GraduationCap size={19} /></span>
            <span className="africa-community-node node-two"><BookOpenCheck size={19} /></span>
            <span className="africa-community-node node-three"><Lightbulb size={19} /></span>
            <span className="africa-community-center"><Users size={30} /></span>
          </div>
          <div className="africa-community-copy">
            <span className="africa-kicker">COMMUNITY &amp; COLLABORATION</span>
            <h2 id="community-heading">When we learn together, we rise together.</h2>
            <p>
              Grow through learning groups, teacher networks, and youth-led innovation.
              Learners, educators, and creators can connect around shared interests, exchange
              practical ideas, and collaborate on projects. Topic-based groups and discussions
              are designed to make it easier to ask questions, share resources, and learn from
              one another in a welcoming space.
            </p>
            <p>
              Join a study conversation, discover educator networks, or share a resource with
              people working toward similar goals. Community features are being developed with
              safety, inclusion, and constructive participation in mind.
            </p>
            <Link className="africa-button africa-button-outline" href="/community">
              Find your community <ArrowRight size={17} />
            </Link>
          </div>
        </section>

        <section className="africa-learning" id="partners" aria-labelledby="partners-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">INVESTORS &amp; PARTNERS</span>
            <h2 id="partners-heading">Help build a more connected learning future.</h2>
            <p>
              EDU247 is developing an education ecosystem for learners, educators,
              creators, and institutions. We welcome thoughtful conversations with people
              and organizations who share that long-term ambition.
            </p>
          </div>
          <div className="africa-impact-grid">
            <DesignCard className="africa-impact-card">
              <span className="africa-card-icon"><Sparkles size={22} /></span>
              <h3>Investors</h3>
              <p>
                Learn about the platform vision, the needs it aims to address, and ways to
                explore potential alignment. This is an invitation to start a conversation,
                not an offer or promise of investment returns.
              </p>
              <a className="africa-inline-link" href="mailto:hello@edu247.com?subject=EDU247%20investor%20enquiry">
                Discuss the vision <ArrowUpRight size={16} />
              </a>
            </DesignCard>
            <DesignCard className="africa-impact-card">
              <span className="africa-card-icon"><School size={22} /></span>
              <h3>Institutional partners</h3>
              <p>
                Schools, colleges, and training organizations can explore connected
                learning experiences and discuss their goals, needs, and onboarding plans.
              </p>
              <Link className="africa-inline-link" href="/institutions">
                Explore partnerships <ArrowRight size={16} />
              </Link>
            </DesignCard>
            <DesignCard className="africa-impact-card">
              <span className="africa-card-icon"><Store size={22} /></span>
              <h3>Creators and vendors</h3>
              <p>
                Educators and creators can prepare useful resources to share, while vendors
                can learn about offering relevant products through the planned marketplace.
              </p>
              <Link className="africa-inline-link" href="/marketplace">
                Explore creator opportunities <ArrowRight size={16} />
              </Link>
            </DesignCard>
          </div>
        </section>

        <section className="africa-learning" id="what-we-offer" aria-labelledby="offers-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">WHAT WE OFFER</span>
            <h2 id="offers-heading">One ecosystem. More ways to grow.</h2>
            <p>
              Bring learning, teaching, useful resources, and community together in
              experiences designed for people and institutions.
            </p>
          </div>
          <div className="africa-impact-grid">
            {platformOffers.map(({ title, description, icon: Icon, href, action }) => (
              <DesignCard className="africa-impact-card" key={title}>
                <span className="africa-card-icon"><Icon size={22} /></span>
                <h3>{title}</h3>
                <p>{description}</p>
                <Link className="africa-inline-link" href={href}>
                  {action} <ArrowRight size={16} />
                </Link>
              </DesignCard>
            ))}
          </div>
        </section>

        <section className="africa-measures" id="impact" aria-labelledby="measures-heading">
          <div className="africa-section-heading">
            <span className="africa-kicker">MEASURE WHAT MATTERS</span>
            <h2 id="measures-heading">A future of opportunity, built together.</h2>
            <p>These are the measures EDU247 is working toward and will report as the platform grows.</p>
          </div>
          <div className="africa-measure-grid">
            {impactMeasures.map(({ label, icon: Icon }) => (
              <DesignCard as="div" className="africa-measure-card" key={label}>
                <Icon size={22} />
                <strong>{label}</strong>
                <span>Impact measure in development</span>
              </DesignCard>
            ))}
          </div>
        </section>

        <section className="africa-newsletter" id="get-started" aria-labelledby="newsletter-title">
          <div>
            <span className="africa-kicker">GROW WITH US</span>
            <h2 id="newsletter-title">A brighter future begins with learning.</h2>
            <p>Follow the EDU247 journey and hear about new learning opportunities.</p>
          </div>
          <DesignCard className="africa-newsletter-panel">
            <form className="africa-newsletter-form" onSubmit={handleNewsletterSubmit}>
              <label htmlFor="newsletter-email">Get updates from EDU247</label>
              <div>
                <input id="newsletter-email" name="email" type="email" placeholder="Your email address" autoComplete="email" required />
                <button className="africa-button africa-button-primary" type="submit">Subscribe <ArrowRight size={17} /></button>
              </div>
            </form>
            {newsletterMessage ? (
              <p className="africa-newsletter-message" aria-live="polite" role="status">{newsletterMessage}</p>
            ) : (
              <span className="africa-newsletter-note">Preview only: your email won&apos;t be sent or saved.</span>
            )}
          </DesignCard>
        </section>
        <SiteFooter />
      </div>
    </main>
  );
}
