"use client";

import { useState, type FormEvent } from "react";
import Link from "next/link";
import { ArrowRight, CircleAlert } from "lucide-react";
import { AuthShell } from "../components/AuthShell";
import { PasswordField } from "../components/PasswordField";
import { SocialAuthOptions } from "../components/SocialAuthOptions";

export default function LoginPage() {
  const [message, setMessage] = useState("");

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setMessage("Sign-in is not connected yet. Your details were not sent or saved. Please check back when EDU247 account services are available.");
  }

  return (
    <AuthShell
      eyebrow="WELCOME BACK"
      title="Sign in to EDU247"
      description="Pick up where your learning journey left off."
    >
      <div className="auth-notice" role="status">
        Account sign-in is not connected yet. This form is a preview and will not send your credentials.
      </div>
      <SocialAuthOptions action="sign in" />
      <form className="auth-form" onSubmit={handleSubmit}>
        <div className="auth-field">
          <label htmlFor="login-email">Email address</label>
          <input id="login-email" name="email" type="email" autoComplete="email" placeholder="you@example.com" required />
        </div>
        <PasswordField id="login-password" label="Password" autoComplete="current-password" />
        <div className="auth-form-row">
          <label className="auth-checkbox"><input name="remember" type="checkbox" /> <span>Remember me</span></label>
          <Link href="/auth/forgot-password">Forgot password?</Link>
        </div>
        {message && (
          <div className="auth-message" role="alert">
            <CircleAlert size={17} />
            <span>{message}</span>
          </div>
        )}
        <button className="auth-submit" type="submit">Sign in <ArrowRight size={17} /></button>
      </form>
      <p className="auth-switch">New to EDU247? <Link href="/auth/register">Create an account</Link></p>
    </AuthShell>
  );
}
