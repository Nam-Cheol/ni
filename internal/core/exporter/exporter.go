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

const (
	HyperRunTarget  = "hyper-run"
	NambaAITarget   = "namba-ai"
	OuroborosTarget = "ouroboros"
	SpecKitTarget   = "spec-kit"
	validExportText = HyperRunTarget + ", " + NambaAITarget + ", " + OuroborosTarget + ", " + SpecKitTarget
)

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

	var docs []exportDoc
	switch selectedTarget.Name {
	case HyperRunTarget:
		docs = hyperRunDocs(c, verification.Lockfile, "sha256:"+lockHash)
	case NambaAITarget:
		docs = nambaAIDocs(c, verification.Lockfile, "sha256:"+lockHash)
	case OuroborosTarget:
		docs = ouroborosDocs(c, verification.Lockfile, "sha256:"+lockHash)
	case SpecKitTarget:
		docs = specKitDocs(c, verification.Lockfile, "sha256:"+lockHash)
	default:
		return Result{}, fmt.Errorf("unsupported export target %q (valid: %s)", selectedTarget.Name, validExportText)
	}

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

func nambaAIDocs(c contract.Contract, l lock.Lockfile, sourceLockHash string) []exportDoc {
	return []exportDoc{
		{name: "planning.md", content: renderNambaPlanning(c)},
		{name: "ni-lock-summary.md", content: renderNambaLockSummary(c, l, sourceLockHash)},
		{name: "capability-map.md", content: renderNambaCapabilityMap(c)},
		{name: "evaluation-map.md", content: renderNambaEvaluationMap(c)},
		{name: "risk-map.md", content: renderNambaRiskMap(c)},
		{name: "suggested-spec-boundaries.md", content: renderNambaSpecBoundaries(c)},
	}
}

func ouroborosDocs(c contract.Contract, l lock.Lockfile, sourceLockHash string) []exportDoc {
	return []exportDoc{
		{name: "ouroboros-seed-notes.md", content: renderTargetSeedNotes(c, l, sourceLockHash, targetSeedOptions{
			title:       "Ouroboros Seed Notes",
			targetName:  "Ouroboros",
			description: "downstream recursive planning tools",
			targetNonGoals: []string{
				"Do not implement interview, crystallize, execute, evaluate, or evolve inside NI.",
				"Do not create an Ouroboros lifecycle, queue, or owned runtime state.",
				"Do not execute downstream agents.",
			},
		})},
	}
}

