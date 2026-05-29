package main

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"--help"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "ni is a project intent compiler") {
		t.Fatalf("help output did not describe ni: %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "ni version") {
		t.Fatalf("help output did not mention version command: %q", stdout.String())
	}
}

func TestVersion(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"version"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if got := strings.TrimSpace(stdout.String()); got != "0.0.0-dev" {
		t.Fatalf("expected version 0.0.0-dev, got %q", got)
	}
}

func TestInit(t *testing.T) {
	dir := t.TempDir()
	var stdout bytes.Buffer

	code := run([]string{"init", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	required := []string{
		"docs/plan/00_project_brief.md",
		".ni/project.json",
		".ni/contract.json",
		".ni/session.json",
		".ni/pressure.json",
		".ni/harness.candidates.json",
		".ni/readiness.rules.json",
		".ni/readiness.profiles.json",
	}
	for _, path := range required {
		if _, err := os.Stat(filepath.Join(dir, path)); err != nil {
			t.Fatalf("expected %s to exist: %v", path, err)
		}
	}
	if !strings.Contains(stdout.String(), "initialized ni planning workspace") {
		t.Fatalf("expected init summary, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "readiness profile: prototype") {
		t.Fatalf("expected prototype profile summary, got %q", stdout.String())
	}
}

func TestInitWithProfile(t *testing.T) {
	dir := t.TempDir()
	var stdout bytes.Buffer

	code := run([]string{"init", "--dir", dir, "--profile", "concept"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	data, err := os.ReadFile(filepath.Join(dir, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("reading contract: %v", err)
	}
	if !strings.Contains(string(data), `"readiness_profile": "concept"`) {
		t.Fatalf("expected concept readiness profile, got %q", string(data))
	}
	if !strings.Contains(stdout.String(), "readiness profile: concept") {
		t.Fatalf("expected concept profile summary, got %q", stdout.String())
	}
}

func TestInitWithProductTypeAndSurface(t *testing.T) {
	dir := t.TempDir()
	var stdout bytes.Buffer

	code := run([]string{"init", "--dir", dir, "--product-type", "conversation_product", "--surface", "conversation", "--interaction-mode", "human_to_human"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	data, err := os.ReadFile(filepath.Join(dir, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("reading contract: %v", err)
	}
	if !strings.Contains(string(data), `"product_type": "conversation_product"`) {
		t.Fatalf("expected conversation product type, got %q", string(data))
	}
	if !strings.Contains(string(data), `"delivery_surfaces": ["conversation"]`) {
		t.Fatalf("expected conversation surface, got %q", string(data))
	}
	if !strings.Contains(stdout.String(), "product type: conversation_product") {
		t.Fatalf("expected product type summary, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "delivery surfaces: conversation") {
		t.Fatalf("expected surface summary, got %q", stdout.String())
	}
}

func TestStatus(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "BLOCKED") {
		t.Fatalf("expected blocked status for template project, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "profile: prototype") {
		t.Fatalf("expected active profile in status output, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "product type: software") {
		t.Fatalf("expected product type in status output, got %q", stdout.String())
	}
}

func TestStatusJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"status": "BLOCKED"`) {
		t.Fatalf("expected JSON blocked status, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), `"profile": "prototype"`) {
		t.Fatalf("expected JSON profile, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), `"product_type": "software"`) {
		t.Fatalf("expected JSON product type, got %q", stdout.String())
	}
	if strings.Contains(stdout.String(), `"next_questions"`) {
		t.Fatalf("did not expect next questions without flag, got %q", stdout.String())
	}
}

func TestStatusNextQuestionsText(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--next-questions"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "question R009 OQ-001:") {
		t.Fatalf("expected blocker open question prompt, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "What answer would resolve it") {
		t.Fatalf("expected deterministic next question, got %q", stdout.String())
	}
}

func TestStatusNextQuestionsJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--json", "--next-questions"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload struct {
		NextQuestions []struct {
			RuleID     string   `json:"rule_id"`
			References []string `json:"references"`
			Question   string   `json:"question"`
		} `json:"next_questions"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("decoding status JSON: %v\n%s", err, stdout.String())
	}
	if len(payload.NextQuestions) == 0 {
		t.Fatalf("expected next questions, got %q", stdout.String())
	}
	found := false
	for _, question := range payload.NextQuestions {
		if question.RuleID == "R009" && contains(question.References, "OQ-001") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected R009 question for OQ-001, got %#v", payload.NextQuestions)
	}
}

func TestStatusNextQuestionsJSONIncludesEmptyWhenReady(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--json", "--next-questions"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("decoding status JSON: %v\n%s", err, stdout.String())
	}
	raw, ok := payload["next_questions"]
	if !ok {
		t.Fatalf("expected next_questions when requested, got %q", stdout.String())
	}
	if string(raw) != "[]" {
		t.Fatalf("expected empty next_questions array, got %s in %q", raw, stdout.String())
	}
}

func TestStatusProofTextWithNextQuestions(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--proof", "--next-questions"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	out := stdout.String()
	for _, want := range []string{
		"NI Intent Readiness: BLOCKED",
		"Blockers:",
		"R014 Project purpose is missing.",
		"Why it matters: ni cannot lock intent until it knows what reality the project is meant to change.",
		"Next: describe the project in one or two sentences: what should change, for whom, and why it matters.",
		"OQ-001 is marked as blocker.",
		"Why it matters: open blocker questions mean required intent is still unresolved.",
		"Next: answer or defer the blocker question, or keep it blocking with an explicit reason.",
		"R015 Actors or outcomes are missing.",
		"Why it matters: ni cannot judge readiness without knowing who uses or operates the product and what successful use looks like for them.",
		"Next: list the primary actors and the outcome each one expects.",
		"R016 Delivery surface is missing.",
		"Why it matters: downstream handoff depends on knowing whether the product is delivered as a CLI, web app, conversation, document, workflow, research protocol, human service, or another surface.",
		"Next: choose the likely delivery surface, or mark it deferred with an explicit reason.",
		"Deferrals:",
		"Warnings:",
		"Passed checks:",
		"Execution must not start.",
		"Next questions:",
		"1. ",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in proof output, got %q", want, out)
		}
	}
	for _, forbidden := range []string{
		"this deterministic readiness rule affects whether the plan can be trusted.",
		"update planning docs and .ni/contract.json together to resolve this rule.",
	} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("fresh-workspace proof output should not contain generic fallback %q, got %q", forbidden, out)
		}
	}
}

func TestStatusProofTextGroupsDeferralsWarningsAndPassedChecks(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyWithDeferralsContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--proof"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	out := stdout.String()
	for _, want := range []string{
		"NI Intent Readiness: READY_WITH_DEFERRALS",
		"Blockers:\n- None.",
		"Deferrals:",
		"OQ-001 remains open.",
		"Warnings:",
		"DEC-001 is deferred.",
		"Why it matters: downstream work must avoid depending on this decision.",
		"Passed checks:",
		"Required docs exist.",
		"Docs and contract are synchronized.",
		"Execution may proceed only after lock; deferrals remain explicit.",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in proof output, got %q", want, out)
		}
	}
}

func TestStatusProofJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--proof", "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload struct {
		Proof []map[string]json.RawMessage `json:"proof"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("decoding status proof JSON: %v\n%s", err, stdout.String())
	}
	var ruleID string
	if len(payload.Proof) == 1 {
		if err := json.Unmarshal(payload.Proof[0]["rule_id"], &ruleID); err != nil {
			t.Fatalf("decoding proof rule id: %v\n%#v", err, payload.Proof[0])
		}
	}
	if len(payload.Proof) != 1 || ruleID != "READY" {
		t.Fatalf("expected ready proof item, got %#v in %q", payload.Proof, stdout.String())
	}
	for _, unstable := range []string{"category", "why", "next"} {
		if _, ok := payload.Proof[0][unstable]; ok {
			t.Fatalf("proof JSON should stay stable without %q, got %#v", unstable, payload.Proof[0])
		}
	}
}

func TestStatusProofTextShowsSyncDiagnosticFields(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	decisionDoc := "# Decision log\n\n## DEC-001: Decision\n\nStatus: rejected\n"
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "11_decision_log.md"), []byte(decisionDoc), 0o644); err != nil {
		t.Fatalf("writing conflicting decision doc: %v", err)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--proof"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	out := stdout.String()
	for _, want := range []string{
		"ID: DEC-001",
		"Location: docs/plan/11_decision_log.md:3",
		"Problem: DEC-001 status is \"rejected\" but .ni/contract.json says \"accepted\".",
		"Suggested repair: Update the stale source so the docs status and contract status match.",
		"Blocks ni-end: true",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in proof output, got %q", want, out)
		}
	}
}

func TestStatusProofTextShowsFirstRunSyncDiagnostic(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	updateContractForCLI(t, dir, func(payload map[string]any) {
		project := payload["project"].(map[string]any)
		project["purpose"] = "TODO"
	})

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--proof"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	out := stdout.String()
	for _, want := range []string{
		"NI Intent Readiness: BLOCKED",
		"Project purpose is documented but missing from .ni/contract.json.",
		"ID: SYNC-014",
		"Location: docs/plan/00_project_brief.md",
		"Problem: Project purpose is documented but missing from .ni/contract.json.",
		"Why it matters: ni cannot safely lock a plan when the human-readable purpose and machine-readable contract disagree.",
		"Suggested repair: Record the documented project purpose in .ni/contract.json, or mark the purpose unresolved with an explicit blocker question.",
		"Blocks ni-end: true",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in proof output, got %q", want, out)
		}
	}
	for _, forbidden := range []string{
		"this deterministic readiness rule affects whether the plan can be trusted.",
		"update planning docs and .ni/contract.json together to resolve this rule.",
	} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("first-run sync diagnostic should not use generic fallback %q, got %q", forbidden, out)
		}
	}
}

func TestStatusJSONIncludesFirstRunSyncDiagnostic(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "08_delivery_operation.md"), []byte("# Delivery and operation\n\n## Delivery surfaces\n\n- conversation\n\n## Initial delivery\n\nConversation planning happens before lock.\n"), 0o644); err != nil {
		t.Fatalf("writing mismatched delivery doc: %v", err)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--json", "--proof"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload struct {
		Issues []struct {
			RuleID         string `json:"rule_id"`
			SyncDiagnostic *struct {
				ID              string `json:"id"`
				Location        string `json:"location"`
				Problem         string `json:"problem"`
				WhyItMatters    string `json:"why_it_matters"`
				SuggestedRepair string `json:"suggested_repair"`
				BlocksEnd       bool   `json:"blocks_ni_end"`
			} `json:"sync_diagnostic"`
		} `json:"issues"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("decoding status JSON: %v\n%s", err, stdout.String())
	}
	var found bool
	for _, issue := range payload.Issues {
		if issue.RuleID == "R012" && issue.SyncDiagnostic != nil && issue.SyncDiagnostic.ID == "SYNC-016" {
			found = true
			if issue.SyncDiagnostic.Location == "" || issue.SyncDiagnostic.Problem == "" || issue.SyncDiagnostic.WhyItMatters == "" || issue.SyncDiagnostic.SuggestedRepair == "" || !issue.SyncDiagnostic.BlocksEnd {
				t.Fatalf("expected stable first-run sync diagnostic fields, got %#v", issue.SyncDiagnostic)
			}
		}
	}
	if !found {
		t.Fatalf("expected SYNC-016 R012 diagnostic in %q", stdout.String())
	}
}

func TestStatusJSONInvalidContractStructuredError(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), []byte("{\n"), 0o644); err != nil {
		t.Fatalf("writing invalid contract: %v", err)
	}

	var stdout, stderr bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--json"}, &stdout, &stderr)
	if code != exitInvalidContract {
		t.Fatalf("expected exit code %d, got %d", exitInvalidContract, code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr for JSON error, got %q", stderr.String())
	}
	envelope := decodeErrorEnvelope(t, stdout.Bytes())
	if envelope.Error.Code != "invalid_contract" || envelope.Error.ExitCode != exitInvalidContract {
		t.Fatalf("expected invalid_contract envelope, got %#v", envelope)
	}
	if envelope.Error.Details["command"] != "status" {
		t.Fatalf("expected status command detail, got %#v", envelope)
	}
	if !strings.Contains(envelope.Error.Message, "malformed contract JSON") {
		t.Fatalf("expected malformed contract message, got %#v", envelope)
	}
}

func TestJSONUsageErrorEnvelope(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run([]string{"status", "--json", "--bad"}, &stdout, &stderr)
	if code != exitUsageError {
		t.Fatalf("expected exit code %d, got %d", exitUsageError, code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr for JSON usage error, got %q", stderr.String())
	}
	envelope := decodeErrorEnvelope(t, stdout.Bytes())
	if envelope.Error.Code != "usage_error" || envelope.Error.ExitCode != exitUsageError {
		t.Fatalf("expected usage_error envelope, got %#v", envelope)
	}
}

func TestEndBlocked(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stderr bytes.Buffer
	code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitReadinessBlocked {
		t.Fatalf("expected exit code %d, got %d", exitReadinessBlocked, code)
	}
	if !strings.Contains(stderr.String(), "readiness is BLOCKED") {
		t.Fatalf("expected blocked error, got %q", stderr.String())
	}
}

func TestEndBlockedByDocsContractSync(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	path := filepath.Join(dir, "docs", "plan", "02_capabilities.md")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading capability doc: %v", err)
	}
	data = append(data, []byte("\n## CAP-999: Docs-only capability\n\nThis section is missing from the contract.\n")...)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("writing capability doc: %v", err)
	}

	var stderr bytes.Buffer
	code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitReadinessBlocked {
		t.Fatalf("expected exit code %d, got %d", exitReadinessBlocked, code)
	}
	if !strings.Contains(stderr.String(), "readiness is BLOCKED") {
		t.Fatalf("expected blocked error, got %q", stderr.String())
	}
}

func TestEndBlockedByFirstRunDocsContractSync(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "01_actors_outcomes.md"), []byte("# Actors and outcomes\n\n## Actors\n\n- TODO\n\n## Outcomes\n\n- TODO\n"), 0o644); err != nil {
		t.Fatalf("writing stale actors doc: %v", err)
	}

	var statusOut bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--proof"}, &statusOut, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected status exit code 0, got %d", code)
	}
	if !strings.Contains(statusOut.String(), "ID: SYNC-015") {
		t.Fatalf("expected SYNC-015 in status proof, got %q", statusOut.String())
	}

	var stderr bytes.Buffer
	code = run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitReadinessBlocked {
		t.Fatalf("expected exit code %d, got %d", exitReadinessBlocked, code)
	}
	if !strings.Contains(stderr.String(), "readiness is BLOCKED") {
		t.Fatalf("expected blocked error, got %q", stderr.String())
	}
}

