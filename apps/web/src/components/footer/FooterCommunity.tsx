import FooterLinkGroup from "./FooterLinkGroup";

export default function FooterCommunity() {
  return (
    <FooterLinkGroup
      title="Community"
      links={[
        { label: "Learners", href: "/learn" },
        { label: "Partners", href: "/institutions" },
        { label: "Beta Testers", href: "/community" },
        { label: "Free Courses", href: "/learn" },
        { label: "Blog", href: "/#learning" },
        { label: "The EDU247 Podcast", href: "/community" },
        { label: "Tech Blog", href: "/developers" },
        { label: "Social Impact", href: "/#why-edu247" },
      ]}
    />
  );
}
