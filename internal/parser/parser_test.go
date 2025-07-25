package parser

import (
	"errors"
	"github.com/gkampitakis/go-snaps/snaps"
	"os"
	"strings"
	"testing"
)

func TestParserSnapshot(t *testing.T) {
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

			parser := NewParser(string(data))
			rootNode, err := parser.Parse()

			if err != nil {
				// if error is of type ParseError, we can provide more context
				var parseErr *ParserError
				if errors.As(err, &parseErr) {
					t.Fatalf("\n%s:%d:%d: %s", file.Name(), parseErr.Line, parseErr.Col, parseErr.Message)
				} else {
					t.Fatalf("failed to parse %s: %v", file.Name(), err)
				}
			}

			snaps.MatchSnapshot(t, rootNode)
		})
	}
}
