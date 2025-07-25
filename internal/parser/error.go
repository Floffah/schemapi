package parser

type ParserError struct {
	Message string
	Line    int
	Col     int
}

func (e *ParserError) Error() string {
	return e.Message
}
