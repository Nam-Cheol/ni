package initui

import (
	"fmt"
	"io"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"ni/internal/core/docstore"
)

type Config struct {
	Dir           string
	DefaultName   string
	ExistingFiles []string
	Input         io.Reader
	Output        io.Writer
}

type Result struct {
	Intent    docstore.GuidedIntent
	Confirmed bool
	Canceled  bool
	Choice    ExistingChoice
}

type ExistingChoice string

const (
	ExistingChoiceMissing ExistingChoice = "missing"
	ExistingChoiceKeep    ExistingChoice = "keep"
	ExistingChoiceAbort   ExistingChoice = "abort"
)

type field struct {
	Label string
	Value string
	Hint  string
}

type stage int

const (
	stageFields stage = iota
	stageConfirm
	stageExisting
	stageDone
)

type Model struct {
	dir            string
	fields         []field
	cursor         int
	stage          stage
	confirmCursor  int
	existingCursor int
	existingFiles  []string
	confirmed      bool
	canceled       bool
	choice         ExistingChoice
}

func NewModel(cfg Config) Model {
	defaultName := strings.TrimSpace(cfg.DefaultName)
	if defaultName == "" {
		defaultName = "my-project"
	}
	m := Model{
		dir:           cfg.Dir,
		existingFiles: append([]string(nil), cfg.ExistingFiles...),
		choice:        ExistingChoiceMissing,
		fields: []field{
			{Label: "Project name", Value: defaultName, Hint: "Name the planning workspace."},
			{Label: "One-sentence project goal", Hint: "What should change, for whom, and why?"},
			{Label: "Target users / audience", Hint: "Who depends on this plan?"},
			{Label: "Downstream agent task", Hint: "What may happen only after lock and handoff?"},
			{Label: "Constraints / non-goals", Value: "Do not execute downstream work before the plan is locked.", Hint: "Keep hard boundaries visible."},
			{Label: "Success criteria", Hint: "How will the accepted plan be judged?"},
			{Label: "Known blockers or open questions", Hint: "Blocking uncertainty prevents lock."},
			{Label: "Deferrals, if any", Value: "None recorded yet.", Hint: "Explicitly name deferred scope."},
		},
	}
	if len(m.existingFiles) > 0 {
		m.stage = stageExisting
	}
	return m
}

func Run(cfg Config) (Result, error) {
	m := NewModel(cfg)
	opts := []tea.ProgramOption{}
	if cfg.Input != nil {
		opts = append(opts, tea.WithInput(cfg.Input))
	}
	if cfg.Output != nil {
		opts = append(opts, tea.WithOutput(cfg.Output))
	}
	final, err := tea.NewProgram(m, opts...).Run()
	if err != nil {
		return Result{}, err
	}
	model, ok := final.(Model)
	if !ok {
		return Result{}, fmt.Errorf("unexpected init TUI model %T", final)
	}
	return model.Result(), nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "ctrl+c", "q":
		m.canceled = true
		m.stage = stageDone
		return m, tea.Quit
	}

	switch m.stage {
	case stageExisting:
		return m.updateExisting(key)
	case stageFields:
		return m.updateFields(key)
	case stageConfirm:
		return m.updateConfirm(key)
	default:
		return m, nil
	}
}

func (m Model) updateExisting(key tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up", "left":
		if m.existingCursor > 0 {
			m.existingCursor--
		}
	case "down", "right":
		if m.existingCursor < 2 {
			m.existingCursor++
		}
	case "esc":
		m.canceled = true
		m.stage = stageDone
		return m, tea.Quit
	case "enter":
		switch m.existingCursor {
		case 0:
			m.choice = ExistingChoiceMissing
			m.confirmed = true
		case 1:
			m.choice = ExistingChoiceKeep
			m.canceled = true
		default:
			m.choice = ExistingChoiceAbort
			m.canceled = true
		}
		m.stage = stageDone
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) updateFields(key tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up", "left":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "right", "enter":
		if m.cursor < len(m.fields)-1 {
			m.cursor++
		} else {
			m.stage = stageConfirm
		}
	case "esc":
		if m.cursor > 0 {
			m.cursor--
		} else {
			m.canceled = true
			m.stage = stageDone
			return m, tea.Quit
		}
	case "backspace":
		value := []rune(m.fields[m.cursor].Value)
		if len(value) > 0 {
			m.fields[m.cursor].Value = string(value[:len(value)-1])
		}
	default:
		text := key.Key().Text
		if text == "" && len(key.String()) == 1 {
			text = key.String()
		}
		if text != "" {
			m.fields[m.cursor].Value += text
		}
	}
	return m, nil
}

