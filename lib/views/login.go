// Package login has the view for the login screen
package login

import (
		"fmt"
		tea "github.com/charmbracelet/bubbletea"
		models "github.com/tdx-client-tui/lib/models/model"
)

type view struct {
	//
	url string // URL for the TDX API
	bearerToken string // Bearer token for the TDX API
	viewFunctions: models.viewFunctions // functions for the view
}

func construct() {
	// Initialize the view with a default value
	v := view{
		viewFunctions: models.viewFunctions{
			init:   init,
			update: update,
			view:   view,
		},
	}

	fmt.Println("Login view initialized")
	return v
}

func init() tea.Cmd {
	// Initialize the view
	fmt.Println("Login view initialized")
	return nil
}

func update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Update the view
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return nil, tea.Quit
		}
	}
	return nil, nil
}

func view() string {
	// Render the view
	return "Login View"
}
