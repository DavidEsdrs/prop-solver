package expressions

import "fmt"

type Op int

const (
	NOT Op = 1 << iota
	AND
	OR
	XOR
	IMPLIES
	IF_AND_ONLY_IF
	IDENT
)

type Term struct {
	name  string
	value bool
}

type Expression struct {
	termA     Term
	termB     Term
	operation Op
}

func (exp Expression) String() string {
	switch exp.operation {
	case AND:
		return fmt.Sprintf("%v ^ %v", exp.termA.name, exp.termB.name)
	case OR:
		return fmt.Sprintf("%v \\/ %v", exp.termA.name, exp.termB.name)
	case XOR:
		return fmt.Sprintf("%v <> %v", exp.termA.name, exp.termB.name)
	case IMPLIES:
		return fmt.Sprintf("%v -> %v", exp.termA.name, exp.termB.name)
	case IF_AND_ONLY_IF:
		return fmt.Sprintf("%v <-> %v", exp.termA.name, exp.termB.name)
	default:
		panic("eval failed")
	}
}

func (exp Expression) Eval() bool {
	switch exp.operation {
	case AND:
		return exp.termA.value && exp.termB.value
	case OR:
		return exp.termA.value || exp.termB.value
	case XOR:
		return exp.termA.value != exp.termB.value
	case IMPLIES:
		return !exp.termA.value || exp.termB.value
	case IF_AND_ONLY_IF:
		return exp.termA.value == exp.termB.value
	default:
		panic("eval failed")
	}
}

func GetConnectiveType(str string) Op {
	switch str {
	case "~", "!":
		return NOT
	case "^", "/\\":
		return AND
	case "\\/":
		return OR
	case "<>", "!=", "/V":
		return XOR
	case "<->", "<=>":
		return IF_AND_ONLY_IF
	case "->", "=>":
		return IMPLIES
	default:
		panic("unknown connective")
	}
}
