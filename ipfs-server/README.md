# EDU247 IPFS Server

The EDU247 IPFS Server is the planned content-addressed storage and file
distribution service for the EDU247 ecosystem. It is intended to provide a
consistent way for services to store, retrieve, and verify files such as
datasets, model checkpoints, curriculum content, student submissions, and
generated learning materials.

> **Implementation status:** This repository currently contains the service
> directory scaffold only. The API, IPFS node integration, storage backends,
> security controls, monitoring, and tests described below are planned; they
> are not implemented or available to run yet.

## Purpose

The service is intended to support:

- AI datasets, model checkpoints, and generated outputs
- Curriculum files, lesson plans, and learning materials
- Student submissions, projects, and voice recordings
- Optional email attachments and large files
- Outputs from creator, recommendation, documentation, code-review, and
  security-auditor services

Each uploaded object can be addressed by its Content Identifier (CID). A CID
allows a client to verify that retrieved bytes match the referenced content.
It does **not** make content private, guarantee that a copy remains available,
or make unsafe content safe.

## Service layout

The service follows an independently deployable microservice layout:

```text
ipfs-server/
├── api/             # REST endpoints, request handling, and access middleware
├── node/            # IPFS daemon, libp2p, peers, blockstore, and datastore
├── storage/         # Pinned/unpinned objects, temporary blocks, snapshots, GC
├── cid/             # CID generation, validation, resolution, and mapping
├── distribution/    # Replication, sharding, caching, and load balancing
├── adapters/        # Database, S3/MinIO, Redis, cluster, and external adapters
├── pipeline/        # File ingestion, preprocessing, indexing, and retrieval
├── security/        # Access control, encryption, signatures, and audit
├── monitoring/      # Metrics, logs, events, and dashboards
├── config/          # Environment-based settings and profiles
├── tests/           # Unit, integration, performance tests, and mocks
├── scripts/         # Setup, backup, maintenance, and utility scripts
├── docs/            # Architecture, protocols, API reference, and guides
└── README.md
```

## Planned capabilities

| Capability | Intended behavior |
| --- | --- |
| Content addressing | Identify content by CID and verify retrieved bytes |
| Pinning | Keep explicitly selected content available on configured nodes |
| Retrieval | Fetch content from a local node or configured IPFS peers/gateways |
| Replication | Configure copies across approved nodes or an IPFS Cluster |
| Storage adapters | Integrate with approved databases, S3-compatible storage, and Redis |
| File pipeline | Validate and process uploads before indexing or distribution |
| Observability | Report node health, peer count, storage use, replication, and latency |

Availability depends on pinning and network configuration. IPFS is not a
backup policy by itself: services remain responsible for retention,
replication requirements, and recovery procedures.

## Planned API

The following routes are a proposed API contract; they are not implemented in
this scaffold:

| Method | Route | Purpose |
| --- | --- | --- |
| `POST` | `/api/upload` | Upload content and return its CID and metadata |
| `GET` | `/api/file/{cid}` | Retrieve content by CID |
| `POST` | `/api/pin/{cid}` | Request that content be pinned |
| `DELETE` | `/api/pin/{cid}` | Request that content be unpinned |
| `GET` | `/api/cid/{cid}` | Validate or resolve a CID and its metadata |

Before implementation, define request/response schemas, size limits,
authentication and authorization behavior, error responses, and API versioning
in `docs/api-reference/`.

## File flow

The intended upload flow is:

1. Authenticate and authorize the requesting service or user.
2. Validate size, media type, and applicable content-safety requirements.
3. Apply encryption where required before storing content on shared or public
   infrastructure.
4. Ingest content, calculate/verify its CID, and store metadata.
5. Apply the requested pinning and replication policy.
6. Return the CID and storage status to the caller.

Retrieval should verify the content against its CID and enforce authorization
before returning private or restricted content.

## Security and privacy

IPFS content may be discoverable by peers or gateways depending on deployment
and configuration. Do not upload sensitive or personal data in plaintext to a
public or shared network. Encrypt content before upload when confidentiality
is required, and manage keys outside this service.

Planned controls include:

- JWT-based service/user authentication and role-based authorization
- CID and metadata validation
- Encryption and signature verification where required
- Audit events for access and administrative operations
- Secret values supplied through an approved secret manager or environment

These controls are **not present yet**. AES-256, JWT access control, role
policies, and audit logging must not be treated as implemented guarantees until
they are delivered and tested.

## Integration with EDU247 services

Services should access storage through a dedicated adapter rather than
depending directly on this service's internal implementation. The proposed
service-side integration point is:

```text
<service>/adapters/ipfs/
```

Adapters should use the versioned IPFS API contract and handle timeouts,
authorization, retries, and storage status explicitly. Email attachments are
optional and must follow email retention and privacy requirements.

## Configuration and operations

Configuration will be environment-based and should cover at least:

- IPFS API endpoint and network mode
- Storage paths and resource limits
- Pinning and replication policies
- Authentication and secret references
- Logging, metrics, and health-check settings

Do not commit secrets or private keys. Operational scripts are intended to live
under `scripts/` and should document their effects before being used against
production data.

## Testing

The planned test suite includes:

- Unit tests for CID handling, validation, access control, and pipeline logic
- Integration tests with mocked or disposable IPFS nodes and adapters
- Performance tests for upload, retrieval, and replication behavior
- Failure tests for unavailable peers, storage limits, and invalid content

No service tests or executable service code are present in the current
scaffold.

## Documentation

Service documentation is organized under `docs/`:

```text
docs/
├── architecture/
├── protocols/
├── developer-guide/
└── api-reference/
```

## Development roadmap

1. Define and version the API schemas and access-control model.
2. Implement a minimal local-node upload, CID verification, and retrieval flow.
3. Add pinning, storage adapters, and integration tests.
4. Add encryption/key-management integration, audit events, and monitoring.
5. Validate replication, recovery, performance, and deployment procedures.

The IPFS Server is a planned shared storage capability. Do not treat it as a
required production dependency until its implementation, security controls, and
operational guarantees have been reviewed and deployed.
