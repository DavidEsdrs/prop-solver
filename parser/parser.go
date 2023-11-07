package parser

import (
	"math"
	"strings"

	"github.com/DavidEsdrs/prop-solver/expressions"
	"github.com/DavidEsdrs/prop-solver/lexer"
	"github.com/DavidEsdrs/prop-solver/tree"
)

func Parse(tokens []*lexer.Token) *tree.Node[*tree.ParseTreeNode] {
	if len(tokens) == 0 {
		return nil
	}

	if len(tokens) == 1 {
		return &tree.Node[*tree.ParseTreeNode]{
			Value: &tree.ParseTreeNode{
				Str:  tokens[0].TStr,
				Type: expressions.GetConnectiveType(tokens[0].TStr),
			},
		}
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
				root = &tree.Node[*tree.ParseTreeNode]{
					Value: &tree.ParseTreeNode{
						Str:  tokens[i].TStr,
						Type: expressions.GetConnectiveType(tokens[i].TStr),
					},
				}
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

func getPropStr(tokens []*lexer.Token) string {
	var builder strings.Builder

	for _, n := range tokens {
		builder.WriteString(n.TStr)
	}

	return builder.String()
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
