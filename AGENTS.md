# AGENT.md

## Purpose

This file is the operating contract for the implementation agent.

The implementation agent turns an upstream source-of-truth handoff package into working software through small, reviewable implementation slices.

The implementation agent does not invent product truth.

## Core rule

Do not begin normal implementation until first-run setup is complete.

If first-run setup is not complete, follow `FIRST_RUN.md` before following `SESSION.md`.

## First-run gate

Before running any normal implementation session, determine whether this repository has completed first-run setup.

First-run setup is complete only if:

- upstream source-of-truth handoff files are present
- the canonical product document is identified
- `ROADMAP.md` has been replaced with a project-specific roadmap
- `ROADMAP.md` was derived from `docs/IMPLEMENTATION_PLAN.md`
- blocked or deferred slices are visible in `ROADMAP.md`
- safety-sensitive slices reference human gates
- `SESSION.md` has been rewritten for the first implementation slice
- durable setup decisions are recorded in `DECISIONS.md`
- active setup risks are recorded in `RISKS.md`
- unresolved setup questions are recorded in `docs/OPEN_QUESTIONS.md`

If first-run setup is not complete, stop normal implementation and execute `FIRST_RUN.md`.

Do not implement product features during first run unless `FIRST_RUN.md` has produced a valid first implementation slice and `SESSION.md` explicitly authorizes that slice.

The generic template `ROADMAP.md` is not implementation truth.

The implementation stream owns `ROADMAP.md`, but it must be derived from the upstream source-of-truth handoff, primarily `docs/IMPLEMENTATION_PLAN.md`.

Constrain roadmap derivation with:

- `docs/IMPLEMENTATION_HANDOFF.md`
- `README.md` or `PRODUCT_BRIEF.md`
- `DECISIONS.md`
- `RISKS.md`
- `docs/OPEN_QUESTIONS.md`

## Required upstream handoff package

A normal implementation project should contain:

- `README.md` or `PRODUCT_BRIEF.md`
- `docs/IMPLEMENTATION_PLAN.md`
- `docs/IMPLEMENTATION_HANDOFF.md`
- `docs/OPEN_QUESTIONS.md`
- `DECISIONS.md`
- `RISKS.md`

If these files are missing, incomplete, or contradictory, do not silently guess.

Record the issue in `docs/OPEN_QUESTIONS.md` and classify it as `BLOCKING` or `NON-BLOCKING`.

## Source priority

When implementation questions arise, use sources in this order:

1. `SESSION.md`
2. `ROADMAP.md`
3. `docs/IMPLEMENTATION_PLAN.md`
4. `docs/IMPLEMENTATION_HANDOFF.md`
5. canonical product document: `README.md` or `PRODUCT_BRIEF.md`
6. `DECISIONS.md`
7. `RISKS.md`
8. `docs/OPEN_QUESTIONS.md`
9. relevant files in `SKILLS/`
10. source code and tests

If a lower-priority file conflicts with a higher-priority file, pause the affected work and record a question.

Do not resolve durable contradictions by guessing.

## Normal implementation session loop

After first-run setup is complete:

1. Read `AGENT.md`.
2. Read `SESSION.md`.
3. Read the current slice in `ROADMAP.md`.
4. Read source-of-truth files relevant to the slice.
5. Read only the relevant files in `SKILLS/`.
6. Confirm no blocking questions affect the slice.
7. Implement only the current slice.
8. Add or update tests for the slice.
9. Run the relevant validation commands.
10. Update durable project files as needed.
11. Write a completed session note under `SESSIONS/`.
12. Update `SESSION.md` for the next slice or handoff.
13. Commit, push, and open a pull request if the environment supports it.
14. Stop for human review before merge.

## Session boundaries

Each implementation session must have a narrow scope.

A session should normally correspond to one roadmap slice.

Do not expand the session scope because related work is nearby.

Do not perform opportunistic rewrites.

Do not refactor unrelated areas unless the current slice requires it and the reason is recorded.

Do not implement future roadmap slices early.

## Branching and pull requests

Use a branch for each implementation session when the environment supports Git.

Recommended branch format:

```text
session/<slice-id>-<short-name>
```

Examples:

```text
session/001-repo-bootstrap
session/002-config-schema
session/003-domain-model
```

Commits should be small and directly related to the current slice.

Pull requests should summarize:

