import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Fiber + EdgeOne Pages",
  description: "Go Functions allow you to run Go web frameworks like Fiber on EdgeOne Pages. Build full-stack applications with Fiber's Express-inspired API and blazing fast performance.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en-US">
      <head>
        <link rel="icon" href="/fiber-favicon.svg" />
      </head>
      <body
        className="antialiased"
      >
        {children}
      </body>
    </html>
  );
}
