"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { ArrowLeft, ArrowRight, Check, GraduationCap, BookOpenCheck, Code2, Store, ShieldCheck, Settings, Sparkles, type LucideIcon } from "lucide-react";
import { AuthShell } from "../components/AuthShell";

type Role = "student" | "teacher" | "developer" | "vendor" | "reviewer" | "administrator";

type ProfileDraft = {
  name: string;
  email: string;
};

type OnboardingDraft = {
  role: Role;
  interests: string[];
  weeklyGoal: string;
};

const roles: { id: Role; label: string; description: string; icon: LucideIcon }[] = [
  { id: "student", label: "Student / learner", description: "Explore courses and grow my skills", icon: GraduationCap },
  { id: "teacher", label: "Teacher / educator", description: "Create lessons and support learners", icon: BookOpenCheck },
  { id: "developer", label: "Developer", description: "Build, contribute and collaborate", icon: Code2 },
  { id: "vendor", label: "Education vendor", description: "Share useful learning products", icon: Store },
  { id: "reviewer", label: "Security reviewer", description: "Help keep the platform trustworthy", icon: ShieldCheck },
  { id: "administrator", label: "Institution / administrator", description: "Manage a learning community", icon: Settings },
];

const interests = [
  "Artificial intelligence",
  "Technology",
  "Teaching & learning",
  "Creative skills",
  "Career development",
  "Science & STEM",
  "Entrepreneurship",
  "Digital creativity",
];

const weeklyGoals = [
  { id: "1-2", label: "1–2 hours", detail: "A steady start" },
  { id: "3-4", label: "3–4 hours", detail: "Build a rhythm" },
  { id: "5+", label: "5+ hours", detail: "Make more time" },
];

