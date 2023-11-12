package eval

import (
	"math"

	"github.com/DavidEsdrs/prop-solver/lexer"
)

type Evaluable struct {
	props   []string
	results map[string][]bool
}

func (e *Evaluable) Props() []string {
	return e.props
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
	props := map[string][]bool{}
	for _, n := range tokens {
		if n.TType == lexer.IDENT {
			props[n.TStr] = []bool{}
		}
	}
	res := []string{}
	for str := range props {
		res = append(res, str)
	}
	return res
}

func (e *Evaluable) Evaluate() {
	for _, n := range e.props {
		e.results[n] = []bool{}
	}
	propsQuant := len(e.results)
	possibleResults := int(math.Pow(2, float64(propsQuant)))
	result := e.generateResultArrays(possibleResults)
	result = fillResultArrays(result, possibleResults)
	e.fillResults(result)
}

func (e *Evaluable) fillResults(idents [][]bool) {
	for i := range idents {
		e.results[e.props[i]] = idents[i]
	}
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
