package feedback

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ni/internal/core/lock"
	"ni/internal/core/pressure"
)

const Schema = "ni.feedback.v0"

type Entry struct {
	Schema                   string   `json:"schema"`
	IngestedAt               string   `json:"ingested_at"`
	SourceTarget             string   `json:"source_target"`
	RelatedLockHash          string   `json:"related_lock_hash,omitempty"`
	RelatedCapabilities      []string `json:"related_capabilities"`
	ObservedBlockers         []string `json:"observed_blockers"`
	ValidationGaps           []string `json:"validation_gaps"`
	RecurringFailures        []string `json:"recurring_failures"`
	SuggestedContractChanges []string `json:"suggested_contract_changes"`
	RawSummary               string   `json:"raw_summary"`
}

type AddOptions struct {
	Dir  string
	File string
	Now  time.Time
}

func Add(opts AddOptions) (Entry, error) {
	if strings.TrimSpace(opts.File) == "" {
		return Entry{}, fmt.Errorf("missing feedback file")
	}
	if err := verifyLockIfPresent(opts.Dir); err != nil {
		return Entry{}, err
	}

	entry, err := loadFeedbackFile(opts.File)
	if err != nil {
		return Entry{}, err
	}
	if opts.Now.IsZero() {
		opts.Now = time.Now().UTC()
	}
	entry.Schema = Schema
	entry.IngestedAt = opts.Now.UTC().Format(time.RFC3339)
	entry.normalize()
	if err := entry.Validate(); err != nil {
		return Entry{}, err
	}

	path := StorePath(opts.Dir)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return Entry{}, err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return Entry{}, err
	}
	defer file.Close()

	data, err := json.Marshal(entry)
	if err != nil {
		return Entry{}, err
	}
	if _, err := file.Write(append(data, '\n')); err != nil {
		return Entry{}, err
	}
	if _, err := pressure.AddObserved(opts.Dir, pressure.FeedbackItems(
		entry.SourceTarget,
		entry.IngestedAt,
		entry.RelatedCapabilities,
		entry.ObservedBlockers,
		entry.ValidationGaps,
		entry.RecurringFailures,
		entry.SuggestedContractChanges,
	)); err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func List(dir string) ([]Entry, error) {
	if err := verifyLockIfPresent(dir); err != nil {
		return nil, err
	}

	path := StorePath(dir)
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var entries []Entry
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		var entry Entry
		if err := json.Unmarshal([]byte(text), &entry); err != nil {
			return nil, fmt.Errorf("malformed feedback at line %d: %w", line, err)
		}
		if err := entry.Validate(); err != nil {
			return nil, fmt.Errorf("invalid feedback at line %d: %w", line, err)
		}
		entry.normalize()
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func FormatText(entries []Entry) string {
	if len(entries) == 0 {
		return "no feedback recorded\n"
	}
	var b strings.Builder
	for _, entry := range entries {
		fmt.Fprintf(&b, "%s\t%s", entry.IngestedAt, entry.SourceTarget)
		if entry.RelatedLockHash != "" {
			fmt.Fprintf(&b, "\tlock:%s", entry.RelatedLockHash)
		}
		if len(entry.RelatedCapabilities) > 0 {
			fmt.Fprintf(&b, "\tcapabilities:%s", strings.Join(entry.RelatedCapabilities, ","))
		}
		if entry.RawSummary != "" {
			fmt.Fprintf(&b, "\t%s", entry.RawSummary)
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func StorePath(dir string) string {
	root := strings.TrimSpace(dir)
	if root == "" {
		root = "."
	}
	return filepath.Join(filepath.Clean(root), ".ni", "feedback.jsonl")
}

func verifyLockIfPresent(dir string) error {
	root := strings.TrimSpace(dir)
	if root == "" {
		root = "."
	}
	lockPath := filepath.Join(filepath.Clean(root), ".ni", "plan.lock.json")
	if _, err := os.Stat(lockPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	verification, err := lock.Verify(root)
	if err != nil {
		return err
	}
	if len(verification.Mismatches) > 0 {
		return fmt.Errorf("BLOCKED: lock hash mismatch for %s", verification.Mismatches[0].Path)
	}
	return nil
}

func (e Entry) Validate() error {
	var missing []string
	if strings.TrimSpace(e.Schema) == "" {
		missing = append(missing, "schema")
	}
	if strings.TrimSpace(e.IngestedAt) == "" {
		missing = append(missing, "ingested_at")
	}
	if strings.TrimSpace(e.SourceTarget) == "" {
		missing = append(missing, "source_target")
	}
	if strings.TrimSpace(e.RawSummary) == "" {
		missing = append(missing, "raw_summary")
	}
	if len(missing) > 0 {
		return fmt.Errorf("feedback missing required field(s): %s", strings.Join(missing, ", "))
	}
	if e.Schema != Schema {
		return fmt.Errorf("unsupported feedback schema %q", e.Schema)
	}
	if _, err := time.Parse(time.RFC3339, e.IngestedAt); err != nil {
		return fmt.Errorf("invalid ingested_at: %w", err)
	}
	return nil
}

func (e *Entry) normalize() {
	if e.RelatedCapabilities == nil {
		e.RelatedCapabilities = []string{}
	}
	if e.ObservedBlockers == nil {
		e.ObservedBlockers = []string{}
	}
	if e.ValidationGaps == nil {
		e.ValidationGaps = []string{}
	}
	if e.RecurringFailures == nil {
		e.RecurringFailures = []string{}
	}
	if e.SuggestedContractChanges == nil {
		e.SuggestedContractChanges = []string{}
	}
}

func loadFeedbackFile(path string) (Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return Entry{}, err
	}
	defer file.Close()

	var entry Entry
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&entry); err != nil {
		return Entry{}, fmt.Errorf("malformed feedback JSON: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return Entry{}, fmt.Errorf("feedback file must contain one JSON object")
	}
	return entry, nil
}
