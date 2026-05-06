# Session 002: Blocking Product Answers

## Objective

Resolve the first set of blocking product-definition questions and update the source-of-truth package.

## Completed

- Recorded `crashsnap` as the public product and command name.
- Recorded that the full README command set remains planned scope and may be implemented in subsets.
- Recorded local-only scope and explicit unredacted handling.
- Recorded required systemd support, optional container support, and no elevated-privilege assumptions.
- Narrowed the remaining blocker to retention policy and `prune` defaults.

## Files changed

- `README.md`
- `DECISIONS.md`
- `RISKS.md`
- `docs/OPEN_QUESTIONS.md`
- `docs/IMPLEMENTATION_PLAN.md`
- `docs/IMPLEMENTATION_HANDOFF.md`
- `REQUIREMENTS.md`
- `NON_GOALS.md`
- `SESSION.md`
- `SESSIONS/002-blocking-product-answers.md`

## Questions asked

- BQ-005: What retention policy should local artifacts follow?

## Questions resolved

- AQ-002: The public product and command name is `crashsnap`.
- AQ-003: All commands in `README.md` remain planned product surface and may be implemented in subsets.
- AQ-004: The product is local-only in scope, and unredacted output is allowed only through explicit operator action.
- AQ-005: Systemd support is required, container support is optional, and the product must not assume elevated privileges are available.

## Decisions

- `crashsnap` is the public name.
- Full command surface remains planned scope.
- Local-only and explicit-unredacted posture is approved.
- Systemd is required; containers are optional; no elevated-privilege assumption is allowed.

## Risks updated

- Retention ambiguity is now the main remaining blocker.
- Local data exposure risk remains important because unredacted local handling is allowed explicitly.

## Next recommended session

`003-retention-policy`