export default function OnboardingPage() {
  const [step, setStep] = useState(0);
  const [profile, setProfile] = useState<ProfileDraft | null>(null);
  const [role, setRole] = useState<Role | "">("");
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);
  const [weeklyGoal, setWeeklyGoal] = useState("");
  const [completed, setCompleted] = useState(false);

  useEffect(() => {
    const rawProfile = sessionStorage.getItem("edu247-profile-draft");
    const rawOnboarding = sessionStorage.getItem("edu247-onboarding-draft");
    if (rawProfile) {
      try {
        const parsed: unknown = JSON.parse(rawProfile);
        if (
          typeof parsed === "object" &&
          parsed !== null &&
          "name" in parsed &&
          "email" in parsed &&
          typeof parsed.name === "string" &&
          typeof parsed.email === "string"
        ) {
          setProfile({ name: parsed.name, email: parsed.email });
        }
      } catch {
        sessionStorage.removeItem("edu247-profile-draft");
      }
    }
    if (rawOnboarding) {
      try {
        const parsed: unknown = JSON.parse(rawOnboarding);
        if (typeof parsed === "object" && parsed !== null) {
          if ("role" in parsed && typeof parsed.role === "string" && roles.some((item) => item.id === parsed.role)) {
            setRole(parsed.role as Role);
          }
          if ("interests" in parsed && Array.isArray(parsed.interests)) {
            setSelectedInterests(parsed.interests.filter((interest): interest is string => typeof interest === "string" && interests.includes(interest)));
          }
          if ("weeklyGoal" in parsed && typeof parsed.weeklyGoal === "string" && weeklyGoals.some((goal) => goal.id === parsed.weeklyGoal)) {
            setWeeklyGoal(parsed.weeklyGoal);
          }
          if ("completed" in parsed && parsed.completed === true) setCompleted(true);
        }
      } catch {
        sessionStorage.removeItem("edu247-onboarding-draft");
      }
    }
  }, []);

  function saveDraft(next: Partial<OnboardingDraft> & { completed?: boolean } = {}) {
    const draft: OnboardingDraft & { completed?: boolean } = {
      role: (next.role ?? role) as Role,
      interests: next.interests ?? selectedInterests,
      weeklyGoal: next.weeklyGoal ?? weeklyGoal,
      ...("completed" in next ? { completed: next.completed } : {}),
    };
    sessionStorage.setItem("edu247-onboarding-draft", JSON.stringify(draft));
  }

  function toggleInterest(interest: string) {
    const next = selectedInterests.includes(interest)
      ? selectedInterests.filter((item) => item !== interest)
      : [...selectedInterests, interest];
    setSelectedInterests(next);
    saveDraft({ interests: next });
  }

  function chooseRole(nextRole: Role) {
    setRole(nextRole);
    saveDraft({ role: nextRole });
  }

  function chooseGoal(nextGoal: string) {
    setWeeklyGoal(nextGoal);
    saveDraft({ weeklyGoal: nextGoal });
  }

  function continueStep() {
    if (step === 0 && !role) return;
    if (step === 1 && selectedInterests.length === 0) return;
    if (step === 2 && !weeklyGoal) return;
    if (step < 2) {
      setStep((current) => current + 1);
      return;
    }
    saveDraft({ completed: true });
    setCompleted(true);
  }

  const displayName = profile?.name.split(" ")[0] ?? "there";

  return (
    <AuthShell
      eyebrow="PERSONALIZE YOUR EXPERIENCE"
      title={completed ? "Your learning space is taking shape." : `Welcome${profile ? `, ${displayName}` : ""}!`}
      description={completed
        ? "Here is the profile you selected for your EDU247 preview."
        : "A few quick choices help us shape a more relevant education experience for you."}
    >
      <div className="auth-notice" role="status">
        Onboarding preview only. Your selections are kept for this browser session; no account is created and no data is sent to a server.
      </div>

      {completed ? (
        <div className="onboarding-complete">
          <span className="onboarding-complete-icon"><Check size={29} /></span>
          <h3>Thanks for helping us make it yours.</h3>
          <p>This preview profile shows the kinds of details EDU247 can use to personalize your learning experience.</p>
          <dl className="onboarding-summary">
            <div><dt>Your role</dt><dd>{roles.find((item) => item.id === role)?.label ?? "Not selected"}</dd></div>
            <div><dt>Your interests</dt><dd>{selectedInterests.join(", ") || "Not selected"}</dd></div>
            <div><dt>Weekly learning goal</dt><dd>{weeklyGoals.find((goal) => goal.id === weeklyGoal)?.label ?? "Not selected"}</dd></div>
          </dl>
          <Link className="auth-submit auth-submit-link" href="/">Explore EDU247 <ArrowRight size={17} /></Link>
          <button className="onboarding-edit" onClick={() => { saveDraft({ completed: false }); setCompleted(false); setStep(0); }} type="button">Edit my choices</button>
        </div>
      ) : (
        <>
          <div className="onboarding-progress" aria-label={`Step ${step + 1} of 3`}>
            <div className="onboarding-progress-label"><span>Step {step + 1} of 3</span><span>{[ "Your role", "Your interests", "Your learning goal" ][step]}</span></div>
            <div className="onboarding-progress-track"><span style={{ width: `${((step + 1) / 3) * 100}%` }} /></div>
          </div>

          {step === 0 && (
            <fieldset className="onboarding-fieldset">
              <legend>Which best describes you?</legend>
              <p>You can explore other parts of EDU247 anytime.</p>
              <div className="onboarding-role-grid">
                {roles.map(({ id, label, description, icon: Icon }) => (
                  <button
                    aria-pressed={role === id}
                    className={`onboarding-role${role === id ? " onboarding-selected" : ""}`}
                    key={id}
                    onClick={() => chooseRole(id)}
                    type="button"
                  >
                    <span className="onboarding-role-icon"><Icon size={19} /></span>
                    <span><strong>{label}</strong><small>{description}</small></span>
                    <span className="onboarding-radio" />
                  </button>
                ))}
              </div>
            </fieldset>
          )}

          {step === 1 && (
            <fieldset className="onboarding-fieldset">
              <legend>What are you interested in?</legend>
              <p>Choose one or more. We&apos;ll use these to shape your preview.</p>
              <div className="onboarding-interest-grid">
                {interests.map((interest) => (
                  <button
                    aria-pressed={selectedInterests.includes(interest)}
                    className={`onboarding-interest${selectedInterests.includes(interest) ? " onboarding-selected" : ""}`}
                    key={interest}
                    onClick={() => toggleInterest(interest)}
                    type="button"
                  >
                    {selectedInterests.includes(interest) && <Check size={15} />}
                    {interest}
                  </button>
                ))}
              </div>
            </fieldset>
          )}

          {step === 2 && (
            <fieldset className="onboarding-fieldset">
              <legend>How much time would you like to set aside?</legend>
              <p>There&apos;s no wrong answer. You can always change your goal.</p>
              <div className="onboarding-goal-list">
                {weeklyGoals.map(({ id, label, detail }) => (
                  <button
                    aria-pressed={weeklyGoal === id}
                    className={`onboarding-goal${weeklyGoal === id ? " onboarding-selected" : ""}`}
                    key={id}
                    onClick={() => chooseGoal(id)}
                    type="button"
                  >
                    <span className="onboarding-radio" />
                    <span><strong>{label} each week</strong><small>{detail}</small></span>
                  </button>
                ))}
              </div>
              <p className="onboarding-name-note"><Sparkles size={15} /> Your learning goal is a guide, not a deadline.</p>
            </fieldset>
          )}

          <div className="onboarding-controls">
            <button
              className="onboarding-back"
              disabled={step === 0}
              onClick={() => setStep((current) => Math.max(0, current - 1))}
              type="button"
            >
              <ArrowLeft size={16} /> Back
            </button>
            <button
              className="auth-submit onboarding-next"
              disabled={(step === 0 && !role) || (step === 1 && selectedInterests.length === 0) || (step === 2 && !weeklyGoal)}
              onClick={continueStep}
              type="button"
            >
              {step === 2 ? "Finish preview" : "Continue"} <ArrowRight size={17} />
            </button>
          </div>
          {!profile && (
            <p className="auth-switch">Want to add your name first? <Link href="/auth/register">Preview account setup</Link></p>
          )}
        </>
      )}
    </AuthShell>
  );
}
