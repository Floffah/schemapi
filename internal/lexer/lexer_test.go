package lexer

import (
	"os"
	"strings"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestLexerSnapshot(t *testing.T) {
	files, err := os.ReadDir("../../tests")

	if err != nil {
		t.Fatalf("failed to read examples directory: %v", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sapi") {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			data, err := os.ReadFile("../../tests/" + file.Name())

			if err != nil {
				t.Fatalf("failed to read %s: %v", file.Name(), err)
			}

			lexer := NewLexer(string(data))
			tokens := lexer.Lex()

			snaps.MatchSnapshot(t, tokens)
		})
	}
}
