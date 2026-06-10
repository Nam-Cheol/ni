package initui

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var ansiPattern = regexp.MustCompile(`\x1b\[[0-9;?]*[ -/]*[@-~]`)

func TestModelInitialStateUsesAltScreen(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	view := m.View()
	if !view.AltScreen {
		t.Fatalf("expected alt screen view")
	}
	if m.stage != stageLanguage {
		t.Fatalf("expected language stage, got %v", m.stage)
	}
	if m.fields[0].Value != "demo" {
		t.Fatalf("expected default project name, got %q", m.fields[0].Value)
	}
}

func TestLanguageSelectionLocalizesDefaults(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = updateForTest(t, m, key(tea.KeyEnter))
	if m.stage != stageFields {
		t.Fatalf("expected fields stage after language selection, got %v", m.stage)
	}
	if m.language != languageKorean {
		t.Fatalf("expected Korean language, got %q", m.language)
	}
	if got := m.fields[4].Value; got != "Plan이 locked 되기 전에는 downstream work를 실행하지 않는다." {
		t.Fatalf("expected Korean safety default, got %q", got)
	}
	if !strings.Contains(m.renderShell(), "planning workspace") {
		t.Fatalf("expected Korean labels in render")
	}
}

func TestUpdateHandlesUpDownAndLeftRight(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)

	m = updateForTest(t, m, key(tea.KeyDown))
	if m.cursor != 1 {
		t.Fatalf("expected down to move cursor to 1, got %d", m.cursor)
	}
	m = updateForTest(t, m, key(tea.KeyRight))
	if m.cursor != 2 {
		t.Fatalf("expected right to move cursor to 2, got %d", m.cursor)
	}
	m = updateForTest(t, m, key(tea.KeyUp))
	if m.cursor != 1 {
		t.Fatalf("expected up to move cursor to 1, got %d", m.cursor)
	}
	m = updateForTest(t, m, key(tea.KeyLeft))
	if m.cursor != 0 {
		t.Fatalf("expected left to move cursor to 0, got %d", m.cursor)
	}
}

func TestUpdateHandlesEnterEscAndCtrlC(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m = updateForTest(t, m, key(tea.KeyEnter))
	if m.cursor != 1 {
		t.Fatalf("expected enter to continue, got cursor %d", m.cursor)
	}
	m = updateForTest(t, m, key(tea.KeyEsc))
	if m.cursor != 0 || m.canceled {
		t.Fatalf("expected esc to go back without canceling, got cursor=%d canceled=%v", m.cursor, m.canceled)
	}
	m = updateForTest(t, m, ctrlKey('c'))
	if !m.canceled {
		t.Fatalf("expected ctrl+c to cancel")
	}
}

func TestFieldsAcceptQAsText(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.cursor = 1

	m = updateForTest(t, m, textKey("q"))
	m = updateForTest(t, m, textKey("d"))
	if got := m.fields[1].Value; got != "qd" {
		t.Fatalf("expected q and d to be typed into field, got %q", got)
	}
	if m.canceled {
		t.Fatalf("did not expect q to cancel while editing a field")
	}
	m = updateForTest(t, m, ctrlKey('u'))
	if got := m.fields[1].Value; got != "" {
		t.Fatalf("expected ctrl+u to clear field, got %q", got)
	}
}

func TestConfirmPathReturnsIntent(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.fields[1].Value = "Compile intent before implementation."
	m.fields[2].Value = "Planning team"
	m.fields[3].Value = "Prepare bounded handoff."
	m.fields[5].Value = "Plan is reviewable."
	m.cursor = len(m.fields) - 1

	m = updateForTest(t, m, key(tea.KeyEnter))
	if m.stage != stageConfirm {
		t.Fatalf("expected confirm stage, got %v", m.stage)
	}
	m = updateForTest(t, m, key(tea.KeyEnter))
	result := m.Result()
	if !result.Confirmed || result.Canceled {
		t.Fatalf("expected confirmed result, got %#v", result)
	}
	if result.Intent.ProjectGoal != "Compile intent before implementation." {
		t.Fatalf("expected intent goal, got %#v", result.Intent)
	}
}

func TestCancelPathWritesNothingSignal(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.cursor = len(m.fields) - 1
	m = updateForTest(t, m, key(tea.KeyEnter))
	m = updateForTest(t, m, key(tea.KeyDown))
	m = updateForTest(t, m, key(tea.KeyEnter))
	if result := m.Result(); !result.Canceled || result.Confirmed {
		t.Fatalf("expected canceled result, got %#v", result)
	}
}