func TestEndCreatesLockfile(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"end", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "locked plan") {
		t.Fatalf("expected lock summary, got %q", stdout.String())
	}
	if _, err := os.Stat(filepath.Join(dir, ".ni", "plan.lock.json")); err != nil {
		t.Fatalf("expected lockfile: %v", err)
	}
}

func TestRunWritesPrompt(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, ".ni", "generated", "goal.prompt.txt")
	var stdout bytes.Buffer
	code := run([]string{"run", "--dir", dir, "--out", out, "--max-chars", "4000"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "compiled prompt") {
		t.Fatalf("expected compile summary, got %q", stdout.String())
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("expected prompt file: %v", err)
	}
	if !strings.Contains(string(data), "Source of truth") {
		t.Fatalf("expected compiled prompt content, got %q", string(data))
	}
}

func TestRunWithTarget(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	for _, targetName := range []string{"generic", "codex", "human-team"} {
		t.Run(targetName, func(t *testing.T) {
			var stdout bytes.Buffer
			code := run([]string{"run", "--dir", dir, "--target", targetName}, &stdout, &bytes.Buffer{})
			if code != 0 {
				t.Fatalf("expected exit code 0, got %d", code)
			}
			if !strings.Contains(stdout.String(), "Target: "+targetName) {
				t.Fatalf("expected target output, got %q", stdout.String())
			}
			if len([]rune(stdout.String())) > 4000 {
				t.Fatalf("expected prompt <= 4000 chars, got %d", len([]rune(stdout.String())))
			}
		})
	}
}

