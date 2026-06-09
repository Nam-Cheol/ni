package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"ni/internal/core/amendment"
	"ni/internal/core/collab"
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
	initui "ni/internal/tui/init"
	"ni/internal/version"
)

type Options struct {
	CommandName string
}

var activeCommandName = "namba-intent"

var initInput io.Reader = os.Stdin

var initInputIsTerminal = func() bool {
	info, err := os.Stdin.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}

func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	return RunWithOptions(args, stdout, stderr, Options{CommandName: "namba-intent"})
}

func RunWithOptions(args []string, stdout io.Writer, stderr io.Writer, options Options) int {
	previousCommandName := activeCommandName
	if options.CommandName != "" {
		activeCommandName = options.CommandName
	}
	defer func() {
		activeCommandName = previousCommandName
	}()
	return run(args, stdout, stderr)
}

func commandName() string {
	if activeCommandName == "" {
		return "namba-intent"
	}
	return activeCommandName
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp(stdout)
		return 0
	}

	switch args[0] {
	case "amend":
		return runAmend(args[1:], stdout, stderr)
	case "conflicts":
		return runConflicts(args[1:], stdout, stderr)
	case "diff":
		return runDiff(args[1:], stdout, stderr)
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
		return exitUsageError
	}
}

func runDiff(args []string, stdout io.Writer, stderr io.Writer) int {
	base, head, jsonOutput, ok := parseCollabArgs(args, stderr, "diff")
	if !ok {
		return exitUsageError
	}
	result, err := collab.Diff(base, head)
	if err != nil {
		return failCommand(stdout, stderr, "diff", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(result); err != nil {
			return failCommand(stdout, stderr, "diff", err, jsonOutput)
		}
		return exitOK
	}
	fmt.Fprint(stdout, collab.FormatDiff(result))
	return exitOK
}

func runConflicts(args []string, stdout io.Writer, stderr io.Writer) int {
	base, head, jsonOutput, ok := parseCollabArgs(args, stderr, "conflicts")
	if !ok {
		return exitUsageError
	}
	result, err := collab.Conflicts(base, head)
	if err != nil {
		return failCommand(stdout, stderr, "conflicts", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(result); err != nil {
			return failCommand(stdout, stderr, "conflicts", err, jsonOutput)
		}
	} else {
		fmt.Fprint(stdout, collab.FormatConflicts(result))
	}
	if len(result.Conflicts) > 0 {
		return exitSemanticConflict
	}
	return exitOK
}

func parseCollabArgs(args []string, stderr io.Writer, command string) (string, string, bool, bool) {
	base := ""
	head := ""
	jsonOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--base":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --base")
				return "", "", false, false
			}
			base = args[i+1]
			i++
		case "--head":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --head")
				return "", "", false, false
			}
			head = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			fmt.Fprintf(stderr, "unknown %s option: %s\n", command, args[i])
			return "", "", false, false
		}
	}
	if base == "" {
		fmt.Fprintln(stderr, "missing --base")
		return "", "", false, false
	}
	if head == "" {
		fmt.Fprintln(stderr, "missing --head")
		return "", "", false, false
	}
	return base, head, jsonOutput, true
}

func runAmend(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		printAmendUsage(stderr)
		return exitUsageError
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
		return exitUsageError
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
				return exitUsageError
			}
			dir = args[i+1]
			i++
		case "--title":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --title")
				return exitUsageError
			}
			title = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown amend create option: %s\n", args[i])
			return exitUsageError
		}
	}

	item, err := amendment.Create(amendment.CreateOptions{Dir: dir, Title: title})
	if err != nil {
		return failCommand(stdout, stderr, "amend create", err, false)
	}
	fmt.Fprintf(stdout, "created amendment %s at %s\n", item.ID, amendment.Path(dir, item.ID))
	return exitOK
}

func runAmendList(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return exitUsageError
			}
			dir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown amend list option: %s\n", args[i])
			return exitUsageError
		}
	}
	items, err := amendment.List(dir)
	if err != nil {
		return failCommand(stdout, stderr, "amend list", err, false)
	}
	fmt.Fprint(stdout, amendment.FormatList(items))
	return exitOK
}

func runAmendShow(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseIDAndDir(args, stderr, "amend show")
	if !ok {
		return exitUsageError
	}
	item, err := amendment.Load(dir, id)
	if err != nil {
		return failCommand(stdout, stderr, "amend show", err, false)
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(item); err != nil {
		return failCommand(stdout, stderr, "amend show", err, false)
	}
	return exitOK
}

func runAmendApply(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseIDAndDir(args, stderr, "amend apply")
	if !ok {
		return exitUsageError
	}
	item, err := amendment.Apply(dir, id, time.Time{})
	if err != nil {
		return failCommand(stdout, stderr, "amend apply", err, false)
	}
	fmt.Fprintf(stdout, "applied amendment %s\n", item.ID)
	return exitOK
}

func runPressure(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintf(stderr, "usage: %[1]s pressure status [--dir <path>] [--json]\n       %[1]s pressure promote <id> [--dir <path>]\n       %[1]s pressure retire <id> [--dir <path>]\n", commandName())
		return exitUsageError
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
		return exitUsageError
	}
}

func runPressureStatus(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := jsonRequested(args)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				return failStructured(stdout, stderr, usageErrorf("missing value for --dir"), jsonOutput)
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown pressure status option: %s", args[i]), jsonOutput)
		}
	}

	ledger, err := pressure.Load(dir)
	if err != nil {
		return failCommand(stdout, stderr, "pressure status", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(ledger); err != nil {
			return failCommand(stdout, stderr, "pressure status", err, jsonOutput)
		}
		return exitOK
	}
	fmt.Fprint(stdout, pressure.FormatText(ledger))
	return exitOK
}

