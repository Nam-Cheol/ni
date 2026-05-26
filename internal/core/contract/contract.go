package contract

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const Schema = "ni.contract.v0"

type Contract struct {
	Schema        string         `json:"schema"`
	Project       Project        `json:"project"`
	NonGoals      []NonGoal      `json:"non_goals"`
	Capabilities  []Capability   `json:"capabilities"`
	Requirements  []Requirement  `json:"requirements"`
	Decisions     []Decision     `json:"decisions"`
	Risks         []Risk         `json:"risks"`
	Evaluations   []Evaluation   `json:"evaluations"`
	Artifacts     []Artifact     `json:"artifacts"`
	OpenQuestions []OpenQuestion `json:"open_questions"`
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

	if err := validateIDs(c); err != nil {
		return err
	}
	return nil
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
