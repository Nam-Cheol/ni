package initui

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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

func TestFieldsTreatQAndDAsTextUntilControlModified(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.cursor = 1

	m = updateForTest(t, m, textKey("d"))
	m = updateForTest(t, m, textKey("q"))
	if got := m.fields[1].Value; got != "dq" {
		t.Fatalf("expected d and q to edit field as text, got %q", got)
	}
	if m.detailsOpen || m.canceled {
		t.Fatalf("plain d/q should not run commands, details=%v canceled=%v", m.detailsOpen, m.canceled)
	}

	m = updateForTest(t, m, ctrlKey('d'))
	if got := m.fields[1].Value; got != "dq" {
		t.Fatalf("expected ctrl+d to leave field text unchanged, got %q", got)
	}
	if !m.detailsOpen {
		t.Fatalf("expected ctrl+d to toggle details while editing a field")
	}

	m = updateForTest(t, m, ctrlKey('q'))
	if !m.canceled || m.stage != stageDone {
		t.Fatalf("expected ctrl+q to quit while editing a field, got stage=%v canceled=%v", m.stage, m.canceled)
	}
}

func TestFieldsAcceptRegularTextAndCtrlU(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.cursor = 1

	m = updateForTest(t, m, textKey("a"))
	m = updateForTest(t, m, textKey("b"))
	if got := m.fields[1].Value; got != "ab" {
		t.Fatalf("expected regular text to be typed into field, got %q", got)
	}
	m = updateForTest(t, m, ctrlKey('u'))
	if got := m.fields[1].Value; got != "" {
		t.Fatalf("expected ctrl+u to clear field, got %q", got)
	}
}

func TestFieldsAcceptBracketedPasteAsAnswerText(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = selectEnglishForTest(t, m)
	m.cursor = 1

	m = updateForTest(t, m, tea.PasteMsg{Content: "첫 줄\r\nsecond line\rthird"})
	if got := m.fields[1].Value; got != "첫 줄\nsecond line\nthird" {
		t.Fatalf("expected pasted content to append as normalized answer text, got %q", got)
	}
	if m.cursor != 1 || m.stage != stageFields {
		t.Fatalf("paste should not advance the answer step, cursor=%d stage=%v", m.cursor, m.stage)
	}

	m = updateForTest(t, m, textKey("!"))
	if got := m.fields[1].Value; got != "첫 줄\nsecond line\nthird!" {
		t.Fatalf("expected typed text after paste to append, got %q", got)
	}
}

