# ROADMAP.md

## Purpose

This roadmap converts the implementation-ready handoff package into small, reviewable implementation slices for `crashsnap`.

It is derived primarily from `docs/IMPLEMENTATION_PLAN.md` and constrained by the canonical product definition in `README.md`, the implementation handoff, recorded decisions, active risks, and any open questions.

## Source documents

- `docs/IMPLEMENTATION_PLAN.md`
- `docs/IMPLEMENTATION_HANDOFF.md`
- `README.md`
- `DECISIONS.md`
- `RISKS.md`
- `docs/OPEN_QUESTIONS.md`

## Roadmap status

- Status: first-run-complete
- Canonical product document: `README.md`
- Current implementation slice: `002a-snapshot-layout`
- Package readiness: implementation-ready
- Blocking questions affecting roadmap derivation: none currently recorded

## Milestones

## Milestone 1: Bootstrap And Contracts

### Goal

Establish a project-specific implementation skeleton, then define the foundational snapshot, manifest, and configuration contracts that later capture modes will use.

### Source basis

- `docs/IMPLEMENTATION_PLAN.md` recommended sequence items 2 and 3
- `README.md` sections covering snapshots, manifests, configuration, and schema-versioned outputs
- `DECISIONS.md` entries for command scope, environment scope, and retention posture

### Exit criteria

- A project-specific code and test scaffold exists.
- Snapshot and manifest structures are defined tightly enough to support one narrow capture flow.
- Configuration loading behavior is defined for the initial implementation path.

### Slices

#### 001-repo-bootstrap: Project bootstrap and command scaffold

- Objective: Choose and wire the implementation toolchain, package layout, baseline test command, and a minimal `crashsnap` executable scaffold without implementing capture behavior.
- Source basis: `docs/IMPLEMENTATION_PLAN.md`, `README.md` synopsis and command surface, `docs/IMPLEMENTATION_OPERATING_MODEL.md`
- Depends on: none
- Expected files or areas: repo root build metadata, source package layout, test harness, minimal CLI entrypoint, example config cleanup
- Tests or validation: toolchain install/build smoke check, minimal CLI help/version smoke test, baseline automated test command wired and run
- Human gates: `HG-001`
- Blocking questions: none
- Out of scope: capture collectors, report rendering, schema commitments, destructive commands

#### 002a-snapshot-layout: Draft snapshot directory layout

- Objective: Define the initial internal snapshot directory structure and fixture shape without introducing manifest semantics yet.
- Source basis: `docs/IMPLEMENTATION_PLAN.md` data model, `README.md` evidence-first sections
- Depends on: `001-repo-bootstrap`
- Expected files or areas: snapshot layout definitions, directory fixtures, internal-draft contract notes
- Tests or validation: fixture layout tests, path-construction tests, rejection tests for invalid snapshot roots
- Human gates: `HG-002`
- Blocking questions: none
- Out of scope: manifest field design, collector status modeling, live capture logic, bundle signing

#### 002b-manifest-core-fields: Draft manifest base model

- Objective: Define the initial manifest model for snapshot identity, schema references, timestamps, and provenance fields.
- Source basis: `docs/IMPLEMENTATION_PLAN.md` data model, `README.md` auditable handling sections, `REQUIREMENTS.md` FR-005 and FR-006
- Depends on: `002a-snapshot-layout`
- Expected files or areas: manifest model definitions, serialization helpers, base-manifest fixtures
- Tests or validation: round-trip serialization tests, required-field validation tests, fixture-based manifest assertions
- Human gates: `HG-002`
- Blocking questions: none
- Out of scope: collector status modeling, report rendering, stable public compatibility promises

#### 002c-manifest-collector-status: Collector outcome metadata

- Objective: Extend the draft manifest with explicit collector success, failure, and missing-evidence metadata needed for partial-success behavior.
- Source basis: `docs/IMPLEMENTATION_PLAN.md` core design rules, `README.md` partial success principle, `REQUIREMENTS.md` FR-004
- Depends on: `002b-manifest-core-fields`
- Expected files or areas: collector status model, failure classification scaffolding, manifest fixtures with mixed outcomes
- Tests or validation: partial-failure metadata assertions, success-and-failure fixture tests, invalid-status rejection tests
- Human gates: `HG-002`
- Blocking questions: none
- Out of scope: live collectors, report rendering, bundle verification