func runPressurePromote(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parsePressureIDAndDir(args, stderr, "promote")
	if !ok {
		return exitUsageError
	}
	item, err := pressure.Promote(dir, id)
	if err != nil {
		return failCommand(stdout, stderr, "pressure promote", err, false)
	}
	fmt.Fprintf(stdout, "promoted %s to %s\n", item.ID, item.Status)
	return exitOK
}

func runPressureRetire(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parsePressureIDAndDir(args, stderr, "retire")
	if !ok {
		return exitUsageError
	}
	item, err := pressure.Retire(dir, id)
	if err != nil {
		return failCommand(stdout, stderr, "pressure retire", err, false)
	}
	fmt.Fprintf(stdout, "retired %s\n", item.ID)
	return exitOK
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
		fmt.Fprintf(stderr, "usage: %[1]s feedback add --file <path> [--dir <path>]\n       %[1]s feedback list [--dir <path>] [--json]\n", commandName())
		return exitUsageError
	}

	switch args[0] {
	case "add":
		return runFeedbackAdd(args[1:], stdout, stderr)
	case "list":
		return runFeedbackList(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown feedback command: %s\n", args[0])
		return exitUsageError
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
				return exitUsageError
			}
			dir = args[i+1]
			i++
		case "--file":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --file")
				return exitUsageError
			}
			file = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown feedback add option: %s\n", args[i])
			return exitUsageError
		}
	}
	if file == "" {
		fmt.Fprintln(stderr, "missing --file")
		return exitUsageError
	}

	entry, err := feedback.Add(feedback.AddOptions{Dir: dir, File: file})
	if err != nil {
		return failCommand(stdout, stderr, "feedback add", err, false)
	}
	fmt.Fprintf(stdout, "recorded feedback from %s at %s\n", entry.SourceTarget, feedback.StorePath(dir))
	return exitOK
}

func runFeedbackList(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := jsonRequested(args)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				return failStructured(stdout, stderr, usageErrorf("missing value for --dir"), jsonOutput)
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown feedback list option: %s", args[i]), jsonOutput)
		}
	}

	entries, err := feedback.List(dir)
	if err != nil {
		return failCommand(stdout, stderr, "feedback list", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(entries); err != nil {
			return failCommand(stdout, stderr, "feedback list", err, jsonOutput)
		}
		return exitOK
	}
	fmt.Fprint(stdout, feedback.FormatText(entries))
	return exitOK
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
				return exitUsageError
			}
			dir = args[i+1]
			i++
		case "--out":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --out")
				return exitUsageError
			}
			out = args[i+1]
			i++
		case "--target":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --target")
				return exitUsageError
			}
			targetName = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown export option: %s\n", args[i])
			return exitUsageError
		}
	}
	if targetName == "" {
		fmt.Fprintln(stderr, "missing --target")
		return exitUsageError
	}

	result, err := exporter.Export(exporter.Options{Dir: dir, OutDir: out, Target: targetName})
	if err != nil {
		return failCommand(stdout, stderr, "export", err, false)
	}
	fmt.Fprintf(stdout, "exported %s seed package at %s\n", targetName, result.OutDir)
	for _, file := range result.Files {
		fmt.Fprintf(stdout, "created %s\n", file)
	}
	return exitOK
}

func runGraph(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := jsonRequested(args)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				return failStructured(stdout, stderr, usageErrorf("missing value for --dir"), jsonOutput)
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown graph option: %s", args[i]), jsonOutput)
		}
	}

	proposal, err := graph.Propose(dir)
	if err != nil {
		return failCommand(stdout, stderr, "graph", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(proposal); err != nil {
			return failCommand(stdout, stderr, "graph", err, jsonOutput)
		}
		return exitOK
	}
	fmt.Fprint(stdout, graph.FormatText(proposal))
	return exitOK
}

func runHarness(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		printHarnessUsage(stderr)
		return exitUsageError
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
		return exitUsageError
	}
}

func runHarnessPlan(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := jsonRequested(args)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				return failStructured(stdout, stderr, usageErrorf("missing value for --dir"), jsonOutput)
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown harness option: %s", args[i]), jsonOutput)
		}
	}

	proposal, err := harness.Plan(dir)
	if err != nil {
		return failCommand(stdout, stderr, "harness", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(proposal); err != nil {
			return failCommand(stdout, stderr, "harness", err, jsonOutput)
		}
		return exitOK
	}
	fmt.Fprint(stdout, harness.FormatText(proposal))
	return exitOK
}

func runHarnessCandidates(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := jsonRequested(args)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				return failStructured(stdout, stderr, usageErrorf("missing value for --dir"), jsonOutput)
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown harness candidates option: %s", args[i]), jsonOutput)
		}
	}

	ledger, err := harness.Candidates(dir)
	if err != nil {
		return failCommand(stdout, stderr, "harness candidates", err, jsonOutput)
	}
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(ledger); err != nil {
			return failCommand(stdout, stderr, "harness candidates", err, jsonOutput)
		}
		return exitOK
	}
	fmt.Fprint(stdout, harness.FormatCandidates(ledger))
	return exitOK
}

func runHarnessPropose(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	pressureID := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return exitUsageError
			}
			dir = args[i+1]
			i++
		case "--from-pressure":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --from-pressure")
				return exitUsageError
			}
			pressureID = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown harness propose option: %s\n", args[i])
			return exitUsageError
		}
	}
	if pressureID == "" {
		fmt.Fprintln(stderr, "missing --from-pressure")
		return exitUsageError
	}

	candidate, err := harness.ProposeFromPressure(dir, pressureID)
	if err != nil {
		return failCommand(stdout, stderr, "harness propose", err, false)
	}
	fmt.Fprintf(stdout, "proposed harness candidate %s from pressure %s\n", candidate.ID, pressureID)
	return exitOK
}

