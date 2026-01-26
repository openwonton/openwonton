<!-- Copyright (c) 2025 OpenWonton Authors. SPDX-License-Identifier: MPL-2.0 -->
# Clean-Room Replacement Notes

This file documents the intended clean-room replacements for files removed due
to upstream BUSL-1.1 licensing. It is not legal advice and does not replace
provenance review; confirm authorship and development history via git log.

## Scope

These replacements target functionality that existed in upstream BUSL-1.1
files and were reintroduced under MPL-2.0 with OpenWonton authorship.

| File | Replacement goal | Source of requirements |
| --- | --- | --- |
| `ui/app/utils/json-to-hcl.js` | Recreate HCL assignment string generation for UI variable flags. | Spec in `CHANGES_FROM_UPSTREAM.md`. |
| `e2e/rescheduling/rescheduling.go` | Provide basic rescheduling e2e coverage. | Minimal behavioral check defined by OpenWonton. |
| `client/pluginmanager/csimanager/testing.go` | Provide CSI manager/volume test doubles. | OpenWonton test needs. |
| `client/structs/csi_test.go` | Validate CSI request conversions and options. | OpenWonton test needs. |
| `nomad/state/iterator_test.go` | Validate slice iterator behavior. | OpenWonton test needs. |
| `helper/backoff_test.go` | Validate backoff behavior (base, clamp). | OpenWonton test needs. |
| `testutil/mock_calls.go` | Provide a basic call counter for tests. | OpenWonton test needs. |

## Notes

- Access-control policy UI was removed to enable a from-scratch redesign.
- The rescheduling replacement is intentionally minimal and does not match the
  full upstream test suite coverage.
