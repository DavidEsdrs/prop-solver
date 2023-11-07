package eval

import (
	"math"

	"github.com/DavidEsdrs/prop-solver/expressions"
)

type Evaluable struct {
	props     []string
	operation expressions.Op
}

func NewEvaluable(simpleProps []string, operation expressions.Op) Evaluable {
	return Evaluable{simpleProps, operation}
}

func (e Evaluable) Eval() ([]bool, []bool, []bool) {
	propsQuant := len(e.props)
	possibleResults := int(math.Pow(2, float64(propsQuant)))
	result := e.generateResultArrays(possibleResults)
	result = fillResultArrays(result, possibleResults)
	if e.operation == expressions.NOT {
		return result[0], nil, e.evaluate(result) // TODO: Get all results
	}
	return result[0], result[1], e.evaluate(result) // TODO: Get all results
}

func (e Evaluable) generateResultArrays(length int) [][]bool {
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

func (e Evaluable) evaluate(input [][]bool) []bool {
	switch e.operation {
	case expressions.AND:
		return and(input[0], input[1])
	case expressions.OR:
		return or(input[0], input[1])
	case expressions.IF_AND_ONLY_IF:
		return ifAndOnlyIf(input[0], input[1])
	case expressions.XOR:
		return xor(input[0], input[1])
	case expressions.IMPLIES:
		return implies(input[0], input[1])
	default:
		return not(input[0])
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