func TestRunCodexAndHumanTeamTargetsArePromptArtifactsOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	cases := []struct {
		target      string
		outName     string
		wantContent []string
		forbidden   []string
	}{
		{
			target:  "codex",
			outName: "codex.prompt.txt",
			wantContent: []string{
				"Codex target prompt",
				"Target: codex (prompt)",
				"Paste this prompt into Codex",
				"Do not ask ni to invoke Codex automatically",
				"ni run only compiled this prompt",
			},
			forbidden: []string{
				filepath.Join(".codex", "commands"),
				"codex-exec.sh",
				"queue",
				"tasks.md",
			},
		},
		{
			target:  "human-team",
			outName: "human-team.prompt.md",
			wantContent: []string{
				"Human-team handoff",
				"Target: human-team (handoff)",
				"PM/dev/design/research team",
				"Execution responsibility stays outside ni",
				"team handoff",
			},
			forbidden: []string{
				"owners.db",
				"team-runtime",
				"queue",
				"tasks.md",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.target, func(t *testing.T) {
			outDir := filepath.Join(dir, "target-output", tc.target)
			outPath := filepath.Join(outDir, tc.outName)
			var stdout bytes.Buffer
			code := run([]string{"run", "--dir", dir, "--target", tc.target, "--out", outPath}, &stdout, &bytes.Buffer{})
			if code != 0 {
				t.Fatalf("expected exit code 0, got %d", code)
			}
			if !strings.Contains(stdout.String(), "compiled prompt at "+outPath) {
				t.Fatalf("expected write summary, got %q", stdout.String())
			}
			assertExportFilesExactly(t, outDir, []string{tc.outName})
			if len([]rune(string(readFileForCLI(t, outPath)))) > 4000 {
				t.Fatalf("expected %s output <= 4000 chars", tc.target)
			}
			assertFileContains(t, outPath, tc.wantContent)
			assertExportOmitsPaths(t, dir, tc.forbidden)
		})
	}
}