func runHarnessValidate(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, evidence, ok := parseHarnessCandidateIDDirEvidence(args, stderr)
	if !ok {
		return exitUsageError
	}
	candidate, err := harness.ValidateCandidate(dir, id, evidence)
	if err != nil {
		return failCommand(stdout, stderr, "harness validate", err, false)
	}
	fmt.Fprintf(stdout, "validated harness candidate %s to %s\n", candidate.ID, candidate.Status)
	return exitOK
}

func runHarnessAccept(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseHarnessCandidateIDAndDir(args, stderr, "accept")
	if !ok {
		return exitUsageError
	}
	candidate, err := harness.AcceptCandidate(dir, id)
	if err != nil {
		return failCommand(stdout, stderr, "harness accept", err, false)
	}
	fmt.Fprintf(stdout, "accepted harness candidate %s as %s\n", candidate.ID, candidate.Status)
	return exitOK
}

func runHarnessRetire(args []string, stdout io.Writer, stderr io.Writer) int {
	id, dir, ok := parseHarnessCandidateIDAndDir(args, stderr, "retire")
	if !ok {
		return exitUsageError
	}
	candidate, err := harness.RetireCandidate(dir, id)
	if err != nil {
		return failCommand(stdout, stderr, "harness retire", err, false)
	}
	fmt.Fprintf(stdout, "retired harness candidate %s\n", candidate.ID)
	return exitOK
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
	fmt.Fprintf(w, "usage: %[1]s harness plan --dir <path> [--json]\n       %[1]s harness candidates [--dir <path>] [--json]\n       %[1]s harness propose --from-pressure <id> [--dir <path>]\n       %[1]s harness validate <candidate-id> --evidence <path> [--dir <path>]\n       %[1]s harness accept <candidate-id> [--dir <path>]\n       %[1]s harness retire <candidate-id> [--dir <path>]\n", commandName())
}

func runInit(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		printInitUsage(stdout)
		return exitOK
	}
	dir := "."
	readinessProfile := profile.Default
	productType := contract.DefaultProductType
	var surfaces []string
	interactionMode := contract.DefaultInteractionMode
	forceInteractive := false
	assumeYes := false
	usedDirFlag := false
	usedPositionalTarget := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return exitUsageError
			}
			dir = args[i+1]
			usedDirFlag = true
			i++
		case "--profile":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --profile")
				return exitUsageError
			}
			readinessProfile = args[i+1]
			if err := profile.Validate(readinessProfile); err != nil {
				fmt.Fprintf(stderr, "invalid --profile value: %s (valid: %s)\n", readinessProfile, profile.NamesText())
				return exitUsageError
			}
			i++
		case "--product-type":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --product-type")
				return exitUsageError
			}
			productType = args[i+1]
			if err := contract.ValidateProductType(productType); err != nil {
				fmt.Fprintf(stderr, "invalid --product-type value: %v\n", err)
				return exitUsageError
			}
			i++
		case "--surface":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --surface")
				return exitUsageError
			}
			surfaces = append(surfaces, args[i+1])
			if err := contract.ValidateDeliverySurfaces(surfaces); err != nil {
				fmt.Fprintf(stderr, "invalid --surface value: %v\n", err)
				return exitUsageError
			}
			i++
		case "--interaction-mode":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --interaction-mode")
				return exitUsageError
			}
			interactionMode = args[i+1]
			if err := contract.ValidateInteractionMode(interactionMode); err != nil {
				fmt.Fprintf(stderr, "invalid --interaction-mode value: %v\n", err)
				return exitUsageError
			}
			i++
		case "--interactive":
			forceInteractive = true
		case "--yes":
			assumeYes = true
		default:
			if strings.HasPrefix(args[i], "--") {
				fmt.Fprintf(stderr, "unknown init option: %s\n", args[i])
				return exitUsageError
			}
			if dir != "." {
				fmt.Fprintf(stderr, "unexpected init argument: %s\n", args[i])
				return exitUsageError
			}
			dir = args[i]
			usedPositionalTarget = true
		}
	}

	root := filepath.Clean(dir)
	inputIsTerminal := initInputIsTerminal()
	if lockExists(root) {
		fmt.Fprintf(stdout, "warning: %s already exists; this project is already locked.\n", filepath.Join(".ni", "plan.lock.json"))
		fmt.Fprintf(stdout, "Use %s status --proof --next-questions, then the amend/relock flow for locked planning changes.\n", commandName())
		fmt.Fprintf(stdout, "No files changed by %s init; the lockfile was not modified.\n", commandName())
		return exitOK
	}

	requiredExisting := existingInitPaths(root)
	if len(requiredExisting) > 0 {
		fmt.Fprintf(stdout, "existing Namba Intent planning files found; %s init will not overwrite them.\n", commandName())
		if inputIsTerminal && !assumeYes && (forceInteractive || (usedPositionalTarget && !usedDirFlag)) {
			choice, err := runExistingInitTUI(root, requiredExisting, stdout)
			if err != nil {
				return failCommand(stdout, stderr, "init", err, false)
			}
			switch choice {
			case initui.ExistingChoiceKeep:
				fmt.Fprintln(stdout, "kept existing files; no files changed.")
				return exitOK
			case initui.ExistingChoiceAbort:
				fmt.Fprintln(stdout, "aborted init; no files changed.")
				return exitOK
			default:
				fmt.Fprintln(stdout, "adding missing files only.")
			}
		} else if forceInteractive && !assumeYes {
			choice, err := promptExistingInitChoice(initInput, stdout)
			if err != nil {
				fmt.Fprintln(stderr, err)
				return exitUsageError
			}
			switch choice {
			case "keep":
				fmt.Fprintln(stdout, "kept existing files; no files changed.")
				return exitOK
			case "abort":
				fmt.Fprintln(stdout, "aborted init; no files changed.")
				return exitOK
			}
		} else {
			fmt.Fprintln(stdout, "adding missing files only.")
		}
	}

	intent := docstore.GuidedIntent{}
	guided := false
	shouldTUI := (forceInteractive || (usedPositionalTarget && !usedDirFlag)) && !assumeYes && len(requiredExisting) == 0 && inputIsTerminal
	shouldPrompt := forceInteractive && !assumeYes && len(requiredExisting) == 0 && !shouldTUI
	if shouldTUI {
		result, err := runGuidedInitTUI(root, stdout)
		if err != nil {
			return failCommand(stdout, stderr, "init", err, false)
		}
		if result.Canceled || !result.Confirmed {
			fmt.Fprintln(stdout, "canceled init; no files written.")
			printInitNextCommands(stdout)
			return exitOK
		}
		intent = result.Intent
		guided = true
	} else if shouldPrompt {
		reader := bufio.NewReader(initInput)
		answers, err := promptGuidedIntent(reader, stdout, filepath.Base(root))
		if err != nil {
			fmt.Fprintln(stderr, err)
			return exitUsageError
		}
		if !confirmGuidedIntent(reader, stdout, answers) {
			fmt.Fprintln(stdout, "aborted init; no files changed.")
			return exitOK
		}
		intent = answers
		guided = true
	}

	result, err := docstore.InitWithOptions(dir, docstore.InitOptions{
		ReadinessProfile: readinessProfile,
		ProductType:      productType,
		DeliverySurfaces: surfaces,
		InteractionMode:  interactionMode,
		Intent:           intent,
	})
	if err != nil {
		return failCommand(stdout, stderr, "init", err, false)
	}
	if result.Locked {
		fmt.Fprintf(stdout, "warning: %s already exists; this project is already locked.\n", filepath.Join(".ni", "plan.lock.json"))
		fmt.Fprintf(stdout, "Use %s status --proof --next-questions, then the amend/relock flow for locked planning changes.\n", commandName())
		fmt.Fprintf(stdout, "No files changed by %s init; the lockfile was not modified.\n", commandName())
		return exitOK
	}

	if len(surfaces) == 0 {
		surfaces = contract.DefaultDeliverySurfaces(productType)
	}
	printInitSummary(stdout, result, readinessProfile, productType, surfaces, interactionMode)
	if guided {
		fmt.Fprintf(stdout, "guided init wrote initial intent draft; run %s status --proof --next-questions next.\n", commandName())
	} else {
		fmt.Fprintf(stdout, "next: use model-user planning conversation, then run %s status --proof --next-questions.\n", commandName())
	}
	printInitNextCommands(stdout)
	return 0
}