#### 003a-config-model: Configuration model foundation

- Objective: Define the initial configuration model and supported keys needed for the first implemented flows, including retention-related settings as data only.
- Source basis: `README.md` configuration keys and locations, `DECISIONS.md` retention decision, `REQUIREMENTS.md` OR-007 and OR-008
- Depends on: `001-repo-bootstrap`
- Expected files or areas: config model definitions, example config alignment, config fixtures
- Tests or validation: valid-config parsing tests, unsupported-key rejection tests, fixture coverage for omitted optional fields
- Human gates: `HG-001`
- Blocking questions: none
- Out of scope: filesystem discovery, config mutation subcommands, active pruning behavior

#### 003b-config-discovery: Configuration discovery and load path

- Objective: Implement configuration file discovery and load precedence without adding mutation helpers.
- Source basis: `README.md` configuration locations, `003a-config-model`
- Depends on: `003a-config-model`
- Expected files or areas: config discovery helpers, precedence rules, temporary-directory fixtures
- Tests or validation: path precedence tests, missing-file behavior tests, load-path command or helper tests
- Human gates: `HG-001`
- Blocking questions: none
- Out of scope: config mutation, retention enforcement, profile expansion

#### 003c-config-retention-validation: Retention-specific validation rules

- Objective: Add validation rules that keep documented retention recommendations separate from active automatic prune policy.
- Source basis: `DECISIONS.md` retention decision, `REQUIREMENTS.md` OR-007 and OR-008
- Depends on: `003a-config-model`, `003b-config-discovery`
- Expected files or areas: retention validation rules, related fixtures, validation error messaging
- Tests or validation: invalid-retention rejection tests, inactive-default assertions, explicit-config acceptance tests
- Human gates: `HG-001`
- Blocking questions: none
- Out of scope: prune command behavior, automatic deletion, config mutation subcommands

## Milestone 2: First Vertical Evidence Flow

### Goal

Prove the evidence-first model end to end with one narrow capture path and offline report generation from saved snapshots.

### Source basis

- `docs/IMPLEMENTATION_PLAN.md` recommended sequence item 3
- `README.md` observe mode, offline report mode, evidence pipeline, and design principles
- `REQUIREMENTS.md` FR-001, FR-003, FR-004, and UR-003

### Exit criteria

- One live collection path can write a snapshot using the defined contract.
- A saved snapshot can be rendered offline without the original target.
- Partial collector failures remain visible and do not discard successful evidence.

### Slices

#### 004a-observe-command-bootstrap: Observe command wiring and snapshot bootstrap

- Objective: Implement the `observe` command path up to PID validation and empty snapshot bootstrap, without collecting procfs evidence yet.
- Source basis: `README.md` observe mode, `DECISIONS.md` environment scope decision, `REQUIREMENTS.md` UR-003
- Depends on: `002a-snapshot-layout`, `002b-manifest-core-fields`, `003b-config-discovery`
- Expected files or areas: observe command wiring, PID validation, snapshot bootstrap, fixture process helpers
- Tests or validation: invalid-PID tests, snapshot-bootstrap assertions, command wiring smoke tests
- Human gates: none
- Blocking questions: none
- Out of scope: procfs collection, report rendering, GDB attach, systemd lookup

#### 004b-observe-procfs-collectors: Minimal non-invasive evidence collection

- Objective: Add the first low-risk procfs-backed collectors to the `observe` path and write their results into the draft snapshot contract.
- Source basis: `README.md` observe mode and low-impact principles, `REQUIREMENTS.md` UR-003
- Depends on: `004a-observe-command-bootstrap`, `002c-manifest-collector-status`
- Expected files or areas: procfs collectors, snapshot population, controlled fixture process helpers
- Tests or validation: integration test against a controlled fixture process, permission-failure coverage, collector-status assertions
- Human gates: none
- Blocking questions: none
- Out of scope: GDB attach, service lookup, container metadata, report formatting

#### 005a-report-text-render: Offline text report regeneration

