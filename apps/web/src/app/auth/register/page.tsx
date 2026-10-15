"use client";

import { useState, type FormEvent } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { ArrowRight, CircleAlert } from "lucide-react";
import { AuthShell } from "../components/AuthShell";
import { PasswordField } from "../components/PasswordField";
import { SocialAuthOptions } from "../components/SocialAuthOptions";

type ProfileDraft = {
  name: string;
  email: string;
};

export default function RegisterPage() {
  const router = useRouter();
  const [message, setMessage] = useState("");

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const name = String(formData.get("name") ?? "").trim();
    const email = String(formData.get("email") ?? "").trim();

    if (!name || !email) {
      setMessage("Please enter your name and a valid email address.");
      return;
    }

    const profile: ProfileDraft = { name, email };
    sessionStorage.setItem("edu247-profile-draft", JSON.stringify(profile));
    router.push("/auth/onboarding");
  }

  return (
    <AuthShell
      eyebrow="YOUR LEARNING JOURNEY STARTS HERE"
      title="Join EDU247"
      description="Create a profile and tell us a little about what you want to learn."
    >
      <div className="auth-notice" role="status">
        Account creation is not connected yet. You can preview the registration and onboarding flow; no account will be created. Your name and email are kept only in this browser session. Your password is never saved.
      </div>
      <SocialAuthOptions action="create an account" />
      <form className="auth-form" onSubmit={handleSubmit}>
        <div className="auth-field">
          <label htmlFor="register-name">Your name</label>
          <input id="register-name" name="name" type="text" autoComplete="name" placeholder="e.g. Alex Morgan" maxLength={80} required />
        </div>
        <div className="auth-field">
          <label htmlFor="register-email">Email address</label>
          <input id="register-email" name="email" type="email" autoComplete="email" placeholder="you@example.com" required />
        </div>
        <PasswordField id="register-password" label="Create a password" autoComplete="new-password" minLength={8} />
        <p className="auth-field-hint">Use at least 8 characters. Your password is not stored or sent.</p>
        {message && (
          <div className="auth-message" role="alert">
            <CircleAlert size={17} />
            <span>{message}</span>
          </div>
        )}
        <button className="auth-submit" type="submit">Continue to onboarding <ArrowRight size={17} /></button>
      </form>
      <p className="auth-switch">Already have an account? <Link href="/auth/login">Sign in</Link></p>
    </AuthShell>
  );
}
