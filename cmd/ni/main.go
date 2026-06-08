package main

import (
	"fmt"
	"os"

	"ni/internal/cli"
)

func main() {
	fmt.Fprintln(os.Stderr, "ni is deprecated; use namba-intent.")
	os.Exit(cli.RunWithOptions(os.Args[1:], os.Stdout, os.Stderr, cli.Options{CommandName: "ni"}))
}
