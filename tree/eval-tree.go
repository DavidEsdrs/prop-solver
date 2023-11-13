package tree

import (
	"github.com/DavidEsdrs/prop-solver/expressions"
)

func EvalTree(t *Tree[*ParseTreeNode]) []bool {
	return interpretate(t.root, t)
}

func interpretate(root *Node[*ParseTreeNode], t *Tree[*ParseTreeNode]) []bool {
	switch {
	case root == nil:
		return []bool{}
	case root.IsLeaf():
		return generateArray(root, t)
	default:
		left := interpretate(root.Left, t)
		right := interpretate(root.Right, t)
		result := eval(root, left, right)
		t.mapping[root.Value.FullQualifiedProp] = result
		return result
	}
}

func generateArray(root *Node[*ParseTreeNode], t *Tree[*ParseTreeNode]) []bool {
	return t.mapping[root.Value.Str]
}

func eval(root *Node[*ParseTreeNode], left, right []bool) []bool {
	switch root.Value.Type {
	case expressions.AND:
		return and(left, right)
	case expressions.OR:
		return or(left, right)
	case expressions.XOR:
		return xor(left, right)
	case expressions.IF_AND_ONLY_IF:
		return ifAndOnlyIf(left, right)
	case expressions.IMPLIES:
		return implies(left, right)
	default:
		return not(right) // the leaf is ALWAYS in the right when the expression is a "not"
	}
}

func not(input []bool) []bool {
	result := make([]bool, len(input))
	for i, n := range input {
		result[i] = !n
	}
	return result
}

func and(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] && b[i]
	}
	return result
}

func or(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] || b[i]
	}
	return result
}

func xor(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] != b[i]
	}
	return result
}

func ifAndOnlyIf(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] == b[i]
	}
	return result
}

func implies(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = !a[i] || b[i]
	}
	return result
}
