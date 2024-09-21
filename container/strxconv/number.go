package strxconv

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func IntToStrings[T constraints.Integer](l []T) []string {
	strs := make([]string, len(l))
	for i := range l {
		if l[i] == 0 {
			strs[i] = "0"
		} else if l[i] < 0 {
			strs[i] = strconv.FormatInt(int64(l[i]), 10)
		} else {
			strs[i] = strconv.FormatUint(uint64(l[i]), 10)
		}
	}
	return strs
}

func FloatToStrings[T constraints.Float](s []T) []string {
	strs := make([]string, len(s))
	for i := range s {
		strs[i] = strconv.FormatFloat(float64(s[i]), 'f', -1, 64)
	}
	return strs
}