func TestReviewShowsWritePlanNextCommandsAndUpdateGuidance(t *testing.T) {
	m := NewModel(Config{Dir: ".", CommandName: "namba-intent", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm

	writePlan := m.renderWritePlan()
	for _, want := range []string{
		"docs/plan/**",
		".ni/contract.json",
		"namba-intent status --proof --next-questions",
		"namba-intent version",
		"never updates automatically",
	} {
		if !strings.Contains(writePlan, want) {
			t.Fatalf("expected write plan to contain %q\n%s", want, writePlan)
		}
	}

	view := m.renderShell()
	for _, want := range []string{"assistant", "choose one", "Write initial intent artifacts", "draft only"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected review view to contain %q\n%s", want, view)
		}
	}
	for _, notWant := range defaultDashboardCopyBanned() {
		if strings.Contains(view, notWant) {
			t.Fatalf("review view exposed dashboard copy %q\n%s", notWant, view)
		}
	}
}

func TestWindowSizeConstrainsRenderedHeight(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.width = 40
	m.height = 12
	m.stage = stageConfirm

	view := m.renderShell()
	if got := len(splitLines(view)); got > 12 {
		t.Fatalf("expected render to fit height, got %d lines\n%s", got, view)
	}
	for _, want := range []string{"Write initial intent artifacts?", "▸ ■ Write initial intent artifacts", "draft only", "↑↓"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected compact view to show %q\n%s", want, view)
		}
	}
}

func TestViewportScrollsWithPageKeys(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 60, Height: 18})

	m = updateForTest(t, m, key(tea.KeyPgDown))
	view := m.renderShell()
	if got := len(splitLines(view)); got > 18 {
		t.Fatalf("expected scrolled render to fit height, got %d lines\n%s", got, view)
	}
	if !strings.Contains(view, "Write initial intent artifacts") {
		t.Fatalf("expected page key render to preserve active composer\n%s", view)
	}
}

func TestTinyModeShowsCoreControlsWithoutScrolling(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 40, Height: 12})

	view := m.renderShell()
	for _, want := range []string{
		"Write initial intent artifacts?",
		"▸ ■ Write initial intent artifacts",
		"draft only",
		"↑↓",
	} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected tiny mode to show %q\n%s", want, view)
		}
	}
	for _, notWant := range append(defaultDashboardCopyBanned(), "details") {
		if strings.Contains(view, notWant) {
			t.Fatalf("expected tiny mode to fold %q\n%s", notWant, view)
		}
	}
}

func TestResponsiveSizesKeepViewportAndBars(t *testing.T) {
	for _, tc := range []struct {
		size tea.WindowSizeMsg
		mode layoutMode
	}{
		{tea.WindowSizeMsg{Width: 120, Height: 40}, layoutWide},
		{tea.WindowSizeMsg{Width: 80, Height: 24}, layoutMedium},
		{tea.WindowSizeMsg{Width: 60, Height: 18}, layoutNarrow},
		{tea.WindowSizeMsg{Width: 40, Height: 12}, layoutTiny},
	} {
		t.Run(fmt.Sprintf("%dx%d", tc.size.Width, tc.size.Height), func(t *testing.T) {
			m := NewModel(Config{Dir: ".", DefaultName: "demo"})
			m = selectEnglishForTest(t, m)
			m.stage = stageConfirm
			m = updateForTest(t, m, tc.size)

			view := m.renderShell()
			if m.layout().mode != tc.mode {
				t.Fatalf("expected layout mode %s, got %s", tc.mode, m.layout().mode)
			}
			if got := len(splitLines(view)); got > tc.size.Height {
				t.Fatalf("expected render to fit %d rows, got %d\n%s", tc.size.Height, got, view)
			}
			for i, line := range splitLines(view) {
				if got := lipgloss.Width(line); got > tc.size.Width {
					t.Fatalf("line %d exceeds width %d with %d cells\n%s", i+1, tc.size.Width, got, view)
				}
			}
			wants := []string{"Write initial intent artifacts", "draft only"}
			if tc.mode != layoutTiny {
				wants = append(wants, "assistant", "choose one")
			}
			for _, want := range wants {
				if !strings.Contains(view, want) {
					t.Fatalf("expected %dx%d view to contain %q\n%s", tc.size.Width, tc.size.Height, want, view)
				}
			}
			for _, notWant := range defaultDashboardCopyBanned() {
				if strings.Contains(view, notWant) {
					t.Fatalf("expected %dx%d view to hide %q\n%s", tc.size.Width, tc.size.Height, notWant, view)
				}
			}
		})
	}
}

