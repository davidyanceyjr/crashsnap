# IMPLEMENTATION_PLAN.md

## Overview

`crashsnap` is currently defined as a Linux command-line crash-diagnostics utility that captures evidence from launched commands, live processes, systemd services, optional container-aware environments, and core dumps, then normalizes, redacts, analyzes, renders, bundles, and verifies that evidence. The implementation must preserve the product's evidence-first model: snapshots are the durable source of truth, and reports are derived views over those snapshots.

This plan is intentionally constrained to the behavior already claimed in `README.md`. The full command set is planned product scope, but implementation may deliver it in subsets. Retention policy is now decided: documented recommendations exist, while automatic pruning remains opt-in until the operator configures retention explicitly.

## Product goals

- Capture repeatable debugging evidence for Linux crashes and incident triage.
- Support both live-target collection and offline report regeneration from saved snapshots.
- Make partial collection useful by preserving successful evidence even when some collectors fail.
- Default toward safer handling of sensitive material through redaction and private artifact generation.
- Keep the product local-only in scope while permitting explicit local unredacted handling.
- Provide a stable machine-readable surface through schema-versioned outputs.

## Safe workflow or operating model

- User selects a capture mode based on target type: launched program, live PID, core dump, or systemd unit.
- The tool collects raw evidence into a snapshot directory.
- Raw evidence is normalized into a structured internal representation.
- Redaction is applied before export unless the operator explicitly opts into a more sensitive path.
- The tool operates with the permissions the current operator actually has; it must report blocked collectors instead of assuming elevated access.
- Analysis produces findings, confidence, symbol quality, and missing-evidence warnings.
- Reports are rendered from the snapshot in human-readable or machine-readable formats.
- Optional bundle and verify flows support handoff and integrity checking.
- Offline `report`, `inspect`, `diff`, `bundle`, and `verify` flows must not require the original live target to still exist.

## Core design rules

- Snapshots are the primary durable artifact; reports are regenerated views.
- Partial collector failure must not discard successful evidence.
- Debugger-backed attachment is explicit because it may stop or perturb the target.
- Non-invasive collection should exist as a separate workflow for latency-sensitive targets.
- Sensitive sections are not universally included by default.
- Redaction is enabled by default.
- Unredacted handling requires explicit operator action.
- Machine-readable output is schema-versioned.
- Systemd support is mandatory.
- Container-aware behavior is optional rather than mandatory.
- The implementation must not assume root, `CAP_SYS_PTRACE`, or equivalent elevated access.
- The implementation must expose missing data, failed collectors, and confidence limitations instead of implying certainty.

## Public surface

Current claimed command surface from `README.md`:

- `run`
- `attach`
- `core`
- `service`
- `observe`
- `report`
- `inspect`
- `diff`
- `bundle`
- `verify`
- `doctor`
- `config`
- `schema`
- `prune`

Current claimed key output surfaces:

- snapshot directories
- rendered reports in `text`, `json`, `yaml`, `md`, and `html`
- manifests with hashes and collection metadata
- optional signed bundles

Current claimed environment integrations:

- `gdb`
- `procfs`
- `journalctl`
- `dmesg`
- `coredumpctl`
- systemd units
- cgroups and namespaces
- debug symbol sources including debuginfod

## Data or state model

Important product-level data domains implied by `README.md`:

- snapshot
  - raw collected artifacts
  - normalized diagnostic data
  - collector success and failure metadata
  - redaction summary
  - schema version references
  - integrity metadata such as hashes
- report
  - derived rendering of snapshot data
  - format-specific output
  - findings and confidence
- bundle
  - packaged snapshot and report artifacts
  - manifest
  - optional signature material
- configuration
  - profiles
  - report defaults
  - redaction settings
  - path and timeout settings

Important state transitions:

- capture -> normalize -> redact -> analyze -> render -> bundle -> verify
- live collection may stop after capture, while later stages can be repeated offline from the snapshot

## Recommended implementation sequence

Because implementation is allowed in subsets, this sequence should be treated as a phased scaffold, not a single-release checklist.

1. Confirm the retention policy for local artifacts and `prune` defaults.
2. Define the snapshot and manifest contracts that all capture modes must converge on.
3. Implement one narrow capture path plus offline report generation to validate the evidence-first model.
4. Add redaction, findings, schema-versioned machine-readable output, and graceful permission-failure reporting.
5. Expand to additional command subsets until the full planned README surface is covered.
6. Add bundle, verify, and opt-in prune behavior with explicit operator-configured retention.

## Open implementation-sensitive areas

- Automatic prune behavior must remain opt-in until the operator configures retention explicitly.
- Recommended retention values must not be mistaken for already-active deletion policy.
