package parser

import (
	"github.com/floffah/schemapi/internal/lexer"
	"strings"
)

func (p *Parser) Parse() (*RootNode, error) {
	root := &RootNode{
		Children: []Node{},
	}

	for !p.expect(lexer.TokenEOF) {
		node, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		if node != nil {
			root.Children = append(root.Children, node)
		} else {
			if p.expect(lexer.TokenEOF) {
				break
			}
			return nil, p.getErrorf("unexpected token: %s", p.curTok.Value)
		}
	}

	return root, nil
}

// --- Blocks ---

func (p *Parser) parseBlock() (Node, error) {
	if p.expect(lexer.TokenKeyword) {
		if p.curTok.Value == "callable" {
			return p.parseCallable()
		}
	}
	return nil, nil
}

func (p *Parser) parseCallable() (Node, error) {
	if !p.expect(lexer.TokenKeyword) || p.curTok.Value != "callable" {
		return nil, p.getErrorf("expected 'callable' but got %s", p.curTok.Value)
	}
	p.next() // advance after confirming 'callable' keyword

	callableType, err := p.parseIdentifier()
	if err != nil {
		return nil, err
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
		return nil, p.getErrorf("expected identifier or path, got %s", p.curTok.Value)
	}

	node := &CallableNode{
		Type:       callableType,
		Identifier: identifier,
		Params:     []Node{},
		Children:   []Node{},
	}

	for !p.expect(lexer.TokenLBrace) {
		if p.expect(lexer.TokenEOF) {
			return nil, p.getError("unexpected EOF before '{'")
		}

		param, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		node.Params = append(node.Params, param)
		p.next()
	}
	p.next() // advance past '{'

	for !p.expect(lexer.TokenRBrace) {
		if p.expect(lexer.TokenEOF) {
			return nil, p.getError("unexpected EOF before '}'")
		}

		childNode, err := p.parseCallableDefinition()
		if err != nil {
			return nil, err
		}

		if childNode != nil {
			node.Children = append(node.Children, childNode)
		} else {
			return nil, p.getErrorf("unexpected token: %s", p.curTok.Value)
		}
	}
	p.next() // advance past '}'

	return node, nil
}

func (p *Parser) parseCallableDefinition() (Node, error) {
	defType, err := p.parseIdentifier()
	if err != nil {
		return nil, err
	}

	node := &CallableDefinitionNode{
		Type:     defType,
		Params:   []Node{},
		Children: nil,
	}

	for !p.expect(lexer.TokenLBrace) {
		if p.expect(lexer.TokenEOF) {
			return nil, p.getError("unexpected EOF before '{'")
		}

		param, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		node.Params = append(node.Params, param)
	}

	dictionary, err := p.parseDictionary()
	if err != nil {
		return nil, err
	}

	node.Children = dictionary

	return node, nil
}

func (p *Parser) parseDictionary() (Node, error) {
	node := &DictionaryNode{
		Entries: []EntryNode{},
	}

	if !p.expect(lexer.TokenLBrace) {
		return nil, p.getErrorf("expected '{' but got %s", p.curTok.Value)
	}
	p.next() // advance past '{'

	for !p.expect(lexer.TokenRBrace) {
		if p.expect(lexer.TokenEOF) {
			return nil, p.getError("unexpected EOF before '}'")
		}

		entry := EntryNode{}

		identifier, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}

		entry.Identifier = identifier

		if !p.expect(lexer.TokenEqual) {
			return nil, p.getErrorf("expected '=' but got %s", p.curTok.Value)
		}

		p.next() // advance past '='

		value, err := p.parseValueOrComplexType()
		if err != nil {
			return nil, err
		}
		entry.Value = value

		node.Entries = append(node.Entries, entry)
	}

	p.next() // advance past '}'

	return node, nil
}

// --- Special Nodes ---

func (p *Parser) parseValue() (Node, error) {
	if p.expect(lexer.TokenIdent) {
		return p.parseIdentifier()
	} else if p.expect(lexer.TokenString) {
		return p.parseString()
	} else if p.expect(lexer.TokenNumber) {
		return p.parseNumber()
	} else if p.expect(lexer.TokenPath) {
		return p.parsePath()
	}

	return nil, p.getErrorf("expected value but got %s", p.curTok.Value)
}

func (p *Parser) parseValueOrComplexType() (Node, error) {
	if p.expect(lexer.TokenLBrace) {
		return p.parseDictionary()
	} else if p.expect(lexer.TokenKeyword) && p.curTok.Value == "callable" {
		return p.parseCallable()
	}

	return p.parseValue()
}

func (p *Parser) parseString() (Node, error) {
	if !p.expect(lexer.TokenString) {
		return nil, p.getErrorf("expected string but got %s", p.curTok.Value)
	}

	node := &StringNode{
		Value: p.curTok.Value,
	}
	p.next()

	return node, nil
}

func (p *Parser) parseNumber() (Node, error) {
	if !p.expect(lexer.TokenNumber) {
		return nil, p.getErrorf("expected number but got %s", p.curTok.Value)
	}

	node := &NumberNode{
		Value: p.curTok.Value,
	}
	p.next()

	return node, nil
}

func (p *Parser) parseIdentifier() (Node, error) {
	if !p.expect(lexer.TokenIdent) {
		return nil, p.getErrorf("expected identifier but got %s", p.curTok.Value)
	}

	node := &IdentifierNode{
		Name: p.curTok.Value,
	}
	p.next()

	return node, nil
}

func (p *Parser) parsePath() (Node, error) {
	if !p.expect(lexer.TokenPath) {
		return nil, p.getErrorf("expected path but got %s", p.curTok.Value)
	}

	path := strings.Clone(p.curTok.Value)
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
