# CRASHSNAP(1)

## NAME

**crashsnap** — collect, preserve, redact, analyze, and render Linux crash diagnostics from live processes, launched commands, systemd services, optional container-aware environments, and core dumps.

## SYNOPSIS

```text
crashsnap [GLOBAL OPTIONS] COMMAND [COMMAND OPTIONS]

crashsnap run [RUN OPTIONS] -- PROGRAM [ARGUMENT ...]
crashsnap attach --pid PID [ATTACH OPTIONS]
crashsnap core --core CORE_FILE --exe EXECUTABLE [CORE OPTIONS]
crashsnap service UNIT [SERVICE OPTIONS]
crashsnap observe --pid PID [OBSERVE OPTIONS]
crashsnap report SNAPSHOT_DIR [REPORT OPTIONS]
crashsnap inspect REPORT_FILE [INSPECT OPTIONS]
crashsnap diff LEFT RIGHT [DIFF OPTIONS]
crashsnap bundle SNAPSHOT_DIR [BUNDLE OPTIONS]
crashsnap verify BUNDLE_FILE [VERIFY OPTIONS]
crashsnap doctor [DOCTOR OPTIONS]
crashsnap config COMMAND [KEY] [VALUE]
crashsnap schema COMMAND [SCHEMA OPTIONS]
crashsnap prune [PRUNE OPTIONS]
```

Compatibility shorthand forms are also accepted:

```text
crashsnap -- PROGRAM [ARGUMENT ...]
crashsnap --pid PID
crashsnap --core CORE_FILE --exe EXECUTABLE
```

## DESCRIPTION

**crashsnap** is a Linux command-line diagnostics utility for collecting repeatable debugging evidence from live processes, launched commands, systemd services, optional container-aware environments, and core dumps. It drives **gdb(1)** when debugger evidence is required, reads Linux runtime metadata from **procfs**, and optionally gathers related host context from **journalctl(1)**, **dmesg(1)**, **coredumpctl(1)**, namespaces, cgroups, build IDs, debug-symbol providers, and Linux security modules.

The program is designed for production incident response, crash triage, postmortem debugging, CI failure capture, and reproducible local review on the same machine. It does not replace GDB. GDB is one collector in a broader evidence pipeline.

A full report may include target metadata, signal and exit analysis, error classification, confidence scoring, thread backtraces, register state, disassembly, memory mappings, loaded libraries, process arguments, environment summaries, discovered configuration files, open file descriptors, systemd unit state, cgroup information, optional container metadata, kernel and journal context, redaction summaries, symbol quality scoring, collection warnings, and artifact integrity metadata.

Reports may be rendered as text, JSON, YAML, Markdown, or HTML. Raw collection artifacts are stored as snapshots and may be normalized, redacted, analyzed, rendered, diffed, bundled, signed, verified, and re-rendered later.

## DESIGN PRINCIPLES

**Safe by default**
: Sensitive sections are disabled unless explicitly requested. Redaction is enabled by default. Generated artifacts are local-only by default.

**Evidence first**
: Snapshots are the source of truth. Reports are generated views over captured evidence.

**Partial success is useful**
: Failure of one collector does not discard successful evidence from other collectors. Reports identify missing sections and failed collectors explicitly.

**Low impact where possible**
: Non-invasive collectors are preferred when debugger attachment is not required. GDB attach behavior is explicit because it can stop the target process.

**Stable automation surface**
: JSON output is versioned by schema. Machine-readable output is suitable for CI, incident pipelines, dashboards, and archival systems.

**Auditable handling**
: Redaction, artifact hashing, collector status, command configuration, and schema versions are recorded in the snapshot manifest.

## EVIDENCE PIPELINE

A complete diagnostic workflow consists of these phases:

```text
capture -> normalize -> redact -> analyze -> render -> bundle -> verify
```

**capture**
: Collect raw process, debugger, core, service, optional container, and host artifacts into a snapshot directory.

**normalize**
: Convert raw collector output into a consistent internal representation.

**redact**
: Apply built-in and user-defined redaction policies before export.

**analyze**
: Classify observed failures, score confidence, evaluate symbol quality, and generate findings.

**render**
: Produce a human-readable or machine-readable report.

**bundle**
: Package reports, manifests, redacted artifacts, schemas, and hashes into an archive.

