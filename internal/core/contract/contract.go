package contract

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"ni/internal/core/profile"
)

const Schema = "ni.contract.v0"

const (
	DefaultProductType     = "software"
	DefaultDeliverySurface = "cli"
	DefaultInteractionMode = "human_to_system"
)

var supportedProductTypes = map[string]struct{}{
	"software":             {},
	"conversation_product": {},
	"research_protocol":    {},
	"operations_process":   {},
	"education_program":    {},
	"document_product":     {},
	"physical_product":     {},
	"mixed":                {},
}

var supportedDeliverySurfaces = map[string]struct{}{
	"web":           {},
	"cli":           {},
	"api":           {},
	"conversation":  {},
	"document":      {},
	"workflow":      {},
	"human_service": {},
	"physical":      {},
}

type Contract struct {
	Schema           string         `json:"schema"`
	ReadinessProfile string         `json:"readiness_profile"`
	ProductType      string         `json:"product_type"`
	DeliverySurfaces []string       `json:"delivery_surfaces"`
	InteractionMode  string         `json:"interaction_mode"`
	Project          Project        `json:"project"`
	NonGoals         []NonGoal      `json:"non_goals"`
	Capabilities     []Capability   `json:"capabilities"`
	Requirements     []Requirement  `json:"requirements"`
	Decisions        []Decision     `json:"decisions"`
	Risks            []Risk         `json:"risks"`
	Evaluations      []Evaluation   `json:"evaluations"`
	Artifacts        []Artifact     `json:"artifacts"`
	OpenQuestions    []OpenQuestion `json:"open_questions"`
}

type Project struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Purpose string `json:"purpose"`
	Status  string `json:"status"`
}

type NonGoal struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Capability struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Status       string   `json:"status"`
	Dependencies []string `json:"dependencies,omitempty"`
	Requirements []string `json:"requirements"`
	Evaluations  []string `json:"evaluations"`
	Risks        []string `json:"risks"`
	Artifacts    []string `json:"artifacts"`
}

type Requirement struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Decision struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Risk struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Severity   string `json:"severity"`
	Status     string `json:"status"`
	Mitigation string `json:"mitigation"`
}

type Evaluation struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Method string `json:"method"`
}

type Artifact struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Kind string `json:"kind"`
}

type OpenQuestion struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Blocker bool   `json:"blocker"`
	Status  string `json:"status"`
}

func LoadFile(path string) (Contract, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Contract{}, err
	}
	return Load(data)
}

func Load(data []byte) (Contract, error) {
	var c Contract
	if err := json.Unmarshal(data, &c); err != nil {
		return Contract{}, fmt.Errorf("malformed contract JSON: %w", err)
	}
	c.applyProductDefaults()
	if err := c.Validate(); err != nil {
		return Contract{}, err
	}
	return c, nil
}

func (c Contract) Validate() error {
	var missing []string
	if c.Schema == "" {
		missing = append(missing, "schema")
	}
	if c.ReadinessProfile == "" {
		missing = append(missing, "readiness_profile")
	}
	if c.Project.ID == "" {
		missing = append(missing, "project.id")
	}
	if c.Project.Name == "" {
		missing = append(missing, "project.name")
	}
	if c.Project.Purpose == "" {
		missing = append(missing, "project.purpose")
	}
	if c.Project.Status == "" {
		missing = append(missing, "project.status")
	}
	if c.NonGoals == nil {
		missing = append(missing, "non_goals")
	}
	if c.Capabilities == nil {
		missing = append(missing, "capabilities")
	}
	if c.Requirements == nil {
		missing = append(missing, "requirements")
	}
	if c.Decisions == nil {
		missing = append(missing, "decisions")
	}
	if c.Risks == nil {
		missing = append(missing, "risks")
	}
	if c.Evaluations == nil {
		missing = append(missing, "evaluations")
	}
	if c.Artifacts == nil {
		missing = append(missing, "artifacts")
	}
	if c.OpenQuestions == nil {
		missing = append(missing, "open_questions")
	}
	if len(missing) > 0 {
		return fmt.Errorf("contract missing required field(s): %s", strings.Join(missing, ", "))
	}
	if c.Schema != Schema {
		return fmt.Errorf("unsupported contract schema %q", c.Schema)
	}
	if err := profile.Validate(c.ReadinessProfile); err != nil {
		return err
	}
	if err := ValidateProductType(c.ProductType); err != nil {
		return err
	}
	if err := ValidateDeliverySurfaces(c.DeliverySurfaces); err != nil {
		return err
	}
	if err := ValidateInteractionMode(c.InteractionMode); err != nil {
		return err
	}

	if err := validateIDs(c); err != nil {
		return err
	}
	return nil
}

