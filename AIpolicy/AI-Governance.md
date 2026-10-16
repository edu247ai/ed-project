# EDU247 AI Governance

## Purpose

AI governance defines who is accountable for AI features and what evidence is
required through their lifecycle. It operationalizes the
[AI Policy](./AI-Policy.md); it is a process requirement, not a claim that a
governance board or approval workflow is already operating.

## Ownership

Every AI feature must have an accountable service owner before user release.
The owner coordinates product, engineering, privacy, security, accessibility,
and educational review appropriate to the feature's impact.

Owners are responsible for maintaining:

- Intended use, user groups, and prohibited uses
- Model/provider, versions, configurations, and material changes
- Input/output data flows and retention arrangements
- Evaluations, limitations, incidents, and release decisions
- Human oversight and operational disablement procedures

## Risk classification

Assess risk based on the potential impact of the use, affected people,
autonomy, data sensitivity, scale, and reversibility—not simply on whether a
model is generative.

At minimum, classify a feature as:

- **Lower impact:** assists with low-stakes drafting or organization where a
  person reviews before relying on the output.
- **Elevated impact:** affects learning pathways, tutoring, moderation,
  recommendations, or decisions that may materially affect a learner.
- **High impact:** could affect grades, access, discipline, eligibility,
  safety, or other important interests.

High-impact features require documented cross-functional review, robust
evaluation, meaningful human decision-making, an appeal or correction path,
and an explicit release decision. Do not deploy them solely on model output.

## Lifecycle gates

### Before development

Define the use, users, risks, data requirements, alternatives, owner, and
success criteria. Do not collect data or procure a model before confirming
that the purpose and data use are approved.

### Before release

Verify the intended behavior and failure modes; evaluate quality, safety,
fairness, privacy, and security; confirm disclosures and human review; record
known limitations; and approve monitoring, incident response, and rollback
plans. Release must be blocked when required controls are missing or
acceptance criteria are not met.

### After release

Monitor appropriate quality and safety indicators, feedback, incidents,
provider/model changes, and data handling. Reassess after material changes and
at a documented interval. Pause or retire a feature when its risks cannot be
managed.

## Change and provider management

Treat changes to a model, provider, prompt, retrieval corpus, training data,
configuration, or intended use as potentially material. Review and re-evaluate
changes proportionate to their likely impact. Contracts with external
providers must address data use, retention, security, incident notification,
and model-training restrictions relevant to EDU247's use.

## Records and review

Maintain review records in the owning service's approved documentation
location. Keep only necessary data and restrict access to sensitive review
materials. Revisit this governance process after incidents, significant
product changes, or changes in applicable requirements.
