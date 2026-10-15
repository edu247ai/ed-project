import Link from "next/link";

export type FooterLink = {
  label: string;
  href: string;
};

type FooterLinkGroupProps = {
  title: string;
  links: FooterLink[];
};

export default function FooterLinkGroup({ title, links }: FooterLinkGroupProps) {
  return (
    <div className="footer-link-group">
      <strong>{title}</strong>
      {links.map(({ label, href }) => (
        href.startsWith("mailto:") ? (
          <a href={href} key={label}>{label}</a>
        ) : (
          <Link href={href} key={label}>{label}</Link>
        )
      ))}
    </div>
  );
}
