# EDU247 AI Policy

## 1. Purpose and scope

This policy sets minimum expectations for designing, building, evaluating,
integrating, and operating AI features in EDU247. It applies to EDU247
employees, contributors, service owners, and vendors who develop or operate AI
capabilities for the platform.

The policy covers models and AI-enabled features used for content generation,
recommendations, tutoring, assessment support, classification, summarization,
and other learner- or educator-facing uses.

> **Policy status:** This document states requirements and intended controls;
> it does not certify that a feature complies or that a safeguard is
> implemented. AI features must not be represented as production-ready until
> their owner has documented and verified the applicable controls.

## 2. Core requirements

EDU247 AI features must:

- Support learners and educators without replacing accountable human judgment.
- Be appropriate for the intended audience, including children and young
  people where applicable.
- Be evaluated for safety, quality, bias, privacy, and reliability before
  release and after material changes.
- Clearly communicate when a user is interacting with AI and explain material
  limitations.
- Use only data that is necessary, authorized, and appropriate for the
  feature's stated purpose.
- Provide a way to report harmful or incorrect output and a process for review.
- Have an accountable service owner who can disable or roll back the feature.

## 3. Prohibited or restricted uses

Unless a separate, documented review approves an appropriate use with
effective safeguards, AI must not:

- Make final high-impact decisions about a person's access, grades, discipline,
  eligibility, or rights without meaningful human review and appeal.
- Infer sensitive traits, diagnose a person's mental or physical health, or
  profile students for unrelated purposes.
- Conduct covert surveillance or emotion recognition in educational settings.
- Generate content that exploits, sexualizes, endangers, or targets minors.
- Be presented as a human, qualified professional, or authoritative source
  when it is not.
- Use personal, student, or confidential information to train or improve a
  model without an approved purpose, lawful basis, consent where required, and
  documented retention/deletion terms.

## 4. Risk review and release

Before enabling an AI feature for users, the service owner must document:

1. Intended use, users, affected groups, and foreseeable misuse.
2. Model/provider, inputs, outputs, data flows, retention, and external
   recipients.
3. Potential impacts, including harmful, biased, misleading, or inaccessible
   outcomes.
4. Evaluation data and measurable acceptance criteria appropriate to the use.
5. Human oversight, user disclosure, feedback, escalation, and disablement
   mechanisms.
6. Security, privacy, and operational risks and the mitigations in place.

Higher-impact uses require stronger evidence and human review. A feature must
not be launched when material risks are unresolved or evaluation results fail
the documented criteria.

## 5. Lifecycle and monitoring

Owners must keep a record of the model/provider version, configuration,
evaluation results, known limitations, and release approval. Re-evaluate after
material model, prompt, data, provider, or product changes. Monitor reported
problems and relevant quality or safety signals, and pause or roll back a
feature when its behavior presents an unacceptable risk.

## 6. Roles and accountability

- **Service owner:** documents intended use, data flows, evaluations,
  safeguards, and operational readiness.
- **Product owner:** ensures user experience, disclosure, and meaningful
  feedback or human-review paths.
- **Engineering and operations:** implement, test, monitor, secure, and support
  the service.
- **Reviewers:** assess privacy, security, safety, accessibility, and
  educational impact within their responsibilities.

The owner of the affected product or service remains accountable for outcomes;
responsibility cannot be delegated to a model or provider.

## 7. Related policies

- [AI Ethics](./AI-Ethics.md)
- [AI Governance](./AI-Governance.md)
- [AI Safety](./AI-Safety.md)
- [AI Transparency](./AI-Transparency.md)
- [Security guidance](../security/README.md)

Review this policy at least annually and after a significant incident,
material product change, or relevant change in legal or institutional
requirements.
