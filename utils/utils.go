package utils

func EqualsSlice[T comparable](sliceA, sliceB []T) bool {
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

type MapFunc[T any, Y any] func(T) Y

// Maps an array from a source arr - uses the given map function f
func Map[T any, Y any](arr []T, f MapFunc[T, Y]) []Y {
	res := []Y{}
	for _, n := range arr {
		res = append(res, f(n))
	}
	return res
}