**verify**
: Validate archive integrity and optional signatures.

## MODES OF OPERATION

### Run Mode

Run mode launches a target program under managed collection.

```sh
crashsnap run -- ./myservice --config ./service.toml
```

By default, run mode captures a snapshot when the target exits with a non-zero status or terminates due to a fatal signal. Use `--on-exit` to collect on all exits.

### Attach Mode

Attach mode connects to an already-running process.

```sh
crashsnap attach --pid 4242 --profile safe --output report.html
```

Attaching a debugger can stop the target process. On production systems this may affect latency, health checks, failover behavior, watchdogs, and service-level objectives. Use `observe` when debugger state is not required.

### Observe Mode

Observe mode collects non-invasive runtime evidence without GDB attachment.

```sh
crashsnap observe --pid 4242 --with-maps --with-limits --with-cgroup
```

Observe mode may collect `/proc` metadata, command line, memory maps, cgroups, resource limits, current working directory, file descriptors, systemd association, and recent host context, subject to permission and redaction policy.

### Core Mode

Core mode analyzes an existing core dump with the executable that produced it.

```sh
crashsnap core --core ./core.4242 --exe ./bin/myservice
```

The executable should match the process image that created the core. Build-ID mismatch, stripped symbols, missing shared libraries, missing debuginfo, or unavailable source paths reduce report quality.

### Service Mode

Service mode collects evidence associated with a systemd unit.

```sh
crashsnap service myservice.service --core latest --profile postmortem
```

Service mode may identify the main PID, recent restarts, unit configuration, drop-ins, journal entries, cgroup path, resource limits, watchdog settings, OOM policy, and latest coredump metadata.

### Offline Report Mode

Offline report mode rebuilds reports from saved snapshots.

```sh
crashsnap report ./snapshots/incident-042 --format md --output incident-042.md
```

This mode is useful when raw evidence was collected during an incident but the final report format or redaction level was not known at collection time.

## QUICK START

Capture a crash report for a command:

```sh
crashsnap run -- ./myservice --config ./service.toml
```

Attach to a running process using the safe incident profile:

```sh
crashsnap attach --pid 31337 --profile safe --output myservice-report.html
```

Collect non-invasive runtime context:

```sh
crashsnap observe --pid 31337 --profile safe --format html --output observe.html
```

Analyze a core dump with full library and memory-map context:

```sh
crashsnap core --core ./core.31337 --exe ./bin/myservice --with-maps --with-libraries
```

Collect latest systemd-coredump evidence for a service:

```sh
crashsnap service myservice.service --core latest --snapshot-out ./incident-001
```

Render a retained snapshot as HTML:

```sh
crashsnap report ./incident-001 --format html --output incident-001.html
```

Create a signed compressed incident bundle:

```sh
crashsnap bundle ./incident-001 --sign --archive-format tar.gz --output incident-001.tar.gz
```

Validate a bundle before handoff:

```sh
crashsnap verify incident-001.tar.gz
```

## COMMANDS

### `run`

Launch a program and collect evidence when a trigger condition is met.

```text
crashsnap run [OPTIONS] -- PROGRAM [ARGUMENT ...]
```

Common options:

```text
--on-fail
        Capture only when the target exits non-zero or receives a fatal signal.
        This is the default.

--on-exit
        Capture when the target exits for any reason.

--signal SIG
        Capture when signal SIG is observed. May be specified more than once.

--follow-fork
        Follow the child after fork or clone when supported.

--no-follow-exec
        Do not continue following the target after exec.

--timeout SECONDS
        Maximum total collection time.

--snapshot-out DIR
        Write snapshot artifacts to DIR.

--report-out FILE
        Write rendered report to FILE.
```

### `attach`

Attach to a running process and collect debugger-backed evidence.

```text
crashsnap attach --pid PID [OPTIONS]
```

Attach-specific options:

```text
--pid PID
        Process ID to inspect.

--stop-timeout SECONDS
        Maximum time the target may remain stopped by debugger operations.

--nonstop
        Request GDB non-stop mode where supported.

--read-only
        Refuse debugger commands that modify target state.

--confirm-attach
        Require interactive confirmation before attaching to a non-child PID.

--force
        Skip interactive safety checks.

--slo-budget MILLISECONDS
        Abort when estimated debugger stop time may exceed the declared budget.
```