func TestPasteOutsideAnswerComposerDoesNotChangeChoices(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})

	m = updateForTest(t, m, tea.PasteMsg{Content: "English"})
	if m.stage != stageLanguage || m.languageCursor != 0 {
		t.Fatalf("paste should not mutate language choice, stage=%v cursor=%d", m.stage, m.languageCursor)
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
			wants := []string{"draft only"}
			if tc.size.Width == 80 {
				wants = append(wants, "Write initial intent artifa")
			} else {
				wants = append(wants, "Write initial intent artifacts")
			}
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
	if layout.shellWidth < layout.width-2 || layout.shellLeft > 1 {
		t.Fatalf("expected wide shell to use terminal canvas, got width=%d left=%d terminal=%d", layout.shellWidth, layout.shellLeft, layout.width)
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
	chatLeft := layout.shellLeft + layout.habitatWidth + layout.habitatGap
	assistantBubbleLeft := borderStartCells(assistantBorder)
	if layout.habitatWidth < 38 || layout.habitatWidth > 42 {
		t.Fatalf("expected wide diorama rail width 38..42, got %d", layout.habitatWidth)
	}
	if assistantBubbleLeft != chatLeft {
		t.Fatalf("assistant bubble starts at %d, want chat column x=%d\n%s", assistantBubbleLeft, chatLeft, rendered)
	}
	if layout.chatWidth*100 < layout.shellWidth*60 {
		t.Fatalf("expected chat column to own most shell width, got chat=%d shell=%d", layout.chatWidth, layout.shellWidth)
	}
	composerLeft := borderStartCells(composerBorder)
	if assistantBubbleLeft != composerLeft {
		t.Fatalf("assistant bubble and composer left differ: %d vs %d\n%s", assistantBubbleLeft, composerLeft, rendered)
	}
	if borderWidth(assistantBorder) != borderWidth(composerBorder) {
		t.Fatalf("assistant and composer width differ\nassistant=%q\ncomposer=%q", assistantBorder, composerBorder)
	}
	if layout.habitatWidth == 0 || layout.habitatGap == 0 {
		t.Fatalf("expected wide layout to reserve habitat rail, got habitat=%d gap=%d", layout.habitatWidth, layout.habitatGap)
	}
	assistantBottom := lineIndexContaining(t, lines, "╰")
	composerTop := lineIndexContainingFrom(t, lines, "╭", assistantBottom)
	if gap := blankLinesBetween(lines, assistantBottom, composerTop); gap > 2 {
		t.Fatalf("expected assistant/composer gap <= 2, got %d\n%s", gap, rendered)
	}
	if maxBlankRun(lines) > 3 {
		t.Fatalf("expected no 4-line blank run in default render\n%s", rendered)
	}
	for _, want := range []string{"▄████▄  ▐██▌  ▄████▄", "██████▄ ▐██▌ ▄██████", "▀██████▄██▄██████▀", "██████████████", "█████████▓▓■▓▓██▓▓■▓▓█████████", "██████████▓▓▓▓▓██▓▓▓▓▓██████████", "██████████■■■■■■■■██████████", "███████▒▒▒▒▒▒▒▒▒▒███████"} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("expected creature diorama to contain %q\n%s", want, rendered)
		}
	}
	for _, notWant := range noisyDioramaCopyBanned() {
		if strings.Contains(rendered, notWant) {
			t.Fatalf("default diorama exposed noisy pattern %q\n%s", notWant, rendered)
		}
	}
	if decorative := countDioramaDecorativeLines(lines, layout); decorative > 4 {
		t.Fatalf("expected sparse diorama accents, got %d decorative lines\n%s", decorative, rendered)
	}
	if countMutedBackgroundGlyphs(lines, layout) > countMascotGlyphs(lines, layout) {
		t.Fatalf("background glyphs should not outnumber mascot glyphs\n%s", rendered)
	}

	m = selectEnglishForTest(t, m)
	rendered = normalizePlainRender(m.renderShell())
	lines = splitLines(rendered)
	userLine := lineContaining(t, lines, "you ·")
	if leadingCells(userLine) < layout.shellLeft || visualEnd(userLine) > layout.shellLeft+layout.shellWidth {
		t.Fatalf("expected user summary inside shell x=%d..%d, got x=%d..%d\n%s", layout.shellLeft, layout.shellLeft+layout.shellWidth, leadingCells(userLine), visualEnd(userLine), rendered)
	}
	mascotLine := lineContaining(t, lines, "██████████████")
	mascotArtSegment := habitatSegment(mascotLine)
	if leadingCells(mascotArtSegment) < layout.shellLeft || visualEnd(mascotArtSegment) > layout.shellLeft+layout.habitatWidth {
		t.Fatalf("expected mascot inside diorama rail x=%d..%d, got x=%d..%d\n%s", layout.shellLeft, layout.shellLeft+layout.habitatWidth, leadingCells(mascotArtSegment), visualEnd(mascotArtSegment), rendered)
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

func TestFieldComposerKeepsGuidanceInAssistantAndTypedInputClean(t *testing.T) {
	m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
	m.animationOn = false
	m = updateForTest(t, m, key(tea.KeyEnter))
	m = updateForTest(t, m, key(tea.KeyEnter))
	m = updateForTest(t, m, textKey("새"))
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 120, Height: 40})

	rendered := normalizePlainRender(m.renderShell())
	if !strings.Contains(rendered, "좋은 답변 형태:") {
		t.Fatalf("expected answer guidance inside assistant message\n%s", rendered)
	}
	if strings.Contains(rendered, "choose one") {
		t.Fatalf("field stage should not keep choice composer label\n%s", rendered)
	}
	composer := blockContaining(rendered, "your answer")
	for _, notWant := range []string{"좋은 답변 형태:", "여기에 입력하세요", "TODO로 남겨도", "프로젝트 목표를 한 문장으로 입력하세요"} {
		if strings.Contains(composer, notWant) {
			t.Fatalf("typed composer should only show user text, but found %q\ncomposer:\n%s\nfull:\n%s", notWant, composer, rendered)
		}
	}
	if !strings.Contains(composer, "> 새") {
		t.Fatalf("expected typed user text in composer\n%s", composer)
	}
}

