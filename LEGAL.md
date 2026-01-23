<!-- Copyright (c) 2025 OpenWonton Authors. SPDX-License-Identifier: MPL-2.0 -->
# Legal and Provenance

This document records OpenWonton legal and provenance details to support clean
forking and license hygiene.

## License

OpenWonton is licensed under the Mozilla Public License, v. 2.0 (MPL-2.0). See
`LICENSE` for the full text. Unless otherwise noted, new files added to this
repository are licensed under MPL-2.0.

## Upstream Provenance

OpenWonton is forked from the MPL-2.0 licensed HashiCorp Nomad releases. Record
the exact upstream reference points below and keep them in sync with
`CHANGES_FROM_UPSTREAM.md`.

- Upstream repository: https://github.com/hashicorp/nomad
- Base release/tag: v1.6.5 (see `version/version.go`)
- Base commit SHA: a7cfff372cd52b02232a02bd3ffb5e5d35f88a0f
- Base release date: 2023-12-13
- Fork date: 2026-01-22

## NOTICE and Third-Party Material

- Upstream NOTICE file: not present in upstream v1.6.5 tag
- Third-party license exceptions:
  - `client/lib/nsutil/netns_linux.go` (Apache-2.0, derived from CNI plugins; see file header)
  - Third-party dependencies are pulled via Go modules and UI package manifests and remain under their respective licenses.

## Known Non-MPL Headers (Needs Remediation)

No BUSL headers detected as of 2026-01-22. Keep running
`rg -n "BUSL-1.1"` to verify before release.

## Audit Notes

- Keep original copyright headers on modified files and add an OpenWonton line.
- Confirm code originates from MPL-2.0 licensed upstream releases and exclude
  any BSL-era code.
- Document notable changes in `CHANGES_FROM_UPSTREAM.md`.
- Use `scripts/audit-mpl-headers.sh` to list Go files missing MPL-2.0 headers.
  The audit skips vendor/, generated outputs (files matching `*.pb.go`,
  `*.generated.go`, `*bindata*.go`, or containing "Code generated"), and the
  Apache-2.0 file `client/lib/nsutil/netns_linux.go`.
- If the audit returns files, add MPL headers or document exceptions below.

### Header Audit Exceptions

- `client/lib/nsutil/netns_linux.go` (Apache-2.0; derived from CNI plugins).
- Generated Go outputs: `*.pb.go`, `*.generated.go`,
  `command/agent/bindata_assetfs.go`, and any file containing "Code generated".

## Nomad Directory Provenance Audit (v1.6.5 base)

This audit summarizes changes under `nomad/` relative to the upstream v1.6.5
base commit.

- Base commit `a7cfff372cd52b02232a02bd3ffb5e5d35f88a0f` is an ancestor of `HEAD`.
- Commits touching `nomad/` after the base:
  - `6fe6e99ff` "first nomad to wonton commit"
  - `3954732d6` "license stuff"
- Diffstat: 264 files changed, 933 insertions(+), 907 deletions(-).
- BSL/BUSL scan: no hits under `nomad/`.
- Import path status: 0 occurrences of `github.com/hashicorp/nomad` and 865
  occurrences of `github.com/openwonton/openwonton` under `nomad/`.
- Detailed categorized diff report: `NOMAD_DIFF_REPORT.md`.

## Trademarks

Nomad is a trademark of HashiCorp, Inc. OpenWonton is not affiliated with,
endorsed by, or sponsored by HashiCorp.
