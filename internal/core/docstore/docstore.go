package docstore

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/profile"
)

type Result struct {
	Root     string
	Created  []string
	Existing []string
}

type templateFile struct {
	path    string
	content string
}

type InitOptions struct {
	ReadinessProfile string
	ProductType      string
	DeliverySurfaces []string
	InteractionMode  string
}

func Init(dir string) (Result, error) {
	return InitWithProfile(dir, profile.Default)
}

func InitWithProfile(dir string, readinessProfile string) (Result, error) {
	return InitWithOptions(dir, InitOptions{ReadinessProfile: readinessProfile})
}

func InitWithOptions(dir string, opts InitOptions) (Result, error) {
	opts = withDefaults(opts)
	if err := profile.Validate(opts.ReadinessProfile); err != nil {
		return Result{}, err
	}
	if err := contract.ValidateProductType(opts.ProductType); err != nil {
		return Result{}, err
	}
	if err := contract.ValidateDeliverySurfaces(opts.DeliverySurfaces); err != nil {
		return Result{}, err
	}
	if err := contract.ValidateInteractionMode(opts.InteractionMode); err != nil {
		return Result{}, err
	}
	root := filepath.Clean(dir)
	files := templateFiles(opts)
	result := Result{Root: root}

	for _, file := range files {
		target := filepath.Join(root, file.path)
		if _, err := os.Stat(target); err == nil {
			result.Existing = append(result.Existing, file.path)
			continue
		} else if !os.IsNotExist(err) {
			return result, err
		}

		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return result, err
		}
		if err := os.WriteFile(target, []byte(file.content), 0o644); err != nil {
			return result, err
		}
		result.Created = append(result.Created, file.path)
	}

	return result, nil
}

func withDefaults(opts InitOptions) InitOptions {
	if opts.ReadinessProfile == "" {
		opts.ReadinessProfile = profile.Default
	}
	if opts.ProductType == "" {
		opts.ProductType = contract.DefaultProductType
	}
	if len(opts.DeliverySurfaces) == 0 {
		opts.DeliverySurfaces = contract.DefaultDeliverySurfaces(opts.ProductType)
	}
	if opts.InteractionMode == "" {
		opts.InteractionMode = contract.DefaultInteractionMode
	}
	return opts
}

func RequiredPaths() []string {
	files := templateFiles(withDefaults(InitOptions{}))
	paths := make([]string, 0, len(files))
	for _, file := range files {
		paths = append(paths, file.path)
	}
	sort.Strings(paths)
	return paths
}

func templateFiles(opts InitOptions) []templateFile {
	planDocs := []templateFile{
		{"docs/plan/00_project_brief.md", "# Project brief\n\n## Product type\n\nTODO\n\n## Delivery surfaces\n\n- TODO\n\n## Purpose\n\nTODO\n\n## Problem\n\nTODO\n\n## Success definition\n\nTODO\n"},
		{"docs/plan/01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: TODO\n- Planning model: TODO\n- CLI: validates readiness and lock state.\n\n## Outcomes\n\n- TODO\n"},
		{"docs/plan/02_capabilities.md", "# Capabilities\n\n## CAP-001: TODO\n\nDescribe the first accepted capability.\n"},
		{"docs/plan/03_interaction_contract.md", "# Interaction contract\n\n## Interaction mode\n\nTODO\n\n## Product interaction\n\nTODO\n\n## User control\n\nTODO\n"},
		{"docs/plan/04_domain_state.md", "# Domain and state model\n\n## Core entities\n\n```text\nproject\ncontract\ncapability\nrequirement\ndecision\nrisk\nevaluation\nartifact\nopen question\nlockfile\nprompt\n```\n"},
		{"docs/plan/05_constraints.md", "# Constraints\n\n## Hard constraints\n\n- Readiness must be rule-based, not model-feeling-based.\n- Prompt output from `ni run` must be 4000 characters or less.\n"},
		{"docs/plan/06_risks_security.md", "# Risks and security\n\n## RISK-001: TODO\n\nSeverity: high\n\nMitigation: TODO\n"},
		{"docs/plan/07_evaluation_contract.md", "# Evaluation contract\n\n## EVAL-001: TODO\n\nMethod: TODO\n"},
		{"docs/plan/08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- TODO\n\n## Initial delivery\n\nTODO\n\n## Operating model\n\n- Planning docs are committed to git.\n- Contract JSON is committed to git.\n"},
		{"docs/plan/09_execution_strategy.md", "# Execution strategy\n\n## v0 execution strategy\n\nDo not execute implementation automatically. Use `ni run` to compile a short prompt after the plan is locked.\n"},
		{"docs/plan/10_open_questions.md", "# Open questions\n\n## OQ-001: TODO\n\nBlocker: true\n\nInitial assumption: TODO\n"},
		{"docs/plan/11_decision_log.md", "# Decision log\n\n## DEC-001: TODO\n\nStatus: accepted\n\nRationale: TODO\n"},
	}

	files := make([]templateFile, 0, len(planDocs)+4)
	files = append(files, planDocs...)
	files = append(files,
		templateFile{".ni/project.json", projectJSON},
		templateFile{".ni/contract.json", contractJSON(opts)},
		templateFile{".ni/readiness.rules.json", readinessRulesJSON},
		templateFile{".ni/readiness.profiles.json", readinessProfilesJSON},
	)
	return files
}