func TestLanguageDownDoesNotWrapOrCorruptChoiceCanvas(t *testing.T) {
	m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
	m.animationOn = false
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 123, Height: 33})
	m = updateForTest(t, m, key(tea.KeyDown))

	rendered := normalizePlainRender(m.renderShell())
	lines := splitLines(rendered)
	if got := len(lines); got != 33 {
		t.Fatalf("expected render to use exactly 33 rows, got %d\n%s", got, rendered)
	}
	for i, line := range lines {
		if got := lipgloss.Width(line); got > 123 {
			t.Fatalf("line %d exceeds 123 cells with %d\n%s", i+1, got, rendered)
		}
	}
	for _, want := range []string{"선택됨: English", "◇ Korean", "▸ ◆ English"} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("expected down-selection render to contain %q\n%s", want, rendered)
		}
	}
	for _, notWant := range []string{"labKorean", "Korean ,", "Korean  ,", "확인하기 English", "선택됨: Korean"} {
		if strings.Contains(rendered, notWant) {
			t.Fatalf("down-selection render contains corrupted overlap %q\n%s", notWant, rendered)
		}
	}
}

func TestResponsiveMascotDensity(t *testing.T) {
	for _, tc := range []struct {
		name string
		size tea.WindowSizeMsg
		want string
	}{
		{"wide", tea.WindowSizeMsg{Width: 120, Height: 40}, "█████████▓▓■▓▓██▓▓■▓▓█████████"},
		{"medium", tea.WindowSizeMsg{Width: 80, Height: 24}, "█████████▓■▓██▓■▓█████████"},
		{"narrow", tea.WindowSizeMsg{Width: 60, Height: 18}, "ni·· assistant"},
		{"tiny", tea.WindowSizeMsg{Width: 40, Height: 12}, "ni▣"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
			m.animationOn = false
			m = updateForTest(t, m, tc.size)
			rendered := normalizePlainRender(m.renderShell())
			if !strings.Contains(rendered, tc.want) {
				t.Fatalf("expected mascot marker %q\n%s", tc.want, rendered)
			}
			if tc.size.Width >= 80 {
				layout := m.layout()
				if layout.mode == layoutWide && (layout.habitatWidth < 38 || layout.habitatWidth > 42) {
					t.Fatalf("expected wide diorama rail width 38..42, got %d\n%s", layout.habitatWidth, rendered)
				}
				if layout.mode == layoutMedium && (layout.habitatWidth < 32 || layout.habitatWidth > 34) {
					t.Fatalf("expected medium diorama rail width 32..34, got %d\n%s", layout.habitatWidth, rendered)
				}
				lines := splitLines(rendered)
				minCreatureLines := 7
				maxCreatureLines := 16
				if layout.mode == layoutMedium {
					minCreatureLines = 10
					maxCreatureLines = 10
				}
				got := countHabitatCreatureLines(lines, layout)
				if got < minCreatureLines || got > maxCreatureLines {
					t.Fatalf("expected %s creature height %d..%d, got %d\n%s", tc.name, minCreatureLines, maxCreatureLines, got, rendered)
				}
				for _, notWant := range noisyDioramaCopyBanned() {
					if strings.Contains(rendered, notWant) {
						t.Fatalf("expected sparse diorama, found noisy pattern %q\n%s", notWant, rendered)
					}
				}
				if decorative := countDioramaDecorativeLines(lines, layout); decorative > 4 {
					t.Fatalf("expected at most 4 decorative diorama lines, got %d\n%s", decorative, rendered)
				}
				if countMutedBackgroundGlyphs(lines, layout) > countMascotGlyphs(lines, layout) {
					t.Fatalf("background glyphs should stay quieter than mascot\n%s", rendered)
				}
				assistantBottom := lineIndexContaining(t, lines, "╰")
				composerTop := lineIndexContainingFrom(t, lines, "╭", assistantBottom)
				if gap := blankLinesBetween(lines, assistantBottom, composerTop); gap > 2 {
					t.Fatalf("expected compact assistant/composer gap, got %d\n%s", gap, rendered)
				}
			}
			if strings.Contains(rendered, ".ni/contract.json") || strings.Contains(rendered, "docs planned:") {
				t.Fatalf("default render leaked details drawer content\n%s", rendered)
			}
		})
	}
}