func printInitUsage(w io.Writer) {
	fmt.Fprintf(w, `usage: %[1]s init [.] [--dir <path>] [--interactive] [--yes] [--profile concept|prototype|mvp|beta|production] [--product-type <type>] [--surface <surface>] [--interaction-mode <mode>]

Create the initial planning docs and .ni skeleton for a Namba Intent workspace.

Options:
  --dir <path>                 Target workspace directory.
  --interactive                Prefer the guided terminal TUI when stdin is a terminal.
  --yes                        Use non-interactive defaults and write without prompting.
  --profile <name>             Readiness profile: concept, prototype, mvp, beta, or production.
  --product-type <type>        Product type recorded in the intent contract.
  --surface <surface>          Delivery surface; repeatable.
  --interaction-mode <mode>    Interaction mode recorded in the intent contract.

Boundary:
  init drafts planning state only. It does not lock the plan, execute agents,
  run shell commands, create queues, automate PRs, or release downstream work.

Next:
  %[1]s status --proof --next-questions
`, commandName())
}

func runGuidedInitTUI(root string, stdout io.Writer) (initui.Result, error) {
	return initui.Run(initui.Config{
		Dir:         root,
		CommandName: commandName(),
		DefaultName: filepath.Base(root),
		Input:       initInput,
		Output:      stdout,
	})
}

func runExistingInitTUI(root string, existing []string, stdout io.Writer) (initui.ExistingChoice, error) {
	result, err := initui.Run(initui.Config{
		Dir:           root,
		CommandName:   commandName(),
		DefaultName:   filepath.Base(root),
		ExistingFiles: existing,
		Input:         initInput,
		Output:        stdout,
	})
	if err != nil {
		return "", err
	}
	return result.Choice, nil
}

func printInitSummary(stdout io.Writer, result docstore.Result, readinessProfile string, productType string, surfaces []string, interactionMode string) {
	fmt.Fprintf(stdout, "initialized Namba Intent planning workspace at %s\n", result.Root)
	fmt.Fprintf(stdout, "target directory: %s\n", result.Root)
	fmt.Fprintf(stdout, "readiness profile: %s\n", readinessProfile)
	fmt.Fprintf(stdout, "product type: %s\n", productType)
	fmt.Fprintf(stdout, "delivery surfaces: %s\n", strings.Join(surfaces, ", "))
	fmt.Fprintf(stdout, "interaction mode: %s\n", interactionMode)
	for _, path := range result.Created {
		fmt.Fprintf(stdout, "created %s\n", path)
	}
	for _, path := range result.Existing {
		fmt.Fprintf(stdout, "exists %s\n", path)
		fmt.Fprintf(stdout, "unchanged %s\n", path)
	}
	if len(result.Created) == 0 {
		fmt.Fprintln(stdout, "created files: none")
	}
	if len(result.Existing) == 0 {
		fmt.Fprintln(stdout, "skipped files: none")
		fmt.Fprintln(stdout, "unchanged files: none")
	}
}

func printInitNextCommands(stdout io.Writer) {
	fmt.Fprintln(stdout, "next suggested commands:")
	fmt.Fprintf(stdout, "- %s status --proof --next-questions\n", commandName())
	fmt.Fprintf(stdout, "- %s end\n", commandName())
	fmt.Fprintf(stdout, "- %s run --max-chars 4000\n", commandName())
}

func lockExists(root string) bool {
	_, err := os.Stat(filepath.Join(root, ".ni", "plan.lock.json"))
	return err == nil
}

