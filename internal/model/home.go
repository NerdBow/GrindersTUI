package model

import (
	"strings"

	"github.com/NerdBow/GrindersTUI/internal/keymap"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type HomeModel struct {
	choices []string
	cursor  int
}

func HomeModelInit() HomeModel {
	return HomeModel{
		choices: []string{"Create Log", "View Log(s)", "Edit Log", "Delete Log", "Sign Out"},
		cursor:  0,
	}
}

func (m HomeModel) Init() tea.Cmd {
	return nil
}

func (m HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.VimBinding.Up):
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil
		case key.Matches(msg, keymap.VimBinding.Down):
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
			return m, nil
		case key.Matches(msg, keymap.VimBinding.Exit):
			return m, tea.Quit
		case key.Matches(msg, keymap.VimBinding.Select):
			return m, nil
		}

	}
	return m, nil
}

func (m HomeModel) View() string {
	var b strings.Builder

	for i := range m.choices {
		if i == m.cursor {
			b.WriteString("> ")
		}

		b.WriteString(m.choices[i])

		if i < len(m.choices)-1 {
			b.WriteRune('\n')
		}
	}
	return b.String()
}