### `observe`

Collect non-invasive process metadata without debugger attachment.

```text
crashsnap observe --pid PID [OPTIONS]
```

Observe mode is appropriate when the target process is latency sensitive or when ptrace is unavailable.

```text
--with-maps
        Include process memory mappings.

--with-open-files
        Include open file descriptor metadata.

--with-limits
        Include rlimit and ulimit values.

--with-cgroup
        Include cgroup metadata.

--with-systemd
        Include associated systemd unit metadata when available.
```

### `core`

Analyze a core dump.

```text
crashsnap core --core CORE_FILE --exe EXECUTABLE [OPTIONS]
```

Core-specific options:

```text
--core CORE_FILE
        Core dump to analyze.

--exe EXECUTABLE
        Executable associated with CORE_FILE.

--sysroot DIR
        Root filesystem used for resolving target paths and libraries.

--debug-file-directory DIR
        Directory containing external debug files.

--build-id-cache DIR
        Build-ID cache used for debuginfo resolution.

--debuginfod
        Enable debuginfod lookup.

--no-debuginfod
        Disable debuginfod lookup.
```

### `service`

Collect evidence for a systemd unit.

```text
crashsnap service UNIT [OPTIONS]
```

Service-specific options:

```text
--pid main|PID
        Use the unit main PID or an explicit PID.

--core latest
        Use the latest coredump associated with UNIT.

--since TIME
        Include journal context since TIME.

--unit-file
        Include the unit file and drop-ins subject to redaction policy.

--status
        Include systemctl status output.
```

### `report`

Render a report from a snapshot directory.

```text
crashsnap report SNAPSHOT_DIR [OPTIONS]
```

Report options:

```text
--format text|json|yaml|html|md
        Render report in the selected format. Default: html.

--output FILE
        Write report to FILE.

--summary-only
        Generate a compact report without raw diagnostic sections.

--title STRING
        Override the report title.

--schema-version VERSION
        Render machine-readable output using the requested schema version.

--validate
        Validate report output against its schema.
```

### `inspect`

Query an existing report.

```text
crashsnap inspect REPORT_FILE [OPTIONS]
```

Inspect options:

```text
--threads
        Show thread summary.

--errors
        Show detected error conditions only.

--findings
        Show findings, confidence, evidence, and recommended actions.

--config
        Show extracted configuration summary.

--section NAME
        Print a single report section.

--jq EXPR
        Apply jq expression to JSON reports.

--validate
        Validate report schema.
```

### `diff`

Compare two reports or snapshots.

```text
crashsnap diff LEFT RIGHT [OPTIONS]
```

Diff output may include differences in binary build IDs, loaded libraries, environment summaries, configuration fingerprints, thread counts, memory maps, resource limits, cgroup settings, fault locations, and journal patterns.

### `bundle`

Package a snapshot, report, schemas, manifest, and integrity metadata.

```text
crashsnap bundle SNAPSHOT_DIR [OPTIONS]
```

Bundle options:

```text
--output FILE
        Write archive to FILE.

--archive-format tar.gz|zip
        Archive format. Default: tar.gz.

--sign
        Sign the bundle manifest when signing support is configured.

--include-raw
        Include raw unredacted artifacts. Requires --redaction off and confirmation.

--manifest
        Emit or include the snapshot manifest.
```

### `verify`

Validate a generated bundle.

```text
crashsnap verify BUNDLE_FILE [OPTIONS]
```

Verification checks archive structure, manifest hashes, schema compatibility, and optional signatures.

### `doctor`

Check whether the host is ready for collection.

```text
crashsnap doctor [OPTIONS]
```

Doctor options:

```text
--pid PID
        Check readiness for attaching to PID.

--core CORE_FILE --exe EXECUTABLE
        Check readiness for core analysis.

--container ID|NAME
        Check namespace and runtime visibility for a container when container-aware collection is in use.

--service UNIT
        Check systemd and coredump readiness for UNIT.

--json
        Emit machine-readable diagnostics.

--fix-suggestions
        Print suggested commands or configuration changes without applying them.
```

