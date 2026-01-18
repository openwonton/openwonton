PLAN: Rename Nomad Fork to OpenWonton

Overview

We are renaming our HashiCorp Nomad fork to OpenWonton to align with the OpenTofu/OpenBao naming convention and to establish a distinct project identity. The primary CLI will become wonton, with compatibility options to minimize migration friction for existing Nomad users.

Goals
	1.	Establish a clear, consistent OpenWonton brand across code, binaries, docs, and releases.
	2.	Provide a first-class CLI named wonton.
	3.	Preserve operational compatibility with upstream Nomad where feasible, including a low-friction migration path for existing scripts and tooling.
	4.	Avoid breaking changes for users unless explicitly documented and versioned.

Non-goals
	•	Re-architecting scheduling behavior, drivers, or protocol semantics.
	•	Changing API compatibility in this effort (unless required by the rename).
	•	Rewriting history of the repo (unless required for legal/security reasons).

Users & Stakeholders
	•	Platform operators currently using Nomad who want a maintained fork.
	•	Internal maintainers who need a clean dev workflow and CI.
	•	Ecosystem integrators (Terraform, Consul integrations, monitoring dashboards).
	•	Packaging/release consumers (Docker, Homebrew, Linux packages).

Success Metrics
	•	Users can install and run wonton and successfully:
	•	start agent/server
	•	run a job
	•	query allocations
	•	Docs and examples consistently reference wonton.
	•	Compatibility path exists for nomad ... scripts (shim/symlink and/or documented alias).
	•	CI produces versioned artifacts named OpenWonton.
	•	No unresolved references to “Nomad” remain in the project branding areas (except where required for compatibility/legal attribution).

⸻

Requirements

R1: Project naming and identity
	•	Project name: OpenWonton
	•	Branding references:
	•	README title, badges, descriptions
	•	Website/docs landing page
	•	GitHub org/repo description
	•	--version output
	•	Release notes template

Acceptance criteria
	•	README.md headline and first paragraph clearly state:
	•	“OpenWonton is a community fork of HashiCorp Nomad” (or your preferred wording)
	•	Compatibility statement (e.g., “Nomad-compatible orchestrator”)

⸻

R2: CLI binary rename
	•	Primary CLI command: wonton
	•	wonton should support the same command surface as nomad (at least at parity for core commands).

Acceptance criteria
	•	wonton version works
	•	wonton agent -dev works
	•	wonton job run works
	•	Help output and completions mention wonton

⸻

R3: Compatibility mode for existing scripts (recommended)

Provide at least one of the following:

Option A (preferred): ship a nomad shim/symlink
	•	Distributions install both:
	•	wonton (real binary)
	•	nomad (symlink or tiny wrapper that execs wonton)
	•	Wrapper prints a one-time warning (configurable) suggesting migration to wonton.

Option B: documented alias-only approach
	•	Document how to alias nomad → wonton in shells and CI environments.
	•	(Less ideal for automation-heavy users.)

Acceptance criteria
	•	Running nomad ... invokes wonton ... successfully (if Option A)
	•	Documentation includes a “Migration: nomad → wonton” section.

⸻

R4: Optional short alias wt
	•	Provide wt as:
	•	a documented shell alias, OR
	•	an optional installed wrapper in an “extras” package.

Rationale
Two-letter commands collide often; make it opt-in.

Acceptance criteria
	•	Documentation includes a safe “wt alias” snippet.
	•	If shipped as a binary, it must be optional and not required.

⸻

R5: Repo/org naming and Go module strategy

This fork likely uses Go modules and import paths.

Decide one of these strategies:

Strategy 1 (fastest, least invasive): keep Go module path unchanged
	•	Keep module github.com/hashicorp/nomad (or current fork path) to reduce disruption.
	•	Rebrand binaries/docs only.
	•	Pros: minimal downstream Go breakage.
	•	Cons: less “pure” identity.

Strategy 2 (cleanest identity): change Go module path to new org
	•	Set module github.com/openwonton/nomad (or .../openwonton).
	•	Update all internal imports.
	•	Pros: clean ownership.
	•	Cons: can break downstream Go consumers (if any).

Acceptance criteria
	•	Document the chosen strategy and its impact.
	•	Repo builds cleanly from scratch with go build ./....

⸻

R6: Artifacts, packaging, and release naming

Update artifact names across:
	•	CI pipelines (GitHub Actions, etc.)
	•	Docker images
	•	Homebrew formula / apt / rpm packaging (if applicable)
	•	Release archives, checksums, SBOMs

Acceptance criteria
	•	Release artifacts use openwonton/wonton naming consistently.
	•	Docker image names and tags documented.
	•	Checksums and signatures remain valid.

