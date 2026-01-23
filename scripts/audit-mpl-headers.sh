#!/usr/bin/env bash
# Copyright (c) 2025 OpenWonton Authors.
# SPDX-License-Identifier: MPL-2.0
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
target="${1:-$root}"
target="${target%/}"

if ! command -v rg >/dev/null 2>&1; then
  echo "rg is required for this audit script" >&2
  exit 1
fi

cd "$root"

ignore_paths=(
  "client/lib/nsutil/netns_linux.go"
)

is_ignored_path() {
  local rel_path="$1"
  local ignored

  for ignored in "${ignore_paths[@]}"; do
    if [[ "$rel_path" == "$ignored" ]]; then
      return 0
    fi
  done

  return 1
}

is_generated_file() {
  local file_path="$1"

  case "$file_path" in
    *.pb.go|*.generated.go|*bindata*.go)
      return 0
      ;;
  esac

  if rg -q -m 1 "Code generated" "$file_path"; then
    return 0
  fi

  return 1
}

target_for_rg="$target"
if [[ "$target" == "$root" ]]; then
  target_for_rg="."
elif [[ "$target" == "$root/"* ]]; then
  target_for_rg="${target#$root/}"
fi

rg --files "$target_for_rg" -g '*.go' -g '!vendor/**' | while read -r file; do
  rel_path="$file"
  if [[ "$rel_path" == "$root/"* ]]; then
    rel_path="${rel_path#$root/}"
  fi
  rel_path="${rel_path#./}"

  if is_ignored_path "$rel_path"; then
    continue
  fi

  if is_generated_file "$file"; then
    continue
  fi

  if ! rg -q "SPDX-License-Identifier: MPL-2.0|Mozilla Public License, v. 2.0" "$file"; then
    echo "$rel_path"
  fi
done