func TestChatShellGeometryContracts(t *testing.T) {
	m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
	m.animationOn = false
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 120, Height: 40})
	layout := m.layout()
	if layout.shellWidth != 104 || layout.shellLeft != 8 {
		t.Fatalf("expected centered 104-col shell at x=8, got width=%d left=%d", layout.shellWidth, layout.shellLeft)
	}

	rendered := normalizePlainRender(m.renderShell())
	lines := splitLines(rendered)
	header := lineContaining(t, lines, "Namba Intent")
	if got := leadingCells(header); got != layout.shellLeft {
		t.Fatalf("expected header to start at shellLeft=%d, got %d\n%s", layout.shellLeft, got, rendered)
	}
	if end := visualEnd(header); end > layout.shellLeft+layout.shellWidth {
		t.Fatalf("expected header to end within shell, end=%d shellEnd=%d\n%s", end, layout.shellLeft+layout.shellWidth, rendered)
	}

	assistantBorder := nthLineContaining(t, lines, "╭", 1)
	composerBorder := nthLineContaining(t, lines, "╭", 2)
	assertShellBox(t, "assistant", assistantBorder, layout)
	assertShellBox(t, "composer", composerBorder, layout)
	if leadingCells(assistantBorder) != leadingCells(composerBorder) {
		t.Fatalf("assistant and composer left differ: %d vs %d\n%s", leadingCells(assistantBorder), leadingCells(composerBorder), rendered)
	}
	if visualEnd(assistantBorder)-leadingCells(assistantBorder) != visualEnd(composerBorder)-leadingCells(composerBorder) {
		t.Fatalf("assistant and composer width differ\nassistant=%q\ncomposer=%q", assistantBorder, composerBorder)
	}

	m = selectEnglishForTest(t, m)
	rendered = normalizePlainRender(m.renderShell())
	lines = splitLines(rendered)
	userLine := lineContaining(t, lines, "you ·")
	if leadingCells(userLine) < layout.shellLeft || visualEnd(userLine) > layout.shellLeft+layout.shellWidth {
		t.Fatalf("expected user summary inside shell x=%d..%d, got x=%d..%d\n%s", layout.shellLeft, layout.shellLeft+layout.shellWidth, leadingCells(userLine), visualEnd(userLine), rendered)
	}
	mascotLine := lineContaining(t, lines, "ni> assistant")
	if leadingCells(mascotLine) != layout.shellLeft {
		t.Fatalf("expected assistant mascot/title inside shell, got x=%d shellLeft=%d\n%s", leadingCells(mascotLine), layout.shellLeft, rendered)
	}

	m = updateForTest(t, m, key(tea.KeyTab))
	open := normalizePlainRender(m.renderShell())
	if strings.Contains(rendered, ".ni/contract.json") {
		t.Fatalf("details drawer leaked while closed\n%s", rendered)
	}
	if !strings.Contains(open, ".ni/contrac") {
		t.Fatalf("details drawer did not open inside shell\n%s", open)
	}
}

func TestPlainRenderSnapshotsCoverResponsiveBreakpoints(t *testing.T) {
	for _, tc := range []struct {
		name string
		size tea.WindowSizeMsg
	}{
		{"120x40", tea.WindowSizeMsg{Width: 120, Height: 40}},
		{"80x24", tea.WindowSizeMsg{Width: 80, Height: 24}},
		{"60x18", tea.WindowSizeMsg{Width: 60, Height: 18}},
		{"40x12", tea.WindowSizeMsg{Width: 40, Height: 12}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
			m.animationOn = false
			m = updateForTest(t, m, tc.size)
			rendered := normalizePlainRender(m.renderShell())
			wants := []string{"Namba Intent", "Korean", "초안", "Enter"}
			if tc.size.Width >= 60 && tc.size.Height >= 16 {
				wants = append(wants, "assistant", "choose one")
			}
			for _, want := range wants {
				if !strings.Contains(rendered, want) {
					t.Fatalf("snapshot precondition missing %q\n%s", want, rendered)
				}
			}
			for _, notWant := range defaultDashboardCopyBanned() {
				if strings.Contains(rendered, notWant) {
					t.Fatalf("snapshot exposed dashboard copy %q\n%s", notWant, rendered)
				}
			}
			path := filepath.Join("testdata", "render", "init_language_"+tc.name+".golden")
			if os.Getenv("UPDATE_RENDER_SNAPSHOTS") == "1" {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatalf("creating snapshot dir: %v", err)
				}
				if err := os.WriteFile(path, []byte(rendered), 0o644); err != nil {
					t.Fatalf("writing snapshot: %v", err)
				}
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("reading snapshot %s: %v", path, err)
			}
			if rendered != string(want) {
				t.Fatalf("render snapshot mismatch for %s\nwant:\n%s\ngot:\n%s", tc.name, string(want), rendered)
			}
		})
	}
}

