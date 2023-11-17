package lexer_test

import (
	"io"
	"strings"
	"testing"

	"github.com/DavidEsdrs/prop-solver/lexer"
	"github.com/DavidEsdrs/prop-solver/utils"
)

func TestLexer(t *testing.T) {
	expectedResult := []string{"p", "/\\", "q", "=>", "t", "!=", "h", "\\/", "v", "<->", "j", "->", "k"}

	l := lexer.NewLexer(strings.NewReader("p /\\ q => t != h \\/ v <-> j -> k"))

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

	strings := utils.Map(tokens, func(tok *lexer.Token) string {
		return tok.TStr
	})

	if !utils.EqualsSlice[string](expectedResult, strings) {
		t.Errorf("results don't match! fail")
	}
}
