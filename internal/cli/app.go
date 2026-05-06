package cli

import (
	"fmt"
	"io"
	"strings"
)

const version = "0.0.0-dev"

var plannedCommands = []string{
	"run",
	"attach",
	"core",
	"service",
	"observe",
	"report",
	"inspect",
	"diff",
	"bundle",
	"verify",
	"doctor",
	"config",
	"schema",
	"prune",
}

func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		writeHelp(stdout)
		return 0
	}

	switch args[0] {
	case "-h", "--help", "help":
		writeHelp(stdout)
		return 0
	case "-v", "--version", "version":
		_, _ = fmt.Fprintf(stdout, "crashsnap %s\n", version)
		return 0
	}

	for _, command := range plannedCommands {
		if args[0] == command {
			_, _ = fmt.Fprintf(stderr, "crashsnap %s: not implemented yet\n", command)
			return 2
		}
	}

	_, _ = fmt.Fprintf(stderr, "unknown command %q\n\n", args[0])
	writeHelp(stderr)
	return 1
}

func writeHelp(w io.Writer) {
	_, _ = fmt.Fprintf(w, `crashsnap collects Linux crash diagnostics.

Usage:
  crashsnap [global options] command [command options]

Planned commands:
  %s

Bootstrap status:
  command scaffold only; collectors and reports are not implemented yet
`, strings.Join(plannedCommands, "\n  "))
}
