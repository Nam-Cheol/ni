package docstore

import (
	"os"
	"path/filepath"
	"sort"
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

func Init(dir string) (Result, error) {
	root := filepath.Clean(dir)
	files := templateFiles()
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

func RequiredPaths() []string {
	files := templateFiles()
	paths := make([]string, 0, len(files))
	for _, file := range files {
		paths = append(paths, file.path)
	}
	sort.Strings(paths)
	return paths
}

func templateFiles() []templateFile {
	planDocs := []templateFile{
		{"docs/plan/00_project_brief.md", "# Project brief\n\n## Purpose\n\nTODO\n\n## Problem\n\nTODO\n\n## Success definition\n\nTODO\n"},
		{"docs/plan/01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: TODO\n- Planning model: TODO\n- CLI: validates readiness and lock state.\n\n## Outcomes\n\n- TODO\n"},
		{"docs/plan/02_capabilities.md", "# Capabilities\n\n## CAP-001: TODO\n\nDescribe the first accepted capability.\n"},
		{"docs/plan/03_interaction_contract.md", "# Interaction contract\n\n## CLI interaction\n\n```text\nni init\nni status\nni end\nni run\n```\n\n## User control\n\nTODO\n"},
		{"docs/plan/04_domain_state.md", "# Domain and state model\n\n## Core entities\n\n```text\nproject\ncontract\ncapability\nrequirement\ndecision\nrisk\nevaluation\nartifact\nopen question\nlockfile\nprompt\n```\n"},
		{"docs/plan/05_constraints.md", "# Constraints\n\n## Hard constraints\n\n- Readiness must be rule-based, not model-feeling-based.\n- Prompt output from `ni run` must be 4000 characters or less.\n"},
		{"docs/plan/06_risks_security.md", "# Risks and security\n\n## RISK-001: TODO\n\nSeverity: high\n\nMitigation: TODO\n"},
		{"docs/plan/07_evaluation_contract.md", "# Evaluation contract\n\n## EVAL-001: TODO\n\nMethod: TODO\n"},
		{"docs/plan/08_delivery_operation.md", "# Delivery and operation\n\n## Initial delivery\n\nTODO\n\n## Operating model\n\n- Planning docs are committed to git.\n- Contract JSON is committed to git.\n"},
		{"docs/plan/09_execution_strategy.md", "# Execution strategy\n\n## v0 execution strategy\n\nDo not execute implementation automatically. Use `ni run` to compile a short prompt after the plan is locked.\n"},
		{"docs/plan/10_open_questions.md", "# Open questions\n\n## OQ-001: TODO\n\nBlocker: true\n\nInitial assumption: TODO\n"},
		{"docs/plan/11_decision_log.md", "# Decision log\n\n## DEC-001: TODO\n\nStatus: accepted\n\nRationale: TODO\n"},
	}

	files := make([]templateFile, 0, len(planDocs)+3)
	files = append(files, planDocs...)
	files = append(files,
		templateFile{".ni/project.json", projectJSON},
		templateFile{".ni/contract.json", contractJSON},
		templateFile{".ni/readiness.rules.json", readinessRulesJSON},
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

const contractJSON = `{
  "schema": "ni.contract.v0",
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
