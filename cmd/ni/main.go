package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"ni/internal/core/amendment"
	"ni/internal/core/contract"
	"ni/internal/core/docstore"
	"ni/internal/core/exporter"
	"ni/internal/core/feedback"
	"ni/internal/core/graph"
	"ni/internal/core/harness"
	"ni/internal/core/lock"
	"ni/internal/core/pressure"
	"ni/internal/core/profile"
	"ni/internal/core/prompt"
	"ni/internal/core/readiness"
	"ni/internal/core/target"
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
	case "amend":
		return runAmend(args[1:], stdout, stderr)
	case "end":
		return runEnd(args[1:], stdout, stderr)
	case "export":
		return runExport(args[1:], stdout, stderr)
	case "feedback":
		return runFeedback(args[1:], stdout, stderr)
	case "graph":
		return runGraph(args[1:], stdout, stderr)
	case "harness":
		return runHarness(args[1:], stdout, stderr)
	case "init":
		return runInit(args[1:], stdout, stderr)
	case "pressure":
		return runPressure(args[1:], stdout, stderr)
	case "run":
		return runRun(args[1:], stdout, stderr)
	case "relock":
		return runRelock(args[1:], stdout, stderr)
	case "status":
		return runStatus(args[1:], stdout, stderr)
	case "targets":
		return runTargets(args[1:], stdout, stderr)
	case "version":
		fmt.Fprintln(stdout, version.Version)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n\n", args[0])
		printHelp(stderr)
		return 1
	}
}

func runAmend(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		printAmendUsage(stderr)
		return 1
	}

	switch args[0] {
	case "create":
		return runAmendCreate(args[1:], stdout, stderr)
	case "list":
		return runAmendList(args[1:], stdout, stderr)
	case "show":
		return runAmendShow(args[1:], stdout, stderr)
	case "apply":
		return runAmendApply(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown amend command: %s\n", args[0])
		return 1
	}
}

func runAmendCreate(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	title := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		case "--title":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --title")
				return 1
			}
			title = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown amend create option: %s\n", args[i])
			return 1
		}
	}

	item, err := amendment.Create(amendment.CreateOptions{Dir: dir, Title: title})
	if err != nil {
		fmt.Fprintf(stderr, "amend create failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "created amendment %s at %s\n", item.ID, amendment.Path(dir, item.ID))
	return 0
}

func runAmendList(args []string, stdout io.Writer, stderr io.Writer) int {
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
			fmt.Fprintf(stderr, "unknown amend list option: %s\n", args[i])
			return 1
		}
	}
	items, err := amendment.List(dir)
	if err != nil {
		fmt.Fprintf(stderr, "amend list failed: %v\n", err)
		return 1
	}
	fmt.Fprint(stdout, amendment.FormatList(items))
	return 0
}

func runAmendShow(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseIDAndDir(args, stderr, "amend show")
	if !ok {
		return 1
	}
	item, err := amendment.Load(dir, id)
	if err != nil {
		fmt.Fprintf(stderr, "amend show failed: %v\n", err)
		return 1
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(item); err != nil {
		fmt.Fprintf(stderr, "amend show failed: %v\n", err)
		return 1
	}
	return 0
}

func runAmendApply(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseIDAndDir(args, stderr, "amend apply")
	if !ok {
		return 1
	}
	item, err := amendment.Apply(dir, id, time.Time{})
	if err != nil {
		fmt.Fprintf(stderr, "amend apply failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "applied amendment %s\n", item.ID)
	return 0
}

func runPressure(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: ni pressure status [--dir <path>] [--json]\n       ni pressure promote <id> [--dir <path>]\n       ni pressure retire <id> [--dir <path>]")
		return 1
	}

	switch args[0] {
	case "status":
		return runPressureStatus(args[1:], stdout, stderr)
	case "promote":
		return runPressurePromote(args[1:], stdout, stderr)
	case "retire":
		return runPressureRetire(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown pressure command: %s\n", args[0])
		return 1
	}
}

func runPressureStatus(args []string, stdout io.Writer, stderr io.Writer) int {
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
			fmt.Fprintf(stderr, "unknown pressure status option: %s\n", args[i])
			return 1
		}
	}

	ledger, err := pressure.Load(dir)
	if err != nil {
		fmt.Fprintf(stderr, "pressure status failed: %v\n", err)
		return 1
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(ledger); err != nil {
			fmt.Fprintf(stderr, "pressure status failed: %v\n", err)
			return 1
		}
		return 0
	}
	fmt.Fprint(stdout, pressure.FormatText(ledger))
	return 0
}

func runPressurePromote(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parsePressureIDAndDir(args, stderr, "promote")
	if !ok {
		return 1
	}
	item, err := pressure.Promote(dir, id)
	if err != nil {
		fmt.Fprintf(stderr, "pressure promote failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "promoted %s to %s\n", item.ID, item.Status)
	return 0
}

func runPressureRetire(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parsePressureIDAndDir(args, stderr, "retire")
	if !ok {
		return 1
	}
	item, err := pressure.Retire(dir, id)
	if err != nil {
		fmt.Fprintf(stderr, "pressure retire failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "retired %s\n", item.ID)
	return 0
}

func parsePressureIDAndDir(args []string, stderr io.Writer, command string) (string, string, bool) {
	dir := "."
	id := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return "", "", false
			}
			dir = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "--") {
				fmt.Fprintf(stderr, "unknown pressure %s option: %s\n", command, args[i])
				return "", "", false
			}
			if id != "" {
				fmt.Fprintf(stderr, "unexpected pressure %s argument: %s\n", command, args[i])
				return "", "", false
			}
			id = args[i]
		}
	}
	if id == "" {
		fmt.Fprintf(stderr, "missing pressure id for %s\n", command)
		return "", "", false
	}
	return id, dir, true
}

func parseIDAndDir(args []string, stderr io.Writer, command string) (string, string, bool) {
	dir := "."
	id := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return "", "", false
			}
			dir = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "--") {
				fmt.Fprintf(stderr, "unknown %s option: %s\n", command, args[i])
				return "", "", false
			}
			if id != "" {
				fmt.Fprintf(stderr, "unexpected %s argument: %s\n", command, args[i])
				return "", "", false
			}
			id = args[i]
		}
	}
	if id == "" {
		fmt.Fprintf(stderr, "missing id for %s\n", command)
		return "", "", false
	}
	return id, dir, true
}

