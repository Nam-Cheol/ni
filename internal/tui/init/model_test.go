package initui

import (
	"fmt"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

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
	if !strings.Contains(m.renderShell(), "프로젝트 이름") {
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
	if got := m.fields[1].Value; got != "q" {
		t.Fatalf("expected q to be typed into field, got %q", got)
	}
	if m.canceled {
		t.Fatalf("did not expect q to cancel while editing a field")
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
	for _, want := range []string{"Main Panel", "Status Panel", "Files Panel", "scroll:"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected review view to contain %q\n%s", want, view)
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
	if !strings.Contains(view, "scroll:") {
		t.Fatalf("expected compact view to expose scroll state\n%s", view)
	}
	if !strings.Contains(view, "no downstream work") {
		t.Fatalf("expected footer to remain visible\n%s", view)
	}
}

func TestViewportScrollsWithPageKeys(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 60, Height: 18})

	before := m.bodyViewport.YOffset()
	m = updateForTest(t, m, key(tea.KeyPgDown))
	if after := m.bodyViewport.YOffset(); after <= before {
		t.Fatalf("expected PgDown to advance viewport offset, before=%d after=%d", before, after)
	}
	view := m.renderShell()
	if got := len(splitLines(view)); got > 18 {
		t.Fatalf("expected scrolled render to fit height, got %d lines\n%s", got, view)
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
		"Status: ready / review",
		"scroll:",
		"no downstream work",
	} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected tiny mode to show %q\n%s", want, view)
		}
	}
	for _, notWant := range []string{"Files Panel", "Next Action Panel"} {
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
			footerText := "No downstream work runs"
			if tc.mode == layoutNarrow || tc.mode == layoutTiny {
				footerText = "no downstream work"
			}
			for _, want := range []string{"Main Panel", "scroll:", footerText} {
				if !strings.Contains(view, want) {
					t.Fatalf("expected %dx%d view to contain %q\n%s", tc.size.Width, tc.size.Height, want, view)
				}
			}
		})
	}
}

func TestAnimationTickAdvancesDeterministically(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	before := m.renderShell()
	m = updateForTest(t, m, initTickMsg{})
	if m.frame != 1 {
		t.Fatalf("expected animation frame to advance to 1, got %d", m.frame)
	}
	after := m.renderShell()
	if before == after {
		t.Fatalf("expected tick frame to affect rendered motion")
	}
}

func TestStageAssetsAndProgressAreFunctionalUI(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	view := m.renderShell()
	for _, want := range []string{"language gate", "Step 1/10", "◇", "◆"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected language view to contain %q\n%s", want, view)
		}
	}

	m = selectEnglishForTest(t, m)
	view = m.renderShell()
	for _, want := range []string{"drafting grid", "Capture slot", "▰", "▱"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected drafting view to contain %q\n%s", want, view)
		}
	}

	m.stage = stageConfirm
	view = m.renderShell()
	for _, want := range []string{"review scan", "checksum waits", "■ Write initial intent artifacts"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected review view to contain %q\n%s", want, view)
		}
	}
}

func TestFilesPanelCollapsesByLayoutMode(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)

	for _, tc := range []struct {
		mode     layoutMode
		width    int
		maxItems int
		wantMore bool
	}{
		{layoutWide, 120, 5, true},
		{layoutMedium, 80, 4, true},
		{layoutNarrow, 60, 3, true},
		{layoutTiny, 40, 1, true},
	} {
		t.Run(string(tc.mode), func(t *testing.T) {
			layout := m.layout()
			layout.mode = tc.mode
			layout.width = tc.width
			layout.panelWidth = tc.width
			items, hidden := m.filesForPanel(layout)
			if len(items) > tc.maxItems {
				t.Fatalf("expected at most %d file items, got %d", tc.maxItems, len(items))
			}
			if tc.wantMore && hidden == 0 {
				t.Fatalf("expected collapsed file summary for %s layout", tc.mode)
			}
			panel := m.renderFilesPanel(layout)
			if !strings.Contains(panel, "blueprint") || !strings.Contains(panel, "more planned") {
				t.Fatalf("expected blueprint collapsed files panel\n%s", panel)
			}
		})
	}
}

func TestWideLayoutShowsSidePanels(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 120, Height: 40})

	view := m.renderShell()
	for _, want := range []string{"Main Panel", "Status Panel", "Files Panel", "Next Action Panel"} {
		if !strings.Contains(view, want) {
			t.Fatalf("expected wide view to contain %q\n%s", want, view)
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
