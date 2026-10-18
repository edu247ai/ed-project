#!/usr/bin/env bash
set -euo pipefail

log() {
  local level="$1"
  shift
  echo "[$level] $*"
}

ensure_dir() {
  local dir="$1"
  if [ ! -d "$dir" ]; then
    mkdir -p "$dir"
  fi
}
