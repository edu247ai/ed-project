# Developer Guide

## Getting Started

### 1. Install Dependencies
Use the setup script:
scripts/setup/install.sh

Code

### 2. Run the Node
go run cmd/ipfs-node/main.go

Code

### 3. Upload a File
curl -X POST http://localhost:8082/upload \
-H "X-API-Key: super-secret-key" \
-d '{"data":"aGVsbG8gd29ybGQ="}'

Code

### 4. Download a File
curl "http://localhost:8082/download?cid=<cid>" \
-H "X-API-Key: super-secret-key"

Code

## Directory Structure
ipfs-server/
api/
adapters/
cid/
config/
distribution/
docs/
monitoring/
node/
pipeline/
scripts/
security/
storage/
tests/

Code

## Extending the System
- Add new adapters (S3, Redis, DB)
- Add new pipeline processors
- Add new security modules
- Add new distribution strategies