func TestMascotAssetsCarryStageState(t *testing.T) {
	for _, tc := range []struct {
		name  string
		setup func(Model) Model
		want  []string
	}{
		{"language", func(m Model) Model { m.stage = stageLanguage; return m }, []string{"▄██▄ ▐█▌ ▄██▄", "▀███▄██▄███▀", "█████████▓■▓██▓■▓█████████", "█████████▓▓▓██▓▓▓█████████", "█████████■■■■█████████", "██████▒▒▒▒▒▒██████"}},
		{"field", func(m Model) Model { m.stage = stageFields; return m }, []string{"▄██▄ ▐█▌ ▄██▄", "▀███▄██▄███▀", "█████████▓■▓██▓■▓█████████", "█████████■■▪■█████████", "██████▒▒▒▒▒▒██████"}},
		{"existing", func(m Model) Model { m.stage = stageExisting; return m }, []string{"▄██▄ ▐█▌ ▄██▄", "█████████▓▓■██▓▓■█████████", "██████▒▒▒▒▒▒██████"}},
		{"confirm", func(m Model) Model { m.stage = stageConfirm; return m }, []string{"▄██▄ ▐█▌ ▄██▄", "█████████▓▓■██▓▓■█████████", "█████████■■■■█████████"}},
		{"done", func(m Model) Model { m.stage = stageDone; m.confirmed = true; return m }, []string{"▄██▄ ▐█▌ ▄██▄", "█████████▓■▓██▓■▓█████████", "█████▒▒▒▒▒▒▒██████"}},
		{"cancelled", func(m Model) Model { m.stage = stageDone; m.canceled = true; return m }, []string{"░▒░ ▐▒▌ ░▒░", "▒▒▓▓▓▒▒▓▓▓▒▒", "▒▒■■■■▒▒"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
			m.animationOn = false
			m = tc.setup(m)
			lines := m.stageAsset(m.layout()).Lines
			if len(lines) != 10 && !m.canceled {
				t.Fatalf("expected fixed 10-line compact creature mascot, got %d lines\n%s", len(lines), strings.Join(lines, "\n"))
			}
			for i, line := range lines {
				if got := lipgloss.Width(line); got != creatureWidth {
					t.Fatalf("expected %s mascot line %d width=%d, got %d: %q", tc.name, i+1, creatureWidth, got, line)
				}
			}
			asset := strings.Join(lines, "\n")
			for _, want := range tc.want {
				if !strings.Contains(asset, want) {
					t.Fatalf("expected %s mascot to contain %q\n%s", tc.name, want, asset)
				}
			}
			for _, forbidden := range []string{"o.o", "ask", "doc", "____", "/|", "(\\"} {
				if strings.Contains(asset, forbidden) {
					t.Fatalf("expected creature mascot, not old line robot marker %q\n%s", forbidden, asset)
				}
			}
			if strings.Contains(asset, "…") {
				t.Fatalf("mascot should never be truncated with ellipsis\n%s", asset)
			}
		})
	}
}