Doctor checks may include GDB availability, GDB Python support, ptrace policy, `kernel.yama.ptrace_scope`, current process ownership, dumpability, `hidepid` procfs options, SELinux or AppArmor confinement, systemd-coredump configuration, `ulimit -c`, `core_pattern`, writable snapshot directories, available disk space, namespace mappings, container runtime visibility when applicable, debuginfod availability, debug symbol availability, and temporary directory restrictions.

### `config`

Manage configuration.

```text
crashsnap config init
crashsnap config list
crashsnap config get KEY
crashsnap config set KEY VALUE
crashsnap config validate
```

### `schema`

Inspect schemas used for machine-readable output.

```text
crashsnap schema list
crashsnap schema show report.v1
crashsnap schema validate report.json
```

### `prune`

Remove old snapshots and reports according to retention policy.

```text
crashsnap prune [OPTIONS]
```

Prune options:

```text
--dry-run
        Show what would be removed.

--older-than DURATION
        Remove artifacts older than DURATION.

--max-total-size SIZE
        Prune until stored artifacts are below SIZE.
```

## GLOBAL OPTIONS

```text
-h, --help
        Display help and exit.

-V, --version
        Display version information and exit.

-q, --quiet
        Reduce non-error output.

-v, --verbose
        Increase diagnostic logging. May be repeated.

--no-color
        Disable ANSI color output. Also implied when NO_COLOR is set.

--log-file FILE
        Write internal logs to FILE.

--tmpdir DIR
        Use DIR for temporary files and intermediate artifacts.

--profile NAME
        Apply a named collection profile.

--format text|json|yaml|html|md
        Select output format. Default: html for reports, text for terminal summaries.

-o, --output FILE
        Write rendered output to FILE.

--snapshot-out DIR
        Write raw snapshot artifacts to DIR.

--overwrite
        Overwrite existing output files.

--timeout SECONDS
        Set maximum total collection time.

--timeout-gdb SECONDS
        Set maximum GDB collector time.

--timeout-symbols SECONDS
        Set maximum symbol-resolution time.

--max-report-size SIZE
        Limit rendered report size.

--max-snapshot-size SIZE
        Limit snapshot size.

--json
        Emit JSON output where supported.

--yaml
        Emit YAML output where supported.

--text
        Emit text output where supported.
```

## REPORT CONTENT OPTIONS

```text
--with-backtrace
        Include GDB backtraces. Enabled by default for run, attach, and core modes.

--with-registers
        Include register dumps.

--with-disassembly
        Include disassembly near the current instruction pointer.

--with-maps
        Include memory mappings.

--with-libraries
        Include loaded library inventory.

--with-args
        Include command-line arguments. Enabled by default with redaction.

--with-env
        Include environment variables. Disabled by default.

--with-config
        Attempt to discover and include configuration files. Disabled by default.

--config-path FILE|DIR
        Explicitly include a configuration file or directory. May be repeated.

--with-journal
        Include recent journal entries related to the target. Disabled by default.

--with-dmesg
        Include recent kernel log messages. Disabled by default.

--with-open-files
        Include open file descriptors. Disabled by default.

--with-cwd
        Include current working directory.

--with-limits
        Include rlimit and ulimit values.

--with-cgroup
        Include cgroup metadata and container metadata when applicable.

--with-systemd
        Include systemd unit metadata when available.

--with-selinux
        Include SELinux or AppArmor context when available.
```

## REDACTION AND PRIVACY

Redaction is enabled by default. The default policy redacts common secret names, token-like values, private keys, authorization headers, passwords, cloud credentials, selected filesystem paths, and high-risk configuration fields.

```text
--redaction default|strict|off
        Select redaction policy. Default: default.

--redact-env PATTERN
        Redact matching environment variable names. May be repeated.

--redact-key PATTERN
        Redact matching configuration keys. May be repeated.

--redact-path PATH
        Redact sensitive file paths. May be repeated.

--redaction-fail abort|warn|allow
        Behavior when redaction fails. Default: abort.

--unsafe-include-secrets
        Permit raw sensitive material in exported output. Requires --redaction off.
```

Sensitive report sections such as full environment variables, full configuration files, journal excerpts, kernel logs, and open file paths require explicit opt-in.

A report includes a redaction summary:

```text
Redaction Summary
    policy: default
    environment variables redacted: 12
    configuration keys redacted: 9
    paths redacted: 3
    journal values redacted: 4
    failures: 0
```

