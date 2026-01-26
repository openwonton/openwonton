<!-- Copyright (c) 2025 OpenWonton Authors. SPDX-License-Identifier: MPL-2.0 -->
# Compatibility Notes

OpenWonton prioritizes compatibility with HashiCorp Nomad MPL-2.0 releases to
keep migration friction low. Where possible, configuration formats, API shapes,
and CLI behavior are preserved.

## Compatibility Highlights

- **CLI**: `wonton` is the primary CLI; a `nomad` shim is provided for existing
  scripts and automation.
- **Config and API**: Nomad-compatible configuration keys and API semantics are
  preserved unless explicitly documented.
- **Directory naming**: The `nomad/` directory name is retained for API and
  configuration compatibility. This name does not imply endorsement or
  affiliation with HashiCorp.

## Binary Compatibility

- **Linux glibc baseline**: Official linux/amd64 builds target glibc 2.31
  (built on Ubuntu 20.04). They are expected to run on Debian 11/Ubuntu 20.04
  and newer; older glibc-based distributions may not be supported by the
  official binaries.
- **RPM-based distributions**: RPM packages are built from the same linux/amd64
  artifact. Expect compatibility with RHEL 9/UBI 9 and other RPM distros with
  glibc 2.31+; RHEL 8/UBI 8 and older are not supported by the official
  binaries.

## Future Compatibility Work

- TODO: Consider renaming `nomad/` to `compat/nomad/` or `legacy/nomad/` with a
  documented migration path.
