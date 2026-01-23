<!-- Copyright (c) 2025 OpenWonton Authors. SPDX-License-Identifier: MPL-2.0 -->
# Nomad Directory Diff Report (v1.6.5 base)

This report summarizes changes under `nomad/` relative to the upstream v1.6.5
base commit. It uses heuristic categorization of diff lines to separate import
path rewrites, branding strings, and other functional changes.

## Base Reference

- Base commit: `a7cfff372cd52b02232a02bd3ffb5e5d35f88a0f` (Nomad v1.6.5)
- Commits touching `nomad/` after base:
  - `6fe6e99ff` "first nomad to wonton commit"
  - `3954732d6` "license stuff"
- Diffstat: 264 files changed, 933 insertions(+), 907 deletions(-).

## Categorized Line Changes (heuristic)

| Category | Added | Removed | Notes |
| --- | --- | --- | --- |
| Import path rewrites | 869 | 870 | Lines containing `github.com/openwonton/openwonton` or `github.com/hashicorp/nomad`. |
| Branding strings | 4 | 0 | Lines containing `OpenWonton` or `Nomad`, or string literals with `nomad`/`wonton`. |
| Functional/other | 60 | 37 | Remaining lines after import/branding classification. |

## Files With Functional/Other Line Changes (heuristic)

- `nomad/client_agent_endpoint_test.go`
- `nomad/client_fs_endpoint_test.go`
- `nomad/encrypter_test.go`
- `nomad/rpc_test.go`
- `nomad/state/iterator_test.go`
- `nomad/structs/structs.generated.go`
- `nomad/worker_string_schedulerworkerstatus.go`
- `nomad/worker_string_workerstatus.go`

## Files With Branding Line Changes (heuristic)

- `nomad/state/iterator_test.go`
- `nomad/structs/structs.generated.go`
- `nomad/worker_string_schedulerworkerstatus.go`
- `nomad/worker_string_workerstatus.go`

## Top Modified Subdirectories

- `nomad/structs`: 54 files
- `nomad/state`: 32 files
- `nomad/mock`: 11 files
- `nomad/drainer`: 11 files
- `nomad/stream`: 8 files
- `nomad/volumewatcher`: 6 files
- `nomad/deploymentwatcher`: 6 files

## Method

- Diff source: `git diff a7cfff372cd52b02232a02bd3ffb5e5d35f88a0f -- nomad`.
- Import path lines: contain `github.com/openwonton/openwonton` or
  `github.com/hashicorp/nomad`.
- Branding lines: contain `OpenWonton` or `Nomad`, or string literals with
  `nomad` or `wonton`.
- Functional/other: all remaining added/removed lines.
- Classification precedence: import > branding > functional.

## Notes

- `nomad/structs/structs.generated.go` and `nomad/worker_string_*.go` are
  generated outputs; changes there should be validated by their generators.