## PROFILES

Profiles define repeatable collection policies.

```text
--profile safe
        Conservative incident profile. Redaction enabled, sensitive sections disabled.

--profile full
        Broad collection profile. Redaction remains enabled.

--profile ci
        Machine-readable output for automated test and CI failure capture.

--profile container
        Optional container-aware collection with namespace and rootfs handling.

--profile postmortem
        Core, journal, symbol, and systemd-focused postmortem collection.
```

Example profile configuration:

```toml
[profiles.safe]
format = "html"
redaction = "strict"
include_env = false
include_config = false
include_journal = true
include_dmesg = false
max_frames = 64
max_threads = 128

[profiles.ci]
format = "json"
summary_only = true
include_journal = false
include_dmesg = false
schema_version = "1"
```

## FINDINGS AND CONFIDENCE

Failure classification is heuristic. Each finding records confidence, evidence, and limitations.

Example finding:

```json
{
  "finding": "segmentation_fault",
  "severity": "critical",
  "confidence": "high",
  "source": ["gdb", "core"],
  "evidence": [
    "target terminated with SIGSEGV",
    "fault address was 0x0",
    "top frame resolved to myservice::handle_request"
  ],
  "recommended_action": "inspect null pointer path in request handler"
}
```

Supported confidence values:

```text
high
medium
low
unknown
```

Supported severity values:

```text
critical
high
medium
low
info
```

## SYMBOL QUALITY

Reports include a symbol-quality section to prevent false confidence in incomplete backtraces.

Example:

```text
Symbol Quality
    target build-id: present
    target symbols: partial
    libc symbols: missing
    frame pointers: omitted
    debuginfod: unavailable
    confidence impact: stack traces may be incomplete
```

Symbol resolution may use executable symbols, external debug files, build-ID caches, distro debug packages, and debuginfod.

## REPORT STRUCTURE

A full report may contain:

```text
1. Executive summary
2. Collection status
3. Target identity
4. Command, service, optional container, or core metadata
5. Exit status and signal analysis
6. Findings, confidence, evidence, and recommended actions
7. Symbol quality
8. Thread overview
9. Backtraces
10. Registers and disassembly
11. Memory mappings
12. Shared libraries and build IDs
13. Runtime arguments
14. Environment summary
15. Configuration summary
16. Open file descriptors
17. Resource limits
18. Cgroup and namespace metadata
19. systemd unit context
20. Kernel and journal excerpts
21. Redaction summary
22. Artifact manifest
23. Collection warnings and limitations
```

## SNAPSHOT MANIFEST

Each snapshot contains a manifest describing collected artifacts.

The manifest may include:

```text
collector version
schema version
command line
profile name
configuration hash
host kernel version
distribution metadata
gdb version
target executable path
target build-id
library build-ids
container image digest when applicable
systemd unit name
timezone and locale
artifact list
SHA256 hashes
redaction policy
collector statuses
warnings and failures
```

Generated files are local-only by default:

```text
directories: 0700
files:       0600
archives:    0600
```

## SCHEMAS

Machine-readable output is schema-versioned.

Default schema files:

```text
/usr/share/crashsnap/schema/report-v1.json
/usr/share/crashsnap/schema/snapshot-v1.json
/usr/share/crashsnap/schema/finding-v1.json
/usr/share/crashsnap/schema/manifest-v1.json
```

Use `crashsnap schema list` to show installed schemas.

Use `crashsnap report --schema-version VERSION` to request a specific output contract.

## CONFIGURATION

Configuration is loaded in this order:

```text
1. built-in defaults
2. /etc/crashsnap/config.toml
3. ~/.config/crashsnap/config.toml
4. file specified by CRASHSNAP_CONFIG
5. command-line options
```

Default configuration locations:

```text
~/.config/crashsnap/config.toml
/etc/crashsnap/config.toml
```

Example configuration:

