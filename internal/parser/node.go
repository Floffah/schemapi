package parser

type NodeType int

const (
	NodeRoot NodeType = iota
	NodeCallable
	NodeIdentifier
)

type Node interface {
	NodeType() NodeType
}

type RootNode struct {
	Children []Node
}

func (n *RootNode) NodeType() NodeType {
	return NodeRoot
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
	return NodeIdentifier
}

type CallableNode struct {
	Type       string
	Identifier Node
	Params     []string
	Children   []Node
}

func (n *CallableNode) NodeType() NodeType {
	return NodeCallable
}