func specKitDocs(c contract.Contract, l lock.Lockfile, sourceLockHash string) []exportDoc {
	return []exportDoc{
		{name: "spec-kit-seed-notes.md", content: renderTargetSeedNotes(c, l, sourceLockHash, targetSeedOptions{
			title:       "Spec Kit Seed Notes",
			targetName:  "Spec Kit",
			description: "downstream spec-first tools",
			targetNonGoals: []string{
				"Do not implement slash commands in NI.",
				"Do not create coding task lists as NI core state.",
				"Do not turn NI into a spec-first coding operating system.",
			},
		})},
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

func renderNambaPlanning(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Planning Seed\n\n")
	fmt.Fprintf(&b, "Project: %s\n\n", c.Project.Name)
	fmt.Fprintf(&b, "Purpose: %s\n\n", c.Project.Purpose)
	fmt.Fprintf(&b, "Product type: %s\n\n", c.ProductType)
	fmt.Fprintf(&b, "Delivery surfaces: %s\n\n", strings.Join(c.DeliverySurfaces, ", "))
	b.WriteString("This is namba-ai-oriented planning seed material derived from a locked NI contract.\n\n")
	b.WriteString("Boundary rules:\n\n")
	b.WriteString("- NI exports markdown seed documents only.\n")
	b.WriteString("- NI does not call namba or implement namba run, sync, pr, or land behavior.\n")
	b.WriteString("- NI does not add Codex-only assumptions to this target.\n")
	b.WriteString("- Downstream planning and execution state must remain outside NI and derived from the lock.\n\n")
	writeIDList(&b, "Accepted capabilities", acceptedCapabilities(c.Capabilities))
	writeIDList(&b, "Accepted requirements", acceptedRequirements(c.Requirements))
	writeIDList(&b, "Non-goals", nonGoals(c.NonGoals))
	return b.String()
}

type targetSeedOptions struct {
	title          string
	targetName     string
	description    string
	targetNonGoals []string
}

func renderTargetSeedNotes(c contract.Contract, l lock.Lockfile, sourceLockHash string, opts targetSeedOptions) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", opts.title)
	fmt.Fprintf(&b, "These notes are seed material for %s. They are derived from a locked NI contract and are not an executable workflow.\n\n", opts.description)
	b.WriteString("Boundary rules:\n\n")
	b.WriteString("- NI exports seed notes only.\n")
	fmt.Fprintf(&b, "- NI does not call %s or own %s state.\n", opts.targetName, opts.targetName)
	b.WriteString("- Downstream state must remain outside NI and derived from the lock.\n\n")

	renderLockedContractSummary(&b, c, l, sourceLockHash)
	renderSeedCapabilities(&b, c)
	renderSeedConstraints(&b, c)
	renderSeedRisks(&b, c)
	renderSeedEvaluations(&b, c)
	renderSeedNonGoals(&b, c, opts.targetNonGoals)
	renderSourceTruthReferences(&b, l)
	return b.String()
}

func renderLockedContractSummary(b *strings.Builder, c contract.Contract, l lock.Lockfile, sourceLockHash string) {
	b.WriteString("## Locked Contract Summary\n\n")
	fmt.Fprintf(b, "- project_id: %s\n", c.Project.ID)
	fmt.Fprintf(b, "- project_name: %s\n", c.Project.Name)
	fmt.Fprintf(b, "- purpose: %s\n", c.Project.Purpose)
	fmt.Fprintf(b, "- product_type: %s\n", c.ProductType)
	fmt.Fprintf(b, "- delivery_surfaces: %s\n", strings.Join(c.DeliverySurfaces, ", "))
	fmt.Fprintf(b, "- interaction_mode: %s\n", c.InteractionMode)
	fmt.Fprintf(b, "- readiness_profile: %s\n", c.ReadinessProfile)
	fmt.Fprintf(b, "- locked_readiness_status: %s\n", l.Readiness.Status)
	fmt.Fprintf(b, "- locked_at: %s\n", l.LockedAt)
	fmt.Fprintf(b, "- source_lock_hash: %s\n\n", sourceLockHash)
}

func renderSeedCapabilities(b *strings.Builder, c contract.Contract) {
	b.WriteString("## Capabilities\n\n")
	requirements := requirementTitles(c.Requirements)
	evaluations := evaluationTitles(c.Evaluations)
	risks := riskTitles(c.Risks)
	artifacts := artifactTitles(c.Artifacts)
	wrote := false
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		wrote = true
		fmt.Fprintf(b, "### %s: %s\n\n", capability.ID, capability.Title)
		fmt.Fprintf(b, "- depends_on: %s\n", joinOrNone(capability.Dependencies))
		fmt.Fprintf(b, "- requirements: %s\n", listWithTitles(capability.Requirements, requirements))
		fmt.Fprintf(b, "- evaluations: %s\n", listWithTitles(capability.Evaluations, evaluations))
		fmt.Fprintf(b, "- risks: %s\n", listWithTitles(capability.Risks, risks))
		fmt.Fprintf(b, "- artifacts: %s\n\n", listWithTitles(capability.Artifacts, artifacts))
	}
	if !wrote {
		b.WriteString("No accepted capabilities are present in the locked contract.\n\n")
	}
}

func renderSeedConstraints(b *strings.Builder, c contract.Contract) {
	b.WriteString("## Constraints\n\n")
	b.WriteString("- Verify `.ni/plan.lock.json` exists and locked hashes match before using these notes.\n")
	b.WriteString("- Preserve accepted requirements, risk mitigations, and blocker handling.\n")
	b.WriteString("- Keep downstream state derived and outside NI core.\n")
	b.WriteString("- Do not weaken acceptance criteria to make downstream work appear ready.\n")
	b.WriteString("\nAccepted requirements:\n\n")
	items := acceptedRequirements(c.Requirements)
	if len(items) == 0 {
		b.WriteString("- none\n\n")
		return
	}
	for _, item := range items {
		fmt.Fprintf(b, "- %s: %s\n", item.id, item.title)
	}
	b.WriteString("\n")
}

func renderSeedRisks(b *strings.Builder, c contract.Contract) {
	b.WriteString("## Risks\n\n")
	if len(c.Risks) == 0 {
		b.WriteString("No risks are listed in the locked contract.\n\n")
		return
	}
	linkedCapabilities := linkedCapabilitiesByRisk(c.Capabilities)
	for _, risk := range c.Risks {
		fmt.Fprintf(b, "### %s: %s\n\n", risk.ID, risk.Title)
		fmt.Fprintf(b, "- severity: %s\n", risk.Severity)
		fmt.Fprintf(b, "- status: %s\n", risk.Status)
		fmt.Fprintf(b, "- mitigation: %s\n", risk.Mitigation)
		fmt.Fprintf(b, "- linked_capabilities: %s\n\n", joinOrNone(linkedCapabilities[risk.ID]))
	}
}

