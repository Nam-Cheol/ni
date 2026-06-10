package initui

import (
	"fmt"
	"io"
	"strings"
	"time"

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

type layoutMode string

const (
	layoutWide   layoutMode = "wide"
	layoutMedium layoutMode = "medium"
	layoutNarrow layoutMode = "narrow"
	layoutTiny   layoutMode = "tiny"
)

type language string

const (
	languageKorean  language = "ko"
	languageEnglish language = "en"
)

type initTickMsg time.Time

const initTickInterval = 360 * time.Millisecond

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
	frame          int
	animationOn    bool
	detailsOpen    bool
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
		animationOn:   true,
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
	if !m.animationOn {
		return requestWindowSizeCmd()
	}
	return tea.Batch(requestWindowSizeCmd(), m.animationTick())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = size.Width
		m.height = size.Height
		m.constrainScroll()
		return m, nil
	}
	if _, ok := msg.(initTickMsg); ok {
		if !m.animationOn || m.stage == stageDone || m.canceled {
			return m, nil
		}
		m.frame = (m.frame + 1) % 120
		m.constrainScroll()
		return m, m.animationTick()
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
	case "ctrl+u":
		if m.stage == stageFields {
			m.fields[m.cursor].Value = ""
			return m, nil
		}
	case "tab":
		if m.layout().mode != layoutTiny {
			m.detailsOpen = !m.detailsOpen
			m.resetViewport()
		}
		return m, nil
	case "d":
		if m.stage != stageFields && m.layout().mode != layoutTiny {
			m.detailsOpen = !m.detailsOpen
			m.resetViewport()
			return m, nil
		}
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

func requestWindowSizeCmd() tea.Cmd {
	return func() tea.Msg { return tea.RequestWindowSize() }
}

func (m Model) animationTick() tea.Cmd {
	return tea.Tick(initTickInterval, func(t time.Time) tea.Msg {
		return initTickMsg(t)
	})
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
	width       int
	height      int
	bodyHeight  int
	panelWidth  int
	shellLeft   int
	shellWidth  int
	chatWidth   int
	drawerWidth int
	mode        layoutMode
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
	header := m.renderHeader(layout)
	help := m.renderBottomHelp(layout)
	body := m.bodyViewport.View()
	screen := lipgloss.JoinVertical(lipgloss.Left, header, body, help)
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
	mode := layoutWide
	switch {
	case width < 60 || height < 16:
		mode = layoutTiny
	case width < 80:
		mode = layoutNarrow
	case width < 110:
		mode = layoutMedium
	default:
		mode = layoutWide
	}
	bodyHeight := height - 2
	if bodyHeight < 1 {
		bodyHeight = 1
	}
	horizontalMargin := 2
	maxShellWidth := 104
	switch mode {
	case layoutMedium:
		horizontalMargin = 2
		maxShellWidth = max(24, width-4)
	case layoutNarrow:
		horizontalMargin = 1
		maxShellWidth = max(24, width-2)
	case layoutTiny:
		horizontalMargin = 0
		maxShellWidth = width
	}
	availableWidth := max(24, width-horizontalMargin*2)
	shellWidth := min(availableWidth, maxShellWidth)
	shellLeft := max(0, (width-shellWidth)/2)
	chatWidth := shellWidth
	drawerWidth := 0
	if mode == layoutWide && m.detailsOpen {
		drawerWidth = min(36, max(28, shellWidth/3))
		chatWidth = max(48, shellWidth-drawerWidth-2)
	} else if m.detailsOpen && mode != layoutTiny {
		drawerWidth = shellWidth
	}
	return layoutSpec{
		width:       width,
		height:      height,
		bodyHeight:  bodyHeight,
		panelWidth:  chatWidth,
		shellLeft:   shellLeft,
		shellWidth:  shellWidth,
		chatWidth:   chatWidth,
		drawerWidth: drawerWidth,
		mode:        mode,
	}
}

func (m Model) renderHeader(layout layoutSpec) string {
	left := headerTitleStyle.Render("Namba Intent")
	progress := fmt.Sprintf("init · %d/%d · %s · d details", m.currentStep(), m.totalSteps(), renderProgress(m.currentStep(), m.totalSteps(), m.frame, layout))
	if m.stage == stageDone {
		progress = m.t("init · closed", "init · 종료")
	}
	right := headerMetaStyle.Render(progress)
	if layout.mode == layoutTiny {
		line := compactModeStyle.Render(truncatePlain("Namba Intent  "+progress, layout.shellWidth))
		return headerBarStyle.Width(layout.width).Render(line)
	}
	gap := max(1, layout.shellWidth-lipgloss.Width(left)-lipgloss.Width(right))
	content := left + strings.Repeat(" ", gap) + right
	line := strings.Repeat(" ", layout.shellLeft) + truncatePlain(content, layout.shellWidth)
	return headerBarStyle.Width(layout.width).Render(line)
}

func (m Model) renderBody(layout layoutSpec) string {
	if layout.mode == layoutTiny {
		return m.renderTinyChat(layout)
	}
	chat := m.renderChatShell(layout.withWidth(layout.chatWidth))
	if !m.detailsOpen {
		return renderShellContent(layout, chat)
	}
	if layout.mode == layoutWide {
		body := lipgloss.JoinHorizontal(lipgloss.Top,
			chat,
			"  ",
			m.renderDetailsDrawer(layout.withWidth(layout.drawerWidth)),
		)
		return renderShellContent(layout, body)
	}
	return renderShellContent(layout, lipgloss.JoinVertical(lipgloss.Left, chat, m.renderDetailsDrawer(layout.withWidth(layout.drawerWidth))))
}

func (l layoutSpec) withWidth(width int) layoutSpec {
	l.panelWidth = width
	l.width = width
	l.shellLeft = 0
	l.shellWidth = width
	l.chatWidth = width
	l.drawerWidth = 0
	return l
}

func (m Model) renderChatShell(layout layoutSpec) string {
	inner := layout.withWidth(layout.chatWidth)
	parts := []string{}
	transcript := m.renderChatTranscript(inner)
	parts = append(parts, transcript)
	if chips := m.renderDoneCommandChips(inner); chips != "" {
		parts = append(parts, chips)
	}
	composerParts := []string{}
	switch m.stage {
	case stageLanguage, stageExisting, stageConfirm:
		composerParts = append(composerParts, m.renderChoiceComposer(inner))
	case stageFields:
		composerParts = append(composerParts, m.renderTextComposer(inner))
	}
	composerParts = append(composerParts, m.renderQuietSummary(inner))
	spacerHeight := max(1, layout.bodyHeight-renderLineCount(parts)-renderLineCount(composerParts)-1)
	if layout.mode == layoutNarrow {
		spacerHeight = max(1, min(spacerHeight, 3))
	}
	parts = append(parts, strings.Repeat("\n", spacerHeight))
	parts = append(parts, composerParts...)
	return strings.Join(parts, "\n")
}

func (m Model) renderChatTranscript(layout layoutSpec) string {
	rows := []string{}
	for _, summary := range m.recentUserSummaries(layout) {
		rows = append(rows, m.renderUserSummaryBubble(summary, layout))
	}
	rows = append(rows, m.renderAssistantRow(layout))
	return strings.Join(rows, "\n")
}

func (m Model) renderAssistantRow(layout layoutSpec) string {
	if layout.mode == layoutNarrow {
		avatar := asciiStyle.Render(m.stageAsset(layout).Compact)
		message := "assistant: " + stripNewlines(m.assistantMessage(layout))
		return avatar + " " + secondaryStyle.Render(truncatePlain(message, max(12, layout.panelWidth-lipgloss.Width(m.stageAsset(layout).Compact)-1))) + "\n"
	}
	title := asciiStyle.Render(m.stageAsset(layout).Compact) + " " + sectionStyle.Render("assistant")
	return m.renderChatBox(title, m.assistantMessage(layout), layout.chatWidth, assistantBubbleStyle) + "\n"
}

func (m Model) renderUserSummaryBubble(summary string, layout layoutSpec) string {
	chip := userChipStyle.Render(truncatePlain("you · "+summary, max(12, layout.panelWidth-4)))
	pad := 0
	if layout.mode != layoutNarrow {
		pad = max(0, layout.panelWidth-lipgloss.Width(chip)-2)
	}
	return lipgloss.NewStyle().MarginLeft(pad).Render(chip) + "\n"
}

func (m Model) renderChoiceComposer(layout layoutSpec) string {
	switch m.stage {
	case stageLanguage:
		options := []struct {
			label string
			help  string
		}{
			{"Korean", "한국어 사용자에게 자연스러운 planning 안내"},
			{"English", "English labels, examples, and review guidance"},
		}
		lines := make([]string, 0, len(options))
		for i, option := range options {
			help := option.help
			if layout.mode == layoutNarrow || layout.mode == layoutMedium {
				help = ""
			}
			lines = append(lines, strings.TrimRight(m.renderChoiceLine(i == m.languageCursor, option.label, help, layout), "\n"))
		}
		return m.renderComposerBox("choose one", strings.Join(lines, "\n"), layout)
	case stageExisting:
		options := []string{
			m.t("Add missing files only", "누락된 파일만 추가"),
			m.t("Keep existing and exit", "기존 파일 유지 후 종료"),
			m.t("Abort", "중단"),
		}
		lines := make([]string, 0, len(options))
		for i, option := range options {
			lines = append(lines, strings.TrimRight(m.renderOption(i == m.existingCursor, option, layout.panelWidth-8), "\n"))
		}
		return m.renderComposerBox("choose one", strings.Join(lines, "\n"), layout)
	case stageConfirm:
		options := []string{
			m.t("Write initial intent artifacts", "초기 intent 파일 저장"),
			m.t("Cancel; write nothing", "취소하고 아무것도 쓰지 않기"),
		}
		lines := make([]string, 0, len(options))
		for i, option := range options {
			lines = append(lines, strings.TrimRight(m.renderOption(i == m.confirmCursor, option, layout.panelWidth-8), "\n"))
		}
		return m.renderComposerBox("choose one", strings.Join(lines, "\n"), layout)
	default:
		return ""
	}
}

func (m Model) renderTextComposer(layout layoutSpec) string {
	guide := m.fieldGuide(m.fields[m.cursor].Key)
	value := strings.TrimSpace(m.fields[m.cursor].Value)
	if value == "" {
		value = "> " + m.textPlaceholder(guide)
	} else {
		value = "> " + value
	}
	width := m.composerWidth(layout)
	body := answerStyle.Render(m.wrapComposerText(value, width-8))
	return m.renderComposerBox(m.t("your answer", "your answer"), body, layout)
}

func (m Model) renderDoneCommandChips(layout layoutSpec) string {
	if m.stage != stageDone || !m.confirmed {
		return ""
	}
	commands := []string{
		fmt.Sprintf("%s status --proof --next-questions", m.commandName),
		fmt.Sprintf("%s end", m.commandName),
	}
	lines := make([]string, 0, len(commands))
	for _, command := range commands {
		lines = append(lines, commandChipStyle.Render(truncatePlain(command, max(12, layout.panelWidth-4))))
	}
	return strings.Join(lines, "\n")
}

func (m Model) renderDetailsDrawer(layout layoutSpec) string {
	width := max(28, layout.panelWidth)
	required := docstore.RequiredPaths()
	items, _ := m.filesForPanel(layout.withWidth(width))
	lines := []string{
		fmt.Sprintf("docs planned: %d · contract draft: ready", len(required)),
		"CLI gates stay authoritative",
		"status, end, run",
		"init writes drafts only; no downstream work runs",
		"",
		"files",
	}
	if len(items) == 0 {
		lines = append(lines, "  no existing planning files detected")
	} else {
		for _, item := range items {
			label := "will create"
			if len(m.existingFiles) > 0 {
				label = "preserve"
			}
			lines = append(lines, truncatePlain("  "+label+" · "+item.Path, max(8, width-6)))
		}
	}
	if hidden := len(required) - len(items); hidden > 0 && len(m.existingFiles) == 0 {
		lines = append(lines, fmt.Sprintf("  +%d more in docs/plan and .ni", hidden))
	}
	lines = append(lines, "", "Tab or d closes details")
	return m.renderChatBox(sectionStyle.Render("details"), strings.Join(lines, "\n"), width, detailsDrawerStyle)
}

func (m Model) renderBottomHelp(layout layoutSpec) string {
	var text string
	switch m.stage {
	case stageLanguage:
		text = "↑↓ choose · Enter select · 1/2 quick select · q quit · d details"
	case stageFields:
		text = "type answer · Enter send · Esc back · Ctrl+U clear · q quit · d details"
	case stageConfirm:
		text = "↑↓ choose · Enter confirm · Esc back · q quit · d details"
	case stageDone:
		text = "q quit"
	default:
		text = "↑↓ choose · Enter send · Esc back · q quit · d details"
	}
	if layout.mode == layoutTiny {
		text = "↑↓ · Enter · Esc · q"
	}
	content := truncatePlain(text, layout.shellWidth)
	return helpBarStyle.Width(layout.width).Render(strings.Repeat(" ", layout.shellLeft) + content)
}

func (m Model) renderTinyChat(layout layoutSpec) string {
	lines := []string{
		asciiStyle.Render(m.stageAsset(layout).Compact),
	}
	switch m.stage {
	case stageLanguage:
		lines = append(lines,
			questionStyle.Render(truncatePlain(m.t("Choose intent language.", "Intent 언어를 선택하세요."), layout.width)),
			strings.TrimRight(m.renderChoiceLine(m.languageCursor == 0, "Korean", "", layout), "\n"),
			strings.TrimRight(m.renderChoiceLine(m.languageCursor == 1, "English", "", layout), "\n"),
		)
	case stageExisting:
		lines = append(lines,
			questionStyle.Render(truncatePlain(m.t("Choose a safe existing-file path.", "기존 파일 처리 방식을 선택하세요."), layout.width)),
			strings.TrimRight(m.renderOption(m.existingCursor == 0, m.t("Add missing files only", "누락분만 추가"), layout.width), "\n"),
			strings.TrimRight(m.renderOption(m.existingCursor == 1, m.t("Keep existing and exit", "기존 유지"), layout.width), "\n"),
			strings.TrimRight(m.renderOption(m.existingCursor == 2, m.t("Abort", "중단"), layout.width), "\n"),
		)
	case stageConfirm:
		lines = append(lines,
			questionStyle.Render(truncatePlain(m.t("Write initial intent artifacts?", "초기 intent 파일을 저장할까요?"), layout.width)),
			strings.TrimRight(m.renderOption(m.confirmCursor == 0, m.t("Write initial intent artifacts", "초기 intent 파일 저장"), layout.width), "\n"),
			strings.TrimRight(m.renderOption(m.confirmCursor == 1, m.t("Cancel; write nothing", "취소하고 쓰지 않기"), layout.width), "\n"),
		)
	case stageDone:
		if m.canceled {
			lines = append(lines,
				questionStyle.Render("Cancelled"),
				secondaryStyle.Render("No files written"),
				secondaryStyle.Render(truncatePlain("Run namba-intent init again", layout.width)),
			)
		} else {
			lines = append(lines,
				questionStyle.Render(m.t("Done", "완료")),
				secondaryStyle.Render(truncatePlain(m.t("Initial intent draft is ready.", "초기 intent 초안이 준비됐습니다."), layout.width)),
			)
		}
	default:
		guide := m.fieldGuide(m.fields[m.cursor].Key)
		value := strings.TrimSpace(m.fields[m.cursor].Value)
		if value == "" {
			value = "> " + m.textPlaceholder(guide)
		} else {
			value = "> " + value
		}
		lines = append(lines,
			questionStyle.Render(truncatePlain(guide.Hint, layout.width)),
			answerStyle.Render(truncatePlain(value, layout.width)),
		)
	}
	lines = append(lines, mutedStyle.Render(truncatePlain(m.compactDraftSummary(), layout.width)))
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func (m Model) renderQuietSummary(layout layoutSpec) string {
	if layout.mode == layoutTiny {
		return mutedStyle.Render(truncatePlain(m.compactDraftSummary(), layout.width))
	}
	return mutedStyle.Render(truncatePlain(m.compactDraftSummary()+" · d details", layout.panelWidth))
}

func (m Model) renderChatBox(title string, body string, width int, style lipgloss.Style) string {
	width = max(20, width)
	return style.Width(width).Render(title + "\n" + body)
}

func (m Model) renderComposerBox(title string, body string, layout layoutSpec) string {
	width := m.composerWidth(layout)
	return composerStyle.Width(width).Render(sectionStyle.Render(title) + "\n" + body)
}

func (m Model) composerWidth(layout layoutSpec) int {
	return max(28, layout.chatWidth)
}

func (m Model) avatarWidth(layout layoutSpec) int {
	if layout.mode == layoutNarrow {
		return lipgloss.Width(m.stageAsset(layout).Compact)
	}
	width := 0
	for _, line := range m.stageAsset(layout).Lines {
		width = max(width, lipgloss.Width(line))
	}
	return width
}

func (m Model) recentUserSummaries(layout layoutSpec) []string {
	summaries := []string{}
	if m.stage != stageLanguage {
		summaries = append(summaries, m.t("language: "+m.languageLabel(), "언어: "+m.languageLabel()))
	}
	if m.stage == stageFields && m.cursor > 0 {
		if summary := m.fieldSummary(m.cursor - 1); summary != "" {
			summaries = append(summaries, summary)
		}
	}
	if m.stage == stageConfirm || m.stage == stageDone {
		indexes := []int{max(0, len(m.fields)-2), len(m.fields) - 1}
		if len(m.fields) > 2 {
			indexes = []int{1, 5}
		}
		for _, index := range indexes {
			if summary := m.fieldSummary(index); summary != "" {
				summaries = append(summaries, summary)
			}
		}
	}
	if len(summaries) > 2 {
		summaries = summaries[len(summaries)-2:]
	}
	if layout.mode == layoutNarrow && len(summaries) > 1 {
		return summaries[len(summaries)-1:]
	}
	return summaries
}

func (m Model) fieldSummary(index int) string {
	if index < 0 || index >= len(m.fields) {
		return ""
	}
	value := strings.TrimSpace(m.fields[index].Value)
	if value == "" {
		return ""
	}
	label := m.fieldGuide(m.fields[index].Key).Label
	return truncatePlain(label+": "+value, 72)
}

func (m Model) languageLabel() string {
	if m.language == languageEnglish {
		return "English"
	}
	return "Korean"
}

func (m Model) textPlaceholder(guide fieldGuide) string {
	switch guide.Label {
	case "One-sentence project goal", "프로젝트 목표 한 문장":
		return m.t("enter the project goal in one sentence", "프로젝트 목표를 한 문장으로 입력하세요")
	case "Project name", "프로젝트 이름":
		return m.t("type a short project name", "짧은 프로젝트 이름을 입력하세요")
	default:
		return m.t("type your answer here", "여기에 답을 입력하세요")
	}
}

func (m Model) wrapComposerText(text string, width int) string {
	lines := splitLines(m.wrapText(text, max(8, width)))
	if len(lines) > 3 {
		lines = lines[len(lines)-3:]
	}
	return strings.Join(lines, "\n")
}

func (m Model) typingIndicator() string {
	switch (m.frame / 3) % 3 {
	case 0:
		return "·"
	case 1:
		return "··"
	default:
		return "···"
	}
}

func (m Model) assistantMessage(layout layoutSpec) string {
	typing := m.typingIndicator()
	switch m.stage {
	case stageLanguage:
		return m.wrapText(m.t(
			"Which language should we use for this intent?\nI will use it for labels and review guidance.\nChoose one below. "+typing,
			"먼저 intent를 어떤 언어로 작성할까요?\n문서와 리뷰 안내에 같은 언어를 사용합니다.\n아래에서 하나를 고르세요. "+typing,
		), max(8, layout.panelWidth-8))
	case stageExisting:
		return m.wrapText(m.t(
			"I found existing planning files.\nI will not overwrite them.\nChoose the safest path. "+typing,
			"기존 planning 파일을 찾았습니다.\n덮어쓰지 않습니다.\n안전한 경로를 골라주세요. "+typing,
		), max(8, layout.panelWidth-8))
	case stageConfirm:
		return m.wrapText(m.t(
			"Ready to write the initial intent draft?\nThis still does not decide readiness or lock the plan.\nConfirm only if this draft looks right. "+typing,
			"초기 intent 초안을 저장할까요?\nreadiness 판단이나 lock은 아직 CLI gate의 일입니다.\n초안이 맞으면 확인하세요. "+typing,
		), max(8, layout.panelWidth-8))
	case stageDone:
		if m.canceled {
			return m.wrapText(m.t(
				"Cancelled\nNo files written\nRun namba-intent init again when ready.",
				"Cancelled\nNo files written\n준비되면 namba-intent init을 다시 실행하세요.",
			), max(8, layout.panelWidth-8))
		}
		return m.wrapText(m.t(
			"Initial intent draft is ready.\nUse the command chips below for the next gate.\nKeep answering blockers before locking.",
			"초기 intent 초안이 준비됐습니다.\n아래 command chip으로 다음 gate를 확인하세요.\nlock 전에는 blocker를 계속 답하세요.",
		), max(8, layout.panelWidth-8))
	default:
		guide := m.fieldGuide(m.fields[m.cursor].Key)
		return m.wrapText(guide.Hint+"\n"+guide.Why+"\n"+m.t("Answer in one compact sentence if you can. ", "가능하면 한 문장으로 답하세요. ")+typing, max(8, layout.panelWidth-8))
	}
}

func (m Model) compactDraftSummary() string {
	if m.stage == stageDone && m.confirmed {
		return m.t("draft written", "초안 저장됨")
	}
	if m.stage == stageDone && m.canceled {
		return m.t("no files written", "파일 쓰지 않음")
	}
	return m.t("draft only · no files written", "초안만 · 파일 쓰지 않음")
}

func (m Model) renderChoiceLine(selected bool, label string, help string, layout layoutSpec) string {
	prefix := "  ◇ "
	style := normalStyle
	if selected {
		prefix = m.selectionPulse() + " ◆ "
		style = selectedStyle
	}
	line := prefix + label
	if layout.mode == layoutTiny || layout.mode == layoutNarrow {
		return style.Render(truncatePlain(line, layout.panelWidth-8)) + "\n"
	}
	if strings.TrimSpace(help) == "" {
		return style.Render(truncatePlain(line, layout.panelWidth-8)) + "\n"
	}
	return style.Render(truncatePlain(line, layout.panelWidth-8)) + "\n" +
		"    " + secondaryStyle.Render(m.wrapText(help, max(8, layout.panelWidth-12))) + "\n"
}

func (m Model) renderOption(selected bool, option string, width int) string {
	prefix := "  □ "
	style := normalStyle
	if selected {
		prefix = m.selectionPulse() + " ■ "
		style = selectedStyle
	}
	return style.Render(truncatePlain(prefix+option, max(8, width))) + "\n"
}

func (m Model) wrapText(text string, width int) string {
	width = max(8, width)
	return lipgloss.NewStyle().Width(width).Render(text)
}

func (m Model) totalSteps() int {
	return len(m.fields) + 2
}

func (m Model) currentStep() int {
	switch m.stage {
	case stageLanguage:
		return 1
	case stageExisting:
		return 2
	case stageConfirm, stageDone:
		return m.totalSteps()
	default:
		return m.cursor + 2
	}
}

type stageAsset struct {
	Compact string
	Lines   []string
}

func (m Model) stageAsset(layout layoutSpec) stageAsset {
	eye := []string{"o", "o", "-", "o"}[(m.frame/6)%4]
	spark := []string{"*", "+", "*", "."}[(m.frame/3)%4]
	stamp := []string{"[ ]", "[=]", "[#]", "[=]"}[(m.frame/4)%4]
	switch m.stage {
	case stageLanguage:
		return stageAsset{
			Compact: "ni>",
			Lines: []string{
				"      (\\_/)",
				"  ni  (" + eye + "_" + eye + ") " + spark,
				"      / >▣",
			},
		}
	case stageExisting:
		return stageAsset{
			Compact: "ni!",
			Lines: []string{
				"      (\\_/)",
				"  ni  (o_o)",
				"      / >□",
			},
		}
	case stageConfirm:
		return stageAsset{
			Compact: "ni#",
			Lines: []string{
				"      (\\_/)",
				"  ni  (" + eye + "_" + eye + ")",
				"      / >" + stamp,
			},
		}
	case stageDone:
		if m.canceled {
			return stageAsset{
				Compact: "nix",
				Lines: []string{
					"      (\\_/)",
					"  ni  (._.)",
					"      / >□",
				},
			}
		}
		return stageAsset{
			Compact: "ni✓",
			Lines: []string{
				"      (\\_/)",
				"  ni  (^_^)",
				"      / >▣",
			},
		}
	default:
		return stageAsset{
			Compact: "ni>",
			Lines: []string{
				"      (\\_/)",
				"  ni  (" + eye + "_" + eye + ")",
				"      / >" + spark,
			},
		}
	}
}

type filePanelItem struct {
	Status string
	Path   string
}

func (m Model) filesForPanel(layout layoutSpec) ([]filePanelItem, int) {
	if len(m.existingFiles) > 0 {
		limit := filePanelLimit(layout)
		items := make([]filePanelItem, 0, min(len(m.existingFiles), limit))
		for _, path := range m.existingFiles[:min(len(m.existingFiles), limit)] {
			items = append(items, filePanelItem{Status: "skipped", Path: path})
		}
		return items, max(0, len(m.existingFiles)-limit)
	}
	required := docstore.RequiredPaths()
	limit := filePanelLimit(layout)
	items := make([]filePanelItem, 0, min(len(required), limit))
	for _, path := range required[:min(len(required), limit)] {
		items = append(items, filePanelItem{Status: "planned", Path: path})
	}
	return items, max(0, len(required)-limit)
}

func filePanelLimit(layout layoutSpec) int {
	switch layout.mode {
	case layoutWide:
		return 5
	case layoutMedium:
		return 4
	case layoutNarrow:
		return 3
	default:
		return 1
	}
}

func (m Model) selectionPulse() string {
	return "▸"
}

func splitLines(text string) []string {
	if text == "" {
		return []string{""}
	}
	return strings.Split(text, "\n")
}

func stripNewlines(text string) string {
	fields := strings.Fields(strings.ReplaceAll(text, "\n", " "))
	return strings.Join(fields, " ")
}

func renderShellContent(layout layoutSpec, content string) string {
	if layout.shellLeft <= 0 {
		return content
	}
	pad := strings.Repeat(" ", layout.shellLeft)
	lines := splitLines(content)
	for i, line := range lines {
		lines[i] = truncatePlain(pad+line, layout.width)
	}
	return strings.Join(lines, "\n")
}

func renderLineCount(parts []string) int {
	total := 0
	for _, part := range parts {
		if part == "" {
			continue
		}
		total += len(splitLines(part))
	}
	return total
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
	if width <= 1 {
		return "…"
	}
	var b strings.Builder
	for _, r := range text {
		next := b.String() + string(r) + "…"
		if lipgloss.Width(next) > width {
			break
		}
		b.WriteRune(r)
	}
	if b.Len() == 0 {
		return "…"
	}
	return b.String() + "…"
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

func renderProgress(current int, total int, frame int, layout layoutSpec) string {
	if current < 1 {
		current = 1
	}
	if current > total {
		current = total
	}
	var b strings.Builder
	for i := 1; i <= total; i++ {
		switch {
		case i < current:
			b.WriteString("▰")
		case i == current:
			if frame%2 == 0 {
				b.WriteString("▓")
			} else {
				b.WriteString("█")
			}
		default:
			b.WriteString("▱")
		}
	}
	label := b.String()
	return truncatePlain(label, max(8, layout.panelWidth-8))
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
	brandPrimary       = lipgloss.Color("#ff7e13")
	brandPrimaryDim    = lipgloss.Color("#9a4d16")
	brandPrimaryStrong = lipgloss.Color("#ffb15e")
	textPrimary        = lipgloss.Color("#f8fafc")
	textSecondary      = lipgloss.Color("#cbd5e1")
	textMuted          = lipgloss.Color("#94a3b8")
	surface            = lipgloss.Color("#111118")
	surfaceRaised      = lipgloss.Color("#1b1714")
	borderSubtle       = lipgloss.Color("#4b5563")
	borderActive       = lipgloss.Color("#ff7e13")
	selected           = lipgloss.Color("#7c2d12")
	success            = lipgloss.Color("#86efac")
	warning            = lipgloss.Color("#fbbf24")
	blocked            = lipgloss.Color("#fb7185")
	help               = lipgloss.Color("#e2e8f0")

	baseStyle        = lipgloss.NewStyle().Foreground(textPrimary).Background(surface)
	headerBarStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#160b02")).Background(brandPrimary)
	helpBarStyle     = lipgloss.NewStyle().Foreground(help).Background(lipgloss.Color("#20202a"))
	compactModeStyle = lipgloss.NewStyle().Bold(true)

	headerTitleStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#160b02"))
	headerMetaStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#261204"))
	questionStyle        = lipgloss.NewStyle().Bold(true).Foreground(textPrimary)
	answerStyle          = lipgloss.NewStyle().Bold(true).Foreground(textPrimary)
	secondaryStyle       = lipgloss.NewStyle().Foreground(textSecondary)
	sectionStyle         = lipgloss.NewStyle().Bold(true).Foreground(textMuted)
	selectedStyle        = lipgloss.NewStyle().Bold(true).Foreground(textPrimary).Background(selected)
	labelStyle           = lipgloss.NewStyle().Bold(true).Foreground(textPrimary)
	normalStyle          = baseStyle.Foreground(textSecondary)
	mutedStyle           = lipgloss.NewStyle().Foreground(textMuted).Background(surface)
	progressStyle        = lipgloss.NewStyle().Foreground(brandPrimaryStrong)
	asciiStyle           = lipgloss.NewStyle().Foreground(brandPrimary)
	assistantBubbleStyle = baseStyle.Padding(1, 2).
				Border(lipgloss.RoundedBorder(), true).
				BorderForeground(borderSubtle)
	userBubbleStyle = baseStyle.Padding(1, 2).
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(brandPrimaryDim).
			Background(surfaceRaised)
	userChipStyle = lipgloss.NewStyle().
			Foreground(textSecondary).
			Background(surfaceRaised).
			Padding(0, 1)
	commandChipStyle = lipgloss.NewStyle().
				Foreground(textPrimary).
				Background(surfaceRaised).
				Border(lipgloss.NormalBorder(), true).
				BorderForeground(brandPrimaryDim).
				Padding(0, 1)
	composerStyle = baseStyle.Padding(0, 1).
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(borderActive)
	detailsDrawerStyle = baseStyle.Padding(1, 2).
				Border(lipgloss.NormalBorder(), true).
				BorderForeground(borderSubtle).
				Background(surfaceRaised)
)
