package docstore

import (
	"encoding/json"
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
	Locked   bool
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
	Intent           GuidedIntent
}

type GuidedIntent struct {
	ProjectName         string
	ProjectGoal         string
	TargetUsers         string
	DownstreamAgentTask string
	ConstraintsNonGoals string
	SuccessCriteria     string
	KnownBlockers       string
	Deferrals           string
}

type FileAction string

const (
	FileActionCreate       FileAction = "create"
	FileActionSkipExisting FileAction = "skip_existing"
)

type PlannedFile struct {
	Path   string
	Action FileAction
}

type FilePlan struct {
	Root   string
	Files  []PlannedFile
	Locked bool
}

func Init(dir string) (Result, error) {
	return InitWithProfile(dir, profile.Default)
}

func InitWithProfile(dir string, readinessProfile string) (Result, error) {
	return InitWithOptions(dir, InitOptions{ReadinessProfile: readinessProfile})
}

func InitWithOptions(dir string, opts InitOptions) (Result, error) {
	plan, err := BuildFilePlan(dir, opts)
	if err != nil {
		return Result{}, err
	}
	if plan.Locked {
		return Result{Root: plan.Root, Locked: true}, nil
	}

	opts = withDefaults(opts)
	root := plan.Root
	result := Result{Root: root}
	files := templateFiles(opts)
	byPath := make(map[string]templateFile, len(files))
	for _, file := range files {
		byPath[file.path] = file
	}

	for _, planned := range plan.Files {
		file := byPath[planned.Path]
		target := filepath.Join(root, file.path)
		if planned.Action == FileActionSkipExisting {
			result.Existing = append(result.Existing, file.path)
			continue
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

func BuildFilePlan(dir string, opts InitOptions) (FilePlan, error) {
	opts = withDefaults(opts)
	if err := profile.Validate(opts.ReadinessProfile); err != nil {
		return FilePlan{}, err
	}
	if err := contract.ValidateProductType(opts.ProductType); err != nil {
		return FilePlan{}, err
	}
	if err := contract.ValidateDeliverySurfaces(opts.DeliverySurfaces); err != nil {
		return FilePlan{}, err
	}
	if err := contract.ValidateInteractionMode(opts.InteractionMode); err != nil {
		return FilePlan{}, err
	}
	root := filepath.Clean(dir)
	plan := FilePlan{Root: root}
	if _, err := os.Stat(filepath.Join(root, ".ni", "plan.lock.json")); err == nil {
		plan.Locked = true
		return plan, nil
	} else if !os.IsNotExist(err) {
		return plan, err
	}

	for _, file := range templateFiles(opts) {
		action := FileActionCreate
		if _, err := os.Stat(filepath.Join(root, file.path)); err == nil {
			action = FileActionSkipExisting
		} else if !os.IsNotExist(err) {
			return plan, err
		}
		plan.Files = append(plan.Files, PlannedFile{Path: file.path, Action: action})
	}
	return plan, nil
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
	guided := hasGuidedIntent(opts.Intent)
	intent := withIntentDefaults(opts.Intent)
	surfaces := opts.DeliverySurfaces
	if len(surfaces) == 0 {
		surfaces = contract.DefaultDeliverySurfaces(opts.ProductType)
	}
	surfaceText := markdownList(surfaces)
	blockerTitle := intent.KnownBlockers
	if blockerTitle == "TODO" {
		blockerTitle = "TODO"
	}
	deferrals := intent.Deferrals
	if deferrals == "TODO" {
		deferrals = "None recorded yet."
	}

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
	if guided {
		planDocs = []templateFile{
			{"docs/plan/00_project_brief.md", fmt.Sprintf("# Project brief\n\n## Product type\n\n%s\n\n## Delivery surfaces\n\n%s\n\n## Purpose\n\n%s\n\n## Problem\n\nInitial guided init note: downstream work should wait until intent is explicit, checked, locked, and compiled into a bounded handoff prompt.\n\n## Success definition\n\n%s\n", opts.ProductType, surfaceText, intent.ProjectGoal, intent.SuccessCriteria)},
			{"docs/plan/01_actors_outcomes.md", fmt.Sprintf("# Actors and outcomes\n\n## Actors\n\n- Target users / audience: %s\n- Planning model: drafts planning docs and contract updates from conversation.\n- CLI: validates readiness and lock state.\n\n## Outcomes\n\n- %s\n", intent.TargetUsers, intent.SuccessCriteria)},
			{"docs/plan/02_capabilities.md", fmt.Sprintf("# Capabilities\n\n## CAP-001: %s\n\nThe downstream agent should eventually do this only after `ni status`, `ni end`, and `ni run` produce an authoritative bounded handoff.\n", intent.DownstreamAgentTask)},
			{"docs/plan/03_interaction_contract.md", fmt.Sprintf("# Interaction contract\n\n## Interaction mode\n\n%s\n\n## Product interaction\n\nGuided init captured the first intent draft. Continued authoring happens through model-user planning conversation plus docs/plan/** and .ni/contract.json updates.\n\n## User control\n\nThe user decides whether ambiguous statements become accepted decisions. The CLI remains the authority for readiness, lock creation, and prompt compilation.\n", opts.InteractionMode)},
			{"docs/plan/04_domain_state.md", "# Domain and state model\n\n## Core entities\n\n```text\nproject\ncontract\ncapability\nrequirement\ndecision\nrisk\nevaluation\nartifact\nopen question\nlockfile\nprompt\n```\n"},
			{"docs/plan/05_constraints.md", fmt.Sprintf("# Constraints\n\n## Hard constraints\n\n- Readiness must be rule-based, not model-feeling-based.\n- Prompt output from `ni run` must be 4000 characters or less.\n- ni must not execute downstream work.\n\n## Constraints / non-goals from guided init\n\n%s\n", intent.ConstraintsNonGoals)},
			{"docs/plan/06_risks_security.md", fmt.Sprintf("# Risks and security\n\n## RISK-001: %s\n\nSeverity: high\n\nMitigation: Keep the blocker visible in docs/plan/** and .ni/contract.json until the user resolves it in planning conversation.\n", blockerTitle)},
			{"docs/plan/07_evaluation_contract.md", fmt.Sprintf("# Evaluation contract\n\n## EVAL-001: %s\n\nMethod: Review the locked plan against this success criterion before downstream handoff.\n", intent.SuccessCriteria)},
			{"docs/plan/08_delivery_operation.md", fmt.Sprintf("# Delivery and operation\n\n## Delivery surfaces\n\n%s\n\n## Initial delivery\n\n%s\n\n## Operating model\n\n- Planning docs are committed to git.\n- Contract JSON is committed to git.\n- `ni run` compiles a bounded handoff prompt only; it does not execute downstream work.\n", surfaceText, intent.DownstreamAgentTask)},
			{"docs/plan/09_execution_strategy.md", "# Execution strategy\n\n## v0 execution strategy\n\nDo not execute implementation automatically. Use `ni run` to compile a short prompt after the plan is locked.\n"},
			{"docs/plan/10_open_questions.md", fmt.Sprintf("# Open questions\n\n## OQ-001: %s\n\nBlocker: true\n\nInitial assumption: %s\n\n## Deferrals\n\n%s\n", blockerTitle, blockerTitle, deferrals)},
			{"docs/plan/11_decision_log.md", "# Decision log\n\n## DEC-001: Compile intent before downstream handoff\n\nStatus: accepted\n\nRationale: ni is a pre-runtime Project Intent Compiler. Planning must pass `ni status`, `ni end`, and `ni run` before downstream actors receive a handoff prompt.\n"},
		}
	}

	files := make([]templateFile, 0, len(planDocs)+6)
	files = append(files, planDocs...)
	files = append(files,
		templateFile{".ni/project.json", projectJSON},
		templateFile{".ni/contract.json", contractJSON(opts)},
		templateFile{".ni/session.json", sessionJSON},
		templateFile{".ni/pressure.json", pressureJSON},
		templateFile{".ni/harness.candidates.json", harnessCandidatesJSON},
		templateFile{".ni/readiness.rules.json", readinessRulesJSON},
		templateFile{".ni/readiness.profiles.json", readinessProfilesJSON},
	)
	return files
}

func withIntentDefaults(intent GuidedIntent) GuidedIntent {
	intent.ProjectName = todoIfBlank(intent.ProjectName)
	intent.ProjectGoal = todoIfBlank(intent.ProjectGoal)
	intent.TargetUsers = todoIfBlank(intent.TargetUsers)
	intent.DownstreamAgentTask = todoIfBlank(intent.DownstreamAgentTask)
	intent.ConstraintsNonGoals = todoIfBlank(intent.ConstraintsNonGoals)
	intent.SuccessCriteria = todoIfBlank(intent.SuccessCriteria)
	intent.KnownBlockers = todoIfBlank(intent.KnownBlockers)
	intent.Deferrals = todoIfBlank(intent.Deferrals)
	return intent
}

func hasGuidedIntent(intent GuidedIntent) bool {
	return strings.TrimSpace(intent.ProjectName) != "" ||
		strings.TrimSpace(intent.ProjectGoal) != "" ||
		strings.TrimSpace(intent.TargetUsers) != "" ||
		strings.TrimSpace(intent.DownstreamAgentTask) != "" ||
		strings.TrimSpace(intent.ConstraintsNonGoals) != "" ||
		strings.TrimSpace(intent.SuccessCriteria) != "" ||
		strings.TrimSpace(intent.KnownBlockers) != "" ||
		strings.TrimSpace(intent.Deferrals) != ""
}

func todoIfBlank(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "TODO"
	}
	return value
}

func markdownList(values []string) string {
	var b strings.Builder
	for _, value := range values {
		fmt.Fprintf(&b, "- %s\n", value)
	}
	return strings.TrimRight(b.String(), "\n")
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

const pressureJSON = `{
  "schema": "ni.pressure.v0",
  "items": []
}
`

const harnessCandidatesJSON = `{
  "schema": "ni.harness_candidates.v0",
  "candidates": []
}
`

const sessionJSON = `{
  "schema": "ni.session.v0",
  "active_planning_focus": "Initial planning workspace setup",
  "last_planning_summary": "No planning summary has been recorded yet.",
  "pending_questions": [],
  "recent_decisions": [],
  "recent_risks": [],
  "last_readiness_status": "UNKNOWN",
  "last_readiness_blockers": [],
  "last_updated_docs": [],
  "authority_order": [
    ".ni/plan.lock.json",
    ".ni/contract.json",
    "docs/plan/**",
    ".ni/session.json",
    "chat history"
  ],
  "notes": {
    "planning_aid_only": "Session state is a planning aid, not authority.",
    "locked_docs": "Session state must not override locked docs.",
    "completion": "Session state must not mark docs complete.",
    "raw_transcript": "Raw transcript is not the source of truth."
  }
}
`

func contractJSON(opts InitOptions) string {
	if !hasGuidedIntent(opts.Intent) {
		return fmt.Sprintf(defaultContractJSONTemplate, jsonString(opts.ReadinessProfile), jsonString(opts.ProductType), quoteStrings(opts.DeliverySurfaces), jsonString(opts.InteractionMode))
	}
	intent := withIntentDefaults(opts.Intent)
	return fmt.Sprintf(
		contractJSONTemplate,
		jsonString(opts.ReadinessProfile),
		jsonString(opts.ProductType),
		quoteStrings(opts.DeliverySurfaces),
		jsonString(opts.InteractionMode),
		jsonString(slugID(intent.ProjectName)),
		jsonString(intent.ProjectName),
		jsonString(intent.ProjectGoal),
		jsonString(intent.ConstraintsNonGoals),
		jsonString(intent.DownstreamAgentTask),
		jsonString(intent.SuccessCriteria),
		jsonString(intent.KnownBlockers),
		jsonString(intent.SuccessCriteria),
		jsonString(intent.KnownBlockers),
	)
}

func quoteStrings(values []string) string {
	quoted := make([]string, 0, len(values))
	for _, value := range values {
		quoted = append(quoted, jsonString(value))
	}
	return strings.Join(quoted, ", ")
}

func jsonString(value string) string {
	data, err := json.Marshal(value)
	if err != nil {
		return "\"TODO\""
	}
	return string(data)
}

func slugID(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" || value == "todo" {
		return "todo"
	}
	var b strings.Builder
	lastDash := false
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			lastDash = false
			continue
		}
		if !lastDash {
			b.WriteByte('-')
			lastDash = true
		}
	}
	slug := strings.Trim(b.String(), "-")
	if slug == "" {
		return "todo"
	}
	return slug
}

