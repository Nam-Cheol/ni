package initui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestModelInitialStateUsesAltScreen(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	view := m.View()
	if !view.AltScreen {
		t.Fatalf("expected alt screen view")
	}
	if m.fields[0].Value != "demo" {
		t.Fatalf("expected default project name, got %q", m.fields[0].Value)
	}
}

func TestUpdateHandlesUpDownAndLeftRight(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})

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

func TestUpdateHandlesEnterEscAndQ(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
	m = updateForTest(t, m, key(tea.KeyEnter))
	if m.cursor != 1 {
		t.Fatalf("expected enter to continue, got cursor %d", m.cursor)
	}
	m = updateForTest(t, m, key(tea.KeyEsc))
	if m.cursor != 0 || m.canceled {
		t.Fatalf("expected esc to go back without canceling, got cursor=%d canceled=%v", m.cursor, m.canceled)
	}
	m = updateForTest(t, m, textKey("q"))
	if !m.canceled {
		t.Fatalf("expected q to cancel")
	}
}

func TestConfirmPathReturnsIntent(t *testing.T) {
	m := NewModel(Config{Dir: ".", DefaultName: "demo"})
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
	m.cursor = len(m.fields) - 1
	m = updateForTest(t, m, key(tea.KeyEnter))
	m = updateForTest(t, m, key(tea.KeyDown))
	m = updateForTest(t, m, key(tea.KeyEnter))
	if result := m.Result(); !result.Canceled || result.Confirmed {
		t.Fatalf("expected canceled result, got %#v", result)
	}
}

func TestExistingFileChoices(t *testing.T) {
	m := NewModel(Config{Dir: ".", ExistingFiles: []string{".ni/contract.json"}})
	if m.stage != stageExisting {
		t.Fatalf("expected existing-file stage, got %v", m.stage)
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

func key(code rune) tea.KeyPressMsg {
	return tea.KeyPressMsg(tea.Key{Code: code})
}

func textKey(value string) tea.KeyPressMsg {
	runes := []rune(value)
	return tea.KeyPressMsg(tea.Key{Code: runes[0], Text: value})
}
