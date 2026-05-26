package main

import (
	"fmt"
	"io"
	"os"

	"ni/internal/core/docstore"
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
	case "init":
		return runInit(args[1:], stdout, stderr)
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

func printHelp(w io.Writer) {
	fmt.Fprint(w, `ni is a project intent compiler.

Usage:
  ni --help
  ni init --dir <path>
  ni version

Commands:
  init     Create planning docs and .ni skeleton.
  version  Print the ni version.
`)
}
