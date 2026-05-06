# FIRST_RUN.md

## Purpose

This file defines the required first-run workflow for a new implementation repository.

The first run is not a normal implementation session. Its job is to convert the imported upstream source-of-truth handoff package into a project-specific implementation workspace.

The primary first-run output is `ROADMAP.md`.

`ROADMAP.md` is owned by the implementation stream, but it must be derived from the upstream source-of-truth handoff. Do not begin normal implementation until this first-run workflow is complete.

## Required upstream inputs

The implementation repository must contain the following upstream files before first-run completion:

- `README.md` or `PRODUCT_BRIEF.md`
- `docs/IMPLEMENTATION_PLAN.md`
- `docs/IMPLEMENTATION_HANDOFF.md`
- `docs/OPEN_QUESTIONS.md`
- `DECISIONS.md`
- `RISKS.md`

If any required upstream input is missing, incomplete, stale, or contradictory, do not invent missing truth. Record the issue in `docs/OPEN_QUESTIONS.md` as `BLOCKING` or `NON-BLOCKING`.

## Source priority for deriving ROADMAP.md

Use these files in this priority order:

1. `docs/IMPLEMENTATION_PLAN.md`
2. `docs/IMPLEMENTATION_HANDOFF.md`
3. canonical product document: `README.md` or `PRODUCT_BRIEF.md`
4. `DECISIONS.md`
5. `RISKS.md`
6. `docs/OPEN_QUESTIONS.md`

The primary source is `docs/IMPLEMENTATION_PLAN.md`, especially the section named or equivalent to `Recommended implementation sequence`.

If the implementation plan does not contain a usable implementation sequence, record a `BLOCKING` question unless a safe, narrow bootstrap slice can be derived from the rest of the handoff.

## First-run workflow

### 1. Inspect repository state

Read:

- `AGENT.md`
- `FIRST_RUN.md`
- `README.md` or `PRODUCT_BRIEF.md`
- `docs/IMPLEMENTATION_PLAN.md`
- `docs/IMPLEMENTATION_HANDOFF.md`
- `docs/OPEN_QUESTIONS.md`
- `DECISIONS.md`
- `RISKS.md`

Determine whether the repository is still in template state or has already been customized.

Signs of template state include:

- generic project names
- placeholder roadmap milestones
- empty or generic risks
- empty or generic decisions
- `SESSION.md` still describing template setup
- no project-specific source files
- no clear first implementation slice

Do not proceed as if implementation has already started unless repository state proves it.

### 2. Identify the canonical product document

Use exactly one canonical product document as product truth:

- `README.md`
- `PRODUCT_BRIEF.md`

If both exist, determine which one the upstream handoff identifies as canonical.

If both exist and conflict, record a `BLOCKING` question.

If neither exists, record a `BLOCKING` question.

### 3. Verify handoff readiness

Read `docs/IMPLEMENTATION_HANDOFF.md`.

Determine whether the upstream source-of-truth stream marked the project ready for implementation.

If the handoff lists blocking issues affecting implementation sequence, do not create roadmap slices for the blocked areas.

Blocked areas commonly include:

- public API or CLI contracts
- data model
- persistence behavior
- destructive operations
- authentication or authorization
- external integrations
- deployment assumptions
- migration behavior
- billing or payment behavior
- user-visible defaults
- security-sensitive behavior

### 4. Derive ROADMAP.md

Replace the template `ROADMAP.md` with a project-specific roadmap.

The roadmap must be derived from `docs/IMPLEMENTATION_PLAN.md` and constrained by:

- `docs/IMPLEMENTATION_HANDOFF.md`
- `README.md` or `PRODUCT_BRIEF.md`
- `DECISIONS.md`
- `RISKS.md`
- `docs/OPEN_QUESTIONS.md`

Do not add product behavior unsupported by the upstream handoff.

Do not treat the generic template roadmap as product truth. The default roadmap is only an example structure.

### 5. Required ROADMAP.md structure

`ROADMAP.md` must contain:

```markdown
# ROADMAP.md

## Purpose

## Source documents

## Roadmap status

## Milestones

## Slice backlog

## Blocked or deferred slices

## Human gates

## Roadmap maintenance rules
```

Each milestone must contain:

```markdown
## Milestone N: <name>

### Goal

### Source basis

### Exit criteria

### Slices
```

Each slice must contain:

```markdown
#### <slice-id>: <slice name>

- Objective:
- Source basis:
- Depends on:
- Expected files or areas:
- Tests or validation:
- Human gates:
- Blocking questions:
- Out of scope:
```

Slice IDs should be stable and ordered, for example:

- `001-repo-bootstrap`
- `002-config-schema`
- `003-domain-model`
- `004-first-command-or-endpoint`
- `005-persistence-foundation`
- `006-happy-path-test`

