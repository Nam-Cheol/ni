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
		return Result{}, fmt.Errorf("BLOCKED: lock hash mismatch for %s", verification.Mismatches[0].Path)
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
	var b strings.Builder
	fmt.Fprintf(&b, "Goal: continue work from locked NI plan %s.\n\n", l.LockedAt)
	fmt.Fprintf(&b, "Project: %s - %s\n", c.Project.Name, c.Project.Purpose)
	fmt.Fprintf(&b, "Readiness: %s\n", l.Readiness.Status)
	fmt.Fprintf(&b, "Target: %s (%s)\n", t.Name, t.Artifact)
	fmt.Fprintf(&b, "Source of truth: %s\n\n", strings.Join(l.SourceOfTruth, " > "))
	b.WriteString("Rules:\n")
	b.WriteString("- Treat .ni/plan.lock.json as authoritative.\n")
	b.WriteString("- Verify locked file hashes before using planning docs.\n")
	b.WriteString("- Do not weaken acceptance criteria.\n")
	b.WriteString("- Do not execute shell, Codex, adapters, queues, or automations from ni run.\n\n")
	writeTargetGuidance(&b, t)

	writeCapabilities(&b, c.Capabilities)
	writeRequirements(&b, c.Requirements)
	writeRisks(&b, c.Risks)
	writeOpenQuestions(&b, c.OpenQuestions)

	b.WriteString("Expected output: make the smallest implementation move that satisfies the locked contract, then report changed files and validation evidence.\n")
	return b.String()
}

func writeTargetGuidance(b *strings.Builder, t target.Target) {
	b.WriteString("Target guidance:\n")
	switch t.Name {
	case "codex":
		b.WriteString("- Produce a Codex-ready implementation prompt only; do not invoke codex exec.\n")
		b.WriteString("- Include scope, locked authority, and validation evidence expected from the worker.\n")
	case "human-team":
		b.WriteString("- Produce a human-team handoff with ownership-neutral next steps and validation checks.\n")
		b.WriteString("- Keep execution responsibility outside ni; this artifact is planning seed material.\n")
	case "hyper-run":
		b.WriteString("- Produce Hyper Run-compatible seed guidance only; do not call hyper run.\n")
		b.WriteString("- Do not generate .hyper/goals runtime packets from ni run.\n")
	case "namba-ai":
		b.WriteString("- Produce Namba AI seed guidance only; do not call downstream runtimes.\n")
	case "ouroboros":
		b.WriteString("- Produce Ouroboros seed guidance only; do not call downstream runtimes.\n")
	case "spec-kit":
		b.WriteString("- Produce Spec Kit export guidance only; do not execute specification tooling.\n")
	default:
		b.WriteString("- Produce a generic downstream prompt only; do not execute downstream runtimes.\n")
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
