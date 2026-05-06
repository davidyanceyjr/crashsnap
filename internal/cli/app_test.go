package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run([]string{"--help"}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("expected help exit code 0, got %d", exitCode)
	}

	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}

	if !strings.Contains(stdout.String(), "Planned commands:") {
		t.Fatalf("expected help output to list planned commands, got %q", stdout.String())
	}

	if !strings.Contains(stdout.String(), "command scaffold only") {
		t.Fatalf("expected help output to describe bootstrap status, got %q", stdout.String())
	}
}

func TestRunVersion(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run([]string{"version"}, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("expected version exit code 0, got %d", exitCode)
	}

	if got := stdout.String(); !strings.Contains(got, "crashsnap "+version) {
		t.Fatalf("expected version output, got %q", got)
	}
}

func TestRunPlannedCommandPlaceholder(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run([]string{"doctor"}, &stdout, &stderr)

	if exitCode != 2 {
		t.Fatalf("expected placeholder exit code 2, got %d", exitCode)
	}

	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	if !strings.Contains(stderr.String(), "not implemented yet") {
		t.Fatalf("expected placeholder message, got %q", stderr.String())
	}
}

func TestRunUnknownCommand(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := Run([]string{"unknown-command"}, &stdout, &stderr)

	if exitCode != 1 {
		t.Fatalf("expected unknown command exit code 1, got %d", exitCode)
	}

	if stdout.Len() != 0 {
		t.Fatalf("expected empty stdout, got %q", stdout.String())
	}

	if !strings.Contains(stderr.String(), "unknown command") {
		t.Fatalf("expected unknown command message, got %q", stderr.String())
	}

	if !strings.Contains(stderr.String(), "Planned commands:") {
		t.Fatalf("expected help output on stderr, got %q", stderr.String())
	}
}
