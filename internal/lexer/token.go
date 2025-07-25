package lexer

type TokenType string

const (
	TokenKeyword TokenType = "Keyword"
	TokenIdent   TokenType = "Ident"
	TokenString  TokenType = "String"
	TokenNumber  TokenType = "Number"
	TokenLBrace  TokenType = "LBrace"
	TokenRBrace  TokenType = "RBrace"
	TokenLParen  TokenType = "LParen"
	TokenRParen  TokenType = "RParen"
	TokenEqual   TokenType = "Equal"
	TokenColon   TokenType = "Colon"
	TokenComma   TokenType = "Comma"
	TokenEOF     TokenType = "EOF"
	TokenPath    TokenType = "Path"
	TokenIllegal TokenType = "Illegal"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
}
