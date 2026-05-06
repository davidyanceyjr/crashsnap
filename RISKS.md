# RISKS.md

## R-001: Retention misconfiguration

- Status: Active
- Area: artifact lifecycle
- Severity: High
- Likelihood: Medium
- Description: Implementation could accidentally treat documented retention recommendations as active automatic deletion policy and remove evidence without explicit operator configuration.
- Mitigation: Keep retention configuration and manual prune flows separate, require explicit operator configuration before automatic pruning, and test the unconfigured default path directly.
- Related slices: `003a-config-model`, `003b-config-discovery`, `003c-config-retention-validation`, `019a-prune-dry-run-manual`, `019b-prune-configured-automatic`
- Related questions: none

## R-002: Over-broad initial scope

- Status: Active
- Area: delivery planning
- Severity: High
- Likelihood: High
- Description: The claimed command surface is broad enough that implementation could fail by trying to deliver too many modes and formats at once.
- Mitigation: Preserve phased delivery in `ROADMAP.md`, keep each session to one slice, and treat deferred areas explicitly instead of pulling them into adjacent work.
- Related slices: `001-repo-bootstrap` through `021-diff-snapshot-report`
- Related questions: none

## R-003: Local data exposure

- Status: Active
- Area: security and privacy
- Severity: High
- Likelihood: Medium
- Description: Crash artifacts can contain secrets, file paths, credentials, and environment data that may be exposed too broadly if redaction and explicit-unredacted behavior are implemented weakly.
- Mitigation: Keep local-only handling visible in docs and code, default implemented flows to redaction, and require explicit operator action for unredacted output.
- Related slices: `005a-report-text-render`, `005b-report-json-render`, `007a-redaction-rule-engine`, `007b-redaction-report-integration`, `016a-bundle-create`, `016b-verify-bundle`
- Related questions: none

## R-004: Environment variability and privilege limits

- Status: Active
- Area: collector behavior
- Severity: Medium
- Likelihood: High
- Description: Linux environments differ across ptrace policy, procfs visibility, systemd availability, and installed tooling, so implementation could fail if it silently assumes elevated or uniform access.
- Mitigation: Prefer explicit probe and degrade-gracefully behavior, add environment diagnostics early, and cover permission-failure paths in tests.
- Related slices: `004a-observe-command-bootstrap`, `004b-observe-procfs-collectors`, `009a-doctor-tooling-checks`, `009b-doctor-permission-checks`, `013-service-unit-resolution`, `014-service-context-capture`, `015-attach-mode-explicit-debugger`
- Related questions: none

## R-005: Bootstrap churn before toolchain approval

- Status: Mitigated
- Area: repository foundations
- Severity: Medium
- Likelihood: Low
- Description: The repository currently has no implementation scaffold or Git history, so code layout or dependency choices could still create avoidable churn if the bootstrap slice grows beyond its intended scope.
- Mitigation: `HG-001` is now approved for Go and a single-binary distribution target. Keep slice `001-repo-bootstrap` narrow, prefer the Go standard library initially, and record any additional dependency or layout commitments durably before later feature slices begin.
- Related slices: `001-repo-bootstrap`, `003a-config-model`, `003b-config-discovery`, `003c-config-retention-validation`, `017-config-read-validate`, `018-config-init-set`
- Related questions: none
