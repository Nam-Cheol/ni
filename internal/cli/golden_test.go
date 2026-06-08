package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"

	"ni/internal/core/lock"
)

const updateGoldenEnv = "NI_UPDATE_GOLDEN"

var (
	goldenHashPattern      = regexp.MustCompile(`\b[a-f0-9]{64}\b`)
	goldenTimestampPattern = regexp.MustCompile(`\b\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z\b`)
)

func TestGoldenStatusJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	got := runForGolden(t, []string{"status", "--dir", dir, "--json"}, 0)
	assertGolden(t, "status_json.golden", normalizeGolden(got, dir))
}

func TestGoldenRunTargets(t *testing.T) {
	for _, targetName := range []string{"generic", "codex", "human-team"} {
		t.Run(targetName, func(t *testing.T) {
			dir := newReadyLockedWorkspaceForGolden(t)

			got := runForGolden(t, []string{"run", "--dir", dir, "--target", targetName}, 0)
			assertGolden(t, "run_"+strings.ReplaceAll(targetName, "-", "_")+".golden", normalizeGolden(got, dir))
		})
	}
}

func TestGoldenExportTargets(t *testing.T) {
	for _, targetName := range []string{"hyper-run", "namba-ai", "ouroboros", "spec-kit"} {
		t.Run(targetName, func(t *testing.T) {
			dir := newReadyLockedWorkspaceForGolden(t)
			out := filepath.Join(dir, "export", targetName)

			stdout := runForGolden(t, []string{"export", "--dir", dir, "--target", targetName, "--out", out}, 0)
			got := exportBundleForGolden(t, stdout, out)
			assertGolden(t, "export_"+strings.ReplaceAll(targetName, "-", "_")+".golden", normalizeGolden(got, dir))
		})
	}
}

func TestGoldenConflictsJSON(t *testing.T) {
	got := runForGolden(t, []string{
		"conflicts",
		"--base", collabFixtureForCLI("base.json"),
		"--head", collabFixtureForCLI("conflicting_decision_head.json"),
		"--json",
	}, exitSemanticConflict)
	assertGolden(t, "conflicts_json.golden", normalizeGolden(got, ""))
}

func newReadyLockedWorkspaceForGolden(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	lockedAt := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	if _, err := lock.CreateAt(dir, lockedAt); err != nil {
		t.Fatalf("creating fixed lockfile: %v", err)
	}
	return dir
}

func runForGolden(t *testing.T, args []string, wantCode int) string {
	t.Helper()

	var stdout, stderr bytes.Buffer
	code := run(args, &stdout, &stderr)
	if code != wantCode {
		t.Fatalf("run(%q) exit code = %d, want %d; stdout=%q stderr=%q", strings.Join(args, " "), code, wantCode, stdout.String(), stderr.String())
	}
	if stderr.Len() > 0 {
		t.Fatalf("run(%q) wrote unexpected stderr: %q", strings.Join(args, " "), stderr.String())
	}
	return stdout.String()
}

func exportBundleForGolden(t *testing.T, stdout string, out string) string {
	t.Helper()

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export dir: %v", err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			t.Fatalf("unexpected directory in export bundle: %s", entry.Name())
		}
		names = append(names, entry.Name())
	}
	sort.Strings(names)

	var b strings.Builder
	b.WriteString("stdout:\n")
	b.WriteString(stdout)
	if !strings.HasSuffix(stdout, "\n") {
		b.WriteString("\n")
	}
	b.WriteString("\nfiles:\n")
	for _, name := range names {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("reading exported file %s: %v", name, err)
		}
		b.WriteString("--- ")
		b.WriteString(name)
		b.WriteString(" ---\n")
		b.Write(data)
		if !bytes.HasSuffix(data, []byte("\n")) {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func assertGolden(t *testing.T, name string, got string) {
	t.Helper()

	path := filepath.Join("..", "..", "testdata", "golden", name)
	got = ensureFinalNewline(got)
	if os.Getenv(updateGoldenEnv) == "1" {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("creating golden dir: %v", err)
		}
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatalf("updating golden %s: %v", name, err)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading golden %s: %v; set %s=1 to update", name, err, updateGoldenEnv)
	}
	want := string(data)
	if got != want {
		t.Fatalf("golden %s mismatch; set %s=1 to update\n--- got ---\n%s\n--- want ---\n%s", name, updateGoldenEnv, got, want)
	}
}

func normalizeGolden(text string, dir string) string {
	if dir != "" {
		text = strings.ReplaceAll(text, filepath.ToSlash(dir), "<WORKDIR>")
		text = strings.ReplaceAll(text, dir, "<WORKDIR>")
	}
	text = goldenTimestampPattern.ReplaceAllString(text, "<TIMESTAMP>")
	text = goldenHashPattern.ReplaceAllString(text, "<SHA256>")
	return text
}

func ensureFinalNewline(text string) string {
	if strings.HasSuffix(text, "\n") {
		return text
	}
	return text + "\n"
}
