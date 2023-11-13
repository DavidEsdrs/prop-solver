package prop_solver

import (
	"io"
	"strings"

	"github.com/DavidEsdrs/prop-solver/eval"
	"github.com/DavidEsdrs/prop-solver/lexer"
	"github.com/DavidEsdrs/prop-solver/parser"
	"github.com/DavidEsdrs/prop-solver/tree"
)

func Solve(input string) (*tree.Tree[*tree.ParseTreeNode], []bool) {
	return evaluate(input)
}

func evaluate(input string) (*tree.Tree[*tree.ParseTreeNode], []bool) {
	tokens := scanner(input)
	root := parser.Parse(tokens)

	e := eval.NewEvaluable(tokens)
	e.Evaluate()

	t := tree.NewTree[*tree.ParseTreeNode](root, e.Result())

	return &t, tree.EvalTree(&t)
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