- slice ID
- objective
- files changed
- tests run
- decisions made
- risks changed
- open questions
- out-of-scope work intentionally avoided

Do not merge without human approval.

## Human gates

Human approval is required before work involving:

- deleting data
- destructive CLI commands
- irreversible migrations
- external writes
- production deployment
- authentication
- authorization
- billing or payment flows
- public API contracts
- schema compatibility commitments
- security-sensitive defaults
- license or compliance changes
- broad dependency changes
- irreversible project structure changes

Human gates must be recorded in `docs/HUMAN_GATES.md` and referenced from the relevant roadmap slice.

If the gate is not documented, do not perform the gated work.

## Open questions

Use `docs/OPEN_QUESTIONS.md` for durable unresolved questions.

Use `BLOCKING` when the answer affects correctness, safety, public behavior, data shape, architecture, or irreversible work.

Use `NON-BLOCKING` only when the work can proceed safely with a reversible documented assumption.

Open-question format:

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

A `BLOCKING` question stops the affected work.

Do not bury questions in chat-only prose.

## Decisions

Use `DECISIONS.md` for durable project decisions.

Record a decision when it affects future implementation, architecture, dependencies, behavior, compatibility, deployment, testing, or operations.

Decision format:

```markdown
## D-<number>: <decision title>

- Date:
- Status: Proposed | Accepted | Superseded
- Context:
- Decision:
- Consequences:
- Source:
```

Do not record trivial moment-to-moment choices.

Do not overwrite old decisions. Supersede them with a new entry.

## Risks

Use `RISKS.md` for active risks.

Risk format:

```markdown
## R-<number>: <risk title>

- Status: Active | Mitigated | Accepted
- Area:
- Severity: Low | Medium | High
- Likelihood: Low | Medium | High
- Description:
- Mitigation:
- Related slices:
- Related questions:
```

Update risks when implementation reveals new uncertainty or changes the risk level.

## Skills

Use `SKILLS/*.md` as narrow implementation guidance.

Load only the skills relevant to the current session.

If a skill conflicts with the upstream handoff or `SESSION.md`, stop and record the conflict.

Do not use skills as permission to expand scope.

## Testing and validation

Each implementation slice should define expected validation in `SESSION.md`.

Prefer automated tests.

If automated tests are not possible yet, use documented manual validation.

Record exactly what was run.

Do not claim tests passed unless they were actually run.

If tests cannot be run, say why.

## File update expectations

During a normal implementation session, update files as needed:

- `SESSION.md`
- `DECISIONS.md`
- `RISKS.md`
- `docs/OPEN_QUESTIONS.md`
- `docs/HUMAN_GATES.md`
- `SESSIONS/<session-id>.md`

Do not update durable files just to appear thorough.

Update them when project truth changed.

## Session notes

At the end of each completed implementation session, create a session note under `SESSIONS/`.

Use the naming pattern:

```text
SESSIONS/<session-id>-<slice-id>.md
```

The note should include:

- session ID
- branch
- slice ID
- objective
- summary of changes
- files changed
- tests run
- decisions recorded
- risks recorded
- open questions created or changed
- human gates encountered
- follow-up slice recommendation

## Prohibited behavior

Do not:

- skip `FIRST_RUN.md` when first-run setup is incomplete
- treat template roadmap content as project truth
- invent product behavior absent from the upstream handoff
- implement through unresolved blocking questions
- bypass human gates
- broaden the current slice without recording why
- perform destructive actions without explicit approval
- silently overwrite upstream decisions, risks, or open questions
- claim tests were run when they were not
- hide assumptions in prose instead of durable files
- merge without human review

## Final response format for normal sessions

At the end of a normal implementation session, report:

```markdown
## Session result

### Completed

### Slice

### Files changed

### Tests run

### Decisions recorded

### Risks recorded

### Open questions

### Human gates

### Ready for review?

Yes or no.
```

If not ready, state exactly what remains.

## Summary

The implementation lifecycle is:

```text
FIRST_RUN.md
  derives ROADMAP.md from upstream handoff
  rewrites SESSION.md for slice 001

SESSION.md
  drives one implementation session

ROADMAP.md
  orders implementation slices

SESSIONS/*.md
  preserve completed session history
```

The implementation agent builds software only after the first-run gate has produced a project-specific roadmap and a valid first slice.
