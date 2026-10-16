# EDU247 AI Safety

## Purpose

This document sets expectations for preventing, detecting, and responding to
unsafe AI behavior. It complements the [AI Policy](./AI-Policy.md) and
[AI Governance](./AI-Governance.md). Safety controls must be implemented and
tested for each use; this document alone does not make a system safe.

## Safety by design

For each feature, identify foreseeable misuse, failure modes, affected users,
and potential severity before release. Safeguards should fit the use and may
include:

- Input validation, access control, and limits on sensitive data
- Output checks and age-appropriate content handling
- Retrieval/source constraints for factual or educational answers
- Human review for high-impact or uncertain outputs
- Clear refusal, escalation, and correction paths
- Limits on automated actions and external tool access

Testing should include ordinary use, ambiguous input, known failure modes,
adversarial attempts, and relevant accessibility and language needs.

## Learner and child safety

Where children or young people may use a feature:

- Use age-appropriate language and behavior.
- Do not encourage secrecy, dependency, self-harm, dangerous acts, or
  inappropriate personal relationships.
- Do not request unnecessary personal details or sensitive disclosures.
- Provide a clear path to a trusted adult or human support when the situation
  is beyond the feature's role.
- Do not present generated advice as professional medical, psychological,
  legal, or emergency guidance.

## High-impact outputs

AI output must not be the sole basis for high-stakes grading, discipline,
access, eligibility, safeguarding, or other consequential decisions. Provide
meaningful human review, clear uncertainty and limitations, and a way to
correct inaccurate information or challenge an outcome.

## Incident response

An AI safety incident includes harmful content, material misinformation,
exposure of protected data, discriminatory impact, unsafe tool action, or
failure of a required safeguard.

When an incident is identified:

1. Protect affected users and restrict or pause the relevant feature when
   needed.
2. Preserve necessary evidence securely while minimizing personal data.
3. Notify the accountable service owner and follow the organizational
   incident process.
4. Assess scope, severity, affected people, and whether notification is
   required.
5. Identify the cause, implement and test corrections, and obtain the required
   review before restoring service.
6. Record lessons learned and update evaluations, safeguards, and user
   guidance.

Do not promise automatic incident logging, immediate detection, or user
notification unless those capabilities are implemented and verified.

## Monitoring and stop conditions

Define measurable safety signals and an owner before release. Pause, restrict,
or roll back a feature when a severe incident occurs, a critical safeguard
fails, or monitoring indicates that documented acceptance thresholds are no
longer met. Resume only after the cause is understood and remediation is
verified.
