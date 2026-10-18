#!/usr/bin/env bash
set -euo pipefail

BASE_DIR=${1:-/var/ipfs/prod}

echo "[maintenance] cleaning temp and unpinned..."

rm -rf "$BASE_DIR/storage/temp"/*
echo "[maintenance] temp cleaned."

# optional: run GC binary later
echo "[maintenance] unpinned GC should be triggered by daemon."