func renderSeedEvaluations(b *strings.Builder, c contract.Contract) {
	b.WriteString("## Evaluation Contract\n\n")
	if len(c.Evaluations) == 0 {
		b.WriteString("No evaluations are listed in the locked contract.\n\n")
		return
	}
	linkedCapabilities := linkedCapabilitiesByEvaluation(c.Capabilities)
	for _, evaluation := range c.Evaluations {
		fmt.Fprintf(b, "### %s: %s\n\n", evaluation.ID, evaluation.Title)
		fmt.Fprintf(b, "- method: %s\n", evaluation.Method)
		fmt.Fprintf(b, "- linked_capabilities: %s\n\n", joinOrNone(linkedCapabilities[evaluation.ID]))
	}
	if len(c.Artifacts) > 0 {
		b.WriteString("Evidence references:\n\n")
		for _, artifact := range c.Artifacts {
			fmt.Fprintf(b, "- %s: %s at %s\n", artifact.ID, artifact.Kind, artifact.Path)
		}
		b.WriteString("\n")
	}
}

func renderSeedNonGoals(b *strings.Builder, c contract.Contract, targetNonGoals []string) {
	b.WriteString("## Non-goals\n\n")
	items := nonGoals(c.NonGoals)
	if len(items) == 0 {
		b.WriteString("- none\n\n")
	} else {
		for _, item := range items {
			fmt.Fprintf(b, "- %s: %s\n", item.id, item.title)
		}
		b.WriteString("\n")
	}
	if len(targetNonGoals) == 0 {
		return
	}
	b.WriteString("## Target Non-goals\n\n")
	for _, item := range targetNonGoals {
		fmt.Fprintf(b, "- %s\n", item)
	}
	b.WriteString("\n")
}

func renderSourceTruthReferences(b *strings.Builder, l lock.Lockfile) {
	b.WriteString("## Source-of-Truth References\n\n")
	fmt.Fprintf(b, "Source-of-truth order: %s\n\n", strings.Join(l.SourceOfTruth, " > "))
	b.WriteString("Locked files:\n\n")
	for _, file := range l.Files {
		fmt.Fprintf(b, "- %s sha256:%s\n", file.Path, file.SHA256)
	}
	b.WriteString("\nStop with `BLOCKED` if any locked hash mismatches before downstream work uses these notes.\n")
}

func renderNambaLockSummary(c contract.Contract, l lock.Lockfile, sourceLockHash string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# NI Lock Summary\n\n")
	fmt.Fprintf(&b, "Project ID: %s\n\n", c.Project.ID)
	fmt.Fprintf(&b, "Readiness profile: %s\n\n", c.ReadinessProfile)
	fmt.Fprintf(&b, "Locked readiness status: %s\n\n", l.Readiness.Status)
	fmt.Fprintf(&b, "Locked at: %s\n\n", l.LockedAt)
	fmt.Fprintf(&b, "Source lock hash: %s\n\n", sourceLockHash)
	fmt.Fprintf(&b, "Source of truth: %s\n\n", strings.Join(l.SourceOfTruth, " > "))
	b.WriteString("Locked files:\n\n")
	for _, file := range l.Files {
		fmt.Fprintf(&b, "- %s sha256:%s\n", file.Path, file.SHA256)
	}
	b.WriteString("\nStop with `BLOCKED` if any locked hash mismatches before downstream work uses this seed.\n")
	return b.String()
}

func renderNambaCapabilityMap(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Capability Map\n\n")
	b.WriteString("Capabilities are graph nodes. `depends_on` edges describe prerequisites without requiring a total execution order.\n\n")
	requirements := requirementTitles(c.Requirements)
	evaluations := evaluationTitles(c.Evaluations)
	risks := riskTitles(c.Risks)
	artifacts := artifactTitles(c.Artifacts)
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		fmt.Fprintf(&b, "## %s: %s\n\n", capability.ID, capability.Title)
		fmt.Fprintf(&b, "- status: %s\n", capability.Status)
		fmt.Fprintf(&b, "- depends_on: %s\n", joinOrNone(capability.Dependencies))
		fmt.Fprintf(&b, "- requirements: %s\n", listWithTitles(capability.Requirements, requirements))
		fmt.Fprintf(&b, "- evaluations: %s\n", listWithTitles(capability.Evaluations, evaluations))
		fmt.Fprintf(&b, "- risks: %s\n", listWithTitles(capability.Risks, risks))
		fmt.Fprintf(&b, "- artifacts: %s\n\n", listWithTitles(capability.Artifacts, artifacts))
	}
	return b.String()
}