Avoid vague slice names such as:

- `001-build-app`
- `002-finish-backend`
- `003-polish`
- `004-fix-everything`

### 6. Slice sizing rules

A roadmap slice must be small enough to complete in one focused implementation session.

Prefer vertical, testable slices.

Good:

```text
config file accepted -> parsed -> validated -> used by first command
```

Bad:

```text
build all config functionality
```

A slice should have:

- one clear objective
- limited files or areas
- explicit tests or validation
- clear out-of-scope boundaries
- no hidden milestone crossing

### 7. Human-gated work

Any slice involving safety-sensitive or irreversible behavior must reference a human gate.

Safety-sensitive work includes:

- deleting data
- mutating persistent state
- database migrations
- external writes
- production deployment
- authentication
- authorization
- billing
- payment flows
- destructive CLI commands
- public API contracts
- schema compatibility
- security-sensitive defaults

Record these gates in `docs/HUMAN_GATES.md` and record related risks in `RISKS.md`.

Do not hide human gates inside prose. Attach them to the specific roadmap slice.

### 8. Open-question handling during first run

If a question blocks roadmap creation or slice sequencing, record it in `docs/OPEN_QUESTIONS.md`.

Use this format:

```markdown
## Q-<number>: <question>

- Status: BLOCKING | NON-BLOCKING
- Area:
- Blocks:
- Source:
- Why it matters:
- Safe temporary assumption:
- Reversal cost:
- Needed answer:
```

If the question blocks a specific slice, also list it under that slice in `ROADMAP.md`.

A `BLOCKING` question stops the affected work.

A `NON-BLOCKING` question may proceed only with a documented safe assumption and low reversal cost.

### 9. Rewrite SESSION.md

After `ROADMAP.md` exists, rewrite `SESSION.md` for the first implementation slice.

The first implementation slice should normally be the smallest useful bootstrap slice.

Common first slices:

- `001-repo-bootstrap`
- `001-tooling-bootstrap`
- `001-project-skeleton`
- `001-config-foundation`

`SESSION.md` must include:

- session ID
- branch name
- current slice ID
- objective
- source documents to read
- files expected to change
- relevant skills
- explicit out-of-scope work
- expected tests or validation
- blocking questions
- known risks
- completion criteria

Do not leave `SESSION.md` in generic template state after first run.

### 10. Update durable project files

During first run, update these files as needed:

- `ROADMAP.md`
- `SESSION.md`
- `DECISIONS.md`
- `RISKS.md`
- `docs/OPEN_QUESTIONS.md`
- `docs/HUMAN_GATES.md`

Update `DECISIONS.md` only for durable choices.

Update `RISKS.md` only for meaningful risks.

## First-run stop conditions

Stop first run and do not proceed to implementation if:

- the canonical product document is missing
- `docs/IMPLEMENTATION_PLAN.md` is missing
- `docs/IMPLEMENTATION_HANDOFF.md` says implementation is not ready
- roadmap sequencing depends on unresolved blocking product behavior
- the first implementation slice cannot be safely defined
- safety-sensitive behavior lacks a human gate
- upstream files contradict each other in a way that affects implementation

When stopping, update:

- `docs/OPEN_QUESTIONS.md`
- `SESSION.md`
- `RISKS.md`

Then state what human decision or upstream correction is needed.

## First-run completion criteria

First run is complete only when:

- upstream handoff files are present
- canonical product document is identified
- `ROADMAP.md` has been replaced with a project-specific roadmap
- roadmap slices cite or reference their source basis
- blocked or deferred work is visible
- safety-sensitive slices have human gates
- `SESSION.md` describes the first implementation slice
- durable decisions are recorded in `DECISIONS.md`
- active risks are recorded in `RISKS.md`
- unresolved questions are recorded in `docs/OPEN_QUESTIONS.md`

Only after these conditions are met may normal implementation sessions begin.

## First-run final response format

At the end of first run, report:

```markdown
## First-run result

### Completed

### ROADMAP.md status

### First implementation slice

### Files changed

### Decisions recorded

### Risks recorded

### Blocking questions

### Non-blocking assumptions

### Ready for implementation?

Yes or no.
```

If not ready, state exactly what is blocking implementation.

## Non-goals

First run does not:

- implement product features
- create broad architecture beyond what the upstream handoff supports
- resolve product ambiguity by guessing
- treat template placeholders as real project decisions
- bypass human gates

## Summary

The required sequence is:

```text
import upstream handoff
read docs/IMPLEMENTATION_PLAN.md
derive ROADMAP.md
rewrite SESSION.md for slice 001
record questions, decisions, and risks
then begin normal implementation sessions
```
