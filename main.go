package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/DavidEsdrs/prop-solver/lexer"
)

func main() {
	input := "(p/\\q\\/r)<->s->r"
	l := lexer.NewLexer(strings.NewReader(input))

	strArr := []lexer.Token{}

	for {
		str, err := l.Lex()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		strArr = append(strArr, str)
	}

	for _, n := range strArr {
		fmt.Printf("%#v\n", n)
	}
}