⸻

R7: Docs, examples, and UX polish
	•	Replace nomad CLI usage in docs/examples with wonton.
	•	Keep a compatibility note for users who still type nomad.
	•	Update config examples, environment variables, and service names only if necessary (avoid churn).

Acceptance criteria
	•	Docs build passes
	•	Grep check: no lingering nomad commands in examples unless intentionally shown in migration section.

⸻

R8: Legal, attribution, and licensing
	•	Preserve upstream license files and required notices.
	•	Ensure “fork of HashiCorp Nomad” attribution exists where appropriate.
	•	Avoid implying official affiliation.

Acceptance criteria
	•	LICENSE/NOTICE intact and compliant
	•	README includes attribution and disclaimer

⸻

R9: Internal package and folder naming
	•	Decide whether the internal `nomad/` package and folder remain or are renamed.
	•	If renamed, update imports and any generated paths.
	•	Define explicit exceptions where “nomad” must remain (compatibility, attribution, env vars, etc.).

Acceptance criteria
	•	Decision documented with rationale.
	•	Imports and paths are consistent with the decision.
	•	Repository passes `go test ./...` with the chosen structure.

⸻

Scope

In-scope
	•	Rename CLI to wonton
	•	Update docs, README, help text
	•	Update release artifacts naming
	•	Optional nomad shim for compatibility
	•	CI pipeline updates
	•	Module path decision and implementation (as selected)

Out-of-scope
	•	Feature additions unrelated to rename
	•	Breaking changes to job specs, APIs, driver behaviors
	•	Large refactors

⸻

Implementation Plan

Phase 0: Decision & prep
	•	Choose module strategy (R5)
	•	Choose compatibility approach (R3)
	•	Reserve org/repo names; create placeholder org README

Phase 1: Brand + docs
	•	README, docs landing, badges, project description
	•	Add Migration doc: docs/migration/nomad-to-wonton.md

Phase 1.5: Codebase naming sweep
	•	Inventory `nomad` references in code and docs.
	•	Decide on internal package/folder naming (R9).
	•	Implement directory and import changes if chosen.
	•	Document explicit compatibility exceptions.

Phase 2: CLI rename
	•	Change binary name to wonton
	•	Update command help text and completions
	•	Add nomad shim (if chosen)

Phase 3: CI + releases
	•	Update pipelines to publish wonton artifacts
	•	Update Docker image name and tags
	•	Update release notes automation

Phase 4: Validation + hardening
	•	Integration tests (agent + job run)
	•	Upgrade/migration test (run old scripts via shim)
	•	Docs link checking

Phase 5: Release & comms
	•	Tag first release: e.g., vX.Y.Z-openwonton.1 (or your scheme)
	•	Publish announcement + migration notes

⸻

Testing Plan
	•	Unit tests unchanged + pass
	•	Smoke tests:
	•	wonton agent -dev
	•	submit sample job, verify allocation
	•	wonton status, wonton node status, wonton alloc status
	•	Compatibility test (if shim):
	•	run same commands via nomad and confirm identical behavior
	•	Packaging tests:
	•	verify PATH contains wonton
	•	verify shim behavior and warnings

⸻

Rollout / Backwards Compatibility
	•	First release includes:
	•	wonton as primary
	•	nomad shim available (recommended) for at least N major releases
	•	Deprecation policy:
	•	If you plan to remove nomad shim, announce timeline early and keep it long.

⸻

Risks & Mitigations
	1.	Tooling expects nomad binary
	•	Mitigation: ship shim/symlink; document aliases.
	2.	Go import path breakage (if module path changes)
	•	Mitigation: keep module path or provide a transition guide; consider replace guidance.
	3.	Docs drift / incomplete rename
	•	Mitigation: enforce a CI “grep gate” for nomad command usage outside migration docs.
	4.	User confusion about compatibility
	•	Mitigation: tagline + README clarity + migration page front-and-center.

⸻

Open Questions (to resolve before Phase 2)
	•	Do we keep the Go module path or change it?
	•	Do we rename the internal `nomad/` package/folder, or keep it for compatibility?
	•	If renamed, do we provide a compatibility alias layer for imports/tests?
	•	Do we ship nomad shim by default in all packages?
	•	Do we publish Docker images under openwonton/* or keep compatibility tags too?
	•	What versioning suffix scheme do we want (+openwonton, -openwonton.1, etc.)?

⸻

Deliverables
	•	Updated repos/org naming (as chosen)
	•	wonton CLI binary
	•	Optional nomad shim and optional wt alias
	•	Updated docs + migration guide
	•	Updated CI/release artifacts and packaging
