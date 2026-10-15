type SocialNetwork = {
  name: string;
  href: string;
  paths: string[];
};

const networks: SocialNetwork[] = [
  {
    name: "LinkedIn",
    href: "https://www.linkedin.com/in/edu247ai/",
    paths: [
      "M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.35V9h3.414v1.561h.049c.476-.9 1.637-1.85 3.37-1.85 3.601 0 4.265 2.37 4.265 5.455v6.286zM5.337 7.433a2.062 2.062 0 1 1-.003-4.125 2.062 2.062 0 0 1 .003 4.125zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z",
    ],
  },
  {
    name: "Instagram",
    href: "https://www.instagram.com/edu247ai/",
    paths: [
      "M7.5 2h9A5.5 5.5 0 0 1 22 7.5v9a5.5 5.5 0 0 1-5.5 5.5h-9A5.5 5.5 0 0 1 2 16.5v-9A5.5 5.5 0 0 1 7.5 2Zm0 2A3.5 3.5 0 0 0 4 7.5v9A3.5 3.5 0 0 0 7.5 20h9a3.5 3.5 0 0 0 3.5-3.5v-9A3.5 3.5 0 0 0 16.5 4h-9Z",
      "M12 7a5 5 0 1 1 0 10 5 5 0 0 1 0-10Zm0 2a3 3 0 1 0 0 6 3 3 0 0 0 0-6Z",
    ],
  },
  {
    name: "Facebook",
    href: "https://www.facebook.com/edu247ai/",
    paths: [
      "M13.4 21v-8.2h2.8l.4-3.2h-3.2v-2c0-.9.3-1.5 1.6-1.5h1.7V3.2c-.3 0-1.3-.2-2.5-.2-2.5 0-4.2 1.5-4.2 4.3v2.3H7.2v3.2H10V21h3.4Z",
    ],
  },
  {
    name: "X",
    href: "https://x.com/edu247ai",
    paths: [
      "M18.9 2H22l-6.8 7.8L23.2 22h-6.3L12 14.6 5.6 22H2.4l7.3-8.4L1.9 2h6.5l4.4 6.8L18.9 2Zm-1.1 18h1.7L7.4 3.9H5.6L17.8 20Z",
    ],
  },
  {
    name: "TikTok",
    href: "https://www.tiktok.com/@edu247ai",
    paths: [
      "M12.53.02c1.31-.02 2.61-.01 3.92-.02.08 1.53.63 3.01 1.72 4.1 1.1 1.1 2.59 1.65 4.12 1.73v4.13c-1.44-.05-2.88-.36-4.22-.98-.58-.26-1.12-.59-1.66-.94-.01 2.86.01 5.71-.02 8.56-.08 1.38-.53 2.75-1.34 3.86-1.32 1.92-3.61 3.18-5.95 3.22-1.44.08-2.88-.31-4.1-1.08-2.01-1.22-3.47-3.72-3.37-6.24.08-1.39.55-2.77 1.4-3.89 1.37-1.88 3.68-3.08 6-3.02.21.01.42.02.63.04.01 1.52-.04 3.04-.04 4.56-1.02-.33-2.22-.24-3.08.43-.63.45-1.1 1.15-1.22 1.92-.14.84.09 1.73.62 2.39.54.69 1.38 1.13 2.26 1.15.58.04 1.17-.12 1.67-.42.83-.5 1.4-1.42 1.42-2.4.04-4.31.02-8.62.04-12.93.01-1.46-.01-2.92.02-4.38Z",
    ],
  },
];

type SocialLinksProps = {
  variant: "header" | "footer";
};

export function SocialLinks({ variant }: SocialLinksProps) {
  return (
    <nav className={`social-links social-links-${variant}`} aria-label="EDU247 social media">
      {networks.map(({ name, href, paths }) => (
        <a
          aria-label={`Visit ${name}`}
          href={href}
          key={name}
          rel="noopener noreferrer"
          target="_blank"
          title={name}
        >
          <svg aria-hidden="true" viewBox="0 0 24 24" focusable="false">
            {paths.map((path) => <path d={path} key={path} />)}
          </svg>
        </a>
      ))}
    </nav>
  );
}
