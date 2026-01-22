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
- Use `scripts/audit-mpl-headers.sh` to list Go files under `nomad/` missing
  MPL-2.0 headers.
- Header audit (all Go files excluding vendor/) currently missing MPL headers
  in:
  - `api/deployments_test.go`
  - `client/fingerprint/env_aws_cpu.go`
  - `client/logmon/proto/logmon.pb.go`
  - `client/structs/structs.generated.go`
  - `command/agent/bindata_assetfs.go`
  - `drivers/docker/docklog/proto/docker_logger.pb.go`
  - `drivers/shared/executor/proto/executor.pb.go`
  - `helper/raftutil/msgtypes.go`
  - `plugins/base/proto/base.pb.go`
  - `plugins/device/proto/device.pb.go`
  - `plugins/drivers/proto/driver.pb.go`
  - `plugins/shared/hclspec/hcl_spec.pb.go`
  - `plugins/shared/structs/proto/attribute.pb.go`
  - `plugins/shared/structs/proto/recoverable_error.pb.go`
  - `plugins/shared/structs/proto/stats.pb.go`

## Trademarks

Nomad is a trademark of HashiCorp, Inc. OpenWonton is not affiliated with,
endorsed by, or sponsored by HashiCorp.