func (m Model) updateConfirm(key tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up", "down", "left", "right":
		if m.confirmCursor == 0 {
			m.confirmCursor = 1
		} else {
			m.confirmCursor = 0
		}
	case "esc":
		m.stage = stageFields
	case "enter":
		m.confirmed = m.confirmCursor == 0
		m.canceled = !m.confirmed
		m.stage = stageDone
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() tea.View {
	content := m.render()
	view := tea.NewView(content)
	view.AltScreen = true
	view.WindowTitle = "ni init"
	return view
}

func (m Model) Result() Result {
	return Result{
		Intent: docstore.GuidedIntent{
			ProjectName:         strings.TrimSpace(m.fields[0].Value),
			ProjectGoal:         strings.TrimSpace(m.fields[1].Value),
			TargetUsers:         strings.TrimSpace(m.fields[2].Value),
			DownstreamAgentTask: strings.TrimSpace(m.fields[3].Value),
			ConstraintsNonGoals: strings.TrimSpace(m.fields[4].Value),
			SuccessCriteria:     strings.TrimSpace(m.fields[5].Value),
			KnownBlockers:       strings.TrimSpace(m.fields[6].Value),
			Deferrals:           strings.TrimSpace(m.fields[7].Value),
		},
		Confirmed: m.confirmed,
		Canceled:  m.canceled,
		Choice:    m.choice,
	}
}

func (m Model) render() string {
	switch m.stage {
	case stageExisting:
		return m.renderExisting()
	case stageConfirm:
		return m.renderConfirm()
	default:
		return m.renderFields()
	}
}

func (m Model) renderFields() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("ni init"))
	b.WriteString("\n")
	b.WriteString(subtleStyle.Render("Guided project intent setup. No readiness, lock, prompt, or downstream execution happens here."))
	b.WriteString("\n\n")
	fmt.Fprintf(&b, "%s\n\n", progressStyle.Render(fmt.Sprintf("Step %d/%d", m.cursor+1, len(m.fields)+1)))
	for i, field := range m.fields {
		prefix := "  "
		lineStyle := normalStyle
		if i == m.cursor {
			prefix = "> "
			lineStyle = activeStyle
		}
		value := field.Value
		if value == "" {
			value = subtleStyle.Render("(empty)")
		}
		fmt.Fprintf(&b, "%s%s %s\n", prefix, lineStyle.Render(field.Label+":"), value)
		if i == m.cursor {
			fmt.Fprintf(&b, "    %s\n", subtleStyle.Render(field.Hint))
		}
	}
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Up/Down/Left/Right move  Enter continues  Esc backs out  q quits"))
	return panelStyle.Render(b.String())
}

func (m Model) renderConfirm() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("Review before write"))
	b.WriteString("\n")
	b.WriteString(subtleStyle.Render("Files are written only after you confirm. Existing files are protected by the domain writer."))
	b.WriteString("\n\n")
	for _, line := range []string{
		"Project: " + m.fields[0].Value,
		"Goal: " + m.fields[1].Value,
		"Audience: " + m.fields[2].Value,
		"Downstream task: " + m.fields[3].Value,
		"Success: " + m.fields[5].Value,
	} {
		fmt.Fprintf(&b, "- %s\n", line)
	}
	b.WriteString("\n")
	for i, option := range []string{"Write initial intent artifacts", "Cancel; write nothing"} {
		prefix := "  "
		if i == m.confirmCursor {
			prefix = "> "
		}
		fmt.Fprintf(&b, "%s%s\n", prefix, option)
	}
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Enter selects  Left/Esc returns  q quits"))
	return panelStyle.Render(b.String())
}

func (m Model) renderExisting() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("Existing planning files found"))
	b.WriteString("\n")
	b.WriteString(subtleStyle.Render("ni init will not overwrite existing planning state. Choose a safe path."))
	b.WriteString("\n\n")
	limit := len(m.existingFiles)
	if limit > 6 {
		limit = 6
	}
	for _, path := range m.existingFiles[:limit] {
		fmt.Fprintf(&b, "- %s\n", path)
	}
	if len(m.existingFiles) > limit {
		fmt.Fprintf(&b, "- ... %d more\n", len(m.existingFiles)-limit)
	}
	b.WriteString("\n")
	for i, option := range []string{"Add missing files only", "Keep existing and exit", "Abort"} {
		prefix := "  "
		if i == m.existingCursor {
			prefix = "> "
		}
		fmt.Fprintf(&b, "%s%s\n", prefix, option)
	}
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Up/Down/Left/Right move  Enter selects  Esc/q cancels"))
	return panelStyle.Render(b.String())
}

var (
	panelStyle    = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.NormalBorder(), true).BorderForeground(lipgloss.Color("#6b7280"))
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#0f766e"))
	activeStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#2563eb"))
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#111827"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#6b7280"))
	progressStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7c2d12"))
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#374151"))
)
