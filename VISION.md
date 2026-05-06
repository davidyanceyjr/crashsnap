# VISION.md

## Vision

Provide a Linux incident-diagnostics tool that turns fragmented crash evidence into a repeatable, reviewable artifact rather than a one-off debugging session.

## Why this should exist

Crash investigation often requires manual use of `gdb`, `/proc`, journal logs, coredump tooling, and host metadata sources. The product should reduce that stitching work and preserve the resulting evidence so operators and developers can reason from the same artifact.

## For whom

The current narrative most clearly targets Linux operators, incident responders, and developers involved in crash triage and postmortem debugging.

This user set is still broad and should be narrowed further in a later session.

## Desired outcome

Users can collect diagnostic evidence once, review it safely, and hand it off or revisit it later without repeating the original incident conditions.

## What success looks like

- Evidence capture is reproducible.
- Sensitive output is constrained by default.
- Handoff artifacts are more trustworthy than ad hoc screenshots or shell transcripts.
- The implementation team can build from explicit product truth rather than implied assumptions.
