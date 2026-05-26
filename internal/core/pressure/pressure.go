package pressure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"ni/internal/core/lock"
)

const Schema = "ni.pressure.v0"

const (
	KindPlanningGap      = "planning_gap"
	KindValidationGap    = "validation_gap"
	KindRecurringBlocker = "recurring_blocker"
	KindReusablePattern  = "reusable_pattern"
	KindHarnessCandidate = "harness_candidate"
)

const (
	StatusObserved   = "observed"
	StatusRepeated   = "repeated"
	StatusPromotable = "promotable"
	StatusAccepted   = "accepted"
	StatusRetired    = "retired"
)

type Ledger struct {
	Schema string `json:"schema"`
	Items  []Item `json:"items"`
}

type Item struct {
	ID                     string   `json:"id"`
	Kind                   string   `json:"kind"`
	Status                 string   `json:"status"`
	EvidenceRefs           []string `json:"evidence_refs"`
	RelatedCapabilities    []string `json:"related_capabilities"`
	RelatedRisks           []string `json:"related_risks"`
	ProposedAction         string   `json:"proposed_action"`
	RequiresUserAcceptance bool     `json:"requires_user_acceptance"`
}

func Load(dir string) (Ledger, error) {
	if err := verifyLockIfPresent(dir); err != nil {
		return Ledger{}, err
	}
	return loadFile(StorePath(dir))
}

func AddObserved(dir string, items []Item) (Ledger, error) {
	if len(items) == 0 {
		return Load(dir)
	}
	if err := verifyLockIfPresent(dir); err != nil {
		return Ledger{}, err
	}
	ledger, err := loadFile(StorePath(dir))
	if err != nil {
		return Ledger{}, err
	}

	next := nextID(ledger.Items)
	for _, item := range items {
		item.ID = next
		item.Status = StatusObserved
		item.RequiresUserAcceptance = true
		item.normalize()
		if err := item.Validate(); err != nil {
			return Ledger{}, err
		}
		ledger.Items = append(ledger.Items, item)
		next = incrementID(next)
	}
	if err := Save(dir, ledger); err != nil {
		return Ledger{}, err
	}
	return ledger, nil
}

func Promote(dir string, id string) (Item, error) {
	return updateStatus(dir, id, promoteStatus)
}

func Retire(dir string, id string) (Item, error) {
	return updateStatus(dir, id, func(status string) (string, error) {
		if status == StatusRetired {
			return StatusRetired, nil
		}
		return StatusRetired, nil
	})
}

func Save(dir string, ledger Ledger) error {
	if err := verifyLockIfPresent(dir); err != nil {
		return err
	}
	ledger.normalize()
	if err := ledger.Validate(); err != nil {
		return err
	}
	path := StorePath(dir)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

func FormatText(ledger Ledger) string {
	if len(ledger.Items) == 0 {
		return "no pressure recorded\n"
	}
	items := append([]Item(nil), ledger.Items...)
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	var b strings.Builder
	fmt.Fprintf(&b, "pressure ledger: %d item(s)\n", len(items))
	for _, item := range items {
		fmt.Fprintf(&b, "%s\t%s\t%s\t%s", item.ID, item.Kind, item.Status, item.ProposedAction)
		if len(item.EvidenceRefs) > 0 {
			fmt.Fprintf(&b, "\tevidence:%s", strings.Join(item.EvidenceRefs, ","))
		}
		if len(item.RelatedCapabilities) > 0 {
			fmt.Fprintf(&b, "\tcapabilities:%s", strings.Join(item.RelatedCapabilities, ","))
		}
		if len(item.RelatedRisks) > 0 {
			fmt.Fprintf(&b, "\trisks:%s", strings.Join(item.RelatedRisks, ","))
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
	return filepath.Join(filepath.Clean(root), ".ni", "pressure.json")
}

func FeedbackItems(sourceTarget string, ingestedAt string, relatedCapabilities []string, observedBlockers []string, validationGaps []string, recurringFailures []string, suggestedContractChanges []string) []Item {
	evidenceRef := "feedback:" + strings.TrimSpace(sourceTarget) + ":" + strings.TrimSpace(ingestedAt)
	items := make([]Item, 0, len(observedBlockers)+len(validationGaps)+len(recurringFailures)+len(suggestedContractChanges))
	appendItems := func(kind string, values []string) {
		for _, value := range values {
			value = strings.TrimSpace(value)
			if value == "" {
				continue
			}
			items = append(items, Item{
				Kind:                   kind,
				Status:                 StatusObserved,
				EvidenceRefs:           []string{evidenceRef},
				RelatedCapabilities:    append([]string(nil), relatedCapabilities...),
				RelatedRisks:           []string{},
				ProposedAction:         value,
				RequiresUserAcceptance: true,
			})
		}
	}
	appendItems(KindRecurringBlocker, observedBlockers)
	appendItems(KindValidationGap, validationGaps)
	appendItems(KindRecurringBlocker, recurringFailures)
	appendItems(KindPlanningGap, suggestedContractChanges)
	return items
}

func loadFile(path string) (Ledger, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Ledger{Schema: Schema, Items: []Item{}}, nil
		}
		return Ledger{}, err
	}
	defer file.Close()

	var ledger Ledger
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&ledger); err != nil {
		return Ledger{}, fmt.Errorf("malformed pressure ledger JSON: %w", err)
	}
	ledger.normalize()
	if err := ledger.Validate(); err != nil {
		return Ledger{}, err
	}
	return ledger, nil
}