const projectJSON = `{
  "schema": "ni.project.v0",
  "id": "todo",
  "name": "TODO",
  "created_for": "ni planning workspace",
  "kernel_phase": "planning",
  "commands": [
    "init",
    "status",
    "end",
    "run"
  ],
  "prompt_max_chars": 4000
}
`

func contractJSON(opts InitOptions) string {
	return fmt.Sprintf(contractJSONTemplate, opts.ReadinessProfile, opts.ProductType, quoteStrings(opts.DeliverySurfaces), opts.InteractionMode)
}

func quoteStrings(values []string) string {
	quoted := make([]string, 0, len(values))
	for _, value := range values {
		quoted = append(quoted, fmt.Sprintf("%q", value))
	}
	return strings.Join(quoted, ", ")
}

const contractJSONTemplate = `{
  "schema": "ni.contract.v0",
  "readiness_profile": %q,
  "product_type": %q,
  "delivery_surfaces": [%s],
  "interaction_mode": %q,
  "project": {
    "id": "todo",
    "name": "TODO",
    "purpose": "TODO",
    "status": "draft"
  },
  "non_goals": [
    {
      "id": "NG-001",
      "title": "Do not implement execution adapters before the planning kernel is ready."
    }
  ],
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "TODO",
      "status": "accepted",
      "requirements": [
        "REQ-001"
      ],
      "evaluations": [
        "EVAL-001"
      ],
      "risks": [
        "RISK-001"
      ],
      "artifacts": [
        "ART-001"
      ]
    }
  ],
  "requirements": [
    {
      "id": "REQ-001",
      "title": "TODO",
      "status": "accepted"
    }
  ],
  "decisions": [
    {
      "id": "DEC-001",
      "title": "TODO",
      "status": "accepted"
    }
  ],
  "risks": [
    {
      "id": "RISK-001",
      "title": "TODO",
      "severity": "high",
      "status": "accepted",
      "mitigation": "TODO"
    }
  ],
  "evaluations": [
    {
      "id": "EVAL-001",
      "title": "TODO",
      "method": "TODO"
    }
  ],
  "artifacts": [
    {
      "id": "ART-001",
      "path": "TODO",
      "kind": "TODO"
    }
  ],
  "open_questions": [
    {
      "id": "OQ-001",
      "title": "TODO",
      "blocker": true,
      "status": "open"
    }
  ]
}
`

