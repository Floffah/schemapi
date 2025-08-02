package parser

type NodeType int

const (
	NodeRoot NodeType = iota
	NodeString
	NodeNumber
	NodePath
	NodeIdentifier
	NodeCallable
	NodeCallableDefinition
	NodeDictionary
	NodeEntry
)

type Node interface {
	NodeType() NodeType
}

// --- basic nodes ---

type RootNode struct {
	Children []Node
}

func (n *RootNode) NodeType() NodeType {
	return NodeRoot
}

type StringNode struct {
	Value string
}

func (n *StringNode) NodeType() NodeType {
	return NodeString
}

type NumberNode struct {
	Value string
}

func (n *NumberNode) NodeType() NodeType {
	return NodeNumber
}

type IdentifierNode struct {
	Name string
}

func (n *IdentifierNode) NodeType() NodeType {
	return NodeIdentifier
}

type PathNode struct {
	parts []string
}

func (n *PathNode) NodeType() NodeType {
	return NodePath
}

// --- blocks ---

type CallableNode struct {
	Type       Node
	Identifier Node
	Params     []Node
	Children   []Node
}

func (n *CallableNode) NodeType() NodeType {
	return NodeCallable
}

type CallableDefinitionNode struct {
	Type     Node
	Params   []Node
	Children Node
}

func (n *CallableDefinitionNode) NodeType() NodeType {
	return NodeCallableDefinition
}

type DictionaryNode struct {
	Entries []EntryNode
}

func (n *DictionaryNode) NodeType() NodeType {
	return NodeDictionary
}

type EntryNode struct {
	Identifier Node
	Value      Node
}

func (n *EntryNode) NodeType() NodeType {
	return NodeEntry
}