func TestRunMissingLockfileGenericFailure(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stderr bytes.Buffer
	code := run([]string{"run", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitGenericFailure {
		t.Fatalf("expected exit code %d, got %d", exitGenericFailure, code)
	}
	if !strings.Contains(stderr.String(), "missing lockfile") {
		t.Fatalf("expected missing lockfile error, got %q", stderr.String())
	}
}

func TestRunRejectsUnsupportedTarget(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stderr bytes.Buffer
	code := run([]string{"run", "--dir", dir, "--target", "shell"}, &bytes.Buffer{}, &stderr)
	if code != exitUnsupportedTarget {
		t.Fatalf("expected exit code %d, got %d", exitUnsupportedTarget, code)
	}
	if !strings.Contains(stderr.String(), `unsupported target "shell"`) {
		t.Fatalf("expected unsupported target error, got %q", stderr.String())
	}
}

func TestAmendApplyAndRelockFlow(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	lockPath := filepath.Join(dir, ".ni", "plan.lock.json")
	v1Lock := readFileForCLI(t, lockPath)

	var stdout bytes.Buffer
	code := run([]string{"amend", "create", "--dir", dir, "--title", "Clarify prompt compiler acceptance"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected amend create exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "created amendment AMEND-001") {
		t.Fatalf("expected amendment create summary, got %q", stdout.String())
	}
	completeAmendmentForCLI(t, dir, "AMEND-001")

	stdout.Reset()
	code = run([]string{"amend", "list", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected amend list exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "AMEND-001") || !strings.Contains(stdout.String(), "draft") {
		t.Fatalf("expected draft amendment in list, got %q", stdout.String())
	}

	path := filepath.Join(dir, "docs", "plan", "02_capabilities.md")
	if err := os.WriteFile(path, []byte("# Capabilities\n\n## CAP-001: Prompt compiler acceptance clarified\n\nDescribe the accepted prompt compiler behavior.\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	var stderr bytes.Buffer
	code = run([]string{"run", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitStaleLock {
		t.Fatalf("expected stale run exit code %d, got %d", exitStaleLock, code)
	}
	if !strings.Contains(stderr.String(), "BLOCKED: lock hash mismatch") {
		t.Fatalf("expected stale lock rejection before relock, got %q", stderr.String())
	}

	stderr.Reset()
	code = run([]string{"relock", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitStaleLock {
		t.Fatalf("expected relock without applied amendment exit code %d, got %d", exitStaleLock, code)
	}
	if !strings.Contains(stderr.String(), "without an applied amendment") {
		t.Fatalf("expected explicit amendment rejection, got %q", stderr.String())
	}

	stdout.Reset()
	code = run([]string{"amend", "apply", "--dir", dir, "AMEND-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected amend apply exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "applied amendment AMEND-001") {
		t.Fatalf("expected apply summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"amend", "show", "--dir", dir, "AMEND-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected amend show exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"status": "applied"`) {
		t.Fatalf("expected applied amendment JSON, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"relock", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected relock exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "previous lock archived") {
		t.Fatalf("expected previous lock summary, got %q", stdout.String())
	}

	v2Lock := readFileForCLI(t, lockPath)
	if bytes.Equal(v1Lock, v2Lock) {
		t.Fatal("expected relock to write a new lock")
	}
	var lockPayload struct {
		PreviousLock *struct {
			Path     string `json:"path"`
			SHA256   string `json:"sha256"`
			LockedAt string `json:"locked_at"`
		} `json:"previous_lock"`
	}
	if err := json.Unmarshal(v2Lock, &lockPayload); err != nil {
		t.Fatalf("parsing v2 lock: %v", err)
	}
	if lockPayload.PreviousLock == nil || lockPayload.PreviousLock.Path == "" || lockPayload.PreviousLock.SHA256 == "" {
		t.Fatalf("expected previous lock reference in v2 lock, got %#v", lockPayload.PreviousLock)
	}
	archived := readFileForCLI(t, filepath.Join(dir, lockPayload.PreviousLock.Path))
	if !bytes.Equal(archived, v1Lock) {
		t.Fatal("archived previous lock does not match v1 lock")
	}

	stdout.Reset()
	code = run([]string{"run", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected run after relock exit code 0, got %d", code)
	}
}

func TestRelockRefusesBlockedReadiness(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}
	writeBlockedQuestionForCLI(t, dir)

	var stderr bytes.Buffer
	code := run([]string{"relock", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != exitReadinessBlocked {
		t.Fatalf("expected relock exit code %d, got %d", exitReadinessBlocked, code)
	}
	if !strings.Contains(stderr.String(), "readiness is BLOCKED") {
		t.Fatalf("expected blocked readiness refusal, got %q", stderr.String())
	}
}

func TestFeedbackAddAndListJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	for _, fixture := range []string{"codex.json", "human-team.json"} {
		var stdout bytes.Buffer
		code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", fixture)}, &stdout, &bytes.Buffer{})
		if code != 0 {
			t.Fatalf("expected exit code 0 adding %s, got %d", fixture, code)
		}
		if !strings.Contains(stdout.String(), "recorded feedback from") {
			t.Fatalf("expected feedback add summary, got %q", stdout.String())
		}
	}

	var textOut bytes.Buffer
	code := run([]string{"feedback", "list", "--dir", dir}, &textOut, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(textOut.String(), "codex") || !strings.Contains(textOut.String(), "human-team") {
		t.Fatalf("expected text feedback list to include both fixtures, got %q", textOut.String())
	}

	var stdout bytes.Buffer
	code = run([]string{"feedback", "list", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload []map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid feedback JSON, got %v: %q", err, stdout.String())
	}
	if len(payload) != 2 {
		t.Fatalf("expected two feedback entries, got %#v", payload)
	}
	if payload[0]["source_target"] != "codex" || payload[1]["source_target"] != "human-team" {
		t.Fatalf("expected codex and human-team feedback, got %#v", payload)
	}
}

func TestPressureStatusEmptyState(t *testing.T) {
	dir := t.TempDir()

	var stdout bytes.Buffer
	code := run([]string{"pressure", "status", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "no pressure recorded") {
		t.Fatalf("expected empty pressure status, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"pressure", "status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected pressure JSON, got %v: %q", err, stdout.String())
	}
	if payload["schema"] != "ni.pressure.v0" {
		t.Fatalf("expected pressure schema, got %#v", payload)
	}
	if items, ok := payload["items"].([]any); !ok || len(items) != 0 {
		t.Fatalf("expected no pressure items, got %#v", payload["items"])
	}
}

func TestFeedbackAddCreatesObservedPressure(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code = run([]string{"pressure", "status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var ledger struct {
		Schema string `json:"schema"`
		Items  []struct {
			ID                     string   `json:"id"`
			Kind                   string   `json:"kind"`
			Status                 string   `json:"status"`
			EvidenceRefs           []string `json:"evidence_refs"`
			RelatedCapabilities    []string `json:"related_capabilities"`
			RelatedRisks           []string `json:"related_risks"`
			ProposedAction         string   `json:"proposed_action"`
			RequiresUserAcceptance bool     `json:"requires_user_acceptance"`
		} `json:"items"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &ledger); err != nil {
		t.Fatalf("expected pressure ledger JSON, got %v: %q", err, stdout.String())
	}
	if ledger.Schema != "ni.pressure.v0" {
		t.Fatalf("expected pressure schema, got %q", ledger.Schema)
	}
	if len(ledger.Items) != 4 {
		t.Fatalf("expected four pressure items from fixture, got %#v", ledger.Items)
	}
	seenKinds := map[string]bool{}
	for _, item := range ledger.Items {
		if item.ID == "" || item.ProposedAction == "" {
			t.Fatalf("expected pressure item fields to be set, got %#v", item)
		}
		if item.Status != "observed" {
			t.Fatalf("expected feedback pressure to start observed, got %#v", item)
		}
		if !item.RequiresUserAcceptance {
			t.Fatalf("expected pressure item to require user acceptance, got %#v", item)
		}
		if len(item.EvidenceRefs) != 1 || !strings.HasPrefix(item.EvidenceRefs[0], "feedback:codex:") {
			t.Fatalf("expected feedback evidence ref, got %#v", item)
		}
		if len(item.RelatedCapabilities) != 1 || item.RelatedCapabilities[0] != "CAP-021" {
			t.Fatalf("expected related capability from feedback, got %#v", item)
		}
		if item.RelatedRisks == nil {
			t.Fatalf("expected related_risks field to be present, got %#v", item)
		}
		seenKinds[item.Kind] = true
	}
	for _, kind := range []string{"recurring_blocker", "validation_gap", "planning_gap"} {
		if !seenKinds[kind] {
			t.Fatalf("expected pressure kind %s in %#v", kind, ledger.Items)
		}
	}
}

func TestPressurePromoteRequiresExplicitCommand(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	if code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("feedback add expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"pressure", "promote", "P-001", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "promoted P-001 to repeated") {
		t.Fatalf("expected one-step promotion summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"pressure", "status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"status": "repeated"`) {
		t.Fatalf("expected promoted status in ledger, got %q", stdout.String())
	}
	if strings.Contains(stdout.String(), `"status": "accepted"`) {
		t.Fatalf("pressure became accepted without enough explicit promotes: %q", stdout.String())
	}
}

func TestPressureRetire(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	if code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("feedback add expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"pressure", "retire", "P-001", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "retired P-001") {
		t.Fatalf("expected retire summary, got %q", stdout.String())
	}
}

func TestFeedbackAddIsInert(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	contractPath := filepath.Join(dir, ".ni", "contract.json")
	lockPath := filepath.Join(dir, ".ni", "plan.lock.json")
	beforeContract := readFileForCLI(t, contractPath)
	beforeLock := readFileForCLI(t, lockPath)

	code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if got := readFileForCLI(t, contractPath); !bytes.Equal(got, beforeContract) {
		t.Fatalf("feedback add changed contract.json")
	}
	if got := readFileForCLI(t, lockPath); !bytes.Equal(got, beforeLock) {
		t.Fatalf("feedback add changed plan.lock.json")
	}
}

func TestFeedbackAddBlocksOnLockMismatch(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("changed after lock\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	var stderr bytes.Buffer
	code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &stderr)
	if code != exitStaleLock {
		t.Fatalf("expected exit code %d, got %d", exitStaleLock, code)
	}
	if !strings.Contains(stderr.String(), "BLOCKED: lock hash mismatch") {
		t.Fatalf("expected lock mismatch block, got %q", stderr.String())
	}
}

func TestExportHyperRunCreatesSeedDocsOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "hyper-seed")
	var stdout bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "hyper-run", "--out", out}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "exported hyper-run seed package") {
		t.Fatalf("expected export summary, got %q", stdout.String())
	}

	wantFiles := []string{
		"plan.md",
		"ni-context.md",
		"readiness-expectations.md",
		"evidence-requirements.md",
		"first-run-focus.md",
	}
	for _, name := range wantFiles {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("expected seed doc %s: %v", name, err)
		}
		if len(data) == 0 {
			t.Fatalf("expected seed doc %s to have content", name)
		}
	}

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	if len(entries) != len(wantFiles) {
		t.Fatalf("expected seed docs only, got %d entries", len(entries))
	}

	forbidden := map[string]struct{}{
		filepath.Join(".hyper", "goals", "GOAL-0001"): {},
		"tasks.md":    {},
		"evidence.md": {},
		"review.md":   {},
		"next.md":     {},
	}
	err = filepath.WalkDir(out, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(out, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if _, ok := forbidden[rel]; ok {
			t.Fatalf("export created forbidden runtime packet path %s", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking export directory: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, ".hyper", "goals")); !os.IsNotExist(err) {
		t.Fatalf("expected no .hyper/goals directory, stat err: %v", err)
	}
}

func TestExportHyperRunRejectsExistingRuntimePacketDirectory(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "hyper-seed")
	if err := os.MkdirAll(filepath.Join(out, ".hyper", "goals", "GOAL-0001"), 0o755); err != nil {
		t.Fatalf("creating stale runtime directory: %v", err)
	}

	var stderr bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "hyper-run", "--out", out}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "non-seed entry") {
		t.Fatalf("expected non-seed rejection, got %q", stderr.String())
	}
}

func TestExportRejectsUnsupportedTargetWithExitCode(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stderr bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "codex", "--out", filepath.Join(dir, "export")}, &bytes.Buffer{}, &stderr)
	if code != exitUnsupportedTarget {
		t.Fatalf("expected exit code %d, got %d", exitUnsupportedTarget, code)
	}
	if !strings.Contains(stderr.String(), `unsupported export target "codex"`) {
		t.Fatalf("expected unsupported export target error, got %q", stderr.String())
	}
}

func TestExportRefusesStaleLocksForAllTargets(t *testing.T) {
	for _, targetName := range []string{"hyper-run", "namba-ai", "ouroboros", "spec-kit"} {
		t.Run(targetName, func(t *testing.T) {
			dir := t.TempDir()
			if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
				t.Fatalf("init expected exit code 0, got %d", code)
			}
			writeReadyContractForCLI(t, dir)
			if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
				t.Fatalf("end expected exit code 0, got %d", code)
			}
			if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("changed after lock\n"), 0o644); err != nil {
				t.Fatalf("changing locked doc: %v", err)
			}

			out := filepath.Join(dir, "export", targetName)
			var stderr bytes.Buffer
			code := run([]string{"export", "--dir", dir, "--target", targetName, "--out", out}, &bytes.Buffer{}, &stderr)
			if code != exitStaleLock {
				t.Fatalf("expected stale lock exit code %d, got %d", exitStaleLock, code)
			}
			if !strings.Contains(stderr.String(), "BLOCKED: lock hash mismatch") {
				t.Fatalf("expected stale lock rejection, got %q", stderr.String())
			}
			if _, err := os.Stat(out); !os.IsNotExist(err) {
				t.Fatalf("stale export should not create output dir, stat err: %v", err)
			}
		})
	}
}

func TestExportCommandDoesNotInvokeExternalBinaries(t *testing.T) {
	for _, path := range []string{
		filepath.Join("..", "..", "cmd", "ni", "main.go"),
		filepath.Join("..", "..", "internal", "core", "exporter", "exporter.go"),
	} {
		t.Run(path, func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("reading %s: %v", path, err)
			}
			text := string(data)
			for _, forbidden := range []string{
				`"os/exec"`,
				"exec.Command",
				"exec.CommandContext",
				"os.StartProcess",
				"syscall.Exec",
			} {
				if strings.Contains(text, forbidden) {
					t.Fatalf("export path must not invoke external binaries; found %q in %s", forbidden, path)
				}
			}
		})
	}
}

func TestExportNambaAICreatesSeedDocsOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "namba-seed")
	var stdout bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "namba-ai", "--out", out}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "exported namba-ai seed package") {
		t.Fatalf("expected export summary, got %q", stdout.String())
	}

	wantFiles := []string{
		"planning.md",
		"ni-lock-summary.md",
		"capability-map.md",
		"evaluation-map.md",
		"risk-map.md",
		"suggested-spec-boundaries.md",
	}
	for _, name := range wantFiles {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("expected seed doc %s: %v", name, err)
		}
		if len(data) == 0 {
			t.Fatalf("expected seed doc %s to have content", name)
		}
	}

	planning, err := os.ReadFile(filepath.Join(out, "planning.md"))
	if err != nil {
		t.Fatalf("reading planning seed: %v", err)
	}
	for _, want := range []string{
		"namba-ai-oriented planning seed material",
		"NI does not call namba",
		"run, sync, pr, or land behavior",
		"does not add Codex-only assumptions",
	} {
		if !strings.Contains(string(planning), want) {
			t.Fatalf("expected planning.md to contain %q, got %q", want, string(planning))
		}
	}

	boundaries, err := os.ReadFile(filepath.Join(out, "suggested-spec-boundaries.md"))
	if err != nil {
		t.Fatalf("reading suggested boundaries: %v", err)
	}
	for _, want := range []string{
		"proposal, not a required sequential SPEC chain",
		"candidate graph boundaries",
		"depends_on",
	} {
		if !strings.Contains(string(boundaries), want) {
			t.Fatalf("expected suggested boundaries to contain %q, got %q", want, string(boundaries))
		}
	}

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	if len(entries) != len(wantFiles) {
		t.Fatalf("expected seed docs only, got %d entries", len(entries))
	}
	for _, forbidden := range []string{"run.md", "sync.md", "pr.md", "land.md"} {
		if _, err := os.Stat(filepath.Join(out, forbidden)); !os.IsNotExist(err) {
			t.Fatalf("expected no namba execution file %s, stat err: %v", forbidden, err)
		}
	}
}

func TestExportSpecKitCreatesSeedNotesOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "spec-kit-seed")
	var stdout bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "spec-kit", "--out", out}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "exported spec-kit seed package") {
		t.Fatalf("expected export summary, got %q", stdout.String())
	}

	assertSeedNotesExport(t, out, "spec-kit-seed-notes.md", []string{
		"# Spec Kit Seed Notes",
		"## Locked Contract Summary",
		"## Capabilities",
		"## Constraints",
		"## Risks",
		"## Evaluation Contract",
		"## Non-goals",
		"## Source-of-Truth References",
		"Do not implement slash commands",
		"Do not create coding task lists as NI core state",
	})
}

func TestExportOuroborosCreatesSeedNotesOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "ouroboros-seed")
	var stdout bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "ouroboros", "--out", out}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "exported ouroboros seed package") {
		t.Fatalf("expected export summary, got %q", stdout.String())
	}

	assertSeedNotesExport(t, out, "ouroboros-seed-notes.md", []string{
		"# Ouroboros Seed Notes",
		"## Locked Contract Summary",
		"## Capabilities",
		"## Constraints",
		"## Risks",
		"## Evaluation Contract",
		"## Non-goals",
		"## Source-of-Truth References",
		"Do not implement interview, crystallize, execute, evaluate, or evolve inside NI",
		"Do not execute downstream agents",
	})
}

func TestExportTargetConformanceBoundaries(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	cases := []struct {
		target         string
		seedFiles      []string
		forbiddenPaths []string
		assertContent  func(t *testing.T, out string)
	}{
		{
			target:    "hyper-run",
			seedFiles: []string{"plan.md", "ni-context.md", "readiness-expectations.md", "evidence-requirements.md", "first-run-focus.md"},
			forbiddenPaths: []string{
				filepath.Join(".hyper", "goals"),
				filepath.Join(".hyper", "goals", "GOAL-0001"),
				"tasks.md",
				"evidence.md",
				"review.md",
				"next.md",
			},
		},
		{
			target:    "namba-ai",
			seedFiles: []string{"planning.md", "ni-lock-summary.md", "capability-map.md", "evaluation-map.md", "risk-map.md", "suggested-spec-boundaries.md"},
			forbiddenPaths: []string{
				".namba",
				filepath.Join(".namba", "specs"),
				"SPEC-001.md",
				"SPEC-002.md",
				"SPEC_SEQUENCE.md",
				"specs",
				"tasks.md",
				"run.md",
				"sync.md",
				"pr.md",
				"land.md",
			},
			assertContent: func(t *testing.T, out string) {
				t.Helper()
				assertFileContains(t, filepath.Join(out, "suggested-spec-boundaries.md"), []string{
					"proposal, not a required sequential SPEC chain",
					"candidate graph boundaries",
					"depends_on",
					"Do not interpret this proposal as permission for NI to run namba",
				})
			},
		},
		{
			target:    "ouroboros",
			seedFiles: []string{"ouroboros-seed-notes.md"},
			forbiddenPaths: []string{
				".ouroboros",
				filepath.Join(".ouroboros", "runtime"),
				"execute",
				"execute.md",
				"evaluate",
				"evaluate.md",
				"evolve",
				"evolve.md",
				"runtime",
			},
		},
		{
			target:    "spec-kit",
			seedFiles: []string{"spec-kit-seed-notes.md"},
			forbiddenPaths: []string{
				".specify",
				filepath.Join(".specify", "specs"),
				filepath.Join(".specify", "memory"),
				filepath.Join(".github", "prompts"),
				filepath.Join(".claude", "commands"),
				filepath.Join(".codex", "commands"),
				"slash-commands.md",
				"commands",
				"specify.md",
				"plan.md",
				"tasks.md",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.target, func(t *testing.T) {
			out := filepath.Join(dir, "conformance", tc.target)
			var stdout bytes.Buffer
			code := run([]string{"export", "--dir", dir, "--target", tc.target, "--out", out}, &stdout, &bytes.Buffer{})
			if code != 0 {
				t.Fatalf("expected exit code 0, got %d", code)
			}
			if !strings.Contains(stdout.String(), "exported "+tc.target+" seed package") {
				t.Fatalf("expected export summary, got %q", stdout.String())
			}

			assertExportFilesExactly(t, out, tc.seedFiles)
			assertExportBundleContains(t, out, []string{".ni/plan.lock.json", "sha256:"})
			assertExportOmitsPaths(t, out, tc.forbiddenPaths)
			if tc.assertContent != nil {
				tc.assertContent(t, out)
			}
		})
	}
}

func assertSeedNotesExport(t *testing.T, out string, wantFile string, wantContent []string) {
	t.Helper()

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	if len(entries) != 1 || entries[0].Name() != wantFile {
		t.Fatalf("expected only %s, got %#v", wantFile, entries)
	}
	if entries[0].IsDir() {
		t.Fatalf("expected %s to be a file", wantFile)
	}

	data, err := os.ReadFile(filepath.Join(out, wantFile))
	if err != nil {
		t.Fatalf("reading %s: %v", wantFile, err)
	}
	text := string(data)
	for _, want := range wantContent {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %s to contain %q, got %q", wantFile, want, text)
		}
	}

	for _, forbidden := range []string{
		"tasks.md",
		"execute.md",
		"evaluate.md",
		"evolve.md",
		"interview.md",
		"crystallize.md",
		"slash-commands.md",
	} {
		if _, err := os.Stat(filepath.Join(out, forbidden)); !os.IsNotExist(err) {
			t.Fatalf("expected no executable workflow file %s, stat err: %v", forbidden, err)
		}
	}
}

func assertExportFilesExactly(t *testing.T, out string, wantFiles []string) {
	t.Helper()

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	if len(entries) != len(wantFiles) {
		t.Fatalf("expected %d seed files, got %d entries", len(wantFiles), len(entries))
	}

	want := map[string]struct{}{}
	for _, name := range wantFiles {
		want[name] = struct{}{}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			t.Fatalf("expected seed file, got directory %s", entry.Name())
		}
		if _, ok := want[entry.Name()]; !ok {
			t.Fatalf("unexpected export file %s", entry.Name())
		}
	}
	for _, name := range wantFiles {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("reading expected seed file %s: %v", name, err)
		}
		if len(data) == 0 {
			t.Fatalf("expected seed file %s to have content", name)
		}
	}
}

func assertExportOmitsPaths(t *testing.T, out string, forbiddenPaths []string) {
	t.Helper()

	forbidden := map[string]struct{}{}
	for _, path := range forbiddenPaths {
		forbidden[path] = struct{}{}
	}
	err := filepath.WalkDir(out, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(out, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		for forbiddenPath := range forbidden {
			if rel == forbiddenPath || strings.HasPrefix(rel, forbiddenPath+string(os.PathSeparator)) {
				t.Fatalf("export created forbidden runtime state path %s", rel)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking export directory: %v", err)
	}

	for _, path := range forbiddenPaths {
		if _, err := os.Stat(filepath.Join(out, path)); !os.IsNotExist(err) {
			t.Fatalf("expected no forbidden export path %s, stat err: %v", path, err)
		}
	}
}

func assertExportBundleContains(t *testing.T, out string, wants []string) {
	t.Helper()

	var combined strings.Builder
	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join(out, entry.Name()))
		if err != nil {
			t.Fatalf("reading export file %s: %v", entry.Name(), err)
		}
		combined.Write(data)
		combined.WriteByte('\n')
	}
	text := combined.String()
	for _, want := range wants {
		if !strings.Contains(text, want) {
			t.Fatalf("expected export bundle to contain %q, got %q", want, text)
		}
	}
}

func assertFileContains(t *testing.T, path string, wants []string) {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	text := string(data)
	for _, want := range wants {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %s to contain %q, got %q", path, want, text)
		}
	}
}

func TestTargets(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"targets"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	for _, name := range []string{"generic", "codex", "human-team", "hyper-run", "namba-ai", "ouroboros", "spec-kit"} {
		if !strings.Contains(stdout.String(), name) {
			t.Fatalf("expected target %q in output, got %q", name, stdout.String())
		}
	}
}

func TestTargetsJSON(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"targets", "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload []map[string]string
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON targets, got %v: %q", err, stdout.String())
	}
	if len(payload) != 7 {
		t.Fatalf("expected 7 targets, got %d: %#v", len(payload), payload)
	}
	if payload[0]["name"] != "generic" {
		t.Fatalf("expected built-in target order, got %#v", payload)
	}
}

func TestDiffCommandJSON(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{
		"diff",
		"--base", collabFixtureForCLI("base.json"),
		"--head", collabFixtureForCLI("non_conflicting_parallel_head.json"),
		"--json",
	}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected diff exit code 0, got %d", code)
	}
	var payload struct {
		Schema  string `json:"schema"`
		Changes []struct {
			Kind       string `json:"kind"`
			EntityType string `json:"entity_type"`
			ID         string `json:"id"`
		} `json:"changes"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid diff JSON, got %v: %q", err, stdout.String())
	}
	if payload.Schema != "ni.collaboration.diff.v0" {
		t.Fatalf("expected diff schema, got %#v", payload)
	}
	found := false
	for _, change := range payload.Changes {
		if change.Kind == "added" && change.EntityType == "capability" && change.ID == "CAP-002" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected CAP-002 added change, got %#v", payload.Changes)
	}
}

func TestConflictsCommandExitsNonzeroOnConflict(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{
		"conflicts",
		"--base", collabFixtureForCLI("base.json"),
		"--head", collabFixtureForCLI("conflicting_decision_head.json"),
	}, &stdout, &bytes.Buffer{})
	if code != exitSemanticConflict {
		t.Fatalf("expected conflicts exit code %d, got %d", exitSemanticConflict, code)
	}
	if !strings.Contains(stdout.String(), "collaboration conflicts") || !strings.Contains(stdout.String(), "DEC-001") {
		t.Fatalf("expected conflict text, got %q", stdout.String())
	}
}

func TestConflictsCommandJSONNoConflicts(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{
		"conflicts",
		"--base", collabFixtureForCLI("base.json"),
		"--head", collabFixtureForCLI("non_conflicting_parallel_head.json"),
		"--json",
	}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected conflicts exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"conflicts": []`) {
		t.Fatalf("expected empty conflicts JSON, got %q", stdout.String())
	}
}

func TestGraph(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"graph", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "work graph proposal") {
		t.Fatalf("expected graph output, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "CAP-001") {
		t.Fatalf("expected capability node, got %q", stdout.String())
	}
}

func TestGraphJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"graph", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"nodes"`) {
		t.Fatalf("expected JSON graph output, got %q", stdout.String())
	}
}

func TestHarnessPlan(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"harness", "plan", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "generated harness proposal") {
		t.Fatalf("expected harness output, got %q", stdout.String())
	}
}

func TestHarnessPlanJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"harness", "plan", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"source_lock_hash"`) {
		t.Fatalf("expected JSON harness output, got %q", stdout.String())
	}
}

func TestHarnessCandidateCommands(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}
	writeAcceptedHarnessPressureForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"harness", "propose", "--dir", dir, "--from-pressure", "P-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected propose exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "proposed harness candidate HC-001") {
		t.Fatalf("expected propose summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "accept", "--dir", dir, "HC-001"}, &stdout, &bytes.Buffer{})
	if code == 0 {
		t.Fatal("expected accept before validation to fail")
	}

	evidencePath := filepath.Join(dir, "harness-evidence.txt")
	if err := os.WriteFile(evidencePath, []byte("validated by CLI test\n"), 0o644); err != nil {
		t.Fatalf("writing evidence: %v", err)
	}
	stdout.Reset()
	code = run([]string{"harness", "validate", "--dir", dir, "HC-001", "--evidence", "harness-evidence.txt"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected validate exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "validated harness candidate HC-001 to validated_candidate") {
		t.Fatalf("expected validate summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "accept", "--dir", dir, "HC-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected accept exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "accepted harness candidate HC-001 as user_accepted") {
		t.Fatalf("expected accept summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "candidates", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected candidates exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"active_rule_id": "HC-001"`) {
		t.Fatalf("expected active rule id in candidates JSON, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "retire", "--dir", dir, "HC-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected retire exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "retired harness candidate HC-001") {
		t.Fatalf("expected retire summary, got %q", stdout.String())
	}
}

func TestUnknownCommand(t *testing.T) {
	var stderr bytes.Buffer
	code := run([]string{"serve"}, &bytes.Buffer{}, &stderr)
	if code != exitUsageError {
		t.Fatalf("expected exit code %d, got %d", exitUsageError, code)
	}
	if !strings.Contains(stderr.String(), "unknown command: serve") {
		t.Fatalf("expected unknown command error, got %q", stderr.String())
	}
}

func writeAcceptedHarnessPressureForCLI(t *testing.T, dir string) {
	t.Helper()

	data := []byte(`{
  "schema": "ni.pressure.v0",
  "items": [
    {
      "id": "P-001",
      "kind": "harness_candidate",
      "status": "accepted",
      "evidence_refs": ["cli:test"],
      "related_capabilities": ["CAP-001"],
      "related_risks": ["RISK-001"],
      "proposed_action": "Create an inert downstream harness proposal for CLI tests.",
      "requires_user_acceptance": true
    }
  ]
}
`)
	if err := os.WriteFile(filepath.Join(dir, ".ni", "pressure.json"), data, 0o644); err != nil {
		t.Fatalf("writing pressure ledger: %v", err)
	}
}

func writeReadyContractForCLI(t *testing.T, dir string) {
	t.Helper()

	contract := map[string]any{
		"schema":            "ni.contract.v0",
		"readiness_profile": "prototype",
		"product_type":      "software",
		"delivery_surfaces": []string{"cli"},
		"interaction_mode":  "human_to_system",
		"project": map[string]any{
			"id":      "cli-fixture",
			"name":    "CLI Fixture",
			"purpose": "Exercise ni end.",
			"status":  "draft",
		},
		"non_goals": []any{
			map[string]any{"id": "NG-001", "title": "Do not execute work."},
		},
		"capabilities": []any{
			map[string]any{
				"id":           "CAP-001",
				"title":        "Capability",
				"status":       "accepted",
				"requirements": []string{"REQ-001"},
				"evaluations":  []string{"EVAL-001"},
				"risks":        []string{},
				"artifacts":    []string{"ART-001"},
			},
		},
		"requirements":   []any{map[string]any{"id": "REQ-001", "title": "Requirement", "status": "accepted"}},
		"decisions":      []any{map[string]any{"id": "DEC-001", "title": "Decision", "status": "accepted"}},
		"risks":          []any{},
		"evaluations":    []any{map[string]any{"id": "EVAL-001", "title": "Evaluation", "method": "fixture"}},
		"artifacts":      []any{map[string]any{"id": "ART-001", "path": "docs/plan/", "kind": "planning_docs"}},
		"open_questions": []any{},
	}
	data, err := json.MarshalIndent(contract, "", "  ")
	if err != nil {
		t.Fatalf("marshaling ready contract: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), append(data, '\n'), 0o644); err != nil {
		t.Fatalf("writing ready contract: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("# Project brief\n\n## Product type\n\nsoftware\n\n## Delivery surfaces\n\n- cli\n\n## Purpose\n\nExercise ni end.\n"), 0o644); err != nil {
		t.Fatalf("writing ready project brief: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "01_actors_outcomes.md"), []byte("# Actors and outcomes\n\n## Actors\n\n- User: reviews the CLI fixture.\n- CLI: validates readiness.\n\n## Outcomes\n\n- Exercise ni end.\n"), 0o644); err != nil {
		t.Fatalf("writing ready actors doc: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "06_risks_security.md"), []byte("# Risks and security\n\nNo accepted risks are listed in this fixture.\n"), 0o644); err != nil {
		t.Fatalf("writing ready risk doc: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "08_delivery_operation.md"), []byte("# Delivery and operation\n\n## Delivery surfaces\n\n- cli\n\n## Initial delivery\n\nThe CLI fixture is reviewed before lock.\n"), 0o644); err != nil {
		t.Fatalf("writing ready delivery doc: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "10_open_questions.md"), []byte("# Open questions\n\nNo open questions are listed in this fixture.\n"), 0o644); err != nil {
		t.Fatalf("writing ready open question doc: %v", err)
	}
}

func writeReadyWithDeferralsContractForCLI(t *testing.T, dir string) {
	t.Helper()

	writeReadyContractForCLI(t, dir)
	path := filepath.Join(dir, ".ni", "contract.json")
	data := readFileForCLI(t, path)
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("parsing contract: %v", err)
	}
	payload["decisions"] = []any{
		map[string]any{"id": "DEC-001", "title": "Decision", "status": "deferred"},
	}
	payload["open_questions"] = []any{
		map[string]any{
			"id":      "OQ-001",
			"title":   "Non-blocking question",
			"blocker": false,
			"status":  "open",
		},
	}
	updated, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatalf("marshaling ready-with-deferrals contract: %v", err)
	}
	if err := os.WriteFile(path, append(updated, '\n'), 0o644); err != nil {
		t.Fatalf("writing ready-with-deferrals contract: %v", err)
	}
	decisionDoc := "# Decision log\n\n## DEC-001: Decision\n\nStatus: deferred\n"
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "11_decision_log.md"), []byte(decisionDoc), 0o644); err != nil {
		t.Fatalf("writing ready-with-deferrals decision doc: %v", err)
	}
	openQuestionDoc := "# Open questions\n\n## OQ-001: Non-blocking question\n\nBlocker: false\n\nStatus: open\n"
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "10_open_questions.md"), []byte(openQuestionDoc), 0o644); err != nil {
		t.Fatalf("writing ready-with-deferrals open question doc: %v", err)
	}
}

func updateContractForCLI(t *testing.T, dir string, mutate func(map[string]any)) {
	t.Helper()

	path := filepath.Join(dir, ".ni", "contract.json")
	data := readFileForCLI(t, path)
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("parsing contract: %v", err)
	}
	mutate(payload)
	updated, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatalf("marshaling contract: %v", err)
	}
	if err := os.WriteFile(path, append(updated, '\n'), 0o644); err != nil {
		t.Fatalf("writing contract: %v", err)
	}
}

func completeAmendmentForCLI(t *testing.T, dir string, id string) {
	t.Helper()

	path := filepath.Join(dir, ".ni", "amendments", id+".json")
	data := readFileForCLI(t, path)
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("parsing amendment: %v", err)
	}
	payload["reason"] = "Locked prompt compiler acceptance needs explicit clarification."
	payload["affected_docs"] = []string{"docs/plan/02_capabilities.md"}
	payload["affected_contract_ids"] = []string{"CAP-001", "REQ-001"}
	payload["proposed_changes"] = []string{"Clarify CAP-001 planning text without changing readiness rules."}
	payload["risk_impact"] = "No new high-severity risk; existing pre-runtime boundary remains in force."
	payload["readiness_impact"] = "Readiness remains READY after deterministic status evaluation."
	payload["created_from_feedback_refs"] = []string{}
	payload["created_from_pressure_refs"] = []string{}

	updated, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatalf("marshaling amendment: %v", err)
	}
	if err := os.WriteFile(path, append(updated, '\n'), 0o644); err != nil {
		t.Fatalf("writing completed amendment: %v", err)
	}
}

func writeBlockedQuestionForCLI(t *testing.T, dir string) {
	t.Helper()

	path := filepath.Join(dir, ".ni", "contract.json")
	data := readFileForCLI(t, path)
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("parsing contract: %v", err)
	}
	payload["open_questions"] = []any{
		map[string]any{
			"id":      "OQ-001",
			"title":   "Blocking relock question",
			"blocker": true,
			"status":  "open",
		},
	}
	updated, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatalf("marshaling blocked contract: %v", err)
	}
	if err := os.WriteFile(path, append(updated, '\n'), 0o644); err != nil {
		t.Fatalf("writing blocked contract: %v", err)
	}
}

func readFileForCLI(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	return data
}

func decodeErrorEnvelope(t *testing.T, data []byte) errorEnvelope {
	t.Helper()
	var envelope errorEnvelope
	if err := json.Unmarshal(data, &envelope); err != nil {
		t.Fatalf("expected error envelope JSON, got %v: %q", err, string(data))
	}
	if envelope.Error.Code == "" || envelope.Error.ExitCode == 0 || envelope.Error.Message == "" {
		t.Fatalf("expected populated error envelope, got %#v", envelope)
	}
	return envelope
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func collabFixtureForCLI(name string) string {
	return filepath.Join("..", "..", "internal", "core", "collab", "testdata", name)
}
