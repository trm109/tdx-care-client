// Package model provides the model for the TUI
package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	login state = iota
	home
)

type mainModel struct {
	state  state
	token  string // api token
	cursor int
}

func Construct() tea.Model {
	// Initialize the main model
	return mainModel{
		state:  login,
		token:  "",
		cursor: 0,
	}
}

func (m mainModel) Init() tea.Cmd {
	// returns the current view's init function
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// returns the current view's update function
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m mainModel) View() string {
	// returns the current view's view function
	return "ERR"
}