func existingInitPaths(root string) []string {
	var existing []string
	for _, path := range docstore.RequiredPaths() {
		if _, err := os.Stat(filepath.Join(root, path)); err == nil {
			existing = append(existing, path)
		}
	}
	return existing
}

func promptExistingInitChoice(input io.Reader, output io.Writer) (string, error) {
	reader := bufio.NewReader(input)
	for {
		fmt.Fprint(output, "Choose: [m] add missing files only, [k] keep existing and exit, [a] abort: ")
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			return "", fmt.Errorf("init prompt canceled before choosing existing-file handling")
		}
		switch strings.ToLower(strings.TrimSpace(line)) {
		case "", "m", "missing":
			return "missing", nil
		case "k", "keep":
			return "keep", nil
		case "a", "abort":
			return "abort", nil
		}
		fmt.Fprintln(output, "Please enter m, k, or a.")
	}
}

func promptGuidedIntent(reader *bufio.Reader, output io.Writer, defaultName string) (docstore.GuidedIntent, error) {
	ask := func(label, fallback string) (string, error) {
		if fallback != "" {
			fmt.Fprintf(output, "%s [%s]: ", label, fallback)
		} else {
			fmt.Fprintf(output, "%s: ", label)
		}
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			return "", fmt.Errorf("init prompt canceled at %q", label)
		}
		value := strings.TrimSpace(line)
		if value == "" {
			return fallback, nil
		}
		return value, nil
	}

	fmt.Fprintf(output, "Guided %s init: capture just enough project intent to start planning.\n", commandName())
	fmt.Fprintln(output, "This does not run agents, execute shell commands, or decide readiness.")
	projectName, err := ask("Project name", defaultName)
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	projectGoal, err := ask("One-sentence project goal", "")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	targetUsers, err := ask("Target users / audience", "")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	downstreamTask, err := ask("What downstream agent should eventually do", "")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	constraints, err := ask("Constraints / non-goals", "Do not execute downstream work before the plan is locked.")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	success, err := ask("Success criteria", "")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	blockers, err := ask("Known blockers or open questions", "")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	deferrals, err := ask("Deferrals, if any", "None recorded yet.")
	if err != nil {
		return docstore.GuidedIntent{}, err
	}
	return docstore.GuidedIntent{
		ProjectName:         projectName,
		ProjectGoal:         projectGoal,
		TargetUsers:         targetUsers,
		DownstreamAgentTask: downstreamTask,
		ConstraintsNonGoals: constraints,
		SuccessCriteria:     success,
		KnownBlockers:       blockers,
		Deferrals:           deferrals,
	}, nil
}

func confirmGuidedIntent(reader *bufio.Reader, output io.Writer, intent docstore.GuidedIntent) bool {
	fmt.Fprintln(output, "\nInit summary:")
	fmt.Fprintf(output, "- project: %s\n", intent.ProjectName)
	fmt.Fprintf(output, "- goal: %s\n", intent.ProjectGoal)
	fmt.Fprintf(output, "- audience: %s\n", intent.TargetUsers)
	fmt.Fprintf(output, "- downstream task: %s\n", intent.DownstreamAgentTask)
	fmt.Fprintf(output, "- success criteria: %s\n", intent.SuccessCriteria)
	for {
		fmt.Fprint(output, "Write initial intent artifacts? [y/N]: ")
		line, err := reader.ReadString('\n')
		if err != nil && len(line) == 0 {
			return false
		}
		switch strings.ToLower(strings.TrimSpace(line)) {
		case "y", "yes":
			return true
		case "", "n", "no":
			return false
		}
		fmt.Fprintln(output, "Please enter y or n.")
	}
}

func runEnd(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return exitUsageError
			}
			dir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown end option: %s\n", args[i])
			return exitUsageError
		}
	}

	existingLock, err := lock.CheckExisting(dir)
	if err != nil {
		return failCommand(stdout, stderr, "end", err, false)
	}
	if existingLock.Exists {
		if existingLock.Stale {
			fmt.Fprintf(stdout, "Existing lock is stale; %s end is the CLI-authoritative relock step after changed intent has been reviewed.\n", commandName())
		} else {
			fmt.Fprintf(stdout, "Existing lock is current; %s end will refresh the lock through the CLI readiness flow.\n", commandName())
		}
	}

	lockfile, err := lock.Create(dir)
	if err != nil {
		return failCommand(stdout, stderr, "end", err, false)
	}
	fmt.Fprintf(stdout, "locked plan at %s\n", lockfile.Path)
	fmt.Fprintf(stdout, "status %s\n", lockfile.Readiness.Status)
	return exitOK
}

