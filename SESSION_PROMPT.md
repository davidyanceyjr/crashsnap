# SESSION_PROMPT.md

Use this prompt to start an implementation session.

```markdown
You are the implementation agent for this project.

Read:

1. AGENT.md
2. ROADMAP.md
3. SESSION.md
4. The upstream canonical product document for this repository
5. docs/IMPLEMENTATION_PLAN.md
6. docs/IMPLEMENTATION_HANDOFF.md
7. docs/OPEN_QUESTIONS.md when relevant to the slice
8. Only the SKILLS files listed in SESSION.md

Then perform the current session objective.

Rules:

- Stay inside SESSION.md scope.
- Keep each task within the session under 30% of the entire model context window.
- If the task grows beyond that budget, narrow the slice, summarize durable findings, and continue in a smaller task.
- Treat the imported source-of-truth handoff package as upstream planning truth.
- Create or switch to the session git branch before implementation work begins.
- Do not implement deferred features.
- Add or update tests for the slice.
- Run relevant tests.
- Update SESSION.md with results.
- Create a completed note under SESSIONS/<session-id>.md.
- Stage only the intended session files.
- Commit with a human-readable message.
- Push the session branch.
- Open a pull request.
- Get human approval before merge and cleanup.
- After approval, merge the pull request and clean up the branch state.

Final response must include:

- summary
- files changed
- tests run
- known failures
- next recommended session
```

Additional requirement:

Every open question must be answered or documented if unanswered.

If an open question is discovered:

- classify it as BLOCKING or NON-BLOCKING
- record it in SESSION.md
- do not guess silently
- if answered, record the answer and source
- if unanswered at session end, move or copy it to docs/OPEN_QUESTIONS.md
- update DECISIONS.md or RISKS.md when appropriate
- if the question exposes a gap in the upstream handoff package, record that explicitly
