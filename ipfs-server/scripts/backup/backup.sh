#!/usr/bin/env bash
set -euo pipefail

BASE_DIR=${1:-/var/ipfs/prod}
BACKUP_DIR=${2:-/var/backups/ipfs}
TS=$(date +"%Y%m%d-%H%M%S")

echo "[backup] base: $BASE_DIR"
echo "[backup] target: $BACKUP_DIR/$TS"

mkdir -p "$BACKUP_DIR/$TS"

tar -czf "$BACKUP_DIR/$TS/storage.tar.gz" -C "$BASE_DIR" storage snapshots

echo "[backup] completed at $TS"
