#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
LOG_DIR="${LOG_DIR:-/tmp/openwonton-local-ci}"
MAX_LINES="${MAX_LINES:-2000}"

mkdir -p "$LOG_DIR"

timestamp() {
  date +"%Y%m%d-%H%M%S"
}

print_log_summary() {
  local log_file="$1"
  local lines
  lines=$(wc -l <"$log_file" | tr -d " ")
  echo "log lines: $lines"

  if (( lines <= MAX_LINES )); then
    cat "$log_file"
    return 0
  fi

  echo "log too large; showing matches and tail (set MAX_LINES to change)"
  if command -v rg >/dev/null 2>&1; then
    rg -n "FAIL|error|panic|FATAL|timeout" "$log_file" | tail -n 200 || true
  else
    grep -nE "FAIL|error|panic|FATAL|timeout" "$log_file" | tail -n 200 || true
  fi
  tail -n 200 "$log_file" || true
}

run_cmd() {
  local name="$1"
  shift
  local log_file="$LOG_DIR/${name}-$(timestamp).log"
  echo "==> $name"
  echo "log: $log_file"
  local rc=0
  "$@" >"$log_file" 2>&1 || rc=$?
  print_log_summary "$log_file"
  return "$rc"
}

run_quick() {
  run_cmd "make-check" env GOTOOLCHAIN=go1.24.11 make check
}

run_act() {
  local workflow="$1"
  local job="$2"

  local image="${ACT_IMAGE:-ghcr.io/catthehacker/ubuntu:full-22.04}"
  local arch="${ACT_ARCH:-}"

  if [[ -z "$arch" && "$(uname -m)" == "arm64" ]]; then
    arch="linux/amd64"
  fi

  local args=()
  args+=(-W "$workflow" -j "$job")
  if [[ -n "$image" ]]; then
    args+=(-P "ubuntu-22.04=$image")
  fi
  if [[ -n "$arch" ]]; then
    args+=(--container-architecture "$arch")
  fi
  if [[ -n "${ACT_FLAGS:-}" ]]; then
    read -r -a extra <<< "${ACT_FLAGS}"
    args+=("${extra[@]}")
  fi

  run_cmd "act-$(basename "$workflow" .yaml)-$job" act "${args[@]}"
}

usage() {
  cat <<'EOF'
Usage:
  scripts/local-ci.sh quick
    Runs "make check" locally with logs written to ./.local-ci.

  scripts/local-ci.sh act [workflow] [job]
    Runs a GitHub Actions job via act with logs written to ./.local-ci.
    Defaults: workflow=.github/workflows/checks.yaml job=checks

  scripts/local-ci.sh act-core [job]
    Runs a job from .github/workflows/test-core.yaml via act.
    Default job=checks

Environment:
  LOG_DIR     Directory for logs (default /tmp/openwonton-local-ci)
  MAX_LINES   Max lines to print before truncating (default 2000)
  ACT_IMAGE   Act runner image (default ghcr.io/catthehacker/ubuntu:full-22.04)
  ACT_ARCH    Act container arch (default linux/amd64 on arm64 hosts)
  ACT_FLAGS   Extra flags passed to act
EOF
}

case "${1:-}" in
  quick)
    run_quick
    ;;
  act)
    workflow="${2:-.github/workflows/checks.yaml}"
    job="${3:-checks}"
    run_act "$workflow" "$job"
    ;;
  act-core)
    job="${2:-checks}"
    run_act ".github/workflows/test-core.yaml" "$job"
    ;;
  *)
    usage
    exit 1
    ;;
esac
