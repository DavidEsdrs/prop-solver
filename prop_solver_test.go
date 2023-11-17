package prop_solver_test

import (
	"testing"

	prop_solver "github.com/DavidEsdrs/prop-solver"
)

var (
	andResult     = []bool{true, false, false, false}
	orResult      = []bool{true, true, true, false}
	xorResult     = []bool{false, true, true, false}
	impliesResult = []bool{true, false, true, true}
	ifAndOnlyIf   = []bool{true, false, false, true}
)

func TestAnd(t *testing.T) {
	_, res := prop_solver.Solve("p /\\ q")

	if !equalsSlice[bool](res, andResult) {
		t.Errorf("result doesn't match! fail")
	}
}

func TestOr(t *testing.T) {
	_, res := prop_solver.Solve("p \\/ q")

	if !equalsSlice[bool](res, orResult) {
		t.Errorf("result doesn't match! fail")
	}
}

func TestXor(t *testing.T) {
	_, res := prop_solver.Solve("p != q")

	if !equalsSlice[bool](res, xorResult) {
		t.Errorf("result doesn't match! fail")
	}
}

func TestImplies(t *testing.T) {
	_, res := prop_solver.Solve("p -> q")

	if !equalsSlice[bool](res, impliesResult) {
		t.Errorf("result doesn't match! fail")
	}
}

func TestIfAndOnlyIf(t *testing.T) {
	_, res := prop_solver.Solve("p <-> q")

	if !equalsSlice[bool](res, ifAndOnlyIf) {
		t.Errorf("result doesn't match! fail")
	}
}

func equalsSlice[T comparable](sliceA, sliceB []T) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}
