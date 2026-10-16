# EDU247 AI Transparency

## Purpose

Users should be able to recognize when AI contributes to an experience and
understand what it can and cannot do. This policy complements the
[AI Policy](./AI-Policy.md).

## User-facing disclosure

When AI materially generates, transforms, classifies, or recommends content,
make that use clear in the interface or documentation at the point where it
matters. Explain in plain language:

- What the AI feature does and what it does not do
- Whether an answer or recommendation may be incomplete or incorrect
- Whether a person reviews the output before it is used for a consequential
  decision
- Whether input is sent to an external model/provider or retained, as
  applicable
- How users can provide feedback, request correction, or reach a human

Do not imply that an AI response was written or verified by a human when it
was not. Do not describe generated output as guaranteed accurate or
unbiased.

## Documentation

Each AI-enabled service should document:

- Intended use and users
- Model/provider and relevant version or configuration
- Input and output types and important data flows
- Evaluation approach, applicable acceptance criteria, and known limitations
- Human oversight and escalation behavior
- Monitoring, incident reporting, and disablement process

Keep service documentation aligned with the deployed version. Clearly label
planned behavior separately from implemented and tested functionality.

## Explanations and user control

Where a recommendation or decision materially affects a person, provide a
useful explanation of the factors or rules that influenced it, without
revealing private data, security-sensitive details, or unsupported
explanations. Allow correction of relevant input data and provide human review
or appeal where the impact warrants it.

Users should not be misled into believing an AI feature is mandatory when it
is optional. Any opt-out or alternative process should be explained when
available.

## Generated and synthetic content

Label generated content when users could reasonably mistake it for verified
factual, official, or human-authored material. Provide appropriate source
references when the system relies on sources and can reliably identify them.
Do not fabricate citations or claim provenance that cannot be verified.

## Transparency limitations

Transparency is not a substitute for privacy, safety, fairness, or human
accountability. Do not disclose sensitive personal data or security details in
the name of explainability. Review disclosures with representative users and
update them when model behavior, data use, or product purpose changes.
