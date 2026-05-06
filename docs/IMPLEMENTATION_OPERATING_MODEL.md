# IMPLEMENTATION_OPERATING_MODEL.md

## Purpose

This project uses an agent/session/skills workflow to keep implementation work focused and context-efficient.

## Main idea

```text
Imported source-of-truth package defines implementation truth.
ROADMAP.md decides sequence.
SESSION.md decides current work.
SKILLS/*.md decide implementation behavior.
SESSIONS/*.md preserve session history.
DECISIONS.md prevents policy churn.
RISKS.md keeps active risks visible.
```

## Recommended workflow

1. Import the source-of-truth handoff package.
2. Pick the next roadmap slice.
3. Create or switch to the branch for that slice.
4. Rewrite `SESSION.md` for that slice.
5. Read the upstream handoff files relevant to the slice.
6. Load only the relevant skills.
7. Implement the slice.
8. Run tests.
9. Save a completed session note.
10. Update `SESSION.md`.
11. Update `DECISIONS.md` or `RISKS.md` when needed.
12. Stage only the intended session files.
13. Commit with a human-readable message.
14. Push the branch and open a pull request.
15. Get human approval before merge and cleanup.
16. Merge the pull request and clean up branch state.

## Session sizing

Ideal session:

- one objective
- a narrow file set
- explicit tests
- no unrelated refactors
- no silent milestone crossing

## Context discipline

For most sessions, load only:

- `AGENT.md`
- `ROADMAP.md`
- `SESSION.md`
- the upstream canonical product document
- `docs/IMPLEMENTATION_PLAN.md`
- `docs/IMPLEMENTATION_HANDOFF.md` and `docs/OPEN_QUESTIONS.md` when relevant
- relevant `SKILLS/*.md`
- source files directly involved in the slice

Each task within a session should be sized so active context stays under 30% of the model's entire context window.

If a task starts to approach that limit:

- split the task into a smaller follow-up slice
- compress findings into `SESSION.md`, `DECISIONS.md`, `RISKS.md`, or a session note
- stop loading additional files until the active context is reduced

## Durable file rules

Update `DECISIONS.md` only when a durable architectural, product, workflow, or contract decision is made.

Update `RISKS.md` only when a meaningful project risk is introduced, reduced, or newly discovered.

Create a new file under `SESSIONS/` at the end of every implementation session.
