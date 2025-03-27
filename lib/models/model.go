// Package model provides the model for the TUI
package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/trm109/tdx-care-client/lib/views"
)

//type state int

//const (
//	login state = iota
//	home
//)

type model struct {
	//state state
	view  tea.Model
	token string // api token
}

func Construct() tea.Model {
	// Initialize the main model
	return model{
		view:  views.ConstructLogin(), // Default view
		token: "",
	}
}

func (m model) Init() tea.Cmd {
	// returns the current view's init function
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	// returns the current view's update function
	//switch msg := msg.(type) {
	//case tea.KeyMsg:
	//	switch msg.String() {
	//	case "ctrl+c", "q":
	//		return m, tea.Quit
	//	}
	//}
	m.view, cmd = m.view.Update(msg)
	cmds = append(cmds, cmd)

	//cmds = append(cmds, cmd)
	//m.setViewportContent() // refresh the content on every Update call
	//return m, tea.Batch(cmds...)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// returns the current view's view function
	// Can wrap this view in a lipgloss formatter
	return m.view.View()
}
