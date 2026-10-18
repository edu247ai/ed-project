#!/usr/bin/env bash
set -euo pipefail

echo "[setup] installing system dependencies..."

sudo apt update
sudo apt install -y \
  golang \
  git \
  jq \
  curl

echo "[setup] creating ipfs-server directories..."
mkdir -p /var/ipfs/{dev,prod}

echo "[setup] done."