func updateStatus(dir string, id string, transition func(string) (string, error)) (Item, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return Item{}, fmt.Errorf("missing pressure id")
	}
	ledger, err := Load(dir)
	if err != nil {
		return Item{}, err
	}
	for i := range ledger.Items {
		if ledger.Items[i].ID != id {
			continue
		}
		next, err := transition(ledger.Items[i].Status)
		if err != nil {
			return Item{}, err
		}
		ledger.Items[i].Status = next
		ledger.Items[i].normalize()
		if err := Save(dir, ledger); err != nil {
			return Item{}, err
		}
		return ledger.Items[i], nil
	}
	return Item{}, fmt.Errorf("unknown pressure id %q", id)
}

func promoteStatus(status string) (string, error) {
	switch status {
	case StatusObserved:
		return StatusRepeated, nil
	case StatusRepeated:
		return StatusPromotable, nil
	case StatusPromotable:
		return StatusAccepted, nil
	case StatusAccepted:
		return StatusAccepted, nil
	case StatusRetired:
		return "", fmt.Errorf("retired pressure cannot be promoted")
	default:
		return "", fmt.Errorf("invalid pressure status %q", status)
	}
}

func (l *Ledger) normalize() {
	if l.Schema == "" {
		l.Schema = Schema
	}
	if l.Items == nil {
		l.Items = []Item{}
	}
	for i := range l.Items {
		l.Items[i].normalize()
	}
}

func (l Ledger) Validate() error {
	if strings.TrimSpace(l.Schema) == "" {
		return fmt.Errorf("pressure ledger missing schema")
	}
	if l.Schema != Schema {
		return fmt.Errorf("unsupported pressure ledger schema %q", l.Schema)
	}
	seen := map[string]struct{}{}
	for _, item := range l.Items {
		if err := item.Validate(); err != nil {
			return err
		}
		if _, ok := seen[item.ID]; ok {
			return fmt.Errorf("duplicate pressure id %q", item.ID)
		}
		seen[item.ID] = struct{}{}
	}
	return nil
}

func (i *Item) normalize() {
	if i.EvidenceRefs == nil {
		i.EvidenceRefs = []string{}
	}
	if i.RelatedCapabilities == nil {
		i.RelatedCapabilities = []string{}
	}
	if i.RelatedRisks == nil {
		i.RelatedRisks = []string{}
	}
}

func (i Item) Validate() error {
	var missing []string
	if strings.TrimSpace(i.ID) == "" {
		missing = append(missing, "id")
	}
	if strings.TrimSpace(i.Kind) == "" {
		missing = append(missing, "kind")
	}
	if strings.TrimSpace(i.Status) == "" {
		missing = append(missing, "status")
	}
	if strings.TrimSpace(i.ProposedAction) == "" {
		missing = append(missing, "proposed_action")
	}
	if len(missing) > 0 {
		return fmt.Errorf("pressure item missing required field(s): %s", strings.Join(missing, ", "))
	}
	if !validKind(i.Kind) {
		return fmt.Errorf("invalid pressure kind %q", i.Kind)
	}
	if !validStatus(i.Status) {
		return fmt.Errorf("invalid pressure status %q", i.Status)
	}
	return nil
}

func validKind(kind string) bool {
	switch kind {
	case KindPlanningGap, KindValidationGap, KindRecurringBlocker, KindReusablePattern, KindHarnessCandidate:
		return true
	default:
		return false
	}
}

func validStatus(status string) bool {
	switch status {
	case StatusObserved, StatusRepeated, StatusPromotable, StatusAccepted, StatusRetired:
		return true
	default:
		return false
	}
}

func nextID(items []Item) string {
	max := 0
	for _, item := range items {
		if !strings.HasPrefix(item.ID, "P-") {
			continue
		}
		var n int
		if _, err := fmt.Sscanf(item.ID, "P-%03d", &n); err == nil && n > max {
			max = n
		}
	}
	return fmt.Sprintf("P-%03d", max+1)
}

func incrementID(id string) string {
	var n int
	if _, err := fmt.Sscanf(id, "P-%03d", &n); err != nil {
		return "P-001"
	}
	return fmt.Sprintf("P-%03d", n+1)
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
