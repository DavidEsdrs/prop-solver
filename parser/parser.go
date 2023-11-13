package parser

import (
	"math"
	"strings"

	"github.com/DavidEsdrs/prop-solver/expressions"
	"github.com/DavidEsdrs/prop-solver/lexer"
	"github.com/DavidEsdrs/prop-solver/tree"
)

// Parse generates the Evaluation Tree and returns its root
func Parse(tokens []*lexer.Token) *tree.Node[*tree.ParseTreeNode] {
	if len(tokens) == 0 {
		return nil
	}

	if len(tokens) == 1 {
		return createNode(tokens[0])
	}

	var (
		root            *tree.Node[*tree.ParseTreeNode]
		connectiveIndex int
	)

	lastPrecedence := math.MaxInt
	length := len(tokens)

	for i := length - 1; i >= 0; i-- {
		if tokens[i].TType == lexer.CONN {
			currentPrecedence := getPrecedence(tokens[i].TStr)

			if currentPrecedence < lastPrecedence {
				root = createNode(tokens[i])
				connectiveIndex = i
				lastPrecedence = currentPrecedence
			}
		}
	}

	if root != nil {
		left := Parse(tokens[:connectiveIndex])
		right := Parse(tokens[connectiveIndex+1:])

		root.Left = left
		root.Right = right

		root.Value.FullQualifiedProp = getPropStr(tokens)
	}

	return root
}

func createNode(token *lexer.Token) *tree.Node[*tree.ParseTreeNode] {
	node := tree.Node[*tree.ParseTreeNode]{
		Value: &tree.ParseTreeNode{
			Str: token.TStr,
		},
	}

	if isConnective(token) {
		node.Value.Type = expressions.GetConnectiveType(token.TStr)
	}

	return &node
}

func getPropStr(tokens []*lexer.Token) string {
	var builder strings.Builder

	for _, n := range tokens {
		builder.WriteString(n.TStr)
	}

	return builder.String()
}

func isConnective(tok *lexer.Token) bool {
	prec := getPrecedence(tok.TStr)
	return prec != 0 && prec != math.MaxInt32
}

// returns the precedence of each connective
func getPrecedence(conn string) int {
	switch conn {
	case "<->":
		return 1
	case "->":
		return 2
	case "/\\", "\\/":
		return 3
	case "~":
		return 4
	case "":
		return math.MaxInt32
	default:
		return 0
	}
}
