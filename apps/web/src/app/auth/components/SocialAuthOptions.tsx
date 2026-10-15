"use client";

import { useState } from "react";

const providers = [
  {
    name: "LinkedIn",
    path: "M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.35V9h3.414v1.561h.049c.476-.9 1.637-1.85 3.37-1.85 3.601 0 4.265 2.37 4.265 5.455v6.286zM5.337 7.433a2.062 2.062 0 1 1-.003-4.125 2.062 2.062 0 0 1 .003 4.125zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z",
  },
  {
    name: "GitHub",
    path: "M12 .9a11.1 11.1 0 0 0-3.51 21.63c.55.1.76-.24.76-.53v-2.08c-3.1.68-3.75-1.32-3.75-1.32-.5-1.29-1.24-1.63-1.24-1.63-1.01-.69.08-.68.08-.68 1.12.08 1.71 1.15 1.71 1.15 1 1.71 2.62 1.22 3.26.93.1-.72.39-1.22.71-1.5-2.48-.28-5.09-1.24-5.09-5.52 0-1.22.44-2.22 1.15-3-.12-.28-.5-1.42.11-2.96 0 0 .94-.3 3.05 1.15a10.6 10.6 0 0 1 5.55 0c2.12-1.45 3.05-1.15 3.05-1.15.61 1.54.23 2.68.11 2.96.72.78 1.15 1.78 1.15 3 0 4.29-2.61 5.24-5.1 5.52.4.35.76 1.02.76 2.06V22c0 .29.2.63.77.52A11.1 11.1 0 0 0 12 .9Z",
  },
  {
    name: "Facebook",
    path: "M13.4 21v-8.2h2.8l.4-3.2h-3.2v-2c0-.9.3-1.5 1.6-1.5h1.7V3.2c-.3 0-1.3-.2-2.5-.2-2.5 0-4.2 1.5-4.2 4.3v2.3H7.2v3.2H10V21h3.4Z",
  },
  {
    name: "X",
    path: "M18.9 2H22l-6.8 7.8L23.2 22h-6.3L12 14.6 5.6 22H2.4l7.3-8.4L1.9 2h6.5l4.4 6.8L18.9 2Zm-1.1 18h1.7L7.4 3.9H5.6L17.8 20Z",
  },
] as const;

type SocialAuthOptionsProps = {
  action: "sign in" | "create an account";
};

export function SocialAuthOptions({ action }: SocialAuthOptionsProps) {
  const [message, setMessage] = useState("");

  return (
    <div className="social-auth">
      <div className="social-auth-divider"><span>or continue with</span></div>
      <div className="social-auth-grid">
        {providers.map(({ name, path }) => (
          <button
            aria-label={`Continue with ${name} to ${action}`}
            className="social-auth-button"
            key={name}
            onClick={() => setMessage(`${name} ${action} is not connected yet. No account or session was created.`)}
            title={name}
            type="button"
          >
            <svg aria-hidden="true" viewBox="0 0 24 24" focusable="false">
              <path d={path} />
            </svg>
          </button>
        ))}
      </div>
      {message && <p className="social-auth-message" role="status" aria-live="polite">{message}</p>}
    </div>
  );
}