- Objective: Implement `report` for saved snapshots with initial text rendering only.
- Source basis: `README.md` offline report mode and report formats, `docs/IMPLEMENTATION_PLAN.md` evidence-first model, `REQUIREMENTS.md` FR-003 and FR-005
- Depends on: `004b-observe-procfs-collectors`
- Expected files or areas: report command wiring, snapshot reader, text renderer, output fixtures
- Tests or validation: offline text report generation from fixture snapshots, snapshot-to-report golden tests, missing-field handling tests
- Human gates: none
- Blocking questions: none
- Out of scope: JSON output, HTML, Markdown, YAML, diffing, bundle creation

#### 005b-report-json-render: Offline JSON report regeneration

- Objective: Add initial JSON report rendering on top of the text-report path while keeping the surface explicitly draft until `HG-002` is approved.
- Source basis: `README.md` report formats, `REQUIREMENTS.md` FR-006
- Depends on: `005a-report-text-render`
- Expected files or areas: JSON renderer, version constants, JSON fixtures
- Tests or validation: JSON output tests, schema-version presence assertions, fixture-based serialization tests
- Human gates: `HG-002`
- Blocking questions: none
- Out of scope: public compatibility guarantees, schema command, other report formats

#### 006-partial-failure-reporting: Collector failure visibility

- Objective: Ensure partial collection failures are preserved in snapshots and surfaced clearly in reports and exit behavior.
- Source basis: `README.md` partial success principle, `docs/IMPLEMENTATION_PLAN.md` core design rules, `REQUIREMENTS.md` FR-004 and OR-002
- Depends on: `004b-observe-procfs-collectors`, `005b-report-json-render`, `002c-manifest-collector-status`
- Expected files or areas: manifest status model, reporting logic, failure classification, exit-code handling tests
- Tests or validation: injected collector-failure tests, assertions that successful artifacts remain present, report warnings and limitations tests
- Human gates: none
- Blocking questions: none
- Out of scope: findings analysis, symbol scoring, debugger-backed flows

## Milestone 3: Safety, Analysis, And Machine-Readable Stability

### Goal

Add the first safety-critical behavior: redaction defaults, machine-readable schema surfacing, and permission-preflight diagnostics.

### Source basis

- `docs/IMPLEMENTATION_PLAN.md` recommended sequence item 4
- `README.md` redaction, schema, doctor, and safety-by-default sections
- `DECISIONS.md` local-only and explicit-unredacted posture

### Exit criteria

- Redaction is enabled by default for initial implemented flows.
- JSON report output carries an explicit schema version and discoverable schema metadata.
- Operators can run a diagnostic preflight command for common environment blockers.

### Slices

#### 007a-redaction-rule-engine: Initial redaction rule set

- Objective: Implement the first redaction rule set for implemented evidence fields without changing report output wiring yet.
- Source basis: `README.md` redaction defaults, `DECISIONS.md` data-handling decision, `REQUIREMENTS.md` OR-001 and OR-006
- Depends on: `005b-report-json-render`, `006-partial-failure-reporting`
- Expected files or areas: redaction rules, sensitive-field fixtures, rule evaluation helpers
- Tests or validation: secret fixture redaction tests, strict-versus-default rule tests, rule failure-path coverage
- Human gates: none
- Blocking questions: none
- Out of scope: report integration, config-file discovery, journal excerpts, signed bundles

#### 007b-redaction-report-integration: Redaction summaries and output integration

- Objective: Integrate the initial redaction rules into implemented report flows and record redaction summaries in the draft manifest.
- Source basis: `README.md` redaction defaults and auditable handling, `REQUIREMENTS.md` OR-006
- Depends on: `007a-redaction-rule-engine`
- Expected files or areas: report rendering integration, manifest redaction summary fields, redacted output fixtures
- Tests or validation: explicit-unredacted opt-in tests, manifest redaction summary assertions, redacted report golden tests
- Human gates: none
- Blocking questions: none
- Out of scope: broad config-file discovery, journal excerpts, signed bundles

#### 008-schema-surface: Schema command and versioned JSON surface

