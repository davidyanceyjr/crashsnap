# SESSION.md

## Current session

ID: `007-002b-manifest-core-fields`

Status: ready-with-gate-constraint

## Slice

`002b-manifest-core-fields`

## Objective

Define the initial draft manifest base model for snapshot identity, schema references, timestamps, and provenance fields, while keeping the contract explicitly internal until `HG-002` is approved for public compatibility claims.

## Scope

Implement:

- define the initial draft manifest base model
- add serialization helpers and base-manifest fixtures
- add tests for required fields and round-trip serialization
- document the manifest boundary as internal-only draft contract

## Out of scope

- collector status modeling
- live capture implementation
- report rendering
- redaction
- bundle signing
- public compatibility promises beyond what `HG-002` explicitly approves

## Source priority for this slice

1. `SESSION.md`
2. `ROADMAP.md` slice `002b-manifest-core-fields`
3. `docs/IMPLEMENTATION_PLAN.md`
4. `docs/IMPLEMENTATION_HANDOFF.md`
5. `README.md`
6. `DECISIONS.md`
7. `RISKS.md`
8. `docs/OPEN_QUESTIONS.md`

## Relevant skills

- `SKILLS/core-implementation.md`
- `SKILLS/testing.md`
- `SKILLS/docs-and-decisions.md`

## Acceptance criteria

- a draft manifest base model is defined in code and fixtures
- required-field and round-trip serialization tests cover the happy path and a relevant failure path
- the work stays explicitly internal-draft and does not claim stable external compatibility
- any durable manifest-boundary decision is recorded in `DECISIONS.md`

## Validation plan

- run `go test ./...`
- run targeted manifest-focused tests if added separately
- verify any draft schema or fixture artifacts remain clearly non-stable

## Canonical product document

`README.md`

## Blocking questions

None currently recorded.

## Human gates

- `HG-002` is not yet approved. This session may proceed only on internal draft contract work and must not claim public machine-readable stability.

## Current repo state

Slices `001-repo-bootstrap` and `002a-snapshot-layout` are complete. The repository now contains a Go module, a thin CLI scaffold, draft snapshot layout code and fixtures, baseline tests, project-specific build and test commands, and a local Git repository.

## Next action required

Implement slice `002b-manifest-core-fields` as an internal draft contract, or obtain `HG-002` approval before making any public compatibility claims.
