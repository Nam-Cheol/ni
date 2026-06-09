package initui

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"ni/internal/core/docstore"
)

type Config struct {
	Dir           string
	CommandName   string
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
	Key   string
	Value string
}

type fieldGuide struct {
	Label    string
	Hint     string
	Why      string
	Good     string
	Example  string
	Optional string
	MapsTo   string
}

type stage int

const (
	stageLanguage stage = iota
	stageFields
	stageConfirm
	stageExisting
	stageDone
)

type language string

const (
	languageKorean  language = "ko"
	languageEnglish language = "en"
)

type Model struct {
	commandName    string
	dir            string
	fields         []field
	cursor         int
	stage          stage
	language       language
	languageCursor int
	confirmCursor  int
	existingCursor int
	existingFiles  []string
	width          int
	height         int
	bodyViewport   viewport.Model
	confirmed      bool
	canceled       bool
	choice         ExistingChoice
}

func NewModel(cfg Config) Model {
	defaultName := strings.TrimSpace(cfg.DefaultName)
	if defaultName == "" {
		defaultName = "my-project"
	}
	commandName := strings.TrimSpace(cfg.CommandName)
	if commandName == "" {
		commandName = "namba-intent"
	}
	m := Model{
		commandName:   commandName,
		dir:           cfg.Dir,
		language:      languageKorean,
		width:         100,
		height:        32,
		bodyViewport:  newBodyViewport(100, 27),
		existingFiles: append([]string(nil), cfg.ExistingFiles...),
		choice:        ExistingChoiceMissing,
		fields: []field{
			{Key: "project_name", Value: defaultName},
			{Key: "project_goal"},
			{Key: "target_users"},
			{Key: "downstream_task"},
			{Key: "constraints"},
			{Key: "success"},
			{Key: "blockers"},
			{Key: "deferrals"},
		},
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
	return func() tea.Msg { return tea.RequestWindowSize() }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = size.Width
		m.height = size.Height
		m.constrainScroll()
		return m, nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}
	switch key.String() {
	case "ctrl+c":
		m.canceled = true
		m.stage = stageDone
		return m, tea.Quit
	}

	if next, handled := m.updateViewport(key); handled {
		return next, nil
	}

	switch m.stage {
	case stageLanguage:
		return m.updateLanguage(key)
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

func (m Model) updateViewport(key tea.KeyPressMsg) (tea.Model, bool) {
	switch key.String() {
	case "pgdown", "pgup", "ctrl+d", "ctrl+u":
		m.refreshViewport()
		next, _ := m.bodyViewport.Update(key)
		m.bodyViewport = next
	case "home":
		m.refreshViewport()
		m.bodyViewport.GotoTop()
	case "end":
		m.refreshViewport()
		m.bodyViewport.GotoBottom()
	default:
		return m, false
	}
	m.constrainScroll()
	return m, true
}

func (m *Model) constrainScroll() {
	m.refreshViewport()
}

func (m Model) updateLanguage(key tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up", "left":
		m.languageCursor = 0
	case "down", "right":
		m.languageCursor = 1
	case "1", "k", "K":
		m.languageCursor = 0
		return m.chooseLanguage()
	case "2", "e", "E":
		m.languageCursor = 1
		return m.chooseLanguage()
	case "enter":
		return m.chooseLanguage()
	case "esc", "q":
		m.canceled = true
		m.stage = stageDone
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) chooseLanguage() (tea.Model, tea.Cmd) {
	if m.languageCursor == 1 {
		m.language = languageEnglish
	} else {
		m.language = languageKorean
	}
	m.applyLocalizedDefaults()
	if len(m.existingFiles) > 0 {
		m.stage = stageExisting
	} else {
		m.stage = stageFields
	}
	m.resetViewport()
	return m, nil
}

func (m *Model) applyLocalizedDefaults() {
	if strings.TrimSpace(m.fields[4].Value) == "" {
		m.fields[4].Value = m.boundaryDefault()
	}
	if strings.TrimSpace(m.fields[7].Value) == "" {
		m.fields[7].Value = m.deferralDefault()
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
		m.resetViewport()
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
			m.resetViewport()
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
	case "q":
		m.canceled = true
		m.stage = stageDone
		return m, tea.Quit
	case "esc":
		m.stage = stageFields
		m.resetViewport()
	case "enter":
		m.confirmed = m.confirmCursor == 0
		m.canceled = !m.confirmed
		m.stage = stageDone
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() tea.View {
	content := m.renderShell()
	view := tea.NewView(content)
	view.AltScreen = true
	view.WindowTitle = m.commandName + " init"
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

type layoutSpec struct {
	width      int
	height     int
	bodyHeight int
	panelWidth int
	sideBySide bool
	compact    bool
}

func newBodyViewport(width int, height int) viewport.Model {
	v := viewport.New(viewport.WithWidth(width), viewport.WithHeight(height))
	v.SoftWrap = true
	v.FillHeight = true
	v.MouseWheelEnabled = true
	return v
}

func (m *Model) refreshViewport() {
	layout := m.layout()
	if m.bodyViewport.Width() == 0 || m.bodyViewport.Height() == 0 {
		m.bodyViewport = newBodyViewport(layout.width, layout.bodyHeight)
	}
	m.bodyViewport.SetWidth(layout.width)
	m.bodyViewport.SetHeight(layout.bodyHeight)
	m.bodyViewport.SoftWrap = true
	m.bodyViewport.FillHeight = true
	m.bodyViewport.SetContent(m.renderBody(layout))
}

func (m *Model) resetViewport() {
	m.refreshViewport()
	m.bodyViewport.GotoTop()
}

func (m Model) renderShell() string {
	layout := m.layout()
	m.refreshViewport()
	bodyHeight := m.bodyViewport.Height()
	totalLines := m.bodyViewport.TotalLineCount()
	offset := m.bodyViewport.YOffset()
	header := m.renderHeader(layout)
	footer := m.renderFooter(layout)
	help := m.renderHelpBar(layout, totalLines, bodyHeight, offset)
	body := m.bodyViewport.View()
	screen := lipgloss.JoinVertical(lipgloss.Left, header, body, help, footer)
	return fitToHeight(screen, layout.height)
}

func (m Model) layout() layoutSpec {
	width := m.width
	height := m.height
	if width <= 0 {
		width = 100
	}
	if height <= 0 {
		height = 32
	}
	width = max(32, width)
	height = max(8, height)
	bodyHeight := height - 5
	if height <= 12 {
		bodyHeight = height - 4
	}
	if bodyHeight < 1 {
		bodyHeight = 1
	}
	return layoutSpec{
		width:      width,
		height:     height,
		bodyHeight: bodyHeight,
		panelWidth: max(24, width),
		sideBySide: width >= 92,
		compact:    height <= 14 || width <= 52,
	}
}

func (m Model) renderHeader(layout layoutSpec) string {
	mode := m.modeLabel()
	left := headerTitleStyle.Render("Namba Intent")
	right := mutedStyle.Render(fmt.Sprintf("%s init / %s", m.commandName, mode))
	if layout.compact {
		line := compactModeStyle.Render(truncatePlain("Namba Intent / "+m.commandName+" init / "+mode, layout.width))
		return headerBarStyle.Width(layout.width).Render(line)
	}
	gap := max(1, layout.width-lipgloss.Width(left)-lipgloss.Width(right)-2)
	return headerBarStyle.Width(layout.width).Render(left + strings.Repeat(" ", gap) + right)
}

func (m Model) renderBody(layout layoutSpec) string {
	main := m.renderMainPanel(layout)
	status := m.renderStatusPanel(layout)
	files := m.renderFilesPanel(layout)
	next := m.renderNextActionPanel(layout)
	if !layout.sideBySide {
		return lipgloss.JoinVertical(lipgloss.Left, main, status, files, next)
	}
	if layout.width >= 96 {
		leftWidth := max(38, layout.width/2)
		remaining := max(40, layout.width-leftWidth)
		middleWidth := max(22, remaining/2)
		rightWidth := max(22, layout.width-leftWidth-middleWidth)
		left := m.renderMainPanel(layout.withWidth(leftWidth))
		middle := lipgloss.JoinVertical(lipgloss.Left,
			m.renderStatusPanel(layout.withWidth(middleWidth)),
			m.renderNextActionPanel(layout.withWidth(middleWidth)),
		)
		right := m.renderFilesPanel(layout.withWidth(rightWidth))
		return lipgloss.JoinHorizontal(lipgloss.Top, left, middle, right)
	}
	leftWidth := max(40, layout.width*60/100)
	rightWidth := max(28, layout.width-leftWidth)
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.renderMainPanel(layout.withWidth(leftWidth)),
		lipgloss.JoinVertical(lipgloss.Left,
			m.renderStatusPanel(layout.withWidth(rightWidth)),
			m.renderNextActionPanel(layout.withWidth(rightWidth)),
			m.renderFilesPanel(layout.withWidth(rightWidth)),
		),
	)
}

func (l layoutSpec) withWidth(width int) layoutSpec {
	l.panelWidth = width
	l.width = width
	l.sideBySide = false
	return l
}

func (m Model) renderMainPanel(layout layoutSpec) string {
	switch m.stage {
	case stageLanguage:
		return m.renderLanguagePanel(layout)
	case stageExisting:
		return m.renderExistingPanel(layout)
	case stageConfirm:
		return m.renderConfirmPanel(layout)
	default:
		return m.renderFieldsPanel(layout)
	}
}

func (m Model) renderLanguagePanel(layout layoutSpec) string {
	var b strings.Builder
	b.WriteString(panelHeadingStyle.Render(m.t("Project Intent initialization", "Project Intent 초기화")))
	b.WriteString("\n")
	b.WriteString(m.wrapText(m.t("Start the planning journey in your language. This screen only drafts intent before status, lock, and handoff.", "언어를 선택한 뒤 planning journey를 시작합니다. 이 화면은 status, lock, handoff 전에 intent 초안만 만듭니다."), layout.panelWidth-8))
	b.WriteString("\n\n")
	b.WriteString(progressStyle.Render(renderProgress(1, len(m.fields)+2)))
	b.WriteString("\n\n")
	options := []struct {
		label string
		help  string
	}{
		{"Korean", "한국어 사용자에게 자연스러운 planning 안내를 보여줍니다."},
		{"English", "Use English labels, examples, and review guidance."},
	}
	for i, option := range options {
		b.WriteString(m.renderChoiceLine(i == m.languageCursor, option.label, option.help, layout))
	}
	return m.renderPanel("Main Panel", b.String(), layout.panelWidth, panelStyle)
}

func (m Model) renderFieldsPanel(layout layoutSpec) string {
	var b strings.Builder
	guide := m.fieldGuide(m.fields[m.cursor].Key)
	b.WriteString(panelHeadingStyle.Render(m.t("Project Intent initialization", "Project Intent 초기화")))
	b.WriteString("\n")
	b.WriteString(progressStyle.Render(renderProgress(m.cursor+2, len(m.fields)+2)))
	b.WriteString("\n\n")
	b.WriteString(safetyStyle.Render(m.wrapText(m.t("Safety boundary: the plan must be locked before any downstream work starts.", "안전 경계: plan이 locked 되기 전에는 어떤 downstream 작업도 시작하지 않습니다."), layout.panelWidth-8)))
	b.WriteString("\n\n")
	for i, field := range m.fields {
		fieldLabel := m.fieldGuide(field.Key).Label
		value := strings.TrimSpace(field.Value)
		if value == "" {
			value = m.t("(empty)", "(비어 있음)")
		}
		line := fieldLabel + ": " + value
		if i == m.cursor {
			b.WriteString(selectedStyle.Render("> " + truncatePlain(line, max(8, layout.panelWidth-10))))
		} else {
			b.WriteString(normalStyle.Render("  " + truncatePlain(line, max(8, layout.panelWidth-10))))
		}
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(m.renderPanel(m.t("Field guide", "입력 안내"), m.renderGuide(guide, layout.panelWidth-14), max(24, layout.panelWidth-4), infoPanelStyle))
	return m.renderPanel("Main Panel", b.String(), layout.panelWidth, panelStyle)
}

func (m Model) renderConfirmPanel(layout layoutSpec) string {
	var b strings.Builder
	b.WriteString(panelHeadingStyle.Render(m.t("Review before write", "저장 전 마지막 확인")))
	b.WriteString("\n")
	b.WriteString(progressStyle.Render(renderProgress(len(m.fields)+2, len(m.fields)+2)))
	b.WriteString("\n\n")
	for _, index := range []int{0, 1, 2, 3, 5, 6, 7} {
		guide := m.fieldGuide(m.fields[index].Key)
		value := strings.TrimSpace(m.fields[index].Value)
		if value == "" {
			value = "TODO"
		}
		b.WriteString(labelStyle.Render(guide.Label + ": "))
		b.WriteString(m.wrapText(value, max(8, layout.panelWidth-14)))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	for i, option := range []string{m.t("Write initial intent artifacts", "초기 intent 파일 저장"), m.t("Cancel; write nothing", "취소하고 아무것도 쓰지 않기")} {
		b.WriteString(m.renderOption(i == m.confirmCursor, option, layout.panelWidth-8))
	}
	return m.renderPanel("Main Panel", b.String(), layout.panelWidth, panelStyle)
}

func (m Model) renderExistingPanel(layout layoutSpec) string {
	var b strings.Builder
	b.WriteString(panelHeadingStyle.Render(m.t("Existing planning files found", "기존 planning 파일을 찾았습니다")))
	b.WriteString("\n")
	b.WriteString(m.wrapText(m.t("Namba Intent init will not overwrite existing planning state. Choose a safe path.", "Namba Intent init은 기존 planning state를 덮어쓰지 않습니다. 안전한 경로를 선택하세요."), layout.panelWidth-8))
	b.WriteString("\n\n")
	for i, option := range []string{m.t("Add missing files only", "누락된 파일만 추가"), m.t("Keep existing and exit", "기존 파일 유지 후 종료"), m.t("Abort", "중단")} {
		b.WriteString(m.renderOption(i == m.existingCursor, option, layout.panelWidth-8))
	}
	return m.renderPanel("Main Panel", b.String(), layout.panelWidth, panelStyle)
}

func (m Model) renderStatusPanel(layout layoutSpec) string {
	status, detail := m.statusLine()
	style := statusStyle
	if strings.Contains(status, "blocked") {
		style = blockedStyle
	}
	if strings.Contains(status, "warning") {
		style = warningStyle
	}
	if strings.Contains(status, "ready") || strings.Contains(status, "done") {
		style = successStyle
	}
	lines := []string{
		style.Render(status),
		m.wrapText(detail, layout.panelWidth-8),
		"",
		labelStyle.Render("Detected:") + " " + truncatePlain(m.dir, max(8, layout.panelWidth-18)),
		labelStyle.Render("Mode:") + " " + m.modeLabel(),
		labelStyle.Render("Authority:") + " CLI status/end/run",
	}
	return m.renderPanel("Status Panel", strings.Join(lines, "\n"), layout.panelWidth, panelStyle)
}

func (m Model) renderFilesPanel(layout layoutSpec) string {
	var b strings.Builder
	paths := m.filesForPanel()
	if len(paths) == 0 {
		b.WriteString(m.t("No planning files detected yet.", "아직 감지된 planning 파일이 없습니다."))
	} else {
		for _, item := range paths {
			b.WriteString(item)
			b.WriteString("\n")
		}
	}
	return m.renderPanel("Files Panel", strings.TrimRight(b.String(), "\n"), layout.panelWidth, panelStyle)
}

func (m Model) renderNextActionPanel(layout layoutSpec) string {
	content := strings.Join([]string{
		m.t("Next commands:", "다음 명령:"),
		fmt.Sprintf("- %s status --proof --next-questions", m.commandName),
		fmt.Sprintf("- %s end", m.commandName),
		fmt.Sprintf("- %s run --max-chars 4000", m.commandName),
		"",
		m.t("Update check:", "업데이트 확인:"),
		fmt.Sprintf("- %s version", m.commandName),
	}, "\n")
	switch m.stage {
	case stageLanguage:
		content = m.t("Choose a language, then capture the first project intent draft.", "언어를 선택한 뒤 첫 project intent 초안을 작성하세요.")
	case stageExisting:
		content = strings.Join([]string{
			m.t("Choose how to handle existing files.", "기존 파일 처리 방식을 선택하세요."),
			m.t("No locked planning state is modified by this screen.", "이 화면은 locked planning state를 수정하지 않습니다."),
		}, "\n")
	case stageConfirm:
		content = strings.Join([]string{
			m.t("Confirm write, then run:", "저장 확인 후 실행:"),
			fmt.Sprintf("- %s status --proof --next-questions", m.commandName),
			"",
			m.t("This screen does not decide readiness or lock.", "이 화면은 readiness나 lock을 결정하지 않습니다."),
			fmt.Sprintf("- %s version", m.commandName),
		}, "\n")
	}
	return m.renderPanel("Next Action Panel", content, layout.panelWidth, panelStyle)
}

func (m Model) renderHelpBar(layout layoutSpec, totalLines int, bodyHeight int, offset int) string {
	scroll := "scroll: all visible"
	if totalLines > bodyHeight {
		scroll = fmt.Sprintf("scroll: %d/%d PgUp/PgDn Home/End", offset+1, max(1, totalLines-bodyHeight+1))
	}
	var keys string
	switch m.stage {
	case stageLanguage:
		keys = "↑/↓ choose  Enter select  1 Korean  2 English  q quit"
	case stageFields:
		keys = m.t("↑/↓ move  type edit  Enter next  Esc back  PgUp/PgDn scroll", "↑/↓ 이동  입력 편집  Enter 다음  Esc 이전  PgUp/PgDn 스크롤")
	case stageConfirm:
		keys = m.t("↑/↓ choose  Enter select  Esc back  q quit", "↑/↓ 선택  Enter 실행  Esc 이전  q 종료")
	default:
		keys = m.t("↑/↓ choose  Enter select  Esc/q cancel", "↑/↓ 선택  Enter 실행  Esc/q 취소")
	}
	text := keys + "  |  " + scroll
	if layout.compact {
		text = compactModeStyle.Render(scroll + "  |  " + keys)
	}
	text = truncatePlain(text, layout.width)
	return helpBarStyle.Width(layout.width).Render(text)
}

func (m Model) renderFooter(layout layoutSpec) string {
	summary := m.t("No downstream work runs. Init drafts docs and contract only.", "downstream 작업은 실행하지 않습니다. init은 docs와 contract 초안만 만듭니다.")
	if m.stage == stageDone {
		summary = m.t("Done.", "완료.")
	}
	return footerStyle.Width(layout.width).Render(truncatePlain(summary, layout.width))
}

func (m Model) renderPanel(title string, body string, width int, style lipgloss.Style) string {
	width = max(20, width)
	contentWidth := max(8, width-6)
	titleLine := sectionStyle.Render(title)
	if body != "" {
		body = titleLine + "\n" + body
	} else {
		body = titleLine
	}
	return style.Width(contentWidth).Render(body)
}

func (m Model) renderChoiceLine(selected bool, label string, help string, layout layoutSpec) string {
	prefix := "  "
	style := normalStyle
	if selected {
		prefix = "> "
		style = selectedStyle
	}
	line := prefix + label
	if layout.compact {
		return style.Render(truncatePlain(line, layout.panelWidth-8)) + "\n"
	}
	return style.Render(truncatePlain(line, layout.panelWidth-8)) + "\n" +
		"    " + mutedStyle.Render(m.wrapText(help, max(8, layout.panelWidth-12))) + "\n"
}

func (m Model) renderOption(selected bool, option string, width int) string {
	prefix := "  "
	style := normalStyle
	if selected {
		prefix = "> "
		style = selectedStyle
	}
	return style.Render(truncatePlain(prefix+option, max(8, width))) + "\n"
}

func (m Model) renderGuide(guide fieldGuide, width int) string {
	lines := []string{
		labelStyle.Render(m.t("Prompt", "질문")) + " " + m.wrapText(guide.Hint, width),
		labelStyle.Render(m.t("Why", "왜 필요한가요")) + " " + m.wrapText(guide.Why, width),
		labelStyle.Render(m.t("Good answer", "좋은 답변")) + " " + m.wrapText(guide.Good, width),
		labelStyle.Render(m.t("Example", "예시")) + " " + m.wrapText(guide.Example, width),
		labelStyle.Render(m.t("Can be blank?", "비워도 되나요")) + " " + m.wrapText(guide.Optional, width),
		labelStyle.Render(m.t("Reflected in", "반영되는 곳")) + " " + m.wrapText(guide.MapsTo, width),
	}
	return strings.Join(lines, "\n")
}

func (m Model) wrapText(text string, width int) string {
	width = max(8, width)
	return lipgloss.NewStyle().Width(width).Render(text)
}

func (m Model) statusLine() (string, string) {
	switch m.stage {
	case stageLanguage:
		return "detected / language", m.t("Project directory detected. Language is needed before guided intent capture.", "Project directory를 감지했습니다. guided intent capture 전에 언어를 선택합니다.")
	case stageExisting:
		return "warning / existing files", m.t("Planning files already exist. Init will not overwrite them.", "Planning 파일이 이미 있습니다. init은 덮어쓰지 않습니다.")
	case stageConfirm:
		return "ready / review", m.t("Ready to write initial intent artifacts if you confirm.", "확인하면 초기 intent artifact를 쓸 준비가 되었습니다.")
	case stageDone:
		return "done", m.t("The TUI flow has ended.", "TUI 흐름이 끝났습니다.")
	default:
		return "blocked / drafting", m.t("The plan is not locked. Downstream work must wait for status, end, and run.", "Plan은 아직 locked 상태가 아닙니다. downstream 작업은 status, end, run 이후까지 대기해야 합니다.")
	}
}

func (m Model) modeLabel() string {
	switch m.stage {
	case stageLanguage:
		return m.t("language select", "언어 선택")
	case stageExisting:
		return m.t("existing-file guard", "기존 파일 보호")
	case stageConfirm:
		return m.t("review/write", "검토/저장")
	case stageDone:
		return m.t("done", "완료")
	default:
		return m.t("guided draft", "guided 초안")
	}
}

func (m Model) filesForPanel() []string {
	if len(m.existingFiles) > 0 {
		items := make([]string, 0, len(m.existingFiles))
		for _, path := range m.existingFiles {
			items = append(items, warningStyle.Render("existing")+" "+path)
		}
		return items
	}
	required := docstore.RequiredPaths()
	items := make([]string, 0, len(required))
	for _, path := range required {
		label := mutedStyle.Render("planned")
		if m.stage == stageConfirm {
			label = successStyle.Render("ready")
		}
		items = append(items, label+" "+path)
	}
	return items
}

func splitLines(text string) []string {
	if text == "" {
		return []string{""}
	}
	return strings.Split(text, "\n")
}

func fitToHeight(text string, height int) string {
	lines := splitLines(text)
	if len(lines) <= height {
		return text
	}
	return strings.Join(lines[:height], "\n")
}

func truncatePlain(text string, width int) string {
	width = max(1, width)
	if lipgloss.Width(text) <= width {
		return text
	}
	runes := []rune(text)
	if width <= 1 {
		return string(runes[:min(len(runes), 1)])
	}
	if len(runes) <= width {
		return text
	}
	return string(runes[:width-1]) + "…"
}

func (m Model) renderWritePlan() string {
	lines := []string{
		m.t("Will create or fill missing files:", "생성하거나 누락분을 채울 파일:"),
		"- docs/plan/**",
		"- .ni/contract.json",
		"- .ni/session.json",
		"- .ni/readiness.rules.json, .ni/readiness.profiles.json",
		"",
		m.t("Next commands:", "다음 명령:"),
		fmt.Sprintf("- %s status --proof --next-questions", m.commandName),
		fmt.Sprintf("- %s end", m.commandName),
		fmt.Sprintf("- %s run --max-chars 4000", m.commandName),
		"",
		m.t("Safe update path:", "안전한 update 경로:"),
		fmt.Sprintf("- %s version", m.commandName),
		m.t("- Check the installer/release notes before updating; this init screen never updates automatically.", "- update 전 installer와 release note를 확인하세요. 이 init 화면은 자동 update를 실행하지 않습니다."),
	}
	return strings.Join(lines, "\n")
}

func (m Model) fieldGuide(key string) fieldGuide {
	if m.language == languageEnglish {
		return englishGuide(key)
	}
	return koreanGuide(key)
}

func (m Model) boundaryDefault() string {
	return m.t("Do not execute downstream work before the plan is locked.", "Plan이 locked 되기 전에는 downstream work를 실행하지 않는다.")
}

func (m Model) deferralDefault() string {
	return m.t("None recorded yet.", "아직 기록된 deferral 없음.")
}

func (m Model) t(en string, ko string) string {
	if m.language == languageEnglish {
		return en
	}
	return ko
}

func renderProgress(current int, total int) string {
	if current < 1 {
		current = 1
	}
	if current > total {
		current = total
	}
	filled := current
	empty := total - current
	return fmt.Sprintf("Step %d/%d  [%s%s]", current, total, strings.Repeat("#", filled), strings.Repeat("-", empty))
}

func englishGuide(key string) fieldGuide {
	switch key {
	case "project_name":
		return fieldGuide{
			Label:    "Project name",
			Hint:     "Name the planning workspace.",
			Why:      "The contract needs a stable name so agents can refer to the same intent.",
			Good:     "Short, specific, and recognizable.",
			Example:  "Customer support handoff redesign",
			Optional: "You can leave it as the folder name.",
			MapsTo:   ".ni/contract.json project, .ni/project.json, docs/plan/00_project_brief.md",
		}
	case "project_goal":
		return fieldGuide{
			Label:    "One-sentence project goal",
			Hint:     "What should change, for whom, and why?",
			Why:      "The goal becomes the first draft of project purpose and anchors readiness questions.",
			Good:     "One sentence with outcome, audience, and reason.",
			Example:  "Help support leads prepare safer agent handoffs before any automation runs.",
			Optional: "If unsure, write TODO or unknown; status will keep it blocking.",
			MapsTo:   "docs/plan/00_project_brief.md, .ni/contract.json project.purpose",
		}
	case "target_users":
		return fieldGuide{
			Label:    "Target users / audience",
			Hint:     "Who depends on this plan?",
			Why:      "Actors clarify whose needs, permissions, and outcomes the plan must protect.",
			Good:     "Name primary users and any reviewers or operators.",
			Example:  "Korean-speaking product managers and the AI agent that receives the final handoff.",
			Optional: "You can write unknown; it should be resolved before lock.",
			MapsTo:   "docs/plan/01_actors_outcomes.md",
		}
	case "downstream_task":
		return fieldGuide{
			Label:    "Downstream agent task",
			Hint:     "What may happen only after lock and handoff?",
			Why:      "Namba Intent must know what downstream work it is preventing until the plan is ready.",
			Good:     "Describe the future task without making init execute it.",
			Example:  "After lock, compile a handoff prompt for an agent to implement the onboarding flow.",
			Optional: "You can leave it vague; status will ask for sharper scope.",
			MapsTo:   "docs/plan/02_capabilities.md, docs/plan/08_delivery_operation.md",
		}
	case "constraints":
		return fieldGuide{
			Label:    "Constraints / non-goals",
			Hint:     "Keep hard boundaries visible.",
			Why:      "Constraints stop agents from turning planning into execution too early.",
			Good:     "Use firm boundaries and explicit non-goals.",
			Example:  "Do not add a shell adapter, queue, PR automation, or downstream executor.",
			Optional: "Keep the default if you have no extra constraints yet.",
			MapsTo:   "docs/plan/05_constraints.md, .ni/contract.json non_goals",
		}
	case "success":
		return fieldGuide{
			Label:    "Success criteria",
			Hint:     "How will the accepted plan be judged?",
			Why:      "Every capability needs an evaluation before the plan can lock.",
			Good:     "Observable criteria that a reviewer can check.",
			Example:  "A new user can run status and understand blockers and next questions.",
			Optional: "If unknown, write TODO; it will block readiness.",
			MapsTo:   "docs/plan/07_evaluation_contract.md, .ni/contract.json evaluations",
		}
	case "blockers":
		return fieldGuide{
			Label:    "Known blockers or open questions",
			Hint:     "Blocking uncertainty prevents lock.",
			Why:      "Open blocker questions must stay visible until the user resolves them.",
			Good:     "Name the uncertainty and why it matters.",
			Example:  "Which install path is primary for Windows users?",
			Optional: "You may write none known; later status can still discover blockers.",
			MapsTo:   "docs/plan/10_open_questions.md, docs/plan/06_risks_security.md",
		}
	case "deferrals":
		return fieldGuide{
			Label:    "Deferrals, if any",
			Hint:     "Explicitly name deferred scope.",
			Why:      "Deferrals keep later work from sneaking into the locked plan.",
			Good:     "List work that is intentionally out of this first plan.",
			Example:  "Web GUI, automatic updates, and downstream execution are deferred.",
			Optional: "The default is fine when nothing is deferred yet.",
			MapsTo:   "docs/plan/10_open_questions.md",
		}
	}
	return fieldGuide{Label: key}
}

func koreanGuide(key string) fieldGuide {
	switch key {
	case "project_name":
		return fieldGuide{
			Label:    "프로젝트 이름",
			Hint:     "이 planning workspace를 부를 짧은 이름입니다.",
			Why:      "계약이 같은 의도를 계속 가리키려면 안정적인 이름이 필요합니다.",
			Good:     "짧고 구체적이며 나중에 다시 봐도 알아볼 수 있는 이름.",
			Example:  "고객 문의 handoff 개선",
			Optional: "그대로 두면 현재 폴더 이름을 사용합니다.",
			MapsTo:   ".ni/contract.json project, .ni/project.json, docs/plan/00_project_brief.md",
		}
	case "project_goal":
		return fieldGuide{
			Label:    "프로젝트 목표 한 문장",
			Hint:     "누구에게 무엇이 좋아져야 하고, 왜 필요한가요?",
			Why:      "목표는 project purpose 초안이 되고 readiness 질문의 기준점이 됩니다.",
			Good:     "대상, 변화, 이유가 한 문장 안에 들어간 답변.",
			Example:  "한국어 사용자가 agent에게 일을 맡기기 전에 의도와 경계를 안전하게 정리하게 한다.",
			Optional: "모르면 TODO 또는 모름이라고 써도 됩니다. status가 blocker로 되돌려줍니다.",
			MapsTo:   "docs/plan/00_project_brief.md, .ni/contract.json project.purpose",
		}
	case "target_users":
		return fieldGuide{
			Label:    "대상 사용자 / 독자",
			Hint:     "이 plan에 기대는 사람이나 agent는 누구인가요?",
			Why:      "Actor를 알아야 누구의 요구, 권한, 결과를 보호해야 하는지 정할 수 있습니다.",
			Good:     "주 사용자와 검토자, 운영자, handoff를 받을 agent를 함께 적습니다.",
			Example:  "비개발자 한국어 사용자, planning을 도와주는 모델, 최종 handoff를 받을 구현 agent.",
			Optional: "모르면 모름이라고 적어도 됩니다. lock 전에는 좁혀야 합니다.",
			MapsTo:   "docs/plan/01_actors_outcomes.md",
		}
	case "downstream_task":
		return fieldGuide{
			Label:    "나중에 agent가 할 일",
			Hint:     "lock과 handoff 뒤에야 해도 되는 일은 무엇인가요?",
			Why:      "Namba Intent는 downstream 일이 너무 일찍 시작되지 않도록 그 일을 먼저 이름 붙입니다.",
			Good:     "미래 작업을 설명하되, init이 그 일을 실행하는 것처럼 쓰지 않습니다.",
			Example:  "Plan이 locked 된 뒤 onboarding TUI 개선을 구현할 handoff prompt를 compile한다.",
			Optional: "흐릿하게 적어도 됩니다. status가 더 구체적인 범위를 물어봅니다.",
			MapsTo:   "docs/plan/02_capabilities.md, docs/plan/08_delivery_operation.md",
		}
	case "constraints":
		return fieldGuide{
			Label:    "제약 / 하지 않을 일",
			Hint:     "반드시 지켜야 할 경계와 non-goal입니다.",
			Why:      "제약은 planning이 너무 빨리 실행 도구로 변하는 것을 막습니다.",
			Good:     "단호한 경계와 이번 범위에서 제외할 일을 함께 적습니다.",
			Example:  "웹 GUI, shell adapter, queue, PR automation, downstream executor는 만들지 않는다.",
			Optional: "추가 제약이 없으면 기본 문장을 유지해도 됩니다.",
			MapsTo:   "docs/plan/05_constraints.md, .ni/contract.json non_goals",
		}
	case "success":
		return fieldGuide{
			Label:    "성공 기준",
			Hint:     "이 plan이 괜찮다고 어떻게 판단할 수 있나요?",
			Why:      "모든 capability는 lock 전에 evaluation과 연결되어야 합니다.",
			Good:     "검토자가 실제로 확인할 수 있는 관찰 가능한 기준.",
			Example:  "처음 온 사용자가 status 결과에서 blocker와 다음 질문을 이해할 수 있다.",
			Optional: "모르면 TODO라고 적으세요. readiness를 통과하지 못하게 막아줍니다.",
			MapsTo:   "docs/plan/07_evaluation_contract.md, .ni/contract.json evaluations",
		}
	case "blockers":
		return fieldGuide{
			Label:    "막힌 점 / 열린 질문",
			Hint:     "확실하지 않아서 lock을 막아야 하는 질문입니다.",
			Why:      "Open blocker는 사용자가 해결할 때까지 계속 보이는 상태로 남아야 합니다.",
			Good:     "무엇이 불확실한지와 왜 중요한지를 함께 적습니다.",
			Example:  "Windows 사용자의 primary install/update 경로를 아직 검증하지 못했다.",
			Optional: "없으면 없음이라고 적어도 됩니다. status가 새 blocker를 찾을 수도 있습니다.",
			MapsTo:   "docs/plan/10_open_questions.md, docs/plan/06_risks_security.md",
		}
	case "deferrals":
		return fieldGuide{
			Label:    "나중으로 미룰 일",
			Hint:     "이번 plan 밖으로 명시적으로 미루는 범위입니다.",
			Why:      "Deferral은 나중 작업이 locked plan 안으로 몰래 들어오는 것을 막습니다.",
			Good:     "이번에 하지 않는 일을 분명한 이름으로 적습니다.",
			Example:  "웹 GUI, 자동 update 실행, downstream execution layer는 이번 범위에서 제외한다.",
			Optional: "아직 미룰 일이 없으면 기본값을 그대로 둬도 됩니다.",
			MapsTo:   "docs/plan/10_open_questions.md",
		}
	}
	return fieldGuide{Label: key}
}

var (
	colorInk     = lipgloss.Color("#111827")
	colorMuted   = lipgloss.Color("#64748b")
	colorAccent  = lipgloss.Color("#0f766e")
	colorWarning = lipgloss.Color("#b45309")
	colorBlocked = lipgloss.Color("#9f1239")

	baseStyle        = lipgloss.NewStyle().Foreground(colorInk)
	borderStyle      = lipgloss.NewStyle().BorderForeground(colorMuted)
	panelStyle       = baseStyle.Padding(1, 2).Border(lipgloss.RoundedBorder(), true).Inherit(borderStyle)
	infoPanelStyle   = baseStyle.Padding(1, 2).Border(lipgloss.NormalBorder(), true).Inherit(borderStyle)
	headerBarStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Background(colorAccent)
	footerStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Background(colorMuted)
	helpBarStyle     = lipgloss.NewStyle().Foreground(colorInk).Background(lipgloss.Color("#ffffff"))
	compactModeStyle = lipgloss.NewStyle().Bold(true)

	headerTitleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff"))
	panelHeadingStyle = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)
	sectionStyle      = lipgloss.NewStyle().Bold(true).Foreground(colorMuted)
	focusedStyle      = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)
	selectedStyle     = focusedStyle
	labelStyle        = lipgloss.NewStyle().Bold(true).Foreground(colorInk)
	normalStyle       = baseStyle
	mutedStyle        = lipgloss.NewStyle().Foreground(colorMuted)
	progressStyle     = lipgloss.NewStyle().Foreground(colorWarning)
	safetyStyle       = lipgloss.NewStyle().Foreground(colorBlocked)
	statusStyle       = lipgloss.NewStyle().Bold(true).Foreground(colorMuted)
	successStyle      = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)
	warningStyle      = lipgloss.NewStyle().Bold(true).Foreground(colorWarning)
	blockedStyle      = lipgloss.NewStyle().Bold(true).Foreground(colorBlocked)
)
