package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/DavidEsdrs/prop-solver/lexer"
	"github.com/DavidEsdrs/prop-solver/tree"
)

func main() {
	args := os.Args[1:]
	input := args[0]

	fmt.Printf("input: %v\n", input)

	tokens := scanner(input)
	root := parse(tokens)

	tree := tree.NewTree[*lexer.Token](root)
	tree.InOrderTraversal()
}

func scanner(input string) []*lexer.Token {
	l := lexer.NewLexer(strings.NewReader(input))

	strArr := []*lexer.Token{}

	for {
		str, err := l.Lex()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		strArr = append(strArr, &str)
	}

	return strArr
}

func parse(tokens []*lexer.Token) *tree.Node[*lexer.Token] {
	if len(tokens) == 0 {
		return nil
	}

	if len(tokens) == 1 {
		return &tree.Node[*lexer.Token]{Value: tokens[0]}
	}

	var (
		root            *tree.Node[*lexer.Token]
		connectiveIndex int
	)

	lastPrecedence := math.MaxInt
	length := len(tokens)

	for i := length - 1; i >= 0; i-- {
		if tokens[i].TType == lexer.CONN {
			currentPrecedence := getPrecedence(tokens[i].TStr)

			if currentPrecedence < lastPrecedence {
				root = &tree.Node[*lexer.Token]{Value: tokens[i]}
				connectiveIndex = i
				lastPrecedence = currentPrecedence
			}
		}
	}

	if root != nil {
		left := parse(tokens[:connectiveIndex])
		right := parse(tokens[connectiveIndex+1:])

		root.Left = left
		root.Right = right
	}

	return root
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