```toml
[report]
format = "html"
compress = true
include_env = false
include_args = true
include_config = false
include_journal = false
include_dmesg = false
schema_version = "1"

[collection]
timeout = 60
max_frames = 64
max_threads = 128
max_report_size = "50MiB"
max_snapshot_size = "500MiB"

[timeouts]
total = 60
gdb = 20
journal = 10
dmesg = 5
symbols = 15
container = 10

[gdb]
path = "/usr/bin/gdb"
safe_auto_load = true
nonstop = false
read_only = true
stop_timeout = 5

[paths]
snapshot_root = "/var/lib/crashsnap/snapshots"
report_root = "/var/lib/crashsnap/reports"
schema_root = "/usr/share/crashsnap/schema"

[redaction]
enabled = true
profile = "default"
fail = "abort"

[retention]
# Recommended initial local retention values.
# Automatic prune behavior remains opt-in until the operator configures retention explicitly.
max_snapshots = 50
max_age_days = 14
max_total_size = "10GiB"
```

Supported configuration keys include:

```text
report.format
report.compress
report.include_env
report.include_args
report.include_config
report.include_journal
report.include_dmesg
report.schema_version
report.summary_only
collection.timeout
collection.max_frames
collection.max_threads
collection.max_report_size
collection.max_snapshot_size
timeouts.total
timeouts.gdb
timeouts.journal
timeouts.dmesg
timeouts.symbols
timeouts.container
gdb.path
gdb.safe_auto_load
gdb.nonstop
gdb.read_only
gdb.stop_timeout
paths.snapshot_root
paths.report_root
paths.schema_root
redaction.enabled
redaction.profile
redaction.fail
retention.max_snapshots
retention.max_age_days
retention.max_total_size
```

## ENVIRONMENT

```text
CRASHSNAP_CONFIG
        Path to the main configuration file.

CRASHSNAP_GDB
        Path to the GDB executable.

CRASHSNAP_TMPDIR
        Temporary directory override.

CRASHSNAP_REDACTION
        Redaction policy override: default, strict, or off.

CRASHSNAP_PROFILE
        Default profile name.

DEBUGINFOD_URLS
        debuginfod server list used by GDB and symbol resolution.

NO_COLOR
        Disable colored output.
```

Compatibility variables:

```text
GDBITCH_CONFIG
GDBITCH_GDB
GDBITCH_TMPDIR
GDBITCH_REDACT
```

## FILES

```text
~/.config/crashsnap/config.toml
        Per-user configuration.

/etc/crashsnap/config.toml
        System-wide configuration.

/var/lib/crashsnap/snapshots/
        Default snapshot storage.

/var/lib/crashsnap/reports/
        Default report output directory.

/usr/share/crashsnap/schema/
        Installed JSON schemas.

/var/log/crashsnap/
        Optional internal log directory.
```

## EXIT STATUS

```text
0   Success.
1   General failure.
2   Invalid command-line usage.
3   Target process not found.
4   GDB invocation failed.
5   Snapshot collection timed out.
6   Permission denied.
7   Core file invalid or unreadable.
8   Report generation failed.
9   Redaction failure.
10  Partial success; report generated with missing sections.
11  Schema validation failed.
12  Bundle verification failed.
13  Unsafe export rejected.
14  Artifact size limit exceeded.
15  Unsupported target environment.
```

## SECURITY NOTES

Attaching GDB to a live process may stop the process. Use `observe` or core-based analysis when stop-the-world behavior is unacceptable.

Reports can contain sensitive data, including command-line arguments, environment variables, file paths, configuration values, hostnames, service names, tokens, credentials, customer identifiers, and internal topology. Redaction is enabled by default, but generated reports should still be reviewed before any unredacted local access.

Raw snapshots may contain more sensitive material than rendered reports. Treat snapshot directories as local incident evidence and restrict access.

Use `bundle --sign` and `verify` for local archival integrity checks.

Use `--redaction off`, `--unsafe-include-secrets`, or `--include-raw` only through explicit operator action in a trusted local environment.

## LIMITATIONS

Stripped binaries reduce symbol resolution quality.

Highly optimized builds can produce incomplete or misleading stack traces.

Missing frame pointers and missing unwind data reduce backtrace reliability.

`SIGKILL` and abrupt termination can leave incomplete userspace evidence.

Debugger attachment can perturb process timing or availability.

Kernel policy, namespaces, optional container boundaries, seccomp, SELinux, AppArmor, and ptrace restrictions may block collection.

Core dumps may not contain all memory regions, file-backed mappings, or host context.

Journal and kernel logs may be unavailable, rotated, rate-limited, or permission restricted.

