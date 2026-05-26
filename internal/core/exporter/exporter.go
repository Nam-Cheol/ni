package exporter

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
	"ni/internal/core/target"
)

const HyperRunTarget = "hyper-run"

type Options struct {
	Dir    string
	OutDir string
	Target string
}

type Result struct {
	OutDir string
	Files  []string
}

func Export(opts Options) (Result, error) {
	root := filepath.Clean(defaultString(opts.Dir, "."))
	outDir := strings.TrimSpace(opts.OutDir)
	if outDir == "" {
		return Result{}, fmt.Errorf("missing --out")
	}

	selectedTarget, err := target.Lookup(opts.Target)
	if err != nil {
		return Result{}, err
	}
	if selectedTarget.Name != HyperRunTarget {
		return Result{}, fmt.Errorf("unsupported export target %q (valid: %s)", selectedTarget.Name, HyperRunTarget)
	}

	verification, err := lock.Verify(root)
	if err != nil {
		return Result{}, err
	}
	if len(verification.Mismatches) > 0 {
		return Result{}, fmt.Errorf("BLOCKED: lock hash mismatch for %s", verification.Mismatches[0].Path)
	}

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	if err != nil {
		return Result{}, err
	}
	lockHash, err := fileSHA256(filepath.Join(root, ".ni", "plan.lock.json"))
	if err != nil {
		return Result{}, err
	}

	docs := hyperRunDocs(c, verification.Lockfile, "sha256:"+lockHash)
	allowed := allowedDocNames(docs)
	if err := prepareOutDir(outDir, allowed); err != nil {
		return Result{}, err
	}

	files := make([]string, 0, len(docs))
	for _, doc := range docs {
		path := filepath.Join(outDir, doc.name)
		if err := os.WriteFile(path, []byte(doc.content), 0o644); err != nil {
			return Result{}, err
		}
		files = append(files, doc.name)
	}
	return Result{OutDir: outDir, Files: files}, nil
}

func prepareOutDir(outDir string, allowed map[string]struct{}) error {
	entries, err := os.ReadDir(outDir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(outDir, 0o755)
		}
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			return fmt.Errorf("out directory contains non-seed entry %q", entry.Name())
		}
		if _, ok := allowed[entry.Name()]; !ok {
			return fmt.Errorf("out directory contains non-seed entry %q", entry.Name())
		}
	}
	return nil
}

func allowedDocNames(docs []exportDoc) map[string]struct{} {
	allowed := make(map[string]struct{}, len(docs))
	for _, doc := range docs {
		allowed[doc.name] = struct{}{}
	}
	return allowed
}

type exportDoc struct {
	name    string
	content string
}

func hyperRunDocs(c contract.Contract, l lock.Lockfile, sourceLockHash string) []exportDoc {
	return []exportDoc{
		{name: "plan.md", content: renderPlan(c)},
		{name: "ni-context.md", content: renderContext(c, l, sourceLockHash)},
		{name: "readiness-expectations.md", content: renderReadiness(l)},
		{name: "evidence-requirements.md", content: renderEvidence(c)},
		{name: "first-run-focus.md", content: renderFirstRunFocus(c)},
	}
}

func renderPlan(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Plan\n\n")
	fmt.Fprintf(&b, "Project: %s\n\n", c.Project.Name)
	fmt.Fprintf(&b, "Purpose: %s\n\n", c.Project.Purpose)
	writeIDList(&b, "Accepted capabilities", acceptedCapabilities(c.Capabilities))
	writeIDList(&b, "Accepted requirements", acceptedRequirements(c.Requirements))
	writeIDList(&b, "Non-goals", nonGoals(c.NonGoals))
	return b.String()
}

func renderContext(c contract.Contract, l lock.Lockfile, sourceLockHash string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# NI Context\n\n")
	fmt.Fprintf(&b, "This is Hyper Run-compatible seed material derived from a locked NI plan.\n\n")
	fmt.Fprintf(&b, "Project ID: %s\n\n", c.Project.ID)
	fmt.Fprintf(&b, "Readiness profile: %s\n\n", c.ReadinessProfile)
	fmt.Fprintf(&b, "Locked at: %s\n\n", l.LockedAt)
	fmt.Fprintf(&b, "Source lock hash: %s\n\n", sourceLockHash)
	fmt.Fprintf(&b, "Source of truth: %s\n\n", strings.Join(l.SourceOfTruth, " > "))
	b.WriteString("Boundary rules:\n\n")
	b.WriteString("- NI exports seed documents only.\n")
	b.WriteString("- NI does not run Hyper Run or write Hyper Run goal runtime packets.\n")
	b.WriteString("- Downstream runtime state must stay outside NI and remain derived from the lock.\n")
	return b.String()
}

