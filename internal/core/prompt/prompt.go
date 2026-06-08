package prompt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
	"ni/internal/core/target"
)

const DefaultMaxChars = 4000

type Options struct {
	Dir      string
	Out      string
	MaxChars int
	Target   string
}

type Result struct {
	Prompt string
	Out    string
}

func Compile(opts Options) (Result, error) {
	dir := opts.Dir
	if dir == "" {
		dir = "."
	}
	maxChars := opts.MaxChars
	if maxChars == 0 {
		maxChars = DefaultMaxChars
	}
	if maxChars < 1 {
		return Result{}, fmt.Errorf("max-chars must be greater than 0")
	}

	selectedTarget, err := target.Lookup(opts.Target)
	if err != nil {
		return Result{}, err
	}

	verification, err := lock.Verify(dir)
	if err != nil {
		return Result{}, err
	}
	if len(verification.Mismatches) > 0 {
		return Result{}, fmt.Errorf("BLOCKED: lock hash mismatch for %s. %s", verification.Mismatches[0].Path, lock.StaleRunRecovery)
	}

	c, err := contract.LoadFile(filepath.Join(filepath.Clean(dir), ".ni", "contract.json"))
	if err != nil {
		return Result{}, err
	}

	text := buildPrompt(c, verification.Lockfile, selectedTarget)
	text = limitChars(text, maxChars)
	if opts.Out != "" {
		if err := os.MkdirAll(filepath.Dir(opts.Out), 0o755); err != nil {
			return Result{}, err
		}
		if err := os.WriteFile(opts.Out, []byte(text), 0o644); err != nil {
			return Result{}, err
		}
	}
	return Result{Prompt: text, Out: opts.Out}, nil
}

func buildPrompt(c contract.Contract, l lock.Lockfile, t target.Target) string {
	tmpl := templateFor(t.Name)
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n\n", tmpl.Title)
	fmt.Fprintf(&b, "Goal: %s\n\n", tmpl.Goal)
	fmt.Fprintf(&b, "Project: %s - %s\n", c.Project.Name, c.Project.Purpose)
	fmt.Fprintf(&b, "Readiness: %s\n", l.Readiness.Status)
	fmt.Fprintf(&b, "Target: %s (%s)\n", t.Name, t.Artifact)
	fmt.Fprintf(&b, "Locked at: %s\n\n", l.LockedAt)
	b.WriteString("Authoritative sources:\n")
	b.WriteString("- .ni/plan.lock.json is authoritative for lock state, hashes, and source-of-truth order.\n")
	b.WriteString("- .ni/contract.json carries accepted CAP/REQ/EVAL/RISK IDs and acceptance criteria.\n")
	b.WriteString("- docs/plan/ contains locked planning context; use only when hashes match.\n")
	b.WriteString("- .ni/session.json is a planning aid below locked docs; it must not override contract or docs.\n")
	fmt.Fprintf(&b, "Source of truth: %s\n\n", strings.Join(l.SourceOfTruth, " > "))
	b.WriteString("Rules:\n")
	b.WriteString("- Treat .ni/plan.lock.json as authoritative over .ni/contract.json, docs/plan, session state, and chat.\n")
	b.WriteString("- Verify locked file hashes before using planning docs; on lock mismatch report BLOCKED and stop.\n")
	b.WriteString("- If there are conflicting requirements, report BLOCKED with conflicting sources or IDs and stop.\n")
	b.WriteString("- Do not weaken acceptance criteria, risk mitigations, or blocker handling to proceed.\n")
	b.WriteString("- Every work item must trace to CAP/REQ/EVAL/RISK IDs.\n")
	b.WriteString("- Keep Namba Intent at the pre-runtime boundary; this prompt is seed material, not kernel-owned execution state.\n")
	b.WriteString("- Do not make namba-intent run execute shell, Codex, adapters, queues, PR automation, or agent teams.\n\n")
	writeTemplateSection(&b, tmpl)

	writeCapabilities(&b, c.Capabilities)
	writeRequirements(&b, c.Requirements)
	writeRisks(&b, c.Risks)
	writeOpenQuestions(&b, c.OpenQuestions)

	fmt.Fprintf(&b, "Expected output: %s\n", tmpl.ExpectedOutput)
	return b.String()
}

type template struct {
	Title          string
	Goal           string
	Instructions   []string
	Process        []string
	ExpectedOutput string
}