func TestRenderDoesNotExposeInternalDebugCopy(t *testing.T) {
	for _, tc := range []struct {
		name string
		size tea.WindowSizeMsg
	}{
		{"120x40", tea.WindowSizeMsg{Width: 120, Height: 40}},
		{"80x24", tea.WindowSizeMsg{Width: 80, Height: 24}},
		{"60x18", tea.WindowSizeMsg{Width: 60, Height: 18}},
		{"40x12", tea.WindowSizeMsg{Width: 40, Height: 12}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
			m = updateForTest(t, m, tc.size)
			view := normalizePlainRender(m.renderShell())
			for _, banned := range append(defaultDashboardCopyBanned(), "tick", "frame", "Motion:", "Layout:") {
				if strings.Contains(view, banned) {
					t.Fatalf("render exposed internal copy %q\n%s", banned, view)
				}
			}
		})
	}
}

func TestNoColorRenderKeepsSemanticSelectionAndStatus(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 40, Height: 12})

	view := normalizePlainRender(m.renderShell())
	for _, want := range []string{"▸ ■ Write initial intent artifacts", "  □ Cancel; write nothing", "draft only", "↑↓"} {
		if !strings.Contains(view, want) {
			t.Fatalf("NO_COLOR render lost semantic marker %q\n%s", want, view)
		}
	}
}

func TestCancellationRenderCannotLookSuccessful(t *testing.T) {
	m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, key(tea.KeyDown))
	m = updateForTest(t, m, key(tea.KeyEnter))
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 40, Height: 12})

	view := normalizePlainRender(m.renderShell())
	for _, want := range []string{"Cancelled", "No files written", "init again"} {
		if !strings.Contains(view, want) {
			t.Fatalf("cancellation render missing %q\n%s", want, view)
		}
	}
	for _, banned := range []string{"captured", "Ready for status", "Initial intent artifacts"} {
		if strings.Contains(view, banned) {
			t.Fatalf("cancellation render looked successful via %q\n%s", banned, view)
		}
	}
}

func TestAnimationTickAdvancesDeterministically(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	before := m.renderShell()
	m = updateForTest(t, m, initTickMsg{})
	if m.frame != 1 {
		t.Fatalf("expected animation step to advance to 1, got %d", m.frame)
	}
	after := m.renderShell()
	if before == after {
		t.Fatalf("expected animation step to affect rendered motion")
	}
}

func TestAnimationCanBeDisabledForDeterministicSnapshots(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m.animationOn = false
	before := m.renderShell()
	m = updateForTest(t, m, initTickMsg{})
	after := m.renderShell()
	if m.frame != 0 {
		t.Fatalf("expected disabled animation to keep frame at 0, got %d", m.frame)
	}
	if before != after {
		t.Fatalf("expected disabled animation render to stay deterministic")
	}
}

func TestStageAssetsAndProgressAreFunctionalUI(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	view := m.renderShell()
	for _, want := range []string{"Namba Intent", "init · 1/10", "◇", "◆", "assistant"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected language view to contain %q\n%s", want, view)
		}
	}

	m = selectEnglishForTest(t, m)
	view = m.renderShell()
	for _, want := range []string{"assistant", "your answer", "▰", "▱"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected drafting view to contain %q\n%s", want, view)
		}
	}

	m.stage = stageConfirm
	view = m.renderShell()
	for _, want := range []string{"assistant", "■ Write initial intent artifacts"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected review view to contain %q\n%s", want, view)
		}
	}
}

func TestDetailsDrawerIsHiddenUntilRequested(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 120, Height: 40})

	closed := normalizePlainRender(m.renderShell())
	for _, notWant := range []string{"files\n", "will create · docs/plan", ".ni/contract.json"} {
		if strings.Contains(closed, notWant) {
			t.Fatalf("expected details drawer to be hidden by default via %q\n%s", notWant, closed)
		}
	}

	m = updateForTest(t, m, key(tea.KeyTab))
	open := normalizePlainRender(m.renderShell())
	for _, want := range []string{"details", "files", "CLI gates stay authoritative", ".ni/contrac"} {
		if !strings.Contains(open, want) {
			t.Fatalf("expected open details drawer to contain %q\n%s", want, open)
		}
	}
	for _, notWant := range []string{"Files Panel", "Status Panel", "Next Action Panel"} {
		if strings.Contains(open, notWant) {
			t.Fatalf("details drawer should not revive dashboard copy %q\n%s", notWant, open)
		}
	}
}

