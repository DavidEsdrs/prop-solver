package lexer

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var zeroValueT T
		return zeroValueT, fmt.Errorf("empty array")
	}
	lastItem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return lastItem, nil
}

// returns wheter the tokens has valids separators
//
// don't check if there is separators with no content, for example:
// ["(", ")", "p", "^", "q"] returns true, even with the invalids "(" and ")"
func ValidateSep(tokens []Token) bool {
	stack := Stack[string]{}
	for _, t := range tokens {
		switch t.TType {
		case OPENING:
			stack.Push(t.TStr)
		case CLOSING:
			item, err := stack.Pop()
			if item != "(" || err != nil {
				return false
			}
		}
	}

	return len(stack) == 0
}
