# USER_FLOWS.md

## Primary flow

1. A Linux operator identifies a failed command, running process, service, or core dump that needs diagnosis.
2. The operator chooses an appropriate collection mode based on whether live attachment risk is acceptable.
3. The product captures evidence into a snapshot directory and records collector successes, failures, and metadata.
4. The product normalizes, redacts, and analyzes the captured evidence.
5. The operator renders a report for local review or handoff.
6. The operator optionally bundles and verifies the artifacts for transfer or archival.

## Alternate flows

- The operator uses a non-invasive observe workflow when debugger attachment is too risky.
- The operator skips live capture and starts from an existing core dump.
- The operator re-renders a stored snapshot later in a different report format for a different audience.
- The operator inspects or diffs existing reports instead of recollecting evidence.

## Failure flows

- Collection partially succeeds because some collectors fail due to permissions, symbols, or environment restrictions; the product preserves successful evidence and reports the missing sections.
- Attach-based collection is rejected or aborted because the operator cannot safely accept target interruption.
- Report export is limited by redaction policy or missing required trust-boundary approval.