func runRelock(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --dir")
				return exitUsageError
			}
			dir = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown relock option: %s\n", args[i])
			return exitUsageError
		}
	}

	status := readiness.Evaluate(dir)
	if status.Status == readiness.StatusBlocked {
		return failCommand(stdout, stderr, "relock", fmt.Errorf("readiness is BLOCKED; refusing to relock"), false)
	}

	currentHash, err := lock.CurrentLockHash(dir)
	if err != nil {
		return failCommand(stdout, stderr, "relock", err, false)
	}
	verification, err := lock.Verify(dir)
	if err != nil {
		return failCommand(stdout, stderr, "relock", err, false)
	}
	if len(verification.Mismatches) > 0 {
		ok, err := amendment.HasAppliedForLock(dir, currentHash)
		if err != nil {
			return failCommand(stdout, stderr, "relock", err, false)
		}
		if !ok {
			return failCommand(stdout, stderr, "relock", fmt.Errorf("BLOCKED: lock hash mismatch for %s without an applied amendment", verification.Mismatches[0].Path), false)
		}
	}

	now := time.Now().UTC()
	previous, err := lock.ArchiveCurrentAt(dir, now)
	if err != nil {
		return failCommand(stdout, stderr, "relock", err, false)
	}
	lockfile, err := lock.CreateAtWithPrevious(dir, now, &previous)
	if err != nil {
		return failCommand(stdout, stderr, "relock", err, false)
	}
	fmt.Fprintf(stdout, "relocked plan at %s\n", lockfile.Path)
	fmt.Fprintf(stdout, "previous lock archived at %s\n", filepath.Join(filepath.Clean(dir), previous.Path))
	fmt.Fprintf(stdout, "status %s\n", lockfile.Readiness.Status)
	return exitOK
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
				return exitUsageError
			}
			dir = args[i+1]
			i++
		case "--out":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --out")
				return exitUsageError
			}
			out = args[i+1]
			i++
		case "--max-chars":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --max-chars")
				return exitUsageError
			}
			value, err := strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Fprintf(stderr, "invalid --max-chars value: %s\n", args[i+1])
				return exitUsageError
			}
			maxChars = value
			i++
		case "--target":
			if i+1 >= len(args) {
				fmt.Fprintln(stderr, "missing value for --target")
				return exitUsageError
			}
			targetName = args[i+1]
			i++
		default:
			fmt.Fprintf(stderr, "unknown run option: %s\n", args[i])
			return exitUsageError
		}
	}

	result, err := prompt.Compile(prompt.Options{Dir: dir, Out: out, MaxChars: maxChars, Target: targetName})
	if err != nil {
		return failCommand(stdout, stderr, "run", err, false)
	}
	if out != "" {
		fmt.Fprintf(stdout, "compiled prompt at %s\n", result.Out)
		return exitOK
	}
	fmt.Fprint(stdout, result.Prompt)
	return exitOK
}

func runTargets(args []string, stdout io.Writer, stderr io.Writer) int {
	jsonOutput := jsonRequested(args)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json":
			jsonOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown targets option: %s", args[i]), jsonOutput)
		}
	}

	items := target.List()
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(items); err != nil {
			return failCommand(stdout, stderr, "targets", err, jsonOutput)
		}
		return exitOK
	}
	for _, item := range items {
		fmt.Fprintf(stdout, "%s\t%s\t%s\n", item.Name, item.Artifact, item.Description)
	}
	return exitOK
}

func runStatus(args []string, stdout io.Writer, stderr io.Writer) int {
	dir := "."
	jsonOutput := jsonRequested(args)
	nextQuestions := false
	proofOutput := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 >= len(args) {
				return failStructured(stdout, stderr, usageErrorf("missing value for --dir"), jsonOutput)
			}
			dir = args[i+1]
			i++
		case "--json":
			jsonOutput = true
		case "--next-questions":
			nextQuestions = true
		case "--proof":
			proofOutput = true
		default:
			return failStructured(stdout, stderr, usageErrorf("unknown status option: %s", args[i]), jsonOutput)
		}
	}

	result := readiness.Evaluate(dir)
	if err := invalidContractFromStatus(result); err != nil {
		return failCommand(stdout, stderr, "status", err, jsonOutput)
	}
	if proofOutput {
		proof := readiness.Proof(result)
		result.Proof = &proof
	}
	if nextQuestions {
		questions := readiness.NextQuestions(result)
		result.NextQuestions = &questions
	}
	existingLock, lockErr := lock.CheckExisting(dir)
	staleLockWarning := lockErr == nil && existingLock.Stale
	if jsonOutput {
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(result); err != nil {
			return failCommand(stdout, stderr, "status", err, jsonOutput)
		}
		return exitOK
	}

	if proofOutput {
		if staleLockWarning {
			appendStaleLockProof(&result, existingLock)
		}
		printStatusProof(stdout, result, nextQuestions)
		return exitOK
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
	if staleLockWarning {
		fmt.Fprintln(stdout, staleLockWarningText(existingLock))
	}
	if nextQuestions {
		printNextQuestions(stdout, result)
	}
	return exitOK
}

func printStatusProof(stdout io.Writer, result readiness.Result, nextQuestions bool) {
	fmt.Fprintf(stdout, "NI Intent Readiness: %s\n\n", result.Status)
	blockers, deferrals, warnings := groupProofItems(result)
	printProofItems(stdout, "Blockers", blockers, true)
	fmt.Fprintln(stdout)
	printProofItems(stdout, "Deferrals", deferrals, false)
	fmt.Fprintln(stdout)
	printProofItems(stdout, "Warnings", warnings, false)
	fmt.Fprintln(stdout)
	printPassedChecks(stdout, result)

	switch result.Status {
	case readiness.StatusBlocked:
		fmt.Fprintln(stdout, "\nExecution must not start.")
	case readiness.StatusReadyWithDeferrals:
		fmt.Fprintln(stdout, "\nExecution may proceed only after lock; deferrals remain explicit.")
	default:
		fmt.Fprintln(stdout, "\nExecution may proceed only after lock.")
	}

	if nextQuestions && result.NextQuestions != nil && len(*result.NextQuestions) > 0 {
		fmt.Fprintln(stdout, "\nNext questions:")
		printGroupedQuestions(stdout, *result.NextQuestions)
		if omitted := readiness.OmittedNextQuestionCount(result); omitted > 0 {
			fmt.Fprintf(stdout, "\n%d additional lower-priority question(s) remain after these top %d.\n", omitted, len(*result.NextQuestions))
		}
	}
}

func printNextQuestions(stdout io.Writer, result readiness.Result) {
	if result.NextQuestions == nil || len(*result.NextQuestions) == 0 {
		return
	}
	fmt.Fprintln(stdout, "Next questions:")
	printGroupedQuestions(stdout, *result.NextQuestions)
	if omitted := readiness.OmittedNextQuestionCount(result); omitted > 0 {
		fmt.Fprintf(stdout, "\n%d additional lower-priority question(s) remain after these top %d.\n", omitted, len(*result.NextQuestions))
	}
}