Failure classification is heuristic and should not be treated as ground truth without evidence review.

## TROUBLESHOOTING

### Permission denied while attaching

Check:

```text
kernel.yama.ptrace_scope
process ownership
container namespace restrictions
hidepid procfs mount options
SELinux or AppArmor policy
target dumpability
```

Run:

```sh
crashsnap doctor --pid PID --fix-suggestions
```

### Missing symbols in backtrace

Install or enable debug symbols for:

```text
target binary
libc
libstdc++
runtime libraries
application plugins
JIT runtimes where applicable
```

Also check:

```sh
crashsnap core --core ./core --exe ./service --debuginfod
crashsnap doctor --core ./core --exe ./service
```

### No core dump found

Check:

```text
ulimit -c
kernel.core_pattern
systemd-coredump configuration
process dumpability
service sandboxing
writable core destination
```

For systemd services:

```sh
crashsnap service myservice.service --core latest --status
```

### Report missing environment or config data

This may be intentional. Environment and configuration sections are disabled by default and redacted when enabled.

Use explicit options:

```sh
crashsnap attach --pid PID --with-env --with-config --redaction strict
```

### Attach causes service disruption

Use non-invasive collection first:

```sh
crashsnap observe --pid PID --profile safe
```

For crash analysis, prefer core dumps:

```sh
crashsnap service myservice.service --core latest
```

## EXAMPLES

### Capture a crash report for a binary

```sh
crashsnap run -- ./myservice --config ./service.toml
```

### Capture on every process exit

```sh
crashsnap run --on-exit -- ./batch-job --input ./data.json
```

### Attach safely to a production process

```sh
crashsnap attach --pid 31337 \
  --profile safe \
  --stop-timeout 3 \
  --read-only \
  --output myservice-report.html
```

### Observe a process without debugger attachment

```sh
crashsnap observe --pid 31337 \
  --with-maps \
  --with-limits \
  --with-cgroup \
  --output observe.html
```

### Analyze a core dump with external symbols

```sh
crashsnap core \
  --core ./core.31337 \
  --exe ./bin/myservice \
  --debug-file-directory /usr/lib/debug \
  --with-maps \
  --with-libraries
```

### Use debuginfod during core analysis

```sh
crashsnap core --core ./core --exe ./myservice --debuginfod
```

### Collect systemd service context

```sh
crashsnap service myservice.service \
  --status \
  --since -30m \
  --with-journal \
  --profile postmortem
```

### Include config and journal data while redacting aggressively

```sh
crashsnap attach --pid 31337 \
  --with-config \
  --with-journal \
  --redaction strict \
  --redact-env '^(TOKEN|KEY|SECRET|PASSWORD)$'
```

### Generate stable JSON for CI

```sh
crashsnap run --profile ci --schema-version 1 -- ./test-binary
```

### Inspect only detected findings

```sh
crashsnap inspect report.json --findings
```

### Compare two incidents

```sh
crashsnap diff ./snapshots/incident-041 ./snapshots/incident-042
```

### Bundle and verify an incident archive

```sh
crashsnap bundle ./snapshots/incident-042 --sign --output incident-042.tar.gz
crashsnap verify incident-042.tar.gz
```

### Prune old snapshots

```sh
crashsnap prune --older-than 14d --dry-run
crashsnap prune --older-than 14d
```

## RELATED TOOLS

```text
gdb(1)
coredumpctl(1)
journalctl(1)
dmesg(1)
strace(1)
ltrace(1)
perf(1)
readelf(1)
objdump(1)
systemctl(1)
```

## PRODUCTION READINESS CHECKLIST

Before using **crashsnap** as part of a production incident workflow, verify:

```text
redaction defaults are enabled
artifact directories are private
schema version is pinned for automation
snapshot retention policy is explicitly configured
doctor passes for target hosts
core dump policy is understood
service owners understand attach impact
symbol resolution is available
CI and incident profiles are tested
bundle verification is part of handoff
```

## LICENSE

MIT

## AUTHORS

Maintained by the project contributors.

## SEE ALSO

**gdb(1)**, **proc(5)**, **core(5)**, **systemd-coredump(8)**, **coredumpctl(1)**, **journalctl(1)**, **systemctl(1)**.
