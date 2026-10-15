import FooterLinkGroup from "./FooterLinkGroup";

export default function FooterAbout() {
  return (
    <FooterLinkGroup
      title="About"
      links={[
        { label: "About", href: "/#why-edu247" },
        { label: "Leadership", href: "/institutions" },
        { label: "Careers", href: "mailto:hello@edu247.com?subject=Careers%20at%20EDU247" },
        { label: "Catalog", href: "/learn#categories" },
        { label: "EDU247 Plus", href: "/learn" },
        { label: "Professional Certificates", href: "/#courses" },
        { label: "Degrees", href: "/institutions" },
        { label: "For Enterprise", href: "/institutions" },
        { label: "For Government", href: "/institutions" },
        { label: "For Campus", href: "/institutions" },
        { label: "Become a Partner", href: "/institutions" },
        { label: "Partner Platforms", href: "/marketplace" },
      ]}
    />
  );
}
