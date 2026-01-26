<!-- Copyright (c) 2025 OpenWonton Authors. SPDX-License-Identifier: MPL-2.0 -->
# Changes From Upstream (HashiCorp Nomad)

This file tracks OpenWonton changes relative to the upstream HashiCorp Nomad
MPL-2.0 licensed releases. Keep entries concise and factual.

## Upstream Base

- Upstream repository: https://github.com/hashicorp/nomad
- Base release/tag: v1.6.5 (see `version/version.go`)
- Base commit SHA: a7cfff372cd52b02232a02bd3ffb5e5d35f88a0f
- Base release date: 2023-12-13

## Summary of Changes

| Area | Description | First version/commit |
| --- | --- | --- |
| Branding & docs | Update docs/website/migration guide to OpenWonton naming and add trademark/non-affiliation language. | unreleased (worktree) |
| Legal/compliance | Add `CHANGES_FROM_UPSTREAM.md`, `LEGAL.md`, `COMPATIBILITY.md`, `CONTRIBUTING.md`, `SECURITY.md`, and an MPL header audit script. | unreleased (worktree) |
| License hygiene | Remove BUSL-licensed UI access-control components and BUSL-tagged tests/helpers; reintroduce clean-room MPL replacements. | unreleased (worktree) |
| Go module path | Use `github.com/openwonton/openwonton` (API: `.../api`). | v1.6.5 fork |

## Removed or Deprecated

- BUSL-licensed UI components removed pending redesign: `ui/app/components/role-editor*`, `ui/app/components/token-editor*`, `ui/app/components/access-control-subnav*`, `ui/app/components/job-editor/review.js`, `ui/app/styles/components/access-control.scss`, `ui/mirage/models/token.js`, `ui/mirage/serializers/role.js`, `ui/tests/acceptance/access-control-test.js`, `ui/tests/acceptance/roles-test.js`, `ui/tests/pages/access-control.js`.
- MPL-licensed access-control policy UI removed to rebuild from scratch: `ui/app/routes/access-control/policies/new.js`, `ui/app/routes/access-control/policies/policy.js`, `ui/app/controllers/access-control/policies/policy.js`, `ui/app/templates/access-control/policies/new.hbs`, `ui/app/templates/access-control/policies/policy.hbs`, `ui/app/components/policy-editor.js`, `ui/app/components/policy-editor.hbs`, `ui/tests/acceptance/policies-test.js`, `ui/tests/integration/components/policy-editor-test.js`, access-control sections in `ui/tests/acceptance/token-test.js`, and the `access-control` route/menu wiring.
- BUSL-licensed rescheduling tests removed: `e2e/rescheduling/doc.go`, `e2e/rescheduling/rescheduling_test.go` (replacement in `e2e/rescheduling/rescheduling.go` with reduced coverage).
- BUSL-licensed tests/helpers removed and reintroduced as clean-room MPL replacements at the same paths: `client/pluginmanager/csimanager/testing.go`, `client/structs/csi_test.go`, `nomad/state/iterator_test.go`, `helper/backoff_test.go`, `testutil/mock_calls.go`.

## Clean-room Specifications

### `ui/app/utils/json-to-hcl.js` (implemented)

Purpose
- Convert a map of HCL variable flags (string keys to string values) into an HCL assignment block string.
- Used by UI flows that prefill `_newDefinitionVariables` when starting or editing jobs.

Inputs
- `flags`: an object/map of string keys to values (typically strings). May be `null` or `undefined`.

Outputs
- A single string containing one assignment per key, each line in the form `key = value\n`.
- If there are no entries, return `""`.
- If there is at least one entry, the output ends with a trailing newline.

Ordering
- Sort keys lexicographically (ASCII, ascending) for deterministic output.

Key formatting
- If key matches `^[A-Za-z_][A-Za-z0-9_]*$`, emit as-is.
- Otherwise, emit a quoted string (same escaping rules as values).

Value formatting
- `null`/`undefined` -> empty string -> `""`.
- If value is boolean or number, emit `String(value)` as a literal.
- If value is a string `v`:
  - Let `t = v.trim()` for detection only.
  - If `t` is empty: emit `""`.
  - If `t` is already quoted (`"..."` or `'...'`) with matching quotes: emit `t`.
  - If `t` matches boolean (`true`/`false`, case-insensitive) or numeric literal (`-?\d+(\.\d+)?([eE][+-]?\d+)?`), emit `t`.
  - If `t` starts with `{` and ends with `}`, or starts with `[` and ends with `]`, emit `t` (treat as HCL object/list expression).
  - Otherwise, emit a quoted string using the original `v` (not trimmed) with escapes:
    - `\` -> `\\`
    - `"` -> `\"`
    - newline -> `\n`
    - tab -> `\t`
    - carriage return -> `\r`

Examples
```text
Input:
{ name: "alice", count: "42", enabled: "true", tags: '["a", "b"]', empty: "" }

Output:
count = 42
enabled = true
empty = ""
name = "alice"
tags = ["a", "b"]
```

## Clean-room Implementations (Status)

| File | Status | Notes |
| --- | --- | --- |
| `ui/app/utils/json-to-hcl.js` | implemented | See spec above. |
| `e2e/rescheduling/rescheduling.go` | implemented | Minimal coverage replacement; expand to restore upstream parity as needed. |
| `client/pluginmanager/csimanager/testing.go` | implemented | Replacement for BUSL mock helpers. |
| `client/structs/csi_test.go` | implemented | Replacement for BUSL tests. |
| `nomad/state/iterator_test.go` | implemented | Replacement for BUSL tests. |
| `helper/backoff_test.go` | implemented | Replacement for BUSL tests. |
| `testutil/mock_calls.go` | implemented | Replacement for BUSL helper. |
