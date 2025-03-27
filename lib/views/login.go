// Package loginPage has the view for the login screen
package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type loginModel struct {
	options        []string
	selectedOption int
	cursor         int
}

func ConstructLogin() tea.Model {
	// Initialize the login model
	return loginModel{}
}

func (m loginModel) Init() tea.Cmd {
	// Initialize the view
	fmt.Println("Login view initialized")
	return nil
}
func (m loginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (m loginModel) View() string {
	// Render the view
	return "Login View"
}
