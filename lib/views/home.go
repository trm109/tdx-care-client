// Package homePage provides the home page for the TDX Client TUI.
package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type homeModel struct {
	options        []string
	selectedOption int
	cursor         int
}

func ConstructHome() tea.Model {
	return homeModel{
		options: []string{
			"Option 1",
			"Option 2",
			"Option 3",
		},
	}
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, textinput.Blink
}

func (m homeModel) View() string {
	s := ""
	for i, option := range m.options {
		if i == m.cursor {
			s += fmt.Sprintf("[%s] %s\n", "x", option)
		} else {
			s += fmt.Sprintf("[ ] %s\n", option)
		}
	}
	return s
}
