package prompt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
)

const DefaultMaxChars = 4000

type Options struct {
	Dir      string
	Out      string
	MaxChars int
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

	text := buildPrompt(c, verification.Lockfile)
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

func buildPrompt(c contract.Contract, l lock.Lockfile) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Goal: continue work from locked NI plan %s.\n\n", l.LockedAt)
	fmt.Fprintf(&b, "Project: %s - %s\n", c.Project.Name, c.Project.Purpose)
	fmt.Fprintf(&b, "Readiness: %s\n", l.Readiness.Status)
	fmt.Fprintf(&b, "Source of truth: %s\n\n", strings.Join(l.SourceOfTruth, " > "))
	b.WriteString("Rules:\n")
	b.WriteString("- Treat .ni/plan.lock.json as authoritative.\n")
	b.WriteString("- Verify locked file hashes before using planning docs.\n")
	b.WriteString("- Do not weaken acceptance criteria.\n")
	b.WriteString("- Do not execute shell, Codex, adapters, queues, or automations from ni run.\n\n")

	writeCapabilities(&b, c.Capabilities)
	writeRequirements(&b, c.Requirements)
	writeRisks(&b, c.Risks)
	writeOpenQuestions(&b, c.OpenQuestions)

	b.WriteString("Expected output: make the smallest implementation move that satisfies the locked contract, then report changed files and validation evidence.\n")
	return b.String()
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