const contractJSONTemplate = `{
  "schema": "ni.contract.v0",
  "readiness_profile": %s,
  "product_type": %s,
  "delivery_surfaces": [%s],
  "interaction_mode": %s,
  "project": {
    "id": %s,
    "name": %s,
    "purpose": %s,
    "status": "draft"
  },
  "non_goals": [
    {
      "id": "NG-001",
      "title": %s
    }
  ],
  "capabilities": [
    {
      "id": "CAP-001",
      "title": %s,
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
      "title": %s,
      "status": "accepted"
    }
  ],
  "decisions": [
    {
      "id": "DEC-001",
      "title": "Compile intent before downstream handoff",
      "status": "accepted"
    }
  ],
  "risks": [
    {
      "id": "RISK-001",
      "title": %s,
      "severity": "high",
      "status": "accepted",
      "mitigation": "Keep the blocker visible in docs/plan/** and .ni/contract.json until the user resolves it in planning conversation."
    }
  ],
  "evaluations": [
    {
      "id": "EVAL-001",
      "title": %s,
      "method": "Review the locked plan against this success criterion before downstream handoff."
    }
  ],
  "artifacts": [
    {
      "id": "ART-001",
      "path": "docs/plan/",
      "kind": "planning_docs"
    }
  ],
  "open_questions": [
    {
      "id": "OQ-001",
      "title": %s,
      "blocker": true,
      "status": "open"
    }
  ]
}
`