func runFeedback(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: ni feedback add --file <path> [--dir <path>]\n       ni feedback list [--dir <path>] [--json]")
		return 1
	}

	switch args[0] {
	case "add":
		return runFeedbackAdd(args[1:], stdout, stderr)
	case "list":
		return runFeedbackList(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown feedback command: %s\n", args[0])
		return 1
	}
}

func runFeedbackAdd(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	file := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		case "--file":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --file")
				return 1
			}
			file = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown feedback add option: %s\n", args[i])
			return 1
		}
	}
	if file == "" {
		fmt.Fprintln(stderr, "missing --file")
		return 1
	}

	entry, err := feedback.Add(feedback.AddOptions{Dir: dir, File: file})
	if err != nil {
		fmt.Fprintf(stderr, "feedback add failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "recorded feedback from %s at %s\n", entry.SourceTarget, feedback.StorePath(dir))
	return 0
}

func runFeedbackList(args []string, stdout io.Writer, stderr io.Writer) int {
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
			fmt.Fprintf(stderr, "unknown feedback list option: %s\n", args[i])
			return 1
		}
	}

	entries, err := feedback.List(dir)
	if err != nil {
		fmt.Fprintf(stderr, "feedback list failed: %v\n", err)
		return 1
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(entries); err != nil {
			fmt.Fprintf(stderr, "feedback list failed: %v\n", err)
			return 1
		}
		return 0
	}
	fmt.Fprint(stdout, feedback.FormatText(entries))
	return 0
}

