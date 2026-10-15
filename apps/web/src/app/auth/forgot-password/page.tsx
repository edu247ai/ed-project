"use client";

import { useState, type FormEvent } from "react";
import Link from "next/link";
import { ArrowRight, CircleAlert } from "lucide-react";
import { AuthShell } from "../components/AuthShell";

export default function ForgotPasswordPage() {
  const [message, setMessage] = useState("");

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setMessage("Password recovery is not connected yet. No reset email was sent. Please try again when EDU247 account services are available.");
  }

  return (
    <AuthShell
      eyebrow="ACCOUNT HELP"
      title="Forgot your password?"
      description="Enter the email address associated with your EDU247 account."
    >
      <div className="auth-notice" role="status">
        Password recovery is not connected yet. This preview will not send an email.
      </div>
      <form className="auth-form" onSubmit={handleSubmit}>
        <div className="auth-field">
          <label htmlFor="reset-email">Email address</label>
          <input id="reset-email" name="email" type="email" autoComplete="email" placeholder="you@example.com" required />
        </div>
        {message && (
          <div className="auth-message" role="alert">
            <CircleAlert size={17} />
            <span>{message}</span>
          </div>
        )}
        <button className="auth-submit" type="submit">Continue <ArrowRight size={17} /></button>
      </form>
      <p className="auth-switch"><Link href="/auth/login">Back to sign in</Link></p>
    </AuthShell>
  );
}