func printGroupedQuestions(stdout io.Writer, questions []readiness.NextQuestion) {
	currentGroup := ""
	indexInGroup := 0
	for _, question := range questions {
		group := question.Group
		if group == "" {
			group = "Planning repairs"
		}
		if group != currentGroup {
			if currentGroup != "" {
				fmt.Fprintln(stdout)
			}
			fmt.Fprintf(stdout, "%s:\n", group)
			currentGroup = group
			indexInGroup = 0
		}
		indexInGroup++
		fmt.Fprintf(stdout, "%d. %s: %s\n", indexInGroup, questionLabel(question), question.Question)
		if question.Location != "" {
			fmt.Fprintf(stdout, "   Location: %s\n", question.Location)
		}
		if question.AnswerShape != "" {
			fmt.Fprintf(stdout, "   Answer shape: %s\n", question.AnswerShape)
		}
	}
}

func questionLabel(question readiness.NextQuestion) string {
	if question.Group == "First-run card" {
		return question.RuleID
	}
	if len(question.References) > 0 {
		return question.References[0]
	}
	return question.RuleID
}

func groupProofItems(result readiness.Result) ([]readiness.ProofItem, []readiness.ProofItem, []readiness.ProofItem) {
	var blockers []readiness.ProofItem
	var deferrals []readiness.ProofItem
	var warnings []readiness.ProofItem
	if result.Proof == nil {
		return blockers, deferrals, warnings
	}
	for _, item := range *result.Proof {
		switch item.Severity {
		case "blocker":
			blockers = append(blockers, item)
		case "deferral":
			if item.RuleID == "D001" {
				warnings = append(warnings, item)
			} else {
				deferrals = append(deferrals, item)
			}
		case "warning":
			warnings = append(warnings, item)
		}
	}
	return blockers, deferrals, warnings
}

func appendStaleLockProof(result *readiness.Result, existingLock lock.ExistingLockState) {
	if result.Proof == nil {
		return
	}
	proof := append(*result.Proof, readiness.ProofItem{
		RuleID:   lock.StaleDiagnosticID,
		Severity: "warning",
		Message:  staleLockWarningText(existingLock),
	})
	result.Proof = &proof
}

func staleLockWarningText(existingLock lock.ExistingLockState) string {
	message := lock.StaleStatusWarning
	if len(existingLock.Verification.Mismatches) > 0 {
		message += " First mismatch: " + existingLock.Verification.Mismatches[0].Path + "."
	}
	return message + " " + lock.StaleStatusRecovery
}

func printProofItems(stdout io.Writer, title string, items []readiness.ProofItem, includeNext bool) {
	fmt.Fprintf(stdout, "%s:\n", title)
	if len(items) == 0 {
		fmt.Fprintln(stdout, "- None.")
		return
	}
	for _, item := range items {
		fmt.Fprintf(stdout, "- %s\n", item.Message)
		if item.SyncDiagnostic != nil {
			fmt.Fprintf(stdout, "  ID: %s\n", item.SyncDiagnostic.ID)
			fmt.Fprintf(stdout, "  Location: %s\n", item.SyncDiagnostic.Location)
			fmt.Fprintf(stdout, "  Problem: %s\n", item.SyncDiagnostic.Problem)
			fmt.Fprintf(stdout, "  Why it matters: %s\n", item.SyncDiagnostic.WhyItMatters)
			fmt.Fprintf(stdout, "  Suggested repair: %s\n", item.SyncDiagnostic.SuggestedRepair)
			fmt.Fprintf(stdout, "  Blocks ni-end: %t\n", item.SyncDiagnostic.BlocksEnd)
			continue
		}
		fmt.Fprintf(stdout, "  Why it matters: %s\n", proofWhy(item.RuleID))
		if includeNext {
			fmt.Fprintf(stdout, "  Next: %s\n", proofNext(item.RuleID))
		}
	}
}

func printPassedChecks(stdout io.Writer, result readiness.Result) {
	fmt.Fprintln(stdout, "Passed checks:")
	for _, message := range passedCheckMessages(result) {
		fmt.Fprintf(stdout, "- %s\n", message)
	}
}

func passedCheckMessages(result readiness.Result) []string {
	hasIssue := map[string]bool{}
	for _, issue := range result.Issues {
		hasIssue[issue.RuleID] = true
	}

	var passed []string
	if !hasIssue["R001"] {
		passed = append(passed, "Required docs exist.")
	}
	if !hasIssue["R002"] {
		passed = append(passed, "Contract JSON is valid.")
	}
	if !hasIssue["R011"] {
		passed = append(passed, "Readiness profile definitions are valid.")
	}
	if !hasIssue["R002"] && !hasIssue["R003"] && !hasIssue["R004"] && !hasIssue["R005"] && !hasIssue["R007"] {
		passed = append(passed, "Capability and evaluation traceability rules passed.")
	}
	if !hasIssue["R002"] && !hasIssue["R006"] {
		passed = append(passed, "High-severity risks have mitigation.")
	}
	if !hasIssue["R002"] && !hasIssue["R008"] && !hasIssue["R013"] {
		passed = append(passed, "Decision statuses are valid and accepted decisions do not conflict.")
	}
	if !hasIssue["R002"] && !hasIssue["R009"] {
		passed = append(passed, "No blocker open questions are present.")
	}
	if !hasIssue["R002"] && !hasIssue["R010"] {
		passed = append(passed, "At least one non-goal is recorded.")
	}
	if !hasIssue["R002"] && !hasIssue["R012"] {
		passed = append(passed, "Docs and contract are synchronized.")
	}
	if len(passed) == 0 {
		return []string{"No checks passed before the blocking error."}
	}
	return passed
}

