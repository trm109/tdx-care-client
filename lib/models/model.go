// Package model provides the model for the TUI
package model

import (
	tea "github.com/charmbracelet/bubbletea"

	// import all the views
	// lib/views/login.go
	login "github.com/tdx-client-tui/lib/views/login"
)

type model struct {
	currentView    struct{} // current view
	tdxBearerToken string   // bearer token for TDX API
	//options     []string         // options for the current view
	//cursor      int              // cursor position
	//selected    map[int]struct{} // selected items
}

type viewFunctions struct {
	init   func() tea.Cmd
	update func(msg tea.Msg) (tea.Model, tea.Cmd)
	view   func() string
}

func tdxModel() model {
	// Initialize the model with a default value
	return model{
		currentView: login.View{},
		cursor:      0,
		selected:    make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// returns the current view's init function
	return m.currentView.viewFunctions.init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// returns the current view's update function
	return m.currentView.viewFunctions.update(msg)
}

func (m model) View() string {
	// returns the current view's view function
	return m.currentView.viewFunctions.view()
}
