#!/usr/bin/env bash
# Copyright (c) 2025 OpenWonton Authors.
# SPDX-License-Identifier: MPL-2.0
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
target="${1:-$root/nomad}"

if ! command -v rg >/dev/null 2>&1; then
  echo "rg is required for this audit script" >&2
  exit 1
fi

rg --files "$target" -g '*.go' | while read -r file; do
  if ! rg -q "SPDX-License-Identifier: MPL-2.0|Mozilla Public License, v. 2.0" "$file"; then
    echo "$file"
  fi
done