func (c *Contract) applyProductDefaults() {
	if strings.TrimSpace(c.ProductType) == "" {
		c.ProductType = DefaultProductType
	}
	if len(c.DeliverySurfaces) == 0 {
		c.DeliverySurfaces = DefaultDeliverySurfaces(c.ProductType)
	}
	if strings.TrimSpace(c.InteractionMode) == "" {
		c.InteractionMode = DefaultInteractionMode
	}
}

func DefaultDeliverySurfaces(productType string) []string {
	switch productType {
	case "conversation_product":
		return []string{"conversation"}
	case "research_protocol":
		return []string{"document"}
	case "operations_process":
		return []string{"workflow"}
	case "education_program":
		return []string{"human_service"}
	case "document_product":
		return []string{"document"}
	case "physical_product":
		return []string{"physical"}
	case "mixed":
		return []string{"workflow"}
	default:
		return []string{DefaultDeliverySurface}
	}
}

func ValidateProductType(value string) error {
	if _, ok := supportedProductTypes[value]; !ok {
		return fmt.Errorf("unsupported product_type %q (valid: %s)", value, strings.Join(ProductTypes(), ", "))
	}
	return nil
}

func ValidateDeliverySurfaces(values []string) error {
	if len(values) == 0 {
		return fmt.Errorf("delivery_surfaces must contain at least one surface")
	}
	for _, value := range values {
		if _, ok := supportedDeliverySurfaces[value]; !ok {
			return fmt.Errorf("unsupported delivery surface %q (valid: %s)", value, strings.Join(DeliverySurfaces(), ", "))
		}
	}
	return nil
}

func ValidateInteractionMode(value string) error {
	if !validIdentifier(value) {
		return fmt.Errorf("interaction_mode must be a lowercase identifier")
	}
	return nil
}

func ProductTypes() []string {
	return []string{
		"software",
		"conversation_product",
		"research_protocol",
		"operations_process",
		"education_program",
		"document_product",
		"physical_product",
		"mixed",
	}
}

func DeliverySurfaces() []string {
	return []string{
		"web",
		"cli",
		"api",
		"conversation",
		"document",
		"workflow",
		"human_service",
		"physical",
	}
}

func validIdentifier(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '_' {
			continue
		}
		return false
	}
	return true
}

func ValidateIDPrefix(id string, prefix string) error {
	if id == "" {
		return fmt.Errorf("missing %s id", prefix)
	}
	if !strings.HasPrefix(id, prefix+"-") {
		return fmt.Errorf("id %q must use %s- prefix", id, prefix)
	}
	return nil
}

func validateIDs(c Contract) error {
	for _, item := range c.NonGoals {
		if err := ValidateIDPrefix(item.ID, "NG"); err != nil {
			return err
		}
	}
	for _, item := range c.Capabilities {
		if err := ValidateIDPrefix(item.ID, "CAP"); err != nil {
			return err
		}
		for _, id := range item.Dependencies {
			if err := ValidateIDPrefix(id, "CAP"); err != nil {
				return err
			}
		}
		for _, id := range item.Requirements {
			if err := ValidateIDPrefix(id, "REQ"); err != nil {
				return err
			}
		}
		for _, id := range item.Evaluations {
			if err := ValidateIDPrefix(id, "EVAL"); err != nil {
				return err
			}
		}
		for _, id := range item.Risks {
			if err := ValidateIDPrefix(id, "RISK"); err != nil {
				return err
			}
		}
		for _, id := range item.Artifacts {
			if err := ValidateIDPrefix(id, "ART"); err != nil {
				return err
			}
		}
	}
	for _, item := range c.Requirements {
		if err := ValidateIDPrefix(item.ID, "REQ"); err != nil {
			return err
		}
	}
	for _, item := range c.Decisions {
		if err := ValidateIDPrefix(item.ID, "DEC"); err != nil {
			return err
		}
	}
	for _, item := range c.Risks {
		if err := ValidateIDPrefix(item.ID, "RISK"); err != nil {
			return err
		}
	}
	for _, item := range c.Evaluations {
		if err := ValidateIDPrefix(item.ID, "EVAL"); err != nil {
			return err
		}
	}
	for _, item := range c.Artifacts {
		if err := ValidateIDPrefix(item.ID, "ART"); err != nil {
			return err
		}
	}
	for _, item := range c.OpenQuestions {
		if err := ValidateIDPrefix(item.ID, "OQ"); err != nil {
			return err
		}
	}
	return nil
}
