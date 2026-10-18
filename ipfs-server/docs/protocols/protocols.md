# Protocols

## Content Addressing
The system uses SHA-256 to generate CIDs:
cid = sha256(data)

Code

## Replication Protocol
Nodes replicate blocks using HTTP POST:
POST /replicate?cid=<cid>
Body: raw bytes

Code

## Sharding Protocol
Nodes are selected using FNV-1a hashing:
node = hash(cid) % node_count

Code

## Security Protocols
- AES-256-GCM encryption
- HMAC-SHA256 signatures
- RBAC access control
- Immutable audit logs

## Monitoring Protocols
Metrics exposed via:
GET /dashboard

Code

## Swarm Protocol (future)
Planned libp2p integration:
- TCP/UDP on port 4002
- Peer discovery
- Bitswap-like block exchange