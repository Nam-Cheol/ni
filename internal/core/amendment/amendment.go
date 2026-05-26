package amendment

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"ni/internal/core/lock"
)

const Schema = "ni.amendment.v0"

const (
	StatusDraft   = "draft"
	StatusApplied = "applied"
)

type Amendment struct {
	Schema                  string   `json:"schema"`
	ID                      string   `json:"id"`
	Title                   string   `json:"title"`
	Reason                  string   `json:"reason"`
	AffectedDocs            []string `json:"affected_docs"`
	AffectedContractIDs     []string `json:"affected_contract_ids"`
	ProposedChanges         []string `json:"proposed_changes"`
	RiskImpact              string   `json:"risk_impact"`
	ReadinessImpact         string   `json:"readiness_impact"`
	CreatedFromFeedbackRefs []string `json:"created_from_feedback_refs"`
	CreatedFromPressureRefs []string `json:"created_from_pressure_refs"`
	Status                  string   `json:"status"`
	CreatedAt               string   `json:"created_at"`
	AppliedAt               string   `json:"applied_at,omitempty"`
	SourceLockHash          string   `json:"source_lock_hash,omitempty"`
}

type CreateOptions struct {
	Dir   string
	Title string
	Now   time.Time
}

func Create(opts CreateOptions) (Amendment, error) {
	title := strings.TrimSpace(opts.Title)
	if title == "" {
		return Amendment{}, fmt.Errorf("missing amendment title")
	}
	if opts.Now.IsZero() {
		opts.Now = time.Now().UTC()
	}
	root := cleanDir(opts.Dir)
	items, err := List(root)
	if err != nil {
		return Amendment{}, err
	}
	sourceLockHash, err := sourceLockHashIfPresent(root)
	if err != nil {
		return Amendment{}, err
	}

	item := Amendment{
		Schema:                  Schema,
		ID:                      nextID(items),
		Title:                   title,
		Reason:                  "",
		AffectedDocs:            []string{},
		AffectedContractIDs:     []string{},
		ProposedChanges:         []string{},
		RiskImpact:              "",
		ReadinessImpact:         "",
		CreatedFromFeedbackRefs: []string{},
		CreatedFromPressureRefs: []string{},
		Status:                  StatusDraft,
		CreatedAt:               opts.Now.UTC().Format(time.RFC3339),
		SourceLockHash:          sourceLockHash,
	}
	if err := Save(root, item); err != nil {
		return Amendment{}, err
	}
	return item, nil
}

func List(dir string) ([]Amendment, error) {
	root := cleanDir(dir)
	entries, err := os.ReadDir(StoreDir(root))
	if err != nil {
		if os.IsNotExist(err) {
			return []Amendment{}, nil
		}
		return nil, err
	}
	items := make([]Amendment, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		item, err := Load(root, strings.TrimSuffix(entry.Name(), ".json"))
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	return items, nil
}

func Load(dir string, id string) (Amendment, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return Amendment{}, fmt.Errorf("missing amendment id")
	}
	path := Path(dir, id)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Amendment{}, fmt.Errorf("unknown amendment id %q", id)
		}
		return Amendment{}, err
	}
	var item Amendment
	if err := json.Unmarshal(data, &item); err != nil {
		return Amendment{}, fmt.Errorf("malformed amendment JSON: %w", err)
	}
	item.normalize()
	if err := item.Validate(false); err != nil {
		return Amendment{}, err
	}
	return item, nil
}

func Apply(dir string, id string, now time.Time) (Amendment, error) {
	if now.IsZero() {
		now = time.Now().UTC()
	}
	root := cleanDir(dir)
	item, err := Load(root, id)
	if err != nil {
		return Amendment{}, err
	}
	if item.Status == StatusApplied {
		return item, nil
	}
	if err := item.Validate(true); err != nil {
		return Amendment{}, err
	}
	if item.SourceLockHash != "" {
		currentHash, err := lock.CurrentLockHash(root)
		if err != nil {
			return Amendment{}, err
		}
		if currentHash != item.SourceLockHash {
			return Amendment{}, fmt.Errorf("BLOCKED: amendment %s belongs to a different source lock", item.ID)
		}
	}
	item.Status = StatusApplied
	item.AppliedAt = now.UTC().Format(time.RFC3339)
	if err := Save(root, item); err != nil {
		return Amendment{}, err
	}
	return item, nil
}

func HasAppliedForLock(dir string, lockHash string) (bool, error) {
	items, err := List(dir)
	if err != nil {
		return false, err
	}
	for _, item := range items {
		if item.Status == StatusApplied && item.SourceLockHash == lockHash {
			return true, nil
		}
	}
	return false, nil
}