func renderNambaEvaluationMap(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Evaluation Map\n\n")
	b.WriteString("Evaluations are validation expectations for downstream planning. NI exports this map but does not run evaluations.\n\n")
	linkedCapabilities := linkedCapabilitiesByEvaluation(c.Capabilities)
	for _, evaluation := range c.Evaluations {
		fmt.Fprintf(&b, "## %s: %s\n\n", evaluation.ID, evaluation.Title)
		fmt.Fprintf(&b, "- method: %s\n", evaluation.Method)
		fmt.Fprintf(&b, "- linked_capabilities: %s\n\n", joinOrNone(linkedCapabilities[evaluation.ID]))
	}
	if len(c.Artifacts) > 0 {
		b.WriteString("Evidence locations:\n\n")
		for _, artifact := range c.Artifacts {
			fmt.Fprintf(&b, "- %s: %s at %s\n", artifact.ID, artifact.Kind, artifact.Path)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func renderNambaRiskMap(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Risk Map\n\n")
	b.WriteString("Risks remain contract constraints for downstream planning; mitigations must not be weakened to make work appear ready.\n\n")
	linkedCapabilities := linkedCapabilitiesByRisk(c.Capabilities)
	for _, risk := range c.Risks {
		fmt.Fprintf(&b, "## %s: %s\n\n", risk.ID, risk.Title)
		fmt.Fprintf(&b, "- severity: %s\n", risk.Severity)
		fmt.Fprintf(&b, "- status: %s\n", risk.Status)
		fmt.Fprintf(&b, "- mitigation: %s\n", risk.Mitigation)
		fmt.Fprintf(&b, "- linked_capabilities: %s\n\n", joinOrNone(linkedCapabilities[risk.ID]))
	}
	writeIDList(&b, "Non-goals", nonGoals(c.NonGoals))
	return b.String()
}

func renderNambaSpecBoundaries(c contract.Contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Suggested Spec Boundaries\n\n")
	b.WriteString("This file is a proposal, not a required sequential SPEC chain.\n\n")
	b.WriteString("Use these as candidate graph boundaries for downstream planning. A downstream tool may split, merge, or reorder boundaries when dependency edges and locked acceptance criteria are preserved.\n\n")
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		fmt.Fprintf(&b, "## Boundary candidate: %s\n\n", capability.ID)
		fmt.Fprintf(&b, "- intent: %s\n", capability.Title)
		fmt.Fprintf(&b, "- depends_on: %s\n", joinOrNone(capability.Dependencies))
		fmt.Fprintf(&b, "- acceptance_refs: %s\n", joinOrNone(capability.Requirements))
		fmt.Fprintf(&b, "- validation_refs: %s\n", joinOrNone(capability.Evaluations))
		fmt.Fprintf(&b, "- risk_refs: %s\n", joinOrNone(capability.Risks))
		fmt.Fprintf(&b, "- artifact_refs: %s\n\n", joinOrNone(capability.Artifacts))
	}
	b.WriteString("Do not interpret this proposal as permission for NI to run namba, create mandatory execution order, or own downstream runtime state.\n")
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

func requirementTitles(items []contract.Requirement) map[string]string {
	result := make(map[string]string, len(items))
	for _, item := range items {
		result[item.ID] = item.Title
	}
	return result
}

func evaluationTitles(items []contract.Evaluation) map[string]string {
	result := make(map[string]string, len(items))
	for _, item := range items {
		result[item.ID] = item.Title
	}
	return result
}

func riskTitles(items []contract.Risk) map[string]string {
	result := make(map[string]string, len(items))
	for _, item := range items {
		result[item.ID] = item.Title
	}
	return result
}

func artifactTitles(items []contract.Artifact) map[string]string {
	result := make(map[string]string, len(items))
	for _, item := range items {
		result[item.ID] = item.Kind + " at " + item.Path
	}
	return result
}

func linkedCapabilitiesByEvaluation(items []contract.Capability) map[string][]string {
	result := make(map[string][]string)
	for _, capability := range items {
		if capability.Status != "accepted" {
			continue
		}
		for _, evaluation := range capability.Evaluations {
			result[evaluation] = append(result[evaluation], capability.ID)
		}
	}
	return result
}

func linkedCapabilitiesByRisk(items []contract.Capability) map[string][]string {
	result := make(map[string][]string)
	for _, capability := range items {
		if capability.Status != "accepted" {
			continue
		}
		for _, risk := range capability.Risks {
			result[risk] = append(result[risk], capability.ID)
		}
	}
	return result
}

func listWithTitles(ids []string, titles map[string]string) string {
	if len(ids) == 0 {
		return "none"
	}
	parts := make([]string, 0, len(ids))
	for _, id := range ids {
		if title := titles[id]; title != "" {
			parts = append(parts, id+" ("+title+")")
			continue
		}
		parts = append(parts, id)
	}
	return strings.Join(parts, ", ")
}

func joinOrNone(items []string) string {
	if len(items) == 0 {
		return "none"
	}
	return strings.Join(items, ", ")
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
