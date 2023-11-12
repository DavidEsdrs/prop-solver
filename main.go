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

	for _, n := range evaluate(input) {
		fmt.Printf("%v\n", n)
	}
}

func evaluate(input string) []bool {
	tokens := scanner(input)
	root := parser.Parse(tokens)

	e := eval.NewEvaluable(tokens)
	e.Evaluate()

	t := tree.NewTree[*tree.ParseTreeNode](root, e.Result())

	return tree.EvalTree(&t)
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
