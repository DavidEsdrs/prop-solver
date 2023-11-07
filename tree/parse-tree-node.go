package tree

import "github.com/DavidEsdrs/prop-solver/expressions"

type ConnType int

type ParseTreeNode struct {
	Str               string
	FullQualifiedProp string
	Type              expressions.Op
}

// this functions generates the truth table for the given expression
func (ptn *ParseTreeNode) GenerateTable() {}
