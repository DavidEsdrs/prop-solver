package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	OPENING TokenType = 1 << iota
	CLOSING
	IDENT
	CONN
)

func (t TokenType) String() string {
	switch t {
	case OPENING:
		return "OPENING"
	case CLOSING:
		return "CLOSING"
	case IDENT:
		return "IDENT"
	case CONN:
		return "CONN"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	TStr              string
	TType             TokenType
	FullQualifiedProp string
}

func (t Token) String() string {
	return fmt.Sprintf("%v", t.TStr)
}

func NewToken(str string, TType TokenType) Token {
	return Token{TStr: str, TType: TType}
}

var ErrInvalidConnective = errors.New("invalid connective")

type Lexer struct {
	buf *bufio.Reader
}

func NewLexer(rd io.Reader) Lexer {
	return Lexer{
		buf: bufio.NewReader(rd),
	}
}

// lex return the next token
func (l *Lexer) Lex() (Token, error) {
	for {
		r, _, err := l.buf.ReadRune()

		if err != nil {
			return Token{}, err
		}

		if unicode.IsSpace(r) {
			continue
		} else if unicode.IsLetter(r) {
			return l.lexIdentifier(r)
		} else {
			switch r {
			case '~':
				return NewToken(string(r), CONN), nil
			case '(':
				return NewToken(string(r), OPENING), nil
			case ')':
				return NewToken(string(r), CLOSING), nil
			case '/':
				return l.parseAnd()
			case '\\':
				return l.parseOr()
			case '-':
				return l.parseImplies()
			case '<':
				return l.parseIfAndOnlyIf()
			case '!':
				return l.parseXor(true) // TODO: remove bool
			case '>':
				return l.parseXor(false)
			default:
				break
			}
		}
	}
}

func (l *Lexer) parseAnd() (Token, error) {
	nextR, _, err := l.buf.ReadRune()

	if err != nil {
		return Token{}, err
	}

	if nextR != '\\' {
		return Token{}, ErrInvalidConnective
	}

	return NewToken("/\\", CONN), nil
}

func (l *Lexer) parseOr() (Token, error) {
	nextR, _, err := l.buf.ReadRune()

	if err != nil {
		panic(err)
	}

	if nextR != '/' {
		return Token{}, ErrInvalidConnective
	}

	return NewToken("\\/", CONN), nil
}

func (l *Lexer) parseImplies() (Token, error) {
	nextR, _, err := l.buf.ReadRune()

	if err != nil {
		panic(err)
	}

	if nextR != '>' {
		return Token{}, ErrInvalidConnective
	}

	return NewToken("->", CONN), nil
}

func (l *Lexer) parseIfAndOnlyIf() (Token, error) {
	nextR, _, err := l.buf.ReadRune()

	if err != nil {
		return Token{}, err
	}

	if nextR != '-' {
		return Token{}, ErrInvalidConnective
	}

	nextR, _, err = l.buf.ReadRune()

	if err != nil {
		return Token{}, err
	}

	if nextR != '>' {
		return Token{}, ErrInvalidConnective
	}

	return NewToken("<->", CONN), nil
}

// TODO: remove bool
func (l *Lexer) parseXor(isEx bool) (Token, error) {
	if isEx {
		nextR, _, err := l.buf.ReadRune()

		if err != nil {
			return Token{}, err
		}

		if nextR != '=' {
			return Token{}, ErrInvalidConnective
		}

		return NewToken("!=", CONN), nil
	}

	nextR, _, err := l.buf.ReadRune()

	if err != nil {
		return Token{}, err
	}

	if nextR != '<' {
		return Token{}, ErrInvalidConnective
	}

	return NewToken("><", CONN), nil
}

func (l *Lexer) lexIdentifier(firstRune rune) (Token, error) {
	var builder strings.Builder

	builder.WriteRune(firstRune)

	for {
		r, _, err := l.buf.ReadRune()

		if err != nil {
			if err == io.EOF {
				return NewToken(builder.String(), IDENT), nil
			}
			return Token{}, err
		}

		if !unicode.IsLetter(r) {
			l.buf.UnreadRune()
			return NewToken(builder.String(), IDENT), nil
		}

		builder.WriteRune(r)
	}
}
