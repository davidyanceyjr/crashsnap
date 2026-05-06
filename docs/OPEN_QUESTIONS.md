# OPEN_QUESTIONS.md

## Purpose

This file stores unresolved implementation-visible questions and preserves the answer history that first-run derived from the upstream handoff package.

## Active questions

No blocking or non-blocking open questions are currently recorded.

## Answered / historical

## Q-001: Which document is canonical right now?

- Status: Answered
- Area: product definition
- Blocks: first-run setup
- Source: `README.md`, `docs/IMPLEMENTATION_HANDOFF.md`
- Why it matters: Implementation needs one canonical source of product truth.
- Safe temporary assumption: none
- Reversal cost: medium
- Needed answer: identify the canonical product document

Answer:

- `README.md` is the canonical product source document until a later decision supersedes it.

## Q-002: What is the public product and command name?

- Status: Answered
- Area: product identity
- Blocks: implementation naming
- Source: `README.md`
- Why it matters: The executable, docs, and implementation artifacts need one stable public name.
- Safe temporary assumption: none
- Reversal cost: high
- Needed answer: identify the public product and command name

Answer:

- The public product and command name is `crashsnap`.

## Q-003: What is the command-scope posture?

- Status: Answered
- Area: roadmap scope
- Blocks: phased implementation planning
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`
- Why it matters: The roadmap must know whether the full command list remains planned scope.
- Safe temporary assumption: none
- Reversal cost: high
- Needed answer: confirm whether all commands stay in planned scope

Answer:

- All commands currently documented in `README.md` remain planned product surface.
- Implementation may deliver them in phased subsets rather than one all-at-once release.

## Q-004: What is the data-handling and sharing posture?

- Status: Answered
- Area: privacy and export
- Blocks: redaction and bundle planning
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`
- Why it matters: Sensitive diagnostics handling changes defaults, export behavior, and later security expectations.
- Safe temporary assumption: none
- Reversal cost: high
- Needed answer: define local-only and unredacted handling expectations

Answer:

- The product is local-only in scope.
- Unredacted output is allowed only through explicit operator action.

## Q-005: Which environment scope decisions are resolved?

- Status: Answered
- Area: platform and collector scope
- Blocks: roadmap sequencing
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`
- Why it matters: Systemd, container support, and privilege assumptions materially change implementation scope.
- Safe temporary assumption: none
- Reversal cost: high
- Needed answer: define mandatory versus optional environment support

Answer:

- Systemd support is required.
- Container-aware behavior is optional, not mandatory.
- The implementation must not assume elevated privileges are available and must degrade gracefully when access is insufficient.

## Q-006: What retention policy should local artifacts follow?

- Status: Answered
- Area: retention and destructive behavior
- Blocks: prune planning and config defaults
- Source: `README.md`, `docs/IMPLEMENTATION_PLAN.md`, `REQUIREMENTS.md`
- Why it matters: Retention behavior can destroy evidence if defaults are guessed incorrectly.
- Safe temporary assumption: none
- Reversal cost: high
- Needed answer: define the retention posture and any recommended values

Answer:

- `crashsnap` ships with documented recommended retention values, but no active automatic pruning until the operator configures retention explicitly.
- Approved recommended initial values are `max_age_days = 14`, `max_total_size = "10GiB"`, and `max_snapshots = 50`.
- `prune` should be explicit and support `--dry-run`; documented recommendations must not be treated as silently active deletion behavior.
