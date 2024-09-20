package cal

import "golang.org/x/exp/constraints"

func If[T any](b bool, trueVal, falseVal T) T {
	if b {
		return trueVal
	}
	return falseVal
}

func Max[V constraints.Ordered](v1, v2 V) V {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min[V constraints.Ordered](v1, v2 V) V {
	if v1 < v2 {
		return v1
	}
	return v2
}
