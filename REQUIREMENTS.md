# REQUIREMENTS.md

## Functional requirements

- FR-001: The product must support collecting crash-diagnostic evidence for a launched program and storing that evidence in a durable snapshot artifact.
- FR-002: The product must support collecting evidence from at least one live-process workflow and make debugger-backed attachment explicit when it may stop the target.
- FR-003: The product must support offline report generation from a previously captured snapshot without requiring the original target to still be available.
- FR-004: The product must preserve partial success by recording successful collectors and failed collectors in the resulting snapshot or report.
- FR-005: The product must render reports from captured evidence rather than treating the rendered report as the only durable artifact.
- FR-006: The product must support machine-readable report output with an explicit schema version.
- FR-007: The product must expose findings, confidence, and evidence limitations when enough input data exists to do so.
- FR-008: The product must support the full command surface documented in `README.md`, even if implementation is delivered in subsets over time.

## User requirements

- UR-001: A Linux operator needs to capture actionable crash evidence from a failing command, process, service, or core dump without manually stitching together many low-level tools.
- UR-002: A developer or incident responder needs a reproducible artifact that can be handed off and re-rendered later.
- UR-003: An operator of latency-sensitive systems needs a non-invasive collection path when debugger attachment risk is unacceptable.

## Operational requirements

- OR-001: Sensitive material must not be included in exported output by default.
- OR-002: The product must make collection gaps, permission failures, and symbol-quality limits visible.
- OR-003: The product must behave as a Linux-only tool unless and until platform scope is explicitly expanded.
- OR-004: The product must support systemd-based workflows.
- OR-005: The product must not assume root, `CAP_SYS_PTRACE`, or equivalent elevated privileges are available.
- OR-006: Unredacted output must require explicit operator action and remain in local-only scope.
- OR-007: Automatic pruning must remain disabled until the operator explicitly configures retention.
- OR-008: The documented recommended retention values are `14` days, `10GiB`, and `50` snapshots.

## Constraints

- C-001: `README.md` is the canonical product definition for the current session.
- C-002: The implementation-facing docs must not invent capabilities beyond those already claimed in `README.md`.
C-003: Documented recommended retention values must not be treated as active automatic deletion policy until the operator explicitly configures retention.

## Undefined requirements

- None currently recorded.
