Here’s a practical, repo-internal checklist you can literally run through like a release gate.
If you do all of this, you’re operating at the same legal hygiene level as OpenTofu/OpenBao–class forks.

Progress notes

- DONE: Reintroduced clean-room tests and test doubles removed due to BUSL headers:
  - client/pluginmanager/csimanager/testing.go
  - client/structs/csi_test.go
  - nomad/state/iterator_test.go
  - helper/backoff_test.go
  - e2e/rescheduling/rescheduling.go
- NOTE: Access-control UI components were removed due to BUSL headers and are
  expected to be rebuilt later:
  - ui/app/routes/access-control/*
  - ui/app/controllers/access-control/*
  - ui/app/templates/access-control*
  - ui/app/components/role-editor*
  - ui/app/components/token-editor*
  - ui/app/components/access-control-subnav*
  - ui/app/styles/components/access-control.scss
  - ui/app/abilities/role.js
  - ui/app/utils/json-to-hcl.js
  - ui/app/components/job-editor/review.js
  - ui/mirage/models/token.js
  - ui/mirage/serializers/role.js
  - ui/tests/pages/access-control.js
  - ui/tests/acceptance/access-control-test.js
  - ui/tests/acceptance/roles-test.js

⸻

A. Code provenance & license hygiene (highest priority)

1️⃣ Verify every file’s license origin
	•	Audit all files under nomad/:
	•	Confirm they originate only from MPL-2.0–licensed Nomad releases
	•	No code copied from BSL-era commits
	•	For modified files:
	•	Keep original copyright headers
	•	Add your own copyright line below (not replacing)

Task:
	•	Script or manual grep to ensure every MPL file has:
	•	Mozilla Public License, v. 2.0
	•	correct copyright holder

⸻

2️⃣ Include MPL artifacts at repo root
	•	Ensure these exist:
	•	LICENSE (full MPL-2.0 text)
	•	NOTICE or equivalent (if present upstream)
	•	If you added new files:
	•	Either MPL them
	•	Or clearly mark different licenses

Task:
	•	Compare root licensing files with upstream Nomad MPL release

⸻

3️⃣ Track modifications explicitly
	•	Keep a CHANGES_FROM_UPSTREAM.md or similar:
	•	What you changed
	•	What you removed
	•	What you renamed

This matters if anyone ever alleges license contamination.

⸻

B. Trademark & confusion avoidance (this is where people get burned)

4️⃣ Add an explicit trademark disclaimer

In README.md (top or bottom, not hidden):

“Nomad is a trademark of HashiCorp, Inc.
OpenWonton is not affiliated with, endorsed by, or sponsored by HashiCorp.”

Task:
	•	Make sure this appears in:
	•	README
	•	Website (if any)
	•	Docs landing page (if separate)

⸻

5️⃣ Describe the nomad/ directory defensively

Add a short explanation somewhere obvious:

“The nomad/ directory exists for API and configuration compatibility with Nomad.”

This shows intent, which matters legally.

Task:
	•	Add a COMPATIBILITY.md or README section

⸻

6️⃣ Audit naming & branding

Check for accidental trademark usage:
	•	❌ Repo description: “Nomad but open”
	•	❌ “Official Nomad fork”
	•	❌ Logos referencing Nomad imagery
	•	✅ “Nomad-compatible”
	•	✅ “Forked from Nomad (MPL-licensed versions)”
	•	✅ “Community-maintained continuation”

Task:
	•	Re-read README + GitHub repo description as if you were a lawyer

⸻

C. Binary, CLI, and UX checks (often overlooked)

7️⃣ CLI name strategy

If the binary is still called nomad:
	•	Add --version output that clearly says OpenWonton
	•	Add a startup banner or log line:
	•	“OpenWonton – Nomad-compatible scheduler”

Safer long-term:
	•	Provide openwonton binary
	•	Optional nomad shim with warnings

Task:
	•	Run nomad version and check what it communicates

⸻

8️⃣ Config & API references
	•	Keep Nomad config keys for compatibility
	•	Avoid calling them “Nomad config” in docs
	•	Prefer:
	•	“Nomad-compatible configuration”
	•	“Upstream Nomad format”

Task:
	•	Grep docs for “Nomad is” → rewrite to “Nomad-compatible”

⸻

D. Governance & intent signaling (soft law, still important)

9️⃣ Add a clear project intent statement

Something like:

“OpenWonton exists to maintain and evolve the MPL-licensed codebase for community use.”

This helps counter claims of:
	•	bad faith
	•	user confusion
	•	trademark misuse

Task:
	•	One paragraph in README or GOVERNANCE.md

⸻

🔟 Contributor agreement posture
	•	Decide explicitly:
	•	No CLA (common for forks)
	•	Or lightweight DCO sign-off
	•	Be consistent

This matters if code ownership is ever questioned.

Task:
	•	Add CONTRIBUTING.md
	•	State licensing of contributions clearly

⸻

E. Optional but very strong defensive moves

1️⃣1️⃣ Rename later without breaking users

Plan (don’t necessarily execute now):
	•	nomad/ → compat/nomad/
	•	or legacy/nomad/

Task:
	•	Add a TODO comment or tracking issue
	•	Shows foresight if questioned

⸻

1️⃣2️⃣ Internal legal “paper trail”

Keep:
	•	links to upstream MPL releases
	•	commit SHAs you forked from
	•	dates relative to license change

Task:
	•	Store in LEGAL.md (private or public)

⸻

Bottom line (very blunt)

If you:
	•	stay inside MPL-licensed code
	•	avoid branding confusion
	•	over-disclose non-affiliation
	•	treat “Nomad” as compatibility, not identity

…you are operating in the safest possible posture short of having lawyers on retainer.

If you want, next I can:
	•	generate exact disclaimer text copied in tone from OpenTofu/OpenBao
	•	review a real README snippet line-by-line
	•	give you a “what would trigger a C&D” red-flag list so you know when to panic and when to ignore noise
