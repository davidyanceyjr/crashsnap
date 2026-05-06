# Core Implementation Skill

## Goal

Implement one narrow project slice without expanding scope.

## Constraints

- Keep handlers and orchestration thin.
- Put reusable logic in testable units.
- Prefer explicit contracts over hidden side effects.
- Avoid broad refactors unless the session explicitly authorizes them.
- Preserve backward compatibility unless the session or a recorded decision allows a change.

## Preferred behavior

- define the smallest complete slice
- identify the main entrypoint and supporting units
- implement the minimal end-to-end path
- add or update tests with the change
- record any durable decision or risk discovered

## Required tests

- happy path coverage for the slice
- at least one relevant failure path when behavior can fail meaningfully