func TestCreatureAssetSizesStaySimple(t *testing.T) {
	full := creatureAsset(creatureAsking, assetFull)
	if got := len(full.Lines); got != 16 {
		t.Fatalf("expected full mascot height 16, got %d\n%s", got, strings.Join(full.Lines, "\n"))
	}
	compact := creatureAsset(creatureAsking, assetCompact)
	if got := len(compact.Lines); got != 10 {
		t.Fatalf("expected compact mascot height 10, got %d\n%s", got, strings.Join(compact.Lines, "\n"))
	}
	tiny := creatureAsset(creatureAsking, assetTiny)
	if tiny.Compact != "ni▣" || len(tiny.Lines) != 1 {
		t.Fatalf("expected tiny mascot to collapse to ni glyph, got %#v", tiny)
	}
	for _, asset := range []stageAsset{full, compact} {
		text := strings.Join(asset.Lines, "\n")
		for _, forbidden := range []string{"░▀░", "▐░▌", "░░▒", "▄▒▄", "▀▒░░▒▀"} {
			if strings.Contains(text, forbidden) {
				t.Fatalf("mascot asset should not look like noisy forest/totem pattern %q\n%s", forbidden, text)
			}
		}
	}
}

func TestCreatureBodyUsesEllipseWidthProfile(t *testing.T) {
	for _, tc := range []struct {
		name       string
		size       assetSize
		wantWidths []int
	}{
		{name: "full", size: assetFull, wantWidths: []int{14, 20, 24, 28, 30, 32, 32, 30, 28, 24, 20, 14}},
		{name: "compact", size: assetCompact, wantWidths: []int{10, 16, 22, 26, 26, 22, 18, 12}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := creatureBodyWidths(creatureAsset(creatureAsking, tc.size), tc.size)
			if !reflect.DeepEqual(got, tc.wantWidths) {
				t.Fatalf("expected %s ellipse profile %v, got %v\n%s", tc.name, tc.wantWidths, got, strings.Join(creatureAsset(creatureAsking, tc.size).Lines, "\n"))
			}
		})
	}
}

func TestCreatureMiddleRowsAreWidest(t *testing.T) {
	for _, size := range []assetSize{assetFull, assetCompact} {
		t.Run(string(size), func(t *testing.T) {
			widths := creatureBodyWidths(creatureAsset(creatureAsking, size), size)
			maxWidth := 0
			maxCount := 0
			for _, width := range widths {
				if width > maxWidth {
					maxWidth = width
					maxCount = 1
				} else if width == maxWidth {
					maxCount++
				}
			}
			if maxCount > 2 {
				t.Fatalf("expected no more than two widest rows, got %d in %v", maxCount, widths)
			}
			for i, width := range widths {
				if width == maxWidth && (i < len(widths)/3 || i > len(widths)*2/3) {
					t.Fatalf("widest row should stay in middle area, index=%d widths=%v", i, widths)
				}
			}
		})
	}
}

func TestCreatureTopAndBottomTaper(t *testing.T) {
	for _, size := range []assetSize{assetFull, assetCompact} {
		t.Run(string(size), func(t *testing.T) {
			widths := creatureBodyWidths(creatureAsset(creatureAsking, size), size)
			mid := len(widths) / 2
			if widths[0] >= widths[mid] || widths[len(widths)-1] >= widths[mid] {
				t.Fatalf("expected top/bottom to taper below middle, got %v", widths)
			}
			for i := 1; i <= mid; i++ {
				if widths[i] < widths[i-1] {
					t.Fatalf("expected width to grow toward middle at index %d, got %v", i, widths)
				}
			}
			for i := mid + 1; i < len(widths); i++ {
				if widths[i] > widths[i-1] {
					t.Fatalf("expected width to shrink after middle at index %d, got %v", i, widths)
				}
			}
		})
	}
}

