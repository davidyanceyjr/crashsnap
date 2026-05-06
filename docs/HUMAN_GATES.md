# HUMAN_GATES.md

## Purpose

This document records the approvals that must be obtained before implementation crosses safety-sensitive, destructive, or hard-to-reverse boundaries.

Roadmap slices must reference these gate IDs directly.

## Active gates

## HG-001: Initial toolchain, dependency baseline, and project structure

- Status: approved on 2026-05-06
- Area: repository foundations
- Why gate: slice `001-repo-bootstrap` will choose the initial implementation toolchain, add baseline dependencies, and establish a project structure that is expensive to reverse later.
- Affected slices: `001-repo-bootstrap`, `003a-config-model`, `003b-config-discovery`, `003c-config-retention-validation`, `017-config-read-validate`, `018-config-init-set`
- Approval needed: approve the implementation language, packaging approach, build command, test command, and initial top-level source layout
- Approval recorded: Go with a single-binary distribution target; prefer the Go standard library initially and keep third-party dependencies narrow during bootstrap
- Safe work allowed before approval: documentation-only planning and roadmap maintenance

## HG-002: Public machine-readable contract boundary

- Status: required before public compatibility claims
- Area: schemas and automation surfaces
- Why gate: snapshot manifests, report JSON, and verification metadata can become public compatibility commitments if they are presented as stable too early.
- Affected slices: `002a-snapshot-layout`, `002b-manifest-core-fields`, `002c-manifest-collector-status`, `005b-report-json-render`, `008-schema-surface`, `016a-bundle-create`, `016b-verify-bundle`
- Approval needed: approve which machine-readable fields and schema versions are treated as externally stable versus explicitly provisional
- Safe work allowed before approval: internal draft contracts, fixtures, and tests that are marked non-stable

## HG-003: Debugger-backed attach behavior and safety defaults

- Status: required before attach implementation
- Area: live-target safety
- Why gate: debugger-backed attachment can stop or perturb a live process and involves security-sensitive defaults around ptrace behavior and operator warnings.
- Affected slices: `015-attach-mode-explicit-debugger`
- Approval needed: approve the attach UX, any required confirmation flags, and the default behavior when permissions or safety constraints are not met
- Safe work allowed before approval: non-invasive observe flows and attach adapter prototyping that is not exposed as user behavior

## HG-004: Destructive prune behavior

- Status: required before deletion paths
- Area: artifact deletion
- Why gate: non-dry-run prune operations and configured automatic retention can delete evidence irreversibly.
- Affected slices: `019a-prune-dry-run-manual`, `019b-prune-configured-automatic`
- Approval needed: approve the exact destructive semantics, safeguards, dry-run defaults, and the conditions under which configured automatic pruning may execute
- Safe work allowed before approval: dry-run-only design, retention evaluation logic, and non-destructive previews

## HG-005: Signed bundle and compliance posture

- Status: required before signing support
- Area: integrity and compliance
- Why gate: signed bundle support introduces cryptographic, trust-distribution, and potential compliance expectations that should not be guessed.
- Affected slices: `016a-bundle-create`, `016b-verify-bundle`
- Approval needed: approve whether signing is in scope, what trust model applies, and any compliance or key-management expectations
- Safe work allowed before approval: unsigned bundle creation and manifest hash verification