func templateFor(name string) template {
	switch name {
	case "codex":
		return template{
			Title: "Codex target prompt",
			Goal:  "Paste this prompt into Codex to make the smallest implementation move allowed by the locked NI plan.",
			Instructions: []string{
				"Use the current worktree as evidence; preserve user changes and stay within the locked contract.",
				"Define validation before editing, then implement only the selected work packet.",
				"Do not ask Namba Intent to invoke Codex automatically; namba-intent run only compiled this prompt.",
			},
			Process: []string{
				"Read .ni/plan.lock.json, .ni/contract.json, and relevant docs/plan files.",
				"Choose the smallest coherent packet traced to CAP/REQ/EVAL/RISK IDs.",
				"Edit code or docs only as needed, run validation, and report evidence.",
			},
			ExpectedOutput: "changed files, validation commands/results, remaining risks, and BLOCKED if the lock or requirements conflict.",
		}
	case "human-team":
		return template{
			Title: "Human-team handoff",
			Goal:  "Hand this locked NI plan to a PM/dev/design/research team for coordinated implementation planning.",
			Instructions: []string{
				"PM: maintain scope against accepted capabilities, non-goals, and open blocker questions.",
				"Dev: identify implementation packets and validation mapped to CAP/REQ/EVAL/RISK IDs.",
				"Design/research: confirm user-facing assumptions, evidence needs, and unresolved risks without changing acceptance criteria.",
				"Execution responsibility stays outside ni; this is a handoff artifact, not orchestration.",
			},
			Process: []string{
				"Review .ni/plan.lock.json first, then .ni/contract.json and docs/plan.",
				"Assign next ownership only after confirming no lock mismatch or conflicting requirements.",
				"Record validation evidence and blockers for the team before implementation proceeds.",
			},
			ExpectedOutput: "team handoff with owners, next packets, validation evidence, risks, decisions needed, and BLOCKED if the lock or requirements conflict.",
		}
	case "hyper-run":
		return seedTemplate("Hyper Run", "do not call hyper run or generate .hyper/goals runtime packets from namba-intent run")
	case "namba-ai":
		return seedTemplate("Namba AI", "do not call downstream runtimes")
	case "ouroboros":
		return seedTemplate("Ouroboros", "do not call downstream runtimes")
	case "spec-kit":
		return seedTemplate("Spec Kit", "do not execute specification tooling")
	default:
		return template{
			Title: "Generic target prompt",
			Goal:  "Give any downstream actor the smallest safe implementation move from the locked NI plan.",
			Instructions: []string{
				"Use this as a generic prompt for a downstream worker, not as automatic execution.",
				"Define validation before implementation and keep work within the locked contract.",
				"Prefer the smallest coherent packet that satisfies accepted requirements.",
			},
			Process: []string{
				"Read .ni/plan.lock.json, .ni/contract.json, and relevant docs/plan files.",
				"Trace each selected task to CAP/REQ/EVAL/RISK IDs.",
				"Run validation appropriate to the packet and report evidence.",
			},
			ExpectedOutput: "selected work packet, changed files or handoff notes, validation evidence, and BLOCKED if the lock or requirements conflict.",
		}
	}
}

func seedTemplate(name string, boundary string) template {
	return template{
		Title: name + " seed prompt",
		Goal:  "Prepare downstream-compatible seed guidance from the locked NI plan.",
		Instructions: []string{
			"Produce seed guidance only; " + boundary + ".",
			"Keep runtime state outside ni and derived from the locked contract.",
		},
		Process: []string{
			"Read .ni/plan.lock.json, .ni/contract.json, and relevant docs/plan files.",
			"Trace proposed work to CAP/REQ/EVAL/RISK IDs.",
			"Report validation expectations without executing downstream tooling.",
		},
		ExpectedOutput: "seed guidance, validation expectations, and BLOCKED if the lock or requirements conflict.",
	}
}

func writeTemplateSection(b *strings.Builder, tmpl template) {
	b.WriteString("Target instructions:\n")
	for _, line := range tmpl.Instructions {
		fmt.Fprintf(b, "- %s\n", line)
	}
	b.WriteString("\nProcess:\n")
	for i, line := range tmpl.Process {
		fmt.Fprintf(b, "%d. %s\n", i+1, line)
	}
	b.WriteString("\n")
}

func writeCapabilities(b *strings.Builder, items []contract.Capability) {
	if len(items) == 0 {
		return
	}
	b.WriteString("Accepted capabilities:\n")
	count := min(len(items), 8)
	for i := 0; i < count; i++ {
		item := items[i]
		if item.Status == "accepted" {
			fmt.Fprintf(b, "- %s: %s\n", item.ID, item.Title)
		}
	}
	if len(items) > count {
		fmt.Fprintf(b, "- ... %d more capabilities omitted\n", len(items)-count)
	}
	b.WriteString("\n")
}

func writeRequirements(b *strings.Builder, items []contract.Requirement) {
	if len(items) == 0 {
		return
	}
	b.WriteString("Accepted requirements:\n")
	count := min(len(items), 8)
	for i := 0; i < count; i++ {
		item := items[i]
		if item.Status == "accepted" {
			fmt.Fprintf(b, "- %s: %s\n", item.ID, item.Title)
		}
	}
	if len(items) > count {
		fmt.Fprintf(b, "- ... %d more requirements omitted\n", len(items)-count)
	}
	b.WriteString("\n")
}

func writeRisks(b *strings.Builder, items []contract.Risk) {
	if len(items) == 0 {
		return
	}
	b.WriteString("Risks:\n")
	count := min(len(items), 6)
	for i := 0; i < count; i++ {
		item := items[i]
		fmt.Fprintf(b, "- %s (%s): %s\n", item.ID, item.Severity, item.Mitigation)
	}
	if len(items) > count {
		fmt.Fprintf(b, "- ... %d more risks omitted\n", len(items)-count)
	}
	b.WriteString("\n")
}

func writeOpenQuestions(b *strings.Builder, items []contract.OpenQuestion) {
	if len(items) == 0 {
		return
	}
	b.WriteString("Open questions:\n")
	count := min(len(items), 6)
	for i := 0; i < count; i++ {
		item := items[i]
		if item.Status != "closed" && item.Status != "resolved" {
			fmt.Fprintf(b, "- %s blocker=%t: %s\n", item.ID, item.Blocker, item.Title)
		}
	}
	if len(items) > count {
		fmt.Fprintf(b, "- ... %d more open questions omitted\n", len(items)-count)
	}
	b.WriteString("\n")
}

func limitChars(text string, maxChars int) string {
	if utf8.RuneCountInString(text) <= maxChars {
		return text
	}
	const suffix = "\n[truncated to max-chars]\n"
	if maxChars <= utf8.RuneCountInString(suffix) {
		return string([]rune(text)[:maxChars])
	}
	limit := maxChars - utf8.RuneCountInString(suffix)
	return string([]rune(text)[:limit]) + suffix
}
