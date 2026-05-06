# DECISIONS.md

## D-001: Canonical product source document

- Date: 2026-05-06
- Status: Accepted
- Context: The upstream handoff package needs one canonical product document for implementation truth, and `PRODUCT_BRIEF.md` is not present as an active competing source.
- Decision: `README.md` is the canonical narrative product definition for the current implementation stream.
- Consequences: Implementation planning and slice derivation must stay aligned to `README.md` until a later decision supersedes it.
- Source: `README.md`, `docs/IMPLEMENTATION_HANDOFF.md`

## D-002: Implementation planning does not expand product scope

- Date: 2026-05-06
- Status: Accepted
- Context: The repository already contains a detailed product narrative, but implementation planning must not invent capabilities beyond that handoff.
- Decision: First-run and later implementation planning may translate existing product truth into slices and contracts, but may not add unsupported product behavior.
- Consequences: Any missing or contradictory product truth must be recorded as an open question instead of being guessed in code or planning docs.
- Source: `AGENTS.md`, `FIRST_RUN.md`, `README.md`

## D-003: Product identity

- Date: 2026-05-06
- Status: Accepted
- Context: The public command name must be stable before implementation begins.
- Decision: The public product and command name is `crashsnap`.
- Consequences: Build metadata, executable names, schemas, docs, and tests must use `crashsnap`.
- Source: `README.md`

## D-004: Command-surface posture

- Date: 2026-05-06
- Status: Accepted
- Context: `README.md` claims a broad command surface, but implementation does not need to deliver every command in one session.
- Decision: All commands described in `README.md` remain planned product surface, and implementation may deliver them in phased subsets.
- Consequences: The roadmap must preserve the full planned surface while breaking it into smaller slices.
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`

## D-005: Data handling and sharing posture

- Date: 2026-05-06
- Status: Accepted
- Context: Crash diagnostics can include sensitive material, so the handling posture must be explicit before implementation.
- Decision: The product is local-only in scope, and unredacted output is allowed only through explicit operator action.
- Consequences: Implemented flows must default toward redaction and must not assume remote sharing or implicit unredacted export.
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`

## D-006: Environment scope

- Date: 2026-05-06
- Status: Accepted
- Context: The product spans multiple Linux collection environments, but not all environment claims are equally mandatory.
- Decision: Systemd support is required, container-aware behavior is optional, and the implementation must not assume elevated privileges are available.
- Consequences: Systemd-related slices remain in roadmap scope, container-aware work can be deferred, and collectors must degrade gracefully when permissions are insufficient.
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`

## D-007: Retention and prune defaults

- Date: 2026-05-06
- Status: Accepted
- Context: Crash artifacts are evidence, so retention defaults must avoid silent deletion while still documenting recommended limits.
- Decision: `crashsnap` ships with documented recommended retention values of `14` days, `10GiB`, and `50` snapshots, but no active automatic pruning until the operator configures retention explicitly.
- Consequences: `prune` must remain explicit, should support `--dry-run`, and implemented config behavior must distinguish recommendations from active deletion policy.
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`, `REQUIREMENTS.md`

## D-008: First implementation slice is a gated bootstrap

- Date: 2026-05-06
- Status: Accepted
- Context: The handoff package is implementation-ready, but the repository still lacks project-specific code scaffolding, working build or test commands, and a derived roadmap.
- Decision: First-run ends by authorizing slice `001-repo-bootstrap` as the first implementation session, with actual execution gated on human approval for the initial toolchain, dependency baseline, and project structure.
- Consequences: No product feature work should begin until `HG-001` is approved, and later slices should build on the resulting scaffold instead of re-deciding repository foundations.
- Source: `AGENTS.md`, `FIRST_RUN.md`, `docs/IMPLEMENTATION_PLAN.md`, repository file inventory

## D-009: Initial implementation language and bootstrap toolchain

- Date: 2026-05-06
- Status: Accepted
- Context: Slice `001-repo-bootstrap` requires an approved implementation language, packaging approach, and baseline toolchain before code scaffolding begins.
- Decision: The initial implementation stream will use Go with a single-binary distribution target. The bootstrap slice should prefer the Go standard library where practical and add third-party dependencies only when a roadmap slice requires them.
- Consequences: The bootstrap session should create a Go module, a minimal `crashsnap` executable entrypoint, and a standard Go test command. CLI, config, and output dependencies should remain intentionally narrow until their slices begin.
- Source: human approval on 2026-05-06, `SESSION.md`, `ROADMAP.md`

## D-010: Bootstrap repository layout

- Date: 2026-05-06
- Status: Accepted
- Context: The first implementation slice needed a concrete layout that keeps command wiring thin and reusable logic testable without overcommitting to later architecture.
- Decision: Use a standard Go module rooted at `crashsnap`, place the executable entrypoint in `cmd/crashsnap`, and keep CLI orchestration in `internal/cli` until later slices justify broader package structure.
- Consequences: Early implementation stays close to Go conventions, `go build ./cmd/crashsnap` and `go test ./...` become the baseline workflow, and later domain packages can be added without moving the executable entrypoint.
- Source: `SESSION.md`, `SKILLS/core-implementation.md`, slice `001-repo-bootstrap`

## D-011: Draft snapshot layout boundary

- Date: 2026-05-06
- Status: Accepted
- Context: Slice `002a-snapshot-layout` needed a concrete internal directory shape for snapshots without prematurely committing to manifest fields or a public compatibility contract.
- Decision: The draft snapshot layout consists of a root directory containing `metadata/`, `raw/`, `normalized/`, `reports/`, and `artifacts/`. This boundary is internal-only draft contract and does not yet imply manifest file names or public schema promises.
- Consequences: Later slices can build manifest semantics and collectors against a stable internal directory shape while `HG-002` remains unapproved for public machine-readable compatibility claims.
- Source: `ROADMAP.md`, `SESSION.md`, `README.md`