func TestSunglassesAreInternalNotOuterSilhouette(t *testing.T) {
	for _, size := range []assetSize{assetFull, assetCompact} {
		t.Run(string(size), func(t *testing.T) {
			rows := creatureBodyRows(creatureAsset(creatureAsking, size), size)
			found := false
			for _, row := range rows {
				row = strings.TrimSpace(row)
				if !strings.Contains(row, "▓") {
					continue
				}
				found = true
				runes := []rune(row)
				firstBlack := runeIndex(row, '▓')
				lastBlack := runeLastIndex(row, '▓')
				if firstBlack < 2 || lastBlack > len(runes)-3 {
					t.Fatalf("sunglasses must have orange cheek pixels on both sides: %q", row)
				}
				center := string(runes[len(runes)/2-1 : len(runes)/2+1])
				if center != "██" {
					t.Fatalf("sunglasses should have orange bridge/gap at center, got %q in %q", center, row)
				}
				if strings.Count(row, "▓")*2 >= lipgloss.Width(row) {
					t.Fatalf("sunglasses should not dominate the body row: %q", row)
				}
			}
			if !found {
				t.Fatalf("expected sunglasses rows in creature asset\n%s", strings.Join(creatureAsset(creatureAsking, size).Lines, "\n"))
			}
		})
	}
}

func TestCreatureDoesNotLookLikeHorizontalStripes(t *testing.T) {
	for _, size := range []assetSize{assetFull, assetCompact} {
		t.Run(string(size), func(t *testing.T) {
			widths := creatureBodyWidths(creatureAsset(creatureAsking, size), size)
			unique := map[int]bool{}
			runWidth := -1
			runLength := 0
			for _, width := range widths {
				unique[width] = true
				if width == runWidth {
					runLength++
				} else {
					runWidth = width
					runLength = 1
				}
				if runLength > 2 {
					t.Fatalf("too many repeated body widths; looks striped/rectangular: %v", widths)
				}
			}
			if len(unique) < 5 {
				t.Fatalf("expected varied ellipse widths, got %v", widths)
			}
		})
	}
}

func TestFullHeightChatCanvasContracts(t *testing.T) {
	for _, tc := range []struct {
		name         string
		size         tea.WindowSizeMsg
		minOccupancy int
	}{
		{"120x40", tea.WindowSizeMsg{Width: 120, Height: 40}, 50},
		{"80x24", tea.WindowSizeMsg{Width: 80, Height: 24}, 60},
		{"60x18", tea.WindowSizeMsg{Width: 60, Height: 18}, 60},
		{"40x12", tea.WindowSizeMsg{Width: 40, Height: 12}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := NewModel(Config{Dir: "/tmp/demo", DefaultName: "demo"})
			m.animationOn = false
			m = updateForTest(t, m, tc.size)
			rendered := normalizePlainRender(m.renderShell())
			lines := splitLines(rendered)
			if got := len(lines); got != tc.size.Height {
				t.Fatalf("expected render to use full terminal height %d, got %d\n%s", tc.size.Height, got, rendered)
			}
			helpLine := lines[len(lines)-1]
			if !strings.Contains(helpLine, "Enter") && !strings.Contains(helpLine, "↑↓") {
				t.Fatalf("expected bottom help on final row, got %q\n%s", helpLine, rendered)
			}
			if tc.minOccupancy > 0 {
				top := firstSceneLineIndex(t, lines)
				status := lineIndexContaining(t, lines, "초안만")
				visibleSceneHeight := status - top + 1
				bodyHeight := tc.size.Height - 2
				if visibleSceneHeight*100 < bodyHeight*tc.minOccupancy {
					t.Fatalf("expected scene occupancy >= %d%%, got %d/%d lines\n%s", tc.minOccupancy, visibleSceneHeight, bodyHeight, rendered)
				}
				if gap := blankLinesBetween(lines, lineIndexContaining(t, lines, "╰"), lineIndexContainingFrom(t, lines, "╭", lineIndexContaining(t, lines, "╰"))); gap > 2 {
					t.Fatalf("expected assistant/composer gap <= 2, got %d\n%s", gap, rendered)
				}
				if gap := blankLinesBetween(lines, status, len(lines)-1); gap > 2 {
					t.Fatalf("expected at most 2 blank lines before help, got %d\n%s", gap, rendered)
				}
				if maxBlankRun(lines) > 3 {
					t.Fatalf("expected no 4-line blank run in full-height canvas\n%s", rendered)
				}
			}
		})
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
	for range 3 {
		m = updateForTest(t, m, initTickMsg{})
	}
	if m.frame != 3 {
		t.Fatalf("expected animation step to advance to 3, got %d", m.frame)
	}
	after := m.renderShell()
	if before == after {
		t.Fatalf("expected animation step to affect rendered motion")
	}
	if !strings.Contains(normalizePlainRender(after), "■") {
		t.Fatalf("expected animation frame to pulse the creature smile or sunglasses\n%s", after)
	}
}

