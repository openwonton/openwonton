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

- BUSL-licensed UI access-control components pending redesign: `ui/app/routes/access-control/*`, `ui/app/controllers/access-control/*`, `ui/app/templates/access-control*`, `ui/app/components/role-editor*`, `ui/app/components/token-editor*`, `ui/app/components/access-control-subnav*`, `ui/app/utils/json-to-hcl.js`, `ui/app/components/job-editor/review.js`, `ui/app/styles/components/access-control.scss`, `ui/mirage/models/token.js`, `ui/mirage/serializers/role.js`, `ui/tests/acceptance/access-control-test.js`, `ui/tests/acceptance/roles-test.js`, `ui/tests/pages/access-control.js`.
- BUSL-licensed rescheduling tests: `e2e/rescheduling/doc.go`, `e2e/rescheduling/rescheduling_test.go` (clean-room replacement in `e2e/rescheduling/rescheduling.go`).
- BUSL-licensed tests/helpers removed and reintroduced as clean-room MPL replacements at the same paths: `client/pluginmanager/csimanager/testing.go`, `client/structs/csi_test.go`, `nomad/state/iterator_test.go`, `helper/backoff_test.go`, `testutil/mock_calls.go`.
