import FooterLinkGroup from "./FooterLinkGroup";

export default function FooterMore() {
  return (
    <FooterLinkGroup
      title="More"
      links={[
        { label: "Press", href: "mailto:hello@edu247.com?subject=EDU247%20press%20enquiry" },
        { label: "Investors", href: "/#partners" },
        { label: "Terms", href: "/institutions#trust" },
        { label: "Privacy", href: "/institutions#trust" },
        { label: "Help", href: "mailto:hello@edu247.com?subject=EDU247%20help" },
        { label: "Accessibility", href: "mailto:hello@edu247.com?subject=Accessibility%20feedback" },
        { label: "Contact", href: "mailto:hello@edu247.com" },
        { label: "Articles", href: "/#learning" },
        { label: "Directory", href: "/learn#categories" },
        { label: "Affiliates", href: "mailto:hello@edu247.com?subject=EDU247%20affiliate%20enquiry" },
        { label: "Modern Slavery Statement", href: "/institutions#trust" },
        { label: "Cookies Preference Center", href: "mailto:hello@edu247.com?subject=Cookie%20preferences%20request" },
      ]}
    />
  );
}
