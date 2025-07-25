package parser

import (
	"fmt"
	"github.com/floffah/schemapi/internal/lexer"
)

// Parser holds the state for parsing
type Parser struct {
	lexer  *lexer.Lexer
	curTok lexer.Token
}

func NewParser(input string) *Parser {
	lex := lexer.NewLexer(input)
	return &Parser{lexer: lex, curTok: lex.NextToken()}
}

func (p *Parser) next() lexer.Token {
	p.curTok = p.lexer.NextToken()
	return p.curTok
}

func (p *Parser) current() lexer.Token {
	return p.curTok
}

func (p *Parser) expect(typ lexer.TokenType) bool {
	return p.curTok.Type == typ
}

func (p *Parser) getError(message string) error {
	return &ParserError{
		Message: message,
		Line:    p.curTok.Line,
		Col:     p.curTok.Col - len(p.curTok.Literal),
	}
}

func (p *Parser) getErrorf(format string, args ...interface{}) error {
	return &ParserError{
		Message: fmt.Sprintf(format, args...),
		Line:    p.curTok.Line,
		Col:     p.curTok.Col - len(p.curTok.Literal),
	}
}