func runExport(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	out := ""
	targetName := ""
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
		case "--target":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --target")
				return 1
			}
			targetName = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown export option: %s\n", args[i])
			return 1
		}
	}
	if targetName == "" {
		fmt.Fprintln(stderr, "missing --target")
		return 1
	}

	result, err := exporter.Export(exporter.Options{Dir: dir, OutDir: out, Target: targetName})
	if err != nil {
		fmt.Fprintf(stderr, "export failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "exported %s seed package at %s\n", targetName, result.OutDir)
	for _, file := range result.Files {
		fmt.Fprintf(stdout, "created %s\n", file)
	}
	return 0
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
	if len(args) == 0 {
		printHarnessUsage(stderr)
		return 1
	}
	switch args[0] {
	case "plan":
		return runHarnessPlan(args[1:], stdout, stderr)
	case "candidates":
		return runHarnessCandidates(args[1:], stdout, stderr)
	case "propose":
		return runHarnessPropose(args[1:], stdout, stderr)
	case "validate":
		return runHarnessValidate(args[1:], stdout, stderr)
	case "accept":
		return runHarnessAccept(args[1:], stdout, stderr)
	case "retire":
		return runHarnessRetire(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown harness command: %s\n", args[0])
		return 1
	}
}

func runHarnessPlan(args []string, stdout io.Writer, stderr io.Writer) int {
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

func runHarnessCandidates(args []string, stdout io.Writer, stderr io.Writer) int {
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
			fmt.Fprintf(stderr, "unknown harness candidates option: %s\n", args[i])
			return 1
		}
	}

	ledger, err := harness.Candidates(dir)
	if err != nil {
		fmt.Fprintf(stderr, "harness candidates failed: %v\n", err)
		return 1
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(ledger); err != nil {
			fmt.Fprintf(stderr, "harness candidates failed: %v\n", err)
			return 1
		}
		return 0
	}
	fmt.Fprint(stdout, harness.FormatCandidates(ledger))
	return 0
}

func runHarnessPropose(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	pressureID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return 1
			}
			dir = args[i+1]
			i++
		case "--from-pressure":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --from-pressure")
				return 1
			}
			pressureID = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown harness propose option: %s\n", args[i])
			return 1
		}
	}
	if pressureID == "" {
		fmt.Fprintln(stderr, "missing --from-pressure")
		return 1
	}

	candidate, err := harness.ProposeFromPressure(dir, pressureID)
	if err != nil {
		fmt.Fprintf(stderr, "harness propose failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "proposed harness candidate %s from pressure %s\n", candidate.ID, pressureID)
	return 0
}

func runHarnessValidate(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, evidence, ok := parseHarnessCandidateIDDirEvidence(args, stderr)
	if !ok {
		return 1
	}
	candidate, err := harness.ValidateCandidate(dir, id, evidence)
	if err != nil {
		fmt.Fprintf(stderr, "harness validate failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "validated harness candidate %s to %s\n", candidate.ID, candidate.Status)
	return 0
}

func runHarnessAccept(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseHarnessCandidateIDAndDir(args, stderr, "accept")
	if !ok {
		return 1
	}
	candidate, err := harness.AcceptCandidate(dir, id)
	if err != nil {
		fmt.Fprintf(stderr, "harness accept failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "accepted harness candidate %s as %s\n", candidate.ID, candidate.Status)
	return 0
}

func runHarnessRetire(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseHarnessCandidateIDAndDir(args, stderr, "retire")
	if !ok {
		return 1
	}
	candidate, err := harness.RetireCandidate(dir, id)
	if err != nil {
		fmt.Fprintf(stderr, "harness retire failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "retired harness candidate %s\n", candidate.ID)
	return 0
}

func parseHarnessCandidateIDDirEvidence(args []string, stderr io.Writer) (string, string, string, bool) {
	dir := "."
	evidence := ""
	id := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return "", "", "", false
			}
			dir = args[i+1]
			i++
		case "--evidence":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --evidence")
				return "", "", "", false
			}
			evidence = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "--") {
				fmt.Fprintf(stderr, "unknown harness validate option: %s\n", args[i])
				return "", "", "", false
			}
			if id != "" {
				fmt.Fprintf(stderr, "unexpected harness validate argument: %s\n", args[i])
				return "", "", "", false
			}
			id = args[i]
		}
	}
	if id == "" {
		fmt.Fprintln(stderr, "missing candidate id for validate")
		return "", "", "", false
	}
	if evidence == "" {
		fmt.Fprintln(stderr, "missing --evidence")
		return "", "", "", false
	}
	return id, dir, evidence, true
}

func parseHarnessCandidateIDAndDir(args []string, stderr io.Writer, command string) (string, string, bool) {
	dir := "."
	id := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return "", "", false
			}
			dir = args[i+1]
			i++
		default:
			if strings.HasPrefix(args[i], "--") {
				fmt.Fprintf(stderr, "unknown harness %s option: %s\n", command, args[i])
				return "", "", false
			}
			if id != "" {
				fmt.Fprintf(stderr, "unexpected harness %s argument: %s\n", command, args[i])
				return "", "", false
			}
			id = args[i]
		}
	}
	if id == "" {
		fmt.Fprintf(stderr, "missing candidate id for %s\n", command)
		return "", "", false
	}
	return id, dir, true
}