func renderReadiness(l lock.Lockfile) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Readiness Expectations\n\n")
	fmt.Fprintf(&b, "Locked readiness status: %s\n\n", l.Readiness.Status)
	b.WriteString("Before downstream work starts:\n\n")
	b.WriteString("- Verify `.ni/plan.lock.json` exists in the NI project.\n")
	b.WriteString("- Verify locked hashes still match `.ni/contract.json` and `docs/plan/**`.\n")
	b.WriteString("- Stop with `BLOCKED` on any lock mismatch or conflicting accepted requirement.\n")
	b.WriteString("- Keep acceptance criteria, risk mitigations, and blocker handling unchanged.\n")
	return b.String()
}

func renderEvidence(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Evidence Requirements\n\n")
	writeIDList(&b, "Evaluations", evaluations(c.Evaluations))
	writeIDList(&b, "Artifacts", artifacts(c.Artifacts))
	writeIDList(&b, "Risks and mitigations", risks(c.Risks))
	return b.String()
}

func renderFirstRunFocus(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# First Run Focus\n\n")
	focus := firstAcceptedCapability(c.Capabilities)
	if focus.ID == "" {
		b.WriteString("No accepted capability was found in the locked contract.\n")
		return b.String()
	}
	fmt.Fprintf(&b, "Start with %s: %s\n\n", focus.ID, focus.Title)
	if len(focus.Requirements) > 0 {
		fmt.Fprintf(&b, "Requirements: %s\n\n", strings.Join(focus.Requirements, ", "))
	}
	if len(focus.Evaluations) > 0 {
		fmt.Fprintf(&b, "Evaluations: %s\n\n", strings.Join(focus.Evaluations, ", "))
	}
	if len(focus.Artifacts) > 0 {
		fmt.Fprintf(&b, "Artifacts: %s\n\n", strings.Join(focus.Artifacts, ", "))
	}
	b.WriteString("Treat this as suggested downstream focus, not NI-owned runtime state.\n")
	return b.String()
}

type idTitle struct {
	id    string
	title string
}

func writeIDList(b *strings.Builder, title string, items []idTitle) {
	if len(items) == 0 {
		return
	}
	fmt.Fprintf(b, "## %s\n\n", title)
	for _, item := range items {
		fmt.Fprintf(b, "- %s: %s\n", item.id, item.title)
	}
	b.WriteString("\n")
}

func acceptedCapabilities(items []contract.Capability) []idTitle {
	result := make([]idTitle, 0, len(items))
	for _, item := range items {
		if item.Status == "accepted" {
			result = append(result, idTitle{item.ID, item.Title})
		}
	}
	return result
}

func acceptedRequirements(items []contract.Requirement) []idTitle {
	result := make([]idTitle, 0, len(items))
	for _, item := range items {
		if item.Status == "accepted" {
			result = append(result, idTitle{item.ID, item.Title})
		}
	}
	return result
}

func nonGoals(items []contract.NonGoal) []idTitle {
	result := make([]idTitle, 0, len(items))
	for _, item := range items {
		result = append(result, idTitle{item.ID, item.Title})
	}
	return result
}

func evaluations(items []contract.Evaluation) []idTitle {
	result := make([]idTitle, 0, len(items))
	for _, item := range items {
		result = append(result, idTitle{item.ID, item.Method})
	}
	return result
}

func artifacts(items []contract.Artifact) []idTitle {
	result := make([]idTitle, 0, len(items))
	for _, item := range items {
		result = append(result, idTitle{item.ID, item.Kind + " at " + item.Path})
	}
	return result
}

func risks(items []contract.Risk) []idTitle {
	result := make([]idTitle, 0, len(items))
	for _, item := range items {
		result = append(result, idTitle{item.ID, item.Severity + ": " + item.Mitigation})
	}
	return result
}

func firstAcceptedCapability(items []contract.Capability) contract.Capability {
	for _, item := range items {
		if item.Status == "accepted" {
			return item
		}
	}
	return contract.Capability{}
}

func defaultString(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func fileSHA256(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}