const readinessRulesJSON = `{
  "schema": "ni.readiness.rules.v0",
  "statuses": [
    "BLOCKED",
    "READY_WITH_DEFERRALS",
    "READY"
  ],
  "blocking_rules": [
    {
      "id": "R001",
      "title": "required planning docs exist"
    },
    {
      "id": "R002",
      "title": ".ni/contract.json is valid JSON"
    },
    {
      "id": "R003",
      "title": "at least one capability exists"
    },
    {
      "id": "R004",
      "title": "every accepted capability has at least one linked evaluation"
    },
    {
      "id": "R005",
      "title": "every evaluation has a method"
    },
    {
      "id": "R006",
      "title": "every high-severity risk has mitigation"
    },
    {
      "id": "R007",
      "title": "every accepted capability has at least one artifact or requirement"
    },
    {
      "id": "R008",
      "title": "decision status is accepted, deferred, rejected, or not_applicable"
    },
    {
      "id": "R009",
      "title": "blocker open questions prevent lock"
    },
    {
      "id": "R010",
      "title": "at least one non-goal exists"
    },
    {
      "id": "R011",
      "title": "readiness profile definitions are valid"
    }
  ],
  "required_docs": [
    "docs/plan/00_project_brief.md",
    "docs/plan/01_actors_outcomes.md",
    "docs/plan/02_capabilities.md",
    "docs/plan/03_interaction_contract.md",
    "docs/plan/04_domain_state.md",
    "docs/plan/05_constraints.md",
    "docs/plan/06_risks_security.md",
    "docs/plan/07_evaluation_contract.md",
    "docs/plan/08_delivery_operation.md",
    "docs/plan/09_execution_strategy.md",
    "docs/plan/10_open_questions.md",
    "docs/plan/11_decision_log.md"
  ]
}
`

const readinessProfilesJSON = `{
  "schema": "ni.readiness.profiles.v0",
  "default_profile": "prototype",
  "profiles": [
    {
      "name": "concept",
      "description": "A planning profile for early intent, audience, and uncertainty discovery before full traceability is required.",
      "issue_severity": {
        "R001": "blocker",
        "R002": "blocker",
        "R003": "deferral",
        "R004": "deferral",
        "R005": "deferral",
        "R006": "blocker",
        "R007": "deferral",
        "R008": "blocker",
        "R009": "blocker",
        "R010": "deferral",
        "D001": "deferral",
        "D002": "deferral"
      }
    },
    {
      "name": "prototype",
      "description": "The default planning profile for a small validated plan with capability, risk, evaluation, and boundary traceability.",
      "issue_severity": {
        "R001": "blocker",
        "R002": "blocker",
        "R003": "blocker",
        "R004": "blocker",
        "R005": "blocker",
        "R006": "blocker",
        "R007": "blocker",
        "R008": "blocker",
        "R009": "blocker",
        "R010": "blocker",
        "D001": "deferral",
        "D002": "deferral"
      }
    },
    {
      "name": "mvp",
      "description": "A planning profile for a usable first release plan with the same blocking traceability as prototype readiness.",
      "issue_severity": {
        "R001": "blocker",
        "R002": "blocker",
        "R003": "blocker",
        "R004": "blocker",
        "R005": "blocker",
        "R006": "blocker",
        "R007": "blocker",
        "R008": "blocker",
        "R009": "blocker",
        "R010": "blocker",
        "D001": "deferral",
        "D002": "deferral"
      }
    },
    {
      "name": "beta",
      "description": "A planning profile for broader validation where core traceability blocks lock and unresolved non-blocking questions remain explicit deferrals.",
      "issue_severity": {
        "R001": "blocker",
        "R002": "blocker",
        "R003": "blocker",
        "R004": "blocker",
        "R005": "blocker",
        "R006": "blocker",
        "R007": "blocker",
        "R008": "blocker",
        "R009": "blocker",
        "R010": "blocker",
        "D001": "deferral",
        "D002": "deferral"
      }
    },
    {
      "name": "production",
      "description": "A planning profile for high-confidence lock readiness where deferred decisions and open questions block the plan.",
      "issue_severity": {
        "R001": "blocker",
        "R002": "blocker",
        "R003": "blocker",
        "R004": "blocker",
        "R005": "blocker",
        "R006": "blocker",
        "R007": "blocker",
        "R008": "blocker",
        "R009": "blocker",
        "R010": "blocker",
        "D001": "blocker",
        "D002": "blocker"
      }
    }
  ]
}
`
