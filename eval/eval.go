package eval

import (
	"math"

	"github.com/DavidEsdrs/prop-solver/lexer"
)

type Evaluable struct {
	props   []string
	results map[string][]bool
}

func (e *Evaluable) Result() map[string][]bool {
	return e.results
}

func NewEvaluable(tokens []*lexer.Token) Evaluable {
	return Evaluable{
		props:   getIdentifiers(tokens),
		results: make(map[string][]bool),
	}
}

func getIdentifiers(tokens []*lexer.Token) []string {
	props := []string{}

	for _, n := range tokens {
		if n.TType == lexer.IDENT {
			props = append(props, n.TStr)
		}
	}

	return props
}

func (e *Evaluable) Evaluate() [][]bool {
	idents := e.generate()
	e.fillResults(idents)
	return idents
}

func (e *Evaluable) fillResults(idents [][]bool) {
	for i := range idents {
		e.results[e.props[i]] = idents[i]
	}
}

func (e *Evaluable) generate() [][]bool {
	propsQuant := len(e.props)
	possibleResults := int(math.Pow(2, float64(propsQuant)))
	result := e.generateResultArrays(possibleResults)
	result = fillResultArrays(result, possibleResults)
	return result
}

func (e *Evaluable) generateResultArrays(length int) [][]bool {
	result := make([][]bool, len(e.props))
	for i := range result {
		result[i] = make([]bool, length)
	}
	return result
}

func fillResultArrays(input [][]bool, length int) [][]bool {
	lastLen := length / 2 // always even
	for i := range input {
		fill(input[i], lastLen)
		lastLen /= 2
	}
	return input
}

func fill(input []bool, lengthWithTrues int) {
	count := lengthWithTrues
	filler := true
	for i := range input {
		input[i] = filler

		count--

		if count == 0 {
			count = lengthWithTrues
			filler = !filler
		}
	}
}

func not(input []bool) []bool {
	result := make([]bool, len(input))
	for i, n := range input {
		result[i] = !n
	}
	return result
}

func and(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] && b[i]
	}
	return result
}

func or(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] || b[i]
	}
	return result
}

func xor(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] != b[i]
	}
	return result
}

func ifAndOnlyIf(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = a[i] == b[i]
	}
	return result
}

func implies(a []bool, b []bool) []bool {
	length := len(a)
	result := make([]bool, length)
	for i := 0; i < length; i++ {
		result[i] = !a[i] || b[i]
	}
	return result
}
