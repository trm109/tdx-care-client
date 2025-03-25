// Package main provides the entry point for the TDX Client TUI application.
package main

// Entry point for TDX Client TUI

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the program
	p := tea.NewProgram(models.tdxModel(), tea.WithAltScreen())

	// Start the program and handle any errors
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exiting TDX Client TUI")
}