func Save(dir string, item Amendment) error {
	root := cleanDir(dir)
	item.normalize()
	if err := item.Validate(false); err != nil {
		return err
	}
	if err := os.MkdirAll(StoreDir(root), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(Path(root, item.ID), append(data, '\n'), 0o644)
}

func StoreDir(dir string) string {
	return filepath.Join(cleanDir(dir), ".ni", "amendments")
}

func Path(dir string, id string) string {
	return filepath.Join(StoreDir(dir), strings.TrimSpace(id)+".json")
}

func FormatList(items []Amendment) string {
	if len(items) == 0 {
		return "no amendments recorded\n"
	}
	var b strings.Builder
	for _, item := range items {
		fmt.Fprintf(&b, "%s\t%s\t%s\n", item.ID, item.Status, item.Title)
	}
	return b.String()
}

func (a *Amendment) normalize() {
	if a.Schema == "" {
		a.Schema = Schema
	}
	a.ID = strings.TrimSpace(a.ID)
	a.Title = strings.TrimSpace(a.Title)
	a.Reason = strings.TrimSpace(a.Reason)
	a.RiskImpact = strings.TrimSpace(a.RiskImpact)
	a.ReadinessImpact = strings.TrimSpace(a.ReadinessImpact)
	a.Status = strings.TrimSpace(a.Status)
	if a.Status == "" {
		a.Status = StatusDraft
	}
	a.AffectedDocs = cleanStrings(a.AffectedDocs)
	a.AffectedContractIDs = cleanStrings(a.AffectedContractIDs)
	a.ProposedChanges = cleanStrings(a.ProposedChanges)
	a.CreatedFromFeedbackRefs = cleanStrings(a.CreatedFromFeedbackRefs)
	a.CreatedFromPressureRefs = cleanStrings(a.CreatedFromPressureRefs)
}

func (a Amendment) Validate(requireComplete bool) error {
	var missing []string
	if a.Schema == "" {
		missing = append(missing, "schema")
	}
	if a.ID == "" {
		missing = append(missing, "id")
	}
	if a.Title == "" {
		missing = append(missing, "title")
	}
	if a.Status == "" {
		missing = append(missing, "status")
	}
	if a.CreatedAt == "" {
		missing = append(missing, "created_at")
	}
	if requireComplete {
		if a.Reason == "" {
			missing = append(missing, "reason")
		}
		if len(a.AffectedDocs) == 0 && len(a.AffectedContractIDs) == 0 {
			missing = append(missing, "affected_docs or affected_contract_ids")
		}
		if len(a.ProposedChanges) == 0 {
			missing = append(missing, "proposed_changes")
		}
		if a.RiskImpact == "" {
			missing = append(missing, "risk_impact")
		}
		if a.ReadinessImpact == "" {
			missing = append(missing, "readiness_impact")
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("amendment missing required field(s): %s", strings.Join(missing, ", "))
	}
	if a.Schema != Schema {
		return fmt.Errorf("unsupported amendment schema %q", a.Schema)
	}
	if !strings.HasPrefix(a.ID, "AMEND-") {
		return fmt.Errorf("amendment id %q must use AMEND- prefix", a.ID)
	}
	if a.Status != StatusDraft && a.Status != StatusApplied {
		return fmt.Errorf("invalid amendment status %q", a.Status)
	}
	if _, err := time.Parse(time.RFC3339, a.CreatedAt); err != nil {
		return fmt.Errorf("invalid created_at: %w", err)
	}
	if a.AppliedAt != "" {
		if _, err := time.Parse(time.RFC3339, a.AppliedAt); err != nil {
			return fmt.Errorf("invalid applied_at: %w", err)
		}
	}
	return nil
}

func sourceLockHashIfPresent(dir string) (string, error) {
	hash, err := lock.CurrentLockHash(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return hash, nil
}

func nextID(items []Amendment) string {
	next := 1
	for _, item := range items {
		var n int
		if _, err := fmt.Sscanf(item.ID, "AMEND-%03d", &n); err == nil && n >= next {
			next = n + 1
		}
	}
	return fmt.Sprintf("AMEND-%03d", next)
}

func cleanDir(dir string) string {
	if strings.TrimSpace(dir) == "" {
		return "."
	}
	return filepath.Clean(dir)
}

func cleanStrings(values []string) []string {
	if values == nil {
		return []string{}
	}
	cleaned := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}