- Objective: Expose the initial schema/version surface needed for automation without promising more compatibility than the implementation can support.
- Source basis: `README.md` schema command and stable automation surface, `REQUIREMENTS.md` FR-006, `docs/IMPLEMENTATION_PLAN.md` machine-readable output rule
- Depends on: `002b-manifest-core-fields`, `002c-manifest-collector-status`, `005b-report-json-render`
- Expected files or areas: schema command wiring, version constants, exported schema artifacts or descriptors, compatibility documentation
- Tests or validation: schema command snapshot tests, JSON output schema-version assertions, fixture validation against exported schema surface
- Human gates: `HG-002`
- Blocking questions: none
- Out of scope: broad compatibility guarantees beyond the documented initial contract

#### 009a-doctor-tooling-checks: Tooling and path diagnostics

- Objective: Implement the first `doctor` checks for required tools, writable snapshot paths, and basic environment prerequisites.
- Source basis: `README.md` doctor command examples, `REQUIREMENTS.md` OR-002
- Depends on: `001-repo-bootstrap`, `003b-config-discovery`
- Expected files or areas: doctor command, tooling probes, path probes, deterministic output formatting
- Tests or validation: probe-unit tests, missing-tool tests, unwritable-path tests
- Human gates: none
- Blocking questions: none
- Out of scope: ptrace policy checks, privilege escalation, environment mutation

#### 009b-doctor-permission-checks: Ptrace and visibility diagnostics

- Objective: Add permission-oriented `doctor` checks such as ptrace restrictions and procfs visibility limits after the tooling checks exist.
- Source basis: `README.md` doctor command examples, `DECISIONS.md` environment scope decision, `REQUIREMENTS.md` OR-003 through OR-005
- Depends on: `009a-doctor-tooling-checks`
- Expected files or areas: permission probes, visibility checks, mocked probe tests, deterministic command output
- Tests or validation: permission-probe tests, restricted-visibility tests, manual smoke validation on the current host
- Human gates: none
- Blocking questions: none
- Out of scope: making environment changes automatically, privilege escalation, deployment automation

## Milestone 4: Expanded Capture Modes

### Goal

Add the major capture paths described in `README.md` without collapsing them into one oversized delivery.

### Source basis

- `docs/IMPLEMENTATION_PLAN.md` recommended sequence item 5
- `README.md` run, core, service, and attach mode sections
- `DECISIONS.md` systemd-required and no-privilege-assumption decisions

### Exit criteria

- At least one launched-program path exists.
- Core-dump and systemd service collection are available in scoped initial form.
- Debugger-backed attach behavior is explicit, gated, and tested separately from non-invasive flows.

### Slices

#### 010a-run-command-execution: Managed command launch

- Objective: Implement the `run` command path up to managed process launch, exit capture, and empty snapshot bootstrap.
- Source basis: `README.md` run mode and quick start, `REQUIREMENTS.md` FR-001
- Depends on: `004a-observe-command-bootstrap`, `003b-config-discovery`
- Expected files or areas: run command wiring, process execution wrapper, exit capture, snapshot naming
- Tests or validation: fixture-program execution tests, exit-status capture tests, snapshot-bootstrap assertions
- Human gates: none
- Blocking questions: none
- Out of scope: failure-trigger policy, procfs evidence collection, bundle creation

#### 010b-run-triggered-snapshot: Failure-trigger capture behavior

- Objective: Add failure-trigger conditions and minimal evidence capture to the managed run path.
- Source basis: `README.md` run mode, `REQUIREMENTS.md` FR-001
- Depends on: `010a-run-command-execution`, `004b-observe-procfs-collectors`, `007b-redaction-report-integration`
- Expected files or areas: trigger logic, run-specific snapshot population, failing-fixture-program helpers
- Tests or validation: on-fail versus on-exit behavior tests, failing-fixture integration tests, collector-status assertions
- Human gates: none
- Blocking questions: none
- Out of scope: debugger-backed launched execution, service-specific behavior, bundle creation

#### 011-core-input-validation: Core mode input and snapshot bootstrap