func TestCreatureBlinkAnimationKeepsSilhouette(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m.animationOn = true
	m = updateForTest(t, m, tea.WindowSizeMsg{Width: 120, Height: 40})
	m.frame = 7
	rendered := normalizePlainRender(m.renderShell())
	if !strings.Contains(rendered, "█████████▓▓▓■▓██▓▓▓■▓█████████") {
		t.Fatalf("expected blink frame to sparkle creature sunglasses\n%s", rendered)
	}
	if !strings.Contains(rendered, "██████████████") || !strings.Contains(rendered, "███████▒▒▒▒▒▒▒▒▒▒███████") {
		t.Fatalf("blink should keep creature silhouette stable\n%s", rendered)
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
	for _, want := range []string{"assistant", "your answer", "[##--------]"} {
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

func blockContaining(rendered string, marker string) string {
	lines := splitLines(rendered)
	start := -1
	for i, line := range lines {
		if strings.Contains(line, marker) {
			start = i
			break
		}
	}
	if start < 0 {
		return ""
	}
	for start > 0 && !strings.Contains(lines[start], "╭") {
		start--
	}
	end := start
	for end < len(lines)-1 && !strings.Contains(lines[end], "╰") {
		end++
	}
	return strings.Join(lines[start:end+1], "\n")
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

func lineIndexContaining(t *testing.T, lines []string, want string) int {
	t.Helper()
	return lineIndexContainingFrom(t, lines, want, 0)
}

func lineIndexContainingFrom(t *testing.T, lines []string, want string, after int) int {
	t.Helper()
	for i := after + 1; i < len(lines); i++ {
		if strings.Contains(lines[i], want) {
			return i
		}
	}
	t.Fatalf("could not find line containing %q after %d\n%s", want, after, strings.Join(lines, "\n"))
	return -1
}

func firstSceneLineIndex(t *testing.T, lines []string) int {
	t.Helper()
	for _, marker := range []string{"██████████████", "██████████", "ni·· assistant", "ni▣"} {
		for i, line := range lines {
			if strings.Contains(line, marker) {
				return i
			}
		}
	}
	t.Fatalf("could not find assistant scene start\n%s", strings.Join(lines, "\n"))
	return -1
}

func creatureBodyRows(asset stageAsset, size assetSize) []string {
	leafRows := 0
	switch size {
	case assetFull:
		leafRows = 4
	case assetCompact:
		leafRows = 2
	}
	if leafRows >= len(asset.Lines) {
		return nil
	}
	return asset.Lines[leafRows:]
}

func creatureBodyWidths(asset stageAsset, size assetSize) []int {
	rows := creatureBodyRows(asset, size)
	widths := make([]int, 0, len(rows))
	for _, row := range rows {
		widths = append(widths, lipgloss.Width(strings.TrimSpace(stripANSI(row))))
	}
	return widths
}

func runeIndex(text string, needle rune) int {
	for i, r := range []rune(text) {
		if r == needle {
			return i
		}
	}
	return -1
}

func runeLastIndex(text string, needle rune) int {
	runes := []rune(text)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == needle {
			return i
		}
	}
	return -1
}

func leadingCells(line string) int {
	return lipgloss.Width(line[:len(line)-len(strings.TrimLeft(line, " "))])
}

func visualEnd(line string) int {
	return lipgloss.Width(strings.TrimRight(line, " "))
}

func borderStartCells(line string) int {
	index := strings.Index(line, "╭")
	if index < 0 {
		index = strings.Index(line, "┌")
	}
	if index < 0 {
		return leadingCells(line)
	}
	return lipgloss.Width(line[:index])
}

func borderWidth(line string) int {
	return visualEnd(line) - borderStartCells(line)
}

func blankLinesBetween(lines []string, start int, end int) int {
	count := 0
	for i := start + 1; i < end && i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			count++
		}
	}
	return count
}

func maxBlankRun(lines []string) int {
	maxRun := 0
	current := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			current++
			if current > maxRun {
				maxRun = current
			}
			continue
		}
		current = 0
	}
	return maxRun
}

