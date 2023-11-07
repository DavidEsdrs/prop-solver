package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/DavidEsdrs/prop-solver/eval"
	"github.com/DavidEsdrs/prop-solver/lexer"
	"github.com/DavidEsdrs/prop-solver/parser"
	"github.com/DavidEsdrs/prop-solver/tree"
)

func main() {
	args := os.Args[1:]
	input := args[0]

	fmt.Printf("input: %v\n", input)

	tokens := scanner(input)
	root := parser.Parse(tokens)

	tree := tree.NewTree[*tree.ParseTreeNode](root)
	tree.InOrderTraversal()

	eval.NewEvaluable(tokens)

}

func scanner(input string) []*lexer.Token {
	l := lexer.NewLexer(strings.NewReader(input))

	tokens := []*lexer.Token{}

	for {
		tok, err := l.Lex()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		tokens = append(tokens, &tok)
	}

	return tokens
}
