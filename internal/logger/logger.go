package logger

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/floffah/schemapi/internal/parser"
	"os"
)

func PrintHeader() {
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#10b981")).
			Render(" ★ Schemapi V0.0.0 ★"),
	)
}

func PrintError(message string, err error) {
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ef4444")).
			PaddingTop(1).
			Render(fmt.Sprintf(" ✘ %s", message)),
	)
	if err != nil {
		fmt.Println(
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("#a8a29e")).
				PaddingLeft(3).
				Render(err.Error()),
		)
	}
}

func HandleError(err error) {
	if err != nil {
		PrintError("An unexpected error occurred", err)
		os.Exit(1)
	}
}

func HandleParseError(err error, fileName string) {
	if err != nil {
		var parseErr *parser.ParserError
		if errors.As(err, &parseErr) {
			PrintError(fmt.Sprintf("%s:%d:%d: %s", fileName, parseErr.Line, parseErr.Col, parseErr.Message), nil)
			os.Exit(1)
		} else {
			HandleError(err)
		}
	}
}