- Objective: Implement `core` command wiring, executable/core input validation, and snapshot bootstrap for existing core-file analysis without adding debugger-backed extraction yet.
- Source basis: `README.md` core mode, `docs/IMPLEMENTATION_PLAN.md` core design rules
- Depends on: `002a-snapshot-layout`, `002b-manifest-core-fields`, `005a-report-text-render`
- Expected files or areas: core command wiring, input validation, snapshot bootstrap, mismatch handling tests
- Tests or validation: invalid-path tests, executable/core mismatch tests, snapshot-bootstrap assertions for accepted inputs
- Human gates: none
- Blocking questions: none
- Out of scope: debugger-backed extraction, symbol loading, service discovery, symbol server automation

#### 012-core-collector-foundation: Core evidence extraction and rendering

- Objective: Add the minimal collector integration needed for `core` snapshots to produce an initial offline report from accepted inputs.
- Source basis: `README.md` core mode, `docs/IMPLEMENTATION_PLAN.md` core design rules
- Depends on: `011-core-input-validation`, `007b-redaction-report-integration`, `002c-manifest-collector-status`
- Expected files or areas: collector adapter boundary, core-derived snapshot population, fixture or sampled-core harness
- Tests or validation: fixture or sampled-core tests where feasible, collector-failure coverage, offline report assertions for core snapshots
- Human gates: none
- Blocking questions: none
- Out of scope: live attach, systemd service discovery, deep symbol-source automation

#### 013-service-unit-resolution: Systemd unit targeting

- Objective: Implement the `service` command path up to unit resolution, systemd metadata lookup, and snapshot bootstrap without full service-context capture.
- Source basis: `README.md` service mode, `DECISIONS.md` systemd-required decision, `REQUIREMENTS.md` OR-004
- Depends on: `004a-observe-command-bootstrap`, `010a-run-command-execution`, `002b-manifest-core-fields`
- Expected files or areas: service command wiring, systemctl adapter layer, unit lookup, snapshot bootstrap, tests with fakes
- Tests or validation: unit-resolution tests, missing-unit tests, mocked systemctl command tests
- Human gates: none
- Blocking questions: none
- Out of scope: journal capture, coredumpctl deep integration, attach semantics, container discovery

#### 014-service-context-capture: Scoped systemd service evidence

- Objective: Add the first scoped service-context collectors, such as recent unit state and limited journal or cgroup context, to the bootstrap created by the prior slice.
- Source basis: `README.md` service mode, `DECISIONS.md` systemd-required decision
- Depends on: `013-service-unit-resolution`, `007b-redaction-report-integration`, `002c-manifest-collector-status`
- Expected files or areas: journalctl or cgroup adapter layer, service-context collectors, service fixtures or fakes
- Tests or validation: mocked collector tests, redacted context assertions, partial-failure coverage for unavailable systemd data
- Human gates: none
- Blocking questions: none
- Out of scope: latest-core resolution, container discovery, debugger attach semantics

#### 015-attach-mode-explicit-debugger: Debugger-backed live attach

- Objective: Implement an explicit `attach` path whose UX and behavior make target interruption risks and permission limits impossible to miss.
- Source basis: `README.md` attach mode warnings, `docs/IMPLEMENTATION_PLAN.md` explicit debugger rule, `REQUIREMENTS.md` FR-002 and OR-005
- Depends on: `004b-observe-procfs-collectors`, `009b-doctor-permission-checks`
- Expected files or areas: attach command wiring, debugger adapter boundary, risk prompts or flags, permission handling, gated manual-test notes
- Tests or validation: adapter-unit tests, permission-failure tests, manual validation against a disposable process only
- Human gates: `HG-003`
- Blocking questions: none
- Out of scope: stealth attach behavior, implicit ptrace assumptions, production-target automation

## Milestone 5: Operational Workflows And Deferred Surface

### Goal

Round out the planned command surface with packaging, verification, configuration management, cleanup workflows, and comparison tools.

### Source basis

- `docs/IMPLEMENTATION_PLAN.md` recommended sequence item 6
- `README.md` bundle, verify, config, prune, inspect, and diff sections
- `DECISIONS.md` retention decision

### Exit criteria

- Bundle and verify operate on implemented snapshot formats.
- Config subcommands are usable for the implemented configuration model.
- Prune preserves the explicit opt-in retention posture.
- Inspect and diff cover the implemented report or snapshot formats.

