import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "EDU247 — Learn, Create and Grow",
  description:
    "An AI-powered education ecosystem for learning, teaching, building and educational commerce.",
  icons: {
    icon: [{ url: "/logoimg/edu247-logo.png", type: "image/png" }],
    apple: [{ url: "/logoimg/edu247-logo.png", type: "image/png" }],
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
