# Session 003: Retention Policy

## Objective

Resolve the remaining blocking retention-policy decision and clear the final product-definition blocker.

## Completed

- Recorded approved recommended retention values of `14` days, `10GiB`, and `50` snapshots.
- Recorded that automatic pruning remains opt-in until the operator configures retention explicitly.
- Cleared the last blocking question from the implementation handoff package.

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
- `SESSIONS/003-retention-policy.md`

## Questions asked

- None.

## Questions resolved

- AQ-006: `crashsnap` ships with documented recommended retention values, while automatic pruning remains opt-in until explicitly configured.

## Decisions

- Recommended retention values are approved.
- Automatic pruning requires explicit operator configuration.

## Risks updated

- Retention ambiguity is resolved.
- Retention misconfiguration remains an implementation risk.

## Next recommended session

`004-implementation-phasing`
