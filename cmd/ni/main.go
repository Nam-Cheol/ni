package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/docstore"
	"ni/internal/core/graph"
	"ni/internal/core/harness"
	"ni/internal/core/lock"
	"ni/internal/core/profile"
	"ni/internal/core/prompt"
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
	case "graph":
		return runGraph(args[1:], stdout, stderr)
	case "harness":
		return runHarness(args[1:], stdout, stderr)
	case "init":
		return runInit(args[1:], stdout, stderr)
	case "run":
		return runRun(args[1:], stdout, stderr)
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

func runGraph(args []string, stdout io.Writer, stderr io.Writer) int {
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
			fmt.Fprintf(stderr, "unknown graph option: %s\n", args[i])
			return 1
		}
	}

	proposal, err := graph.Propose(dir)
	if err != nil {
		fmt.Fprintf(stderr, "graph failed: %v\n", err)
		return 1
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(proposal); err != nil {
			fmt.Fprintf(stderr, "graph failed: %v\n", err)
			return 1
		}
		return 0
	}
	fmt.Fprint(stdout, graph.FormatText(proposal))
	return 0
}

func runHarness(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 || args[0] != "plan" {
		fmt.Fprintln(stderr, "usage: ni harness plan --dir <path> [--json]")
		return 1
	}
	dir := "."
	jsonOutput := false
	for i := 1; i < len(args); i++ {
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
			fmt.Fprintf(stderr, "unknown harness option: %s\n", args[i])
			return 1
		}
	}

	proposal, err := harness.Plan(dir)
	if err != nil {
		fmt.Fprintf(stderr, "harness failed: %v\n", err)
		return 1
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(proposal); err != nil {
			fmt.Fprintf(stderr, "harness failed: %v\n", err)
			return 1
		}
		return 0
	}
	fmt.Fprint(stdout, harness.FormatText(proposal))
	return 0
}

func runInit(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	readinessProfile := profile.Default
	productType := contract.DefaultProductType
	var surfaces []string
	interactionMode := contract.DefaultInteractionMode
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		case "--profile":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --profile")
				return 1
			}
			readinessProfile = args[i+1]
			if err := profile.Validate(readinessProfile); err != nil {
				fmt.Fprintf(stderr, "invalid --profile value: %s (valid: %s)\n", readinessProfile, profile.NamesText())
				return 1
			}
			i++
		case "--product-type":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --product-type")
				return 1
			}
			productType = args[i+1]
			if err := contract.ValidateProductType(productType); err != nil {
				fmt.Fprintf(stderr, "invalid --product-type value: %v\n", err)
				return 1
			}
			i++
		case "--surface":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --surface")
				return 1
			}
			surfaces = append(surfaces, args[i+1])
			if err := contract.ValidateDeliverySurfaces(surfaces); err != nil {
				fmt.Fprintf(stderr, "invalid --surface value: %v\n", err)
				return 1
			}
			i++
		case "--interaction-mode":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --interaction-mode")
				return 1
			}
			interactionMode = args[i+1]
			if err := contract.ValidateInteractionMode(interactionMode); err != nil {
				fmt.Fprintf(stderr, "invalid --interaction-mode value: %v\n", err)
				return 1
			}
			i++
		default:
			fmt.Fprintf(stderr, "unknown init option: %s\n", args[i])
			return 1
		}
	}

	result, err := docstore.InitWithOptions(dir, docstore.InitOptions{
		ReadinessProfile: readinessProfile,
		ProductType:      productType,
		DeliverySurfaces: surfaces,
		InteractionMode:  interactionMode,
	})
	if err != nil {
		fmt.Fprintf(stderr, "init failed: %v\n", err)
		return 1
	}

	if len(surfaces) == 0 {
		surfaces = contract.DefaultDeliverySurfaces(productType)
	}
	fmt.Fprintf(stdout, "initialized ni planning workspace at %s\n", result.Root)
	fmt.Fprintf(stdout, "readiness profile: %s\n", readinessProfile)
	fmt.Fprintf(stdout, "product type: %s\n", productType)
	fmt.Fprintf(stdout, "delivery surfaces: %s\n", strings.Join(surfaces, ", "))
	fmt.Fprintf(stdout, "interaction mode: %s\n", interactionMode)
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

func runRun(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	out := ""
	maxChars := prompt.DefaultMaxChars
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		case "--out":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --out")
				return 1
			}
			out = args[i+1]
			i++
		case "--max-chars":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --max-chars")
				return 1
			}
			value, err := strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Fprintf(stderr, "invalid --max-chars value: %s\n", args[i+1])
				return 1
			}
			maxChars = value
			i++
		default:
			fmt.Fprintf(stderr, "unknown run option: %s\n", args[i])
			return 1
		}
	}

	result, err := prompt.Compile(prompt.Options{Dir: dir, Out: out, MaxChars: maxChars})
	if err != nil {
		fmt.Fprintf(stderr, "run failed: %v\n", err)
		return 1
	}
	if out != "" {
		fmt.Fprintf(stdout, "compiled prompt at %s\n", result.Out)
		return 0
	}
	fmt.Fprint(stdout, result.Prompt)
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
	fmt.Fprintf(stdout, "profile: %s\n", result.Profile)
	if result.ProductType != "" {
		fmt.Fprintf(stdout, "product type: %s\n", result.ProductType)
	}
	if len(result.DeliverySurfaces) > 0 {
		fmt.Fprintf(stdout, "delivery surfaces: %s\n", strings.Join(result.DeliverySurfaces, ", "))
	}
	if result.InteractionMode != "" {
		fmt.Fprintf(stdout, "interaction mode: %s\n", result.InteractionMode)
	}
	for _, guidance := range result.Guidance {
		fmt.Fprintf(stdout, "guidance: %s\n", guidance)
	}
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
  ni graph --dir <path> [--json]
  ni harness plan --dir <path> [--json]
  ni init --dir <path> [--profile concept|prototype|mvp|beta|production] [--product-type <type>] [--surface <surface>] [--interaction-mode <mode>]
  ni run --dir <path> [--out <path>] [--max-chars N]
  ni status --dir <path> [--json]
  ni version

Commands:
  end      Lock the accepted planning contract.
  graph    Propose a read-only work graph.
  harness  Propose a generated harness contract.
  init     Create planning docs and .ni skeleton.
  run      Compile a goal prompt from the locked plan.
  status   Validate planning readiness.
  version  Print the ni version.
`)
}
