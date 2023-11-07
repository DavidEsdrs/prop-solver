package eval

import "github.com/DavidEsdrs/prop-solver/expressions"

func Evaluate(a []bool, b []bool, op expressions.Op) []bool {
	switch op {
	case expressions.AND:
		return and(a, b)
	case expressions.OR:
		return or(a, b)
	case expressions.IF_AND_ONLY_IF:
		return ifAndOnlyIf(a, b)
	case expressions.XOR:
		return xor(a, b)
	case expressions.IMPLIES:
		return implies(a, b)
	default:
		return not(a)
	}
}
