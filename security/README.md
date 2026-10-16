# EDU247 Security Guidance

## Purpose and status

This document describes security expectations for EDU247 development. It is
guidance, not evidence that a control has been implemented or independently
validated.

> **Current status:** `security/` currently contains this README only. The
> repository does not implement the cryptography, blockchain audit ledger,
> service authentication, centralized threat monitoring, or incident
> automation previously described here. Treat those capabilities as planned
> until they are implemented, tested, and operated.

For repository vulnerability reporting, see the root
[Security Policy](../SECURITY.md). Its current reporting instructions are a
placeholder and should be completed with a private reporting channel and
response expectations before the project is used with sensitive data.

## Security principles

- **Least privilege:** grant users, services, and automation only the access
  needed for their task.
- **Defense in depth:** do not rely on a gateway, UI, network boundary, or
  single security control as the only protection.
- **Secure defaults:** deny access unless explicitly authorized; validate
  inputs and fail safely.
- **Data minimization:** collect, transmit, retain, and log only what a feature
  needs.
- **Service ownership:** each service enforces authorization for its own
  resources and owns its data.
- **Transparency:** distinguish implemented controls from plans and document
  security assumptions and limitations.

## Secure development baseline

### Identity and authorization

- Authenticate callers at appropriate trust boundaries.
- Enforce authorization on the server and in the service that owns the
  requested resource.
- Do not treat client-side checks or API gateway routing as authorization.
- Validate tokens using a maintained library and an explicit issuer, audience,
  expiry, and key-rotation policy.
- Use service-to-service authentication where services cross trust boundaries.
  Select and document the mechanism before deployment; do not claim mTLS or
  signed-token enforcement unless it is configured and verified.

### Secrets and configuration

- Never commit passwords, API tokens, private keys, or production credentials.
- Load secret values from a protected runtime environment or an approved
  secret manager.
- Commit only safe examples or references; use non-production credentials for
  local development.
- Restrict access to deployment variables and rotate credentials after
  suspected exposure.

### Transport and data protection

- Use HTTPS/TLS at network boundaries and secure internal links according to
  the deployment threat model.
- Encrypt sensitive stored data when required by its classification and
  retention needs.
- Use well-maintained cryptographic libraries and standard protocols; do not
  invent encryption, signing, or key-exchange schemes.
- Do not hard-code a cipher or key size as a substitute for key management,
  authenticated encryption, rotation, and tested recovery procedures.
- Treat content-addressed identifiers and hashes as integrity aids, not
  confidentiality or access-control mechanisms.

### Input handling and dependencies

- Validate inputs at service boundaries and enforce size, type, and rate
  limits appropriate to the endpoint.
- Encode output for its context; avoid unsafe deserialization and shell
  construction from untrusted input.
- Keep dependencies current and review security advisories before release.
- Return errors that help callers recover without revealing secrets,
  implementation details, or private data.

### Logging and audit

- Use structured logs with request or correlation identifiers where available.
- Do not log credentials, tokens, private keys, or sensitive message bodies.
- Restrict log access and define retention based on privacy and operational
  requirements.
- Record security-relevant actions when a service implements an audit feature;
  do not claim that logs are immutable or tamper-proof without a verified
  design.

## AI, files, and personal data

- Do not send personal, student, or confidential data to an external model or
  provider without an approved privacy and security review.
- Treat model prompts, uploaded files, generated content, and retrieved
  documents as untrusted input.
- Apply authorization and retention rules to datasets, checkpoints, generated
  outputs, and attachments.
- A CID, signature, or checksum does not make public content private or safe.
- Document what data a service processes, where it is sent or stored, and how
  it can be deleted.

## Testing and release

As security-sensitive functionality is introduced:

- Add tests for authorization decisions, invalid inputs, failure handling, and
  secret redaction.
- Review dependency and configuration changes.
- Validate deployment permissions and ensure credentials are scoped to the
  required environment.
- Test backup, recovery, and key-rotation procedures before relying on them.
- Record remaining risks and limitations in the owning service documentation.

The root [`tests/`](../tests/) security test directory is currently a
placeholder; no repository-wide security test suite is configured there.

## Incident response

If a credential or sensitive data is exposed:

1. Restrict or revoke the affected credential and contain the exposure.
2. Preserve relevant evidence without copying secrets into public issues or
   logs.
3. Notify the service owner and follow the organization's incident process.
4. Rotate dependent credentials and assess access and data impact.
5. Fix the root cause, validate the remediation, and document follow-up work.

Use the private reporting method documented in the root
[Security Policy](../SECURITY.md) for suspected vulnerabilities; do not publish
unfixed security details in a public issue.

## Keeping this guidance accurate

Update this document when a security control is implemented or its behavior
changes. Describe the enforcement point, operational owner, tests, and known
limitations. Do not list planned controls as active guarantees.
