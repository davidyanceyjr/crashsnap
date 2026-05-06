# ACCEPTANCE_CRITERIA.md

## Product-level acceptance criteria

- AC-001: The product can capture crash-diagnostic evidence for at least one in-scope target type and store it as a reusable snapshot artifact.
- AC-002: A report can be rendered from a saved snapshot without recollecting from the original target.
- AC-003: Reports and snapshots make missing evidence, failed collectors, or low-confidence findings explicit.
- AC-004: Machine-readable output declares its schema version.
- AC-005: Sensitive output is not exported by default.

## Workflow-level acceptance criteria

- The canonical product document, implementation plan, decisions, risks, and open questions do not contradict each other.
- Blocking questions are either answered or explicitly recorded in `docs/OPEN_QUESTIONS.md`.
- Non-blocking assumptions that implementation may carry are visible and reversible.
- `docs/IMPLEMENTATION_HANDOFF.md` accurately summarizes the implementation-facing package.

## Current package acceptance status

- The documentation package meets the workflow-level requirement of making blockers visible.
- The package does not yet meet the standard for trustworthy implementation start because blocking product questions remain unresolved.
