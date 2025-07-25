package parser

import (
	"github.com/floffah/schemapi/internal/lexer"
	"strings"
)

// Parse parses the tokens into an AST
func (p *Parser) Parse() (*RootNode, error) {
	root := &RootNode{
		Children: []Node{},
	}

	for p.curTok.Type != lexer.TokenEOF {
		node, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		if node != nil {
			root.Children = append(root.Children, node)
		} else {
			if p.curTok.Type == lexer.TokenEOF {
				break
			}
			return nil, p.getErrorf("unexpected token: %s", p.curTok.Literal)
		}
	}

	return root, nil
}

func (p *Parser) parseBlock() (Node, error) {
	if p.curTok.Type == lexer.TokenKeyword {
		if p.curTok.Literal == "callable" {
			return p.parseCallable()
		}
	}
	return nil, nil
}

func (p *Parser) parseCallable() (Node, error) {
	if !p.expect(lexer.TokenKeyword) || p.curTok.Literal != "callable" {
		return nil, p.getErrorf("expected 'callable' but got %s", p.curTok.Literal)
	}
	p.next() // advance after confirming 'callable' keyword

	callableType := ""

	if p.expect(lexer.TokenIdent) {
		callableType = p.curTok.Literal
		p.next()
	} else {
		return nil, p.getError("expected callable type")
	}

	var identifier Node = &IdentifierNode{
		Name: "",
	}

	if p.expect(lexer.TokenIdent) {
		i, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		identifier = i
	} else if p.expect(lexer.TokenPath) {
		i, err := p.parsePath()
		if err != nil {
			return nil, err
		}
		identifier = i
	} else {
		return nil, p.getErrorf("expected identifier or path, got %s", p.curTok.Literal)
	}

	node := &CallableNode{
		Type:       callableType,
		Identifier: identifier,
		Params:     []string{},
		Children:   []Node{},
	}

	for !p.expect(lexer.TokenLBrace) {
		if p.curTok.Type == lexer.TokenEOF {
			return nil, p.getError("unexpected EOF before '{'")
		}
		node.Params = append(node.Params, p.curTok.Literal)
		p.next()
	}
	p.next() // advance past '{'

	for !p.expect(lexer.TokenRBrace) {
		if p.curTok.Type == lexer.TokenEOF {
			return nil, p.getError("unexpected EOF before '}'")
		}
		p.next()
	}
	p.next() // advance past '}'

	return node, nil
}

func (p *Parser) parseIdentifier() (Node, error) {
	if p.curTok.Type != lexer.TokenIdent {
		return nil, p.getErrorf("expected identifier but got %s", p.curTok.Literal)
	}

	node := &IdentifierNode{
		Name: p.curTok.Literal,
	}
	p.next()

	return node, nil
}

func (p *Parser) parsePath() (Node, error) {
	if p.curTok.Type != lexer.TokenPath {
		return nil, p.getErrorf("expected path but got %s", p.curTok.Literal)
	}

	path := strings.Clone(p.curTok.Literal)
	path = strings.TrimPrefix(path, "/")

	pathPartsSplit := strings.Split(path, "/")
	var parts []string

	for _, part := range pathPartsSplit {
		parts = append(parts, strings.TrimSpace(part))
	}

	node := &PathNode{
		parts: parts,
	}
	p.next()

	return node, nil
}