const defaultContractJSONTemplate = `{
  "schema": "ni.contract.v0",
  "readiness_profile": %s,
  "product_type": %s,
  "delivery_surfaces": [%s],
  "interaction_mode": %s,
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
    },
    {
      "id": "R012",
      "title": "planning docs and contract are synchronized"
    },
    {
      "id": "R013",
      "title": "accepted decisions do not contradict each other"
    },
    {
      "id": "R014",
      "title": "project purpose is concrete"
    },
    {
      "id": "R015",
      "title": "actors and outcomes are specified"
    },
    {
      "id": "R016",
      "title": "delivery surface is specified"
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
        "R012": "blocker",
        "R013": "blocker",
        "R014": "blocker",
        "R015": "blocker",
        "R016": "blocker",
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
        "R012": "blocker",
        "R013": "blocker",
        "R014": "blocker",
        "R015": "blocker",
        "R016": "blocker",
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
        "R012": "blocker",
        "R013": "blocker",
        "R014": "blocker",
        "R015": "blocker",
        "R016": "blocker",
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
        "R012": "blocker",
        "R013": "blocker",
        "R014": "blocker",
        "R015": "blocker",
        "R016": "blocker",
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
        "R012": "blocker",
        "R013": "blocker",
        "R014": "blocker",
        "R015": "blocker",
        "R016": "blocker",
        "D001": "blocker",
        "D002": "blocker"
      }
    }
  ]
}
`
