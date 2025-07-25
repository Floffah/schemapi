package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#10b981")).
			Render(" ★ Schemapi V0.0.0 ★"),
	)
}
