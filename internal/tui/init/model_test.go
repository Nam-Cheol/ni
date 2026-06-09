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
	if !strings.Contains(view, "No downstream work runs") {
		t.Fatalf("expected footer to remain visible\n%s", view)
	}
}

func TestViewportScrollsWithPageKeys(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 40, Height: 12})

	before := m.bodyViewport.YOffset()
	m = updateForTest(t, m, key(tea.KeyPgDown))
	if after := m.bodyViewport.YOffset(); after <= before {
		t.Fatalf("expected PgDown to advance viewport offset, before=%d after=%d", before, after)
	}
	view := m.renderShell()
	if got := len(splitLines(view)); got > 12 {
		t.Fatalf("expected scrolled render to fit height, got %d lines\n%s", got, view)
	}
}

func TestViewportScrollReachesLowerPanelsInCompactMode(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.stage = stageConfirm
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 40, Height: 12})

	top := m.renderShell()
	if strings.Contains(top, "Next Action Panel") {
		t.Fatalf("expected compact top viewport to hide lower panels before scrolling\n%s", top)
	}

	m = updateForTest(t, m, key(tea.KeyEnd))
	bottom := m.renderShell()
	if !strings.Contains(bottom, "namba-intent version") {
		t.Fatalf("expected End to reveal lower next-action content in compact viewport\n%s", bottom)
	}
	if !strings.Contains(bottom, "scroll:") || !strings.Contains(bottom, "No downstream work runs") {
		t.Fatalf("expected help/footer to remain visible after scrolling\n%s", bottom)
	}
}

func TestResponsiveSizesKeepViewportAndBars(t *testing.T) {
	for _, size := range []tea.WindowSizeMsg{
		{Width: 120, Height: 40},
		{Width: 80, Height: 24},
		{Width: 60, Height: 18},
		{Width: 40, Height: 12},
	} {
		t.Run(fmt.Sprintf("%dx%d", size.Width, size.Height), func(t *testing.T) {
			m := NewModel(Config{Dir: ".", DefaultName: "demo"})
			m = selectEnglishForTest(t, m)
			m.stage = stageConfirm
			m = updateForTest(t, m, size)

			view := m.renderShell()
			if got := len(splitLines(view)); got > size.Height {
				t.Fatalf("expected render to fit %d rows, got %d\n%s", size.Height, got, view)
			}
			for i, line := range splitLines(view) {
				if got := lipgloss.Width(line); got > size.Width {
					t.Fatalf("line %d exceeds width %d with %d cells\n%s", i+1, size.Width, got, view)
				}
			}
			for _, want := range []string{"Main Panel", "scroll:", "No downstream work runs"} {
				if !strings.Contains(view, want) {
					t.Fatalf("expected %dx%d view to contain %q\n%s", size.Width, size.Height, want, view)
				}
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