### Slices

#### 016a-bundle-create: Unsigned bundle creation

- Objective: Implement unsigned bundle creation for the implemented snapshot artifacts without verification yet.
- Source basis: `README.md` bundle command, `docs/IMPLEMENTATION_PLAN.md` bundle flow
- Depends on: `005b-report-json-render`, `007b-redaction-report-integration`, `008-schema-surface`
- Expected files or areas: bundle writer, archive format helpers, fixture archives
- Tests or validation: bundle creation tests, fixture archive assertions, redacted-artifact inclusion tests
- Human gates: `HG-002`, `HG-005`
- Blocking questions: none
- Out of scope: verification, signing, external publication

#### 016b-verify-bundle: Bundle integrity verification

- Objective: Implement verification for unsigned bundles created by the prior slice, including tamper detection and manifest hash checks.
- Source basis: `README.md` verify command, `docs/IMPLEMENTATION_PLAN.md` verify flow
- Depends on: `016a-bundle-create`
- Expected files or areas: verify command, manifest hashing checks, tampered fixture archives
- Tests or validation: bundle round-trip tests, manifest hash verification tests, tamper-detection tests
- Human gates: `HG-002`, `HG-005`
- Blocking questions: none
- Out of scope: signing, remote trust distribution, external publication

#### 017-config-read-validate: Config read and validate subcommands

- Objective: Implement the lowest-risk `config` subcommands for reading discovered configuration and validating it against the supported model.
- Source basis: `README.md` config command section and config path rules, `003a-config-model`, `003b-config-discovery`
- Depends on: `003a-config-model`, `003b-config-discovery`, `003c-config-retention-validation`
- Expected files or areas: `config list`, `config get`, and `config validate` handlers, file IO helpers, config fixtures
- Tests or validation: path precedence tests, valid and invalid config tests, temporary-directory command tests
- Human gates: `HG-001`
- Blocking questions: none
- Out of scope: config mutation, file creation helpers, unsupported keys, remote config distribution

#### 018-config-init-set: Config bootstrap and mutation helpers

- Objective: Implement `config init` and `config set` only after the read and validation path is stable.
- Source basis: `README.md` config command section, `003a-config-model`, `003b-config-discovery`
- Depends on: `017-config-read-validate`
- Expected files or areas: config file creation helpers, mutation handlers, config rewrite logic, temporary-directory fixtures
- Tests or validation: init-file creation tests, set-command update tests, preservation tests for unchanged settings
- Human gates: `HG-001`
- Blocking questions: none
- Out of scope: unsupported keys, policy mutation outside implemented settings, remote config distribution

#### 019a-prune-dry-run-manual: Manual prune planning and dry-run

- Objective: Implement manual prune selection and `--dry-run` reporting without deleting artifacts.
- Source basis: `README.md` prune section, `DECISIONS.md` retention decision, `REQUIREMENTS.md` OR-007 and OR-008
- Depends on: `003c-config-retention-validation`, `016b-verify-bundle`
- Expected files or areas: prune candidate selection, dry-run reporting, retention evaluator, fixture snapshots
- Tests or validation: dry-run tests, candidate-selection tests, assertions that unconfigured installs never prune automatically
- Human gates: `HG-004`
- Blocking questions: none
- Out of scope: non-dry-run deletion, configured automatic pruning, artifact recovery

#### 019b-prune-configured-automatic: Configured deletion paths

- Objective: Add non-dry-run deletion and configured automatic prune behavior only after the dry-run logic is already proven.
- Source basis: `README.md` prune section, `DECISIONS.md` retention decision
- Depends on: `019a-prune-dry-run-manual`
- Expected files or areas: destructive prune path, configured automatic prune integration, safeguard prompts or flags
- Tests or validation: configured-retention enforcement tests, deletion-safeguard tests, assertions that unconfigured installs still never prune automatically
- Human gates: `HG-004`
- Blocking questions: none
- Out of scope: silent background deletion outside explicit configuration, remote artifact lifecycle management

#### 020-inspect-snapshot-report: Offline introspection

