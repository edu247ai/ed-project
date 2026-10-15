import type { Metadata } from "next";
import "./globals.css";

const basePath = process.env.NEXT_PUBLIC_BASE_PATH ?? "";

export const metadata: Metadata = {
  title: "EDU247 — Learn, Create and Grow",
  description:
    "A modern education ecosystem connecting learners, teachers, creators, and institutions with AI-powered learning and opportunity around the world.",
  icons: {
    icon: [{ url: `${basePath}/logoimg/edu247-logo.png`, type: "image/png" }],
    apple: [{ url: `${basePath}/logoimg/edu247-logo.png`, type: "image/png" }],
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