func proofWhy(ruleID string) string {
	switch ruleID {
	case "R001":
		return "the kernel cannot hash or trust a plan whose required source docs are absent."
	case "R002":
		return "the readiness gate cannot evaluate intent without a parseable contract."
	case "R003":
		return "a plan with no capability has no deliverable behavior to verify."
	case "R004":
		return "Namba Intent cannot prove this capability is verifiable."
	case "R005":
		return "evidence is not trustworthy unless the evaluation method is explicit."
	case "R006":
		return "high-severity risks can invalidate downstream work unless mitigation is explicit."
	case "R007":
		return "capabilities must trace to requirements or artifacts before downstream actors can rely on them."
	case "R008":
		return "decision status controls whether downstream actors may depend on the decision."
	case "R009":
		return "open blocker questions mean required intent is still unresolved."
	case "R010":
		return "non-goals keep downstream work inside the accepted boundary."
	case "R011":
		return "readiness severity rules must be trustworthy before status can be trusted."
	case "R012":
		return "docs and contract are both lock sources, so drift makes the plan ambiguous."
	case "R013":
		return "conflicting accepted decisions give downstream actors incompatible instructions."
	case "R014":
		return "Namba Intent cannot lock intent until it knows what reality the project is meant to change."
	case "R015":
		return "Namba Intent cannot judge readiness without knowing who uses or operates the product and what successful use looks like for them."
	case "R016":
		return "downstream handoff depends on knowing whether the product is delivered as a CLI, web app, conversation, document, workflow, research protocol, human service, or another surface."
	case "D001":
		return "downstream work must avoid depending on this decision."
	case "D002":
		return "the question is non-blocking but still unresolved intent."
	case lock.StaleDiagnosticID:
		return "the current planning contract and the existing locked plan no longer match."
	default:
		return "this deterministic readiness rule affects whether the plan can be trusted."
	}
}

func proofNext(ruleID string) string {
	switch ruleID {
	case "R001":
		return "create the missing required planning doc or restore it from the template."
	case "R002":
		return "fix .ni/contract.json so it is valid contract JSON."
	case "R003":
		return "define at least one accepted capability."
	case "R004":
		return "define evidence and link an evaluation."
	case "R005":
		return "add a deterministic Method to the evaluation."
	case "R006":
		return "add mitigation, an owner, or an explicit accepted-risk decision."
	case "R007":
		return "link the capability to at least one requirement or artifact."
	case "R008":
		return "set the decision status to accepted, deferred, rejected, or not_applicable."
	case "R009":
		return "answer or defer the blocker question, or keep it blocking with an explicit reason."
	case "R010":
		return "add at least one explicit non-goal."
	case "R011":
		return "fix .ni/readiness.profiles.json so every required profile and rule severity is valid."
	case "R012":
		return "update docs/plan/** and .ni/contract.json together so the referenced record matches."
	case "R013":
		return "revise, reject, or split one conflicting accepted decision."
	case "R014":
		return "describe the project in one or two sentences: what should change, for whom, and why it matters."
	case "R015":
		return "list the primary actors and the outcome each one expects."
	case "R016":
		return "choose the likely delivery surface, or mark it deferred with an explicit reason."
	default:
		return "update planning docs and .ni/contract.json together to resolve this rule."
	}
}

func invalidContractFromStatus(result readiness.Result) error {
	for _, issue := range result.Issues {
		if issue.RuleID == "R002" {
			return commandError{Code: "invalid_contract", ExitCode: exitInvalidContract, Message: issue.Message}
		}
	}
	return nil
}

func printHelp(w io.Writer) {
	name := commandName()
	fmt.Fprintf(w, `Namba Intent is a Project Intent Compiler for AI Agents.

Don't run the agent yet. Compile the intent first.

Usage:
  %[1]s --help
  %[1]s amend create --title <title> [--dir <path>]
  %[1]s amend list [--dir <path>]
  %[1]s amend show <id> [--dir <path>]
  %[1]s amend apply <id> [--dir <path>]
  %[1]s conflicts --base <path-or-lock> --head <path-or-lock> [--json]
  %[1]s diff --base <path-or-lock> --head <path-or-lock> [--json]
  %[1]s end --dir <path>
  %[1]s export --target hyper-run|namba-ai|ouroboros|spec-kit --out <dir> [--dir <path>]
  %[1]s feedback add --file <path> [--dir <path>]
  %[1]s feedback list [--dir <path>] [--json]
  %[1]s graph --dir <path> [--json]
  %[1]s harness plan --dir <path> [--json]
  %[1]s harness candidates [--dir <path>] [--json]
  %[1]s harness propose --from-pressure <id> [--dir <path>]
  %[1]s harness validate <candidate-id> --evidence <path> [--dir <path>]
  %[1]s harness accept <candidate-id> [--dir <path>]
  %[1]s harness retire <candidate-id> [--dir <path>]
  %[1]s init [.] [--dir <path>] [--interactive] [--yes] [--profile concept|prototype|mvp|beta|production] [--product-type <type>] [--surface <surface>] [--interaction-mode <mode>]
  %[1]s pressure status [--dir <path>] [--json]
  %[1]s pressure promote <id> [--dir <path>]
  %[1]s pressure retire <id> [--dir <path>]
  %[1]s relock --dir <path>
  %[1]s run --dir <path> [--target <target>] [--out <path>] [--max-chars N]
  %[1]s status --dir <path> [--json] [--proof] [--next-questions]
  %[1]s targets [--json]
  %[1]s version

Commands:
  amend   Create, inspect, and apply explicit contract amendments.
  conflicts Detect semantic planning conflicts between two contracts or locked plans.
  diff     Show contract-level changes between two contracts or locked plans.
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
  version  Print the Namba Intent version.

Namba Intent keeps .ni/ for compatibility. run compiles a bounded handoff
prompt; it does not execute agents, shell commands, queues, PR automation, or
release automation.
`, name)
}

func printAmendUsage(w io.Writer) {
	fmt.Fprintf(w, "usage: %[1]s amend create --title <title> [--dir <path>]\n       %[1]s amend list [--dir <path>]\n       %[1]s amend show <id> [--dir <path>]\n       %[1]s amend apply <id> [--dir <path>]\n", commandName())
}