- Objective: Implement scoped `inspect` support for the implemented snapshot or report surfaces.
- Source basis: `README.md` inspect command section, `docs/IMPLEMENTATION_PLAN.md` offline analysis goals
- Depends on: `005a-report-text-render`, `005b-report-json-render`, `008-schema-surface`
- Expected files or areas: inspect command, snapshot or report introspection output, fixture reports or snapshots
- Tests or validation: inspect output tests, fixture coverage for missing-field handling, schema-aware output assertions
- Human gates: none
- Blocking questions: none
- Out of scope: diffing, bundle-aware introspection, unsupported report formats

#### 021-diff-snapshot-report: Offline comparison

- Objective: Implement scoped `diff` support for the implemented snapshot or report surfaces after inspection behavior is already established.
- Source basis: `README.md` diff command section, `docs/IMPLEMENTATION_PLAN.md` offline analysis goals
- Depends on: `020-inspect-snapshot-report`
- Expected files or areas: diff engine, fixture reports or snapshots, comparison formatting
- Tests or validation: snapshot diff fixture tests, schema-aware comparison coverage, changed-versus-unchanged output tests
- Human gates: none
- Blocking questions: none
- Out of scope: bundle-aware diffing, unsupported report formats, cross-version compatibility beyond the documented initial schema rules

## Slice backlog

1. `001-repo-bootstrap`
2. `002a-snapshot-layout`
3. `002b-manifest-core-fields`
4. `002c-manifest-collector-status`
5. `003a-config-model`
6. `003b-config-discovery`
7. `003c-config-retention-validation`
8. `004a-observe-command-bootstrap`
9. `004b-observe-procfs-collectors`
10. `005a-report-text-render`
11. `005b-report-json-render`
12. `006-partial-failure-reporting`
13. `007a-redaction-rule-engine`
14. `007b-redaction-report-integration`
15. `008-schema-surface`
16. `009a-doctor-tooling-checks`
17. `009b-doctor-permission-checks`
18. `010a-run-command-execution`
19. `010b-run-triggered-snapshot`
20. `011-core-input-validation`
21. `012-core-collector-foundation`
22. `013-service-unit-resolution`
23. `014-service-context-capture`
24. `015-attach-mode-explicit-debugger`
25. `016a-bundle-create`
26. `016b-verify-bundle`
27. `017-config-read-validate`
28. `018-config-init-set`
29. `019a-prune-dry-run-manual`
30. `019b-prune-configured-automatic`
31. `020-inspect-snapshot-report`
32. `021-diff-snapshot-report`

## Blocked or deferred slices

- `015-attach-mode-explicit-debugger` is gated by `HG-003` because it introduces debugger-backed target interruption and security-sensitive defaults.
- `016a-bundle-create` and `016b-verify-bundle` may implement unsigned archives first; bundle signing remains deferred behind `HG-005`.
- `019a-prune-dry-run-manual` and `019b-prune-configured-automatic` must not enable active automatic deletion without the approval recorded in `HG-004`.
- Container-aware collection remains deferred until after milestone 4 because `DECISIONS.md` records it as optional rather than mandatory scope.

## Human gates

- `HG-001`: approve the initial implementation toolchain, dependency baseline, and project structure before slice `001-repo-bootstrap`.
- `HG-002`: approve the first public-facing snapshot/report schema compatibility boundary before slices `002a`, `002b`, `002c`, `005b`, `008`, `016a`, or `016b` claim stable machine-readable contracts.
- `HG-003`: approve debugger-backed attach UX and safety defaults before slice `015-attach-mode-explicit-debugger`.
- `HG-004`: approve destructive prune behavior before slice `019a-prune-dry-run-manual` or `019b-prune-configured-automatic` enables non-dry-run deletion or configured automatic pruning.
- `HG-005`: approve signing and compliance posture before slices `016a-bundle-create` or `016b-verify-bundle` add signed bundle support.

## Roadmap maintenance rules

- Keep slice order derived from `docs/IMPLEMENTATION_PLAN.md` unless a narrower prerequisite is required and recorded in `DECISIONS.md`.
- Do not add product behavior not already supported by `README.md` and the implementation handoff package.
- Keep each slice small enough for one focused implementation session with explicit validation.
- Record blocked, gated, or deferred work here rather than hiding it inside session notes.
