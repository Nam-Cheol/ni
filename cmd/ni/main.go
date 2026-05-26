package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"ni/internal/core/docstore"
	"ni/internal/core/lock"
	"ni/internal/core/readiness"
	"ni/internal/version"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp(stdout)
		return 0
	}

	switch args[0] {
	case "end":
		return runEnd(args[1:], stdout, stderr)
	case "init":
		return runInit(args[1:], stdout, stderr)
	case "status":
		return runStatus(args[1:], stdout, stderr)
	case "version":
		fmt.Fprintln(stdout, version.Version)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n\n", args[0])
		printHelp(stderr)
		return 1
	}
}

func runInit(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown init option: %s\n", args[i])
			return 1
		}
	}

	result, err := docstore.Init(dir)
	if err != nil {
		fmt.Fprintf(stderr, "init failed: %v\n", err)
		return 1
	}

	fmt.Fprintf(stdout, "initialized ni planning workspace at %s\n", result.Root)
	for _, path := range result.Created {
		fmt.Fprintf(stdout, "created %s\n", path)
	}
	for _, path := range result.Existing {
		fmt.Fprintf(stdout, "exists %s\n", path)
	}
	return 0
}

func runEnd(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown end option: %s\n", args[i])
			return 1
		}
	}

	lockfile, err := lock.Create(dir)
	if err != nil {
		fmt.Fprintf(stderr, "end failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "locked plan at %s\n", lockfile.Path)
	fmt.Fprintf(stdout, "status %s\n", lockfile.Readiness.Status)
	return 0
}

func runStatus(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown status option: %s\n", args[i])
			return 1
		}
	}

	result := readiness.Evaluate(dir)
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(result); err != nil {
			fmt.Fprintf(stderr, "status failed: %v\n", err)
			return 1
		}
		return 0
	}

	fmt.Fprintln(stdout, result.Status)
	for _, issue := range result.Issues {
		fmt.Fprintf(stdout, "%s %s: %s\n", issue.Severity, issue.RuleID, issue.Message)
	}
	return 0
}

func printHelp(w io.Writer) {
	fmt.Fprint(w, `ni is a project intent compiler.

Usage:
  ni --help
  ni end --dir <path>
  ni init --dir <path>
  ni status --dir <path> [--json]
  ni version

Commands:
  end      Lock the accepted planning contract.
  init     Create planning docs and .ni skeleton.
  status   Validate planning readiness.
  version  Print the ni version.
`)
}