func TestWideLayoutKeepsChatFirstByDefault(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 120, Height: 40})

	view := m.renderShell()
	for _, want := range []string{"assistant", "your answer", "draft only"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected wide view to contain %q\n%s", want, view)
		}
	}
	for _, notWant := range defaultDashboardCopyBanned() {
		if strings.Contains(view, notWant) {
			t.Fatalf("expected wide default view to hide %q\n%s", notWant, view)
		}
	}
}

func TestExistingFileChoices(t *testing.T) {
	m := NewModel(Config{Dir: ".", ExistingFiles: []string{".ni/contract.json"}})
	if m.stage != stageLanguage {
		t.Fatalf("expected language stage, got %v", m.stage)
	}
	m = selectEnglishForTest(t, m)
	if m.stage != stageExisting {
		t.Fatalf("expected existing-file stage after language selection, got %v", m.stage)
	}
	m = updateForTest(t, m, key(tea.KeyDown))
	m = updateForTest(t, m, key(tea.KeyEnter))
	result := m.Result()
	if result.Choice != ExistingChoiceKeep || !result.Canceled {
		t.Fatalf("expected keep choice, got %#v", result)
	}
}

func updateForTest(t *testing.T, m Model, msg tea.Msg) Model {
	t.Helper()
	next, _ := m.Update(msg)
	model, ok := next.(Model)
	if !ok {
		t.Fatalf("expected Model, got %T", next)
	}
	return model
}

func selectEnglishForTest(t *testing.T, m Model) Model {
	t.Helper()
	m = updateForTest(t, m, key(tea.KeyDown))
	m = updateForTest(t, m, key(tea.KeyEnter))
	if m.stage != stageFields && m.stage != stageExisting {
		t.Fatalf("expected language selection to continue, got %v", m.stage)
	}
	if m.language != languageEnglish {
		t.Fatalf("expected English language, got %q", m.language)
	}
	return m
}

func key(code rune) tea.KeyPressMsg {
	return tea.KeyPressMsg(tea.Key{Code: code})
}

func textKey(value string) tea.KeyPressMsg {
	runes := []rune(value)
	return tea.KeyPressMsg(tea.Key{Code: runes[0], Text: value})
}

func ctrlKey(value rune) tea.KeyPressMsg {
	return tea.KeyPressMsg(tea.Key{Code: value, Mod: tea.ModCtrl})
}

func stripANSI(value string) string {
	return ansiPattern.ReplaceAllString(value, "")
}

func normalizePlainRender(value string) string {
	lines := strings.Split(stripANSI(value), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}
	return strings.Join(lines, "\n")
}

func lineContaining(t *testing.T, lines []string, want string) string {
	t.Helper()
	return nthLineContaining(t, lines, want, 1)
}

func nthLineContaining(t *testing.T, lines []string, want string, n int) string {
	t.Helper()
	seen := 0
	for _, line := range lines {
		if strings.Contains(line, want) {
			seen++
			if seen == n {
				return line
			}
		}
	}
	t.Fatalf("could not find line %d containing %q\n%s", n, want, strings.Join(lines, "\n"))
	return ""
}

func leadingCells(line string) int {
	return lipgloss.Width(line[:len(line)-len(strings.TrimLeft(line, " "))])
}

func visualEnd(line string) int {
	return lipgloss.Width(strings.TrimRight(line, " "))
}

func assertShellBox(t *testing.T, name string, line string, layout layoutSpec) {
	t.Helper()
	left := leadingCells(line)
	width := visualEnd(line) - left
	if left != layout.shellLeft {
		t.Fatalf("%s box starts at %d, want shellLeft=%d: %q", name, left, layout.shellLeft, line)
	}
	if width != layout.shellWidth {
		t.Fatalf("%s box width=%d, want shellWidth=%d: %q", name, width, layout.shellWidth, line)
	}
}

func defaultDashboardCopyBanned() []string {
	return []string{
		"Main Panel",
		"Status Panel",
		"Files Panel",
		"Next Action Panel",
		"Detected:",
		"Authority: CLI status/end/run",
		"status/end/run",
		"contract draft:",
		"docs planned:",
		"tick",
		"frame",
		"Motion",
		"Layout:",
	}
}