func printHarnessUsage(w io.Writer) {
	fmt.Fprintln(w, "usage: ni harness plan --dir <path> [--json]\n       ni harness candidates [--dir <path>] [--json]\n       ni harness propose --from-pressure <id> [--dir <path>]\n       ni harness validate <candidate-id> --evidence <path> [--dir <path>]\n       ni harness accept <candidate-id> [--dir <path>]\n       ni harness retire <candidate-id> [--dir <path>]")
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

func runRelock(args []string, stdout io.Writer, stderr io.Writer) int {
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
			fmt.Fprintf(stderr, "unknown relock option: %s\n", args[i])
			return 1
		}
	}

	status := readiness.Evaluate(dir)
	if status.Status == readiness.StatusBlocked {
		fmt.Fprintln(stderr, "relock failed: readiness is BLOCKED; refusing to relock")
		return 1
	}

	currentHash, err := lock.CurrentLockHash(dir)
	if err != nil {
		fmt.Fprintf(stderr, "relock failed: %v\n", err)
		return 1
	}
	verification, err := lock.Verify(dir)
	if err != nil {
		fmt.Fprintf(stderr, "relock failed: %v\n", err)
		return 1
	}
	if len(verification.Mismatches) > 0 {
		ok, err := amendment.HasAppliedForLock(dir, currentHash)
		if err != nil {
			fmt.Fprintf(stderr, "relock failed: %v\n", err)
			return 1
		}
		if !ok {
			fmt.Fprintf(stderr, "relock failed: BLOCKED: lock hash mismatch for %s without an applied amendment\n", verification.Mismatches[0].Path)
			return 1
		}
	}

	now := time.Now().UTC()
	previous, err := lock.ArchiveCurrentAt(dir, now)
	if err != nil {
		fmt.Fprintf(stderr, "relock failed: %v\n", err)
		return 1
	}
	lockfile, err := lock.CreateAtWithPrevious(dir, now, &previous)
	if err != nil {
		fmt.Fprintf(stderr, "relock failed: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "relocked plan at %s\n", lockfile.Path)
	fmt.Fprintf(stdout, "previous lock archived at %s\n", filepath.Join(filepath.Clean(dir), previous.Path))
	fmt.Fprintf(stdout, "status %s\n", lockfile.Readiness.Status)
	return 0
}

func runRun(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	out := ""
	maxChars := prompt.DefaultMaxChars
	targetName := target.Generic
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
		case "--target":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --target")
				return 1
			}
			targetName = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown run option: %s\n", args[i])
			return 1
		}
	}

	result, err := prompt.Compile(prompt.Options{Dir: dir, Out: out, MaxChars: maxChars, Target: targetName})
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

func runTargets(args []string, stdout io.Writer, stderr io.Writer) int {
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown targets option: %s\n", args[i])
			return 1
		}
	}

	items := target.List()
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(items); err != nil {
			fmt.Fprintf(stderr, "targets failed: %v\n", err)
			return 1
		}
		return 0
	}
	for _, item := range items {
		fmt.Fprintf(stdout, "%s\t%s\t%s\n", item.Name, item.Artifact, item.Description)
	}
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
  ni amend create --title <title> [--dir <path>]
  ni amend list [--dir <path>]
  ni amend show <id> [--dir <path>]
  ni amend apply <id> [--dir <path>]
  ni end --dir <path>
  ni export --target hyper-run|namba-ai --out <dir> [--dir <path>]
  ni feedback add --file <path> [--dir <path>]
  ni feedback list [--dir <path>] [--json]
  ni graph --dir <path> [--json]
  ni harness plan --dir <path> [--json]
  ni harness candidates [--dir <path>] [--json]
  ni harness propose --from-pressure <id> [--dir <path>]
  ni harness validate <candidate-id> --evidence <path> [--dir <path>]
  ni harness accept <candidate-id> [--dir <path>]
  ni harness retire <candidate-id> [--dir <path>]
  ni init --dir <path> [--profile concept|prototype|mvp|beta|production] [--product-type <type>] [--surface <surface>] [--interaction-mode <mode>]
  ni pressure status [--dir <path>] [--json]
  ni pressure promote <id> [--dir <path>]
  ni pressure retire <id> [--dir <path>]
  ni relock --dir <path>
  ni run --dir <path> [--target <target>] [--out <path>] [--max-chars N]
  ni status --dir <path> [--json]
  ni targets [--json]
  ni version

Commands:
  amend   Create, inspect, and apply explicit contract amendments.
  end      Lock the accepted planning contract.
  export   Write locked-plan seed artifacts for a downstream target.
  feedback Record and list inert downstream feedback.
  graph    Propose a read-only work graph.
  harness  Manage inert generated harness proposals.
  init     Create planning docs and .ni skeleton.
  pressure Track inert planning pressure without changing readiness rules.
  relock   Create a new lock from an explicitly amended plan.
  run      Compile a goal prompt from the locked plan.
  status   Validate planning readiness.
  targets  List downstream prompt/export targets.
  version  Print the ni version.
`)
}

func printAmendUsage(w io.Writer) {
	fmt.Fprintln(w, "usage: ni amend create --title <title> [--dir <path>]\n       ni amend list [--dir <path>]\n       ni amend show <id> [--dir <path>]\n       ni amend apply <id> [--dir <path>]")
}
