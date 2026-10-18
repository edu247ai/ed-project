# EDU247 IPFS Server

The EDU247 IPFS Server is the planned content-addressed storage and file
distribution service for the EDU247 ecosystem. It is intended to provide a
consistent way for services to store, retrieve, and verify files such as
datasets, model checkpoints, curriculum content, student submissions, and
generated learning materials.

> **Implementation status:** The service now has a buildable Go module and a
> tested local upload/retrieval core. It is **not yet a Kubo-compatible or
> production-ready IPFS daemon**. In particular, libp2p, DHT, Bitswap, remote
> replication, streaming/chunked uploads, production authentication, and
> pipeline encryption/signing are not implemented.

## Implemented local core

- CIDv1 strings using the raw codec and SHA-256 multihashes, encoded as
  lowercase base32.
- A filesystem-backed local store that checks CID/content correspondence on
  writes and verifies content again on reads.
- Atomic block-file writes and separate pinned/unpinned directories. The
  storage API rejects explicit deletion of pinned blocks; the standalone GC
  scaffold is not yet wired to the blockstore API.
- A pipeline upload/download path and basic HTTP `/upload` and `/download`
  routes.

The current API middleware contains a fixed development API key. Do not expose
the HTTP API to an untrusted network or use it for production data. Other
security and distributed-system components in the directory tree are
scaffolding and must not be treated as active protections or functionality.

Run the module checks from the repository root:

```sh
go -C ipfs-server test ./...
```

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

## Capability status

| Capability | Intended behavior |
| --- | --- |
| Content addressing | Local CIDv1/raw/SHA-256 generation and content verification are implemented |
| Pinning | Local pin/unpin operations are implemented; cluster pinning is not |
| Retrieval | Local block retrieval is implemented; peer/gateway retrieval is not |
| Replication | Not implemented |
| Storage adapters | Adapter scaffolds exist; they are not wired as blockstore backends |
| File pipeline | Byte-slice upload/retrieval is implemented; streaming/chunking is not |
| Observability | Basic in-process counters/dashboard scaffolding exists; Prometheus and swarm metrics are not implemented |

Availability depends on pinning and network configuration. IPFS is not a
backup policy by itself: services remain responsible for retention,
replication requirements, and recovery procedures.

## Current API

The current routes are a minimal development interface, not a versioned or
hardened API contract:

| Method | Route | Purpose |
| --- | --- | --- |
| `POST` | `/upload` | Accept a JSON object with a base64-encoded `data` field and return its CID |
| `GET` | `/download?cid={cid}` | Return the CID, byte size, and base64-encoded data |
| `POST` | `/replicate?cid={cid}` | Development replication handler; validation and auth are incomplete |
| `GET` | `/dashboard` | In-process JSON counters when monitoring is initialized |

`/upload` and `/download` currently use the hard-coded development API key
documented above. Do not rely on these routes for authorization, resource
limits, or production deployments.

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

These controls are **not fully wired or production-ready**. AES-GCM, HMAC,
role-policy, and audit packages are not yet enforced by API/pipeline flows and
must not be treated as implemented guarantees.

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

The local CID, pipeline/storage, and HTTP upload/download flows have tests.
Swarm, replication, production security, and performance-under-load behavior
remain unverified and are not covered by the current test suite.

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

1. Replace the development API key with configured role-based access control.
2. Add streaming/chunked ingestion and explicit size/resource limits.
3. Wire encryption, integrity signatures, audit events, and metrics through API and pipeline flows.
4. Add libp2p transport, peer discovery, DHT, and block exchange.
5. Wire backend adapters and validate cluster replication, recovery, and deployment procedures.

Do not treat the IPFS Server as a production storage dependency until its
security controls, distributed behavior, and operational guarantees have been
implemented, reviewed, and deployed.
