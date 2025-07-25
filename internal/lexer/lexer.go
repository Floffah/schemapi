package lexer

import (
	"unicode"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	Line         int
	Col          int
	ch           rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, Line: 1, Col: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}

	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.Line++
		l.Col = 0
	} else {
		l.Col++
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	switch l.ch {
	case '{':
		l.readChar()
		return Token{Type: TokenLBrace, Literal: "{", Line: l.Line, Col: l.Col}
	case '}':
		l.readChar()
		return Token{Type: TokenRBrace, Literal: "}", Line: l.Line, Col: l.Col}
	case '=':
		l.readChar()
		return Token{Type: TokenEqual, Literal: "=", Line: l.Line, Col: l.Col}
	case ':':
		l.readChar()
		return Token{Type: TokenColon, Literal: ":", Line: l.Line, Col: l.Col}
	case ',':
		l.readChar()
		return Token{Type: TokenComma, Literal: ",", Line: l.Line, Col: l.Col}
	case '"', '\'':
		return l.readString()
	case 0:
		return Token{Type: TokenEOF, Literal: "", Line: l.Line, Col: l.Col}
	default:
		if l.ch == '/' {
			path := l.readPath()
			return Token{Type: TokenPath, Literal: path, Line: l.Line, Col: l.Col}
		} else if isLetter(l.ch) {
			ident := l.readIdentifier()
			if isKeyword(ident) {
				return Token{Type: TokenKeyword, Literal: ident, Line: l.Line, Col: l.Col}
			}
			return Token{Type: TokenIdent, Literal: ident, Line: l.Line, Col: l.Col}
		} else if unicode.IsDigit(l.ch) {
			num := l.readNumber()
			return Token{Type: TokenNumber, Literal: num, Line: l.Line, Col: l.Col}
		}
		illegal := l.ch
		l.readChar()
		return Token{Type: TokenIllegal, Literal: string(illegal)}
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) || unicode.IsDigit(l.ch) || l.ch == '.' || l.ch == '_' || l.ch == '/' {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readString() Token {
	quote := l.ch
	l.readChar() // skip opening quote
	pos := l.position
	for l.ch != quote && l.ch != 0 {
		l.readChar()
	}
	str := l.input[pos:l.position]
	l.readChar() // skip closing quote
	return Token{Type: TokenString, Literal: str, Line: l.Line, Col: l.Col}
}

func (l *Lexer) readPath() string {
	pos := l.position
	for l.ch == '/' || isLetter(l.ch) || unicode.IsDigit(l.ch) || l.ch == '-' || l.ch == '_' || l.ch == ':' {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || ch == '.' || ch == '/'
}

func isKeyword(ident string) bool {
	// Only treat truly reserved words as keywords, e.g. 'callable'.
	reserved := []string{"callable"}
	for _, k := range reserved {
		if ident == k {
			return true
		}
	}
	return false
}

func (l *Lexer) Lex() []Token {
	var tokens []Token
	for tok := l.NextToken(); tok.Type != TokenEOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}
	return tokens
}