func countHabitatCreatureLines(lines []string, layout layoutSpec) int {
	count := 0
	for _, line := range lines[1:max(1, len(lines)-1)] {
		segment := habitatSegment(line)
		if strings.Contains(segment, "█") ||
			strings.Contains(segment, "▒▒") ||
			strings.Contains(segment, "▄▀▀▌") ||
			strings.Contains(segment, "▄▀▌") ||
			strings.Contains(segment, "▐██") ||
			strings.Contains(segment, "▀████") ||
			strings.Contains(segment, "▀▀▄") {
			count++
		}
	}
	return count
}

func countDioramaDecorativeLines(lines []string, layout layoutSpec) int {
	count := 0
	for _, line := range lines[1 : len(lines)-1] {
		segment := habitatSegment(line)
		if strings.TrimSpace(segment) == "" || isMascotSegment(segment) {
			continue
		}
		count++
	}
	return count
}

func countMutedBackgroundGlyphs(lines []string, layout layoutSpec) int {
	count := 0
	for _, line := range lines[1 : len(lines)-1] {
		segment := habitatSegment(line)
		if isMascotSegment(segment) {
			continue
		}
		count += strings.Count(segment, "░")
		count += strings.Count(segment, "·")
	}
	return count
}

func countMascotGlyphs(lines []string, layout layoutSpec) int {
	count := 0
	for _, line := range lines[1 : len(lines)-1] {
		segment := habitatSegment(line)
		count += strings.Count(segment, "█")
		count += strings.Count(segment, "▒")
		count += strings.Count(segment, "░")
	}
	return count
}

func isMascotSegment(segment string) bool {
	for _, marker := range []string{"▄▀▀▌", "▄▀▌", "▐██", "█", "▒", "██░", "▀████", "▀▀▄", "░▀▌"} {
		if strings.Contains(segment, marker) {
			return true
		}
	}
	return false
}

func habitatSegment(line string) string {
	cut := len(line)
	for _, marker := range []string{"╭", "│", "╰"} {
		if index := strings.Index(line, marker); index >= 0 && index < cut {
			cut = index
		}
	}
	return line[:cut]
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

func noisyDioramaCopyBanned() []string {
	return []string{
		"░▀░",
		"░░▒░░",
		"▐░▌",
		"░░▄░░ · ▄▄░░▄░░",
		"▄░░░░░░░░░",
	}
}
