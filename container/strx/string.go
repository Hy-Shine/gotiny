package str

import (
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func Contains[T ~string](strs []T, target T) bool {
	for i := range strs {
		if strs[i] == target {
			return true
		}
	}
	return false
}

// Reverse reverses s.
func Reverse(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func Concat(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

func ConcatFunc(strs []string, sep string, f func(string) (string, bool)) string {
	newStrs := make([]string, 0, len(strs))
	for i := range strs {
		if s, b := f(strs[i]); b {
			newStrs = append(newStrs, s)
		}
	}
	return strings.Join(newStrs, sep)
}

func ToNumber[T constraints.Integer | constraints.Float](str string) (T, error) {
	f, err := strconv.ParseFloat(str, 64)
	return T(f), err
}

func ToUpper(s *string) {
	if s != nil {
		*s = strings.ToUpper(*s)
	}
}

func ToLower(s *string) {
	if s != nil {
		*s = strings.ToLower(*s)
	}
}

func IntToStrings[T constraints.Integer](s []T) []string {
	strs := make([]string, len(s))
	for i := range s {
		strs[i] = strconv.FormatInt(int64(s[i]), 10)
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

func ToBytes(str string) []byte {
	if len(str) == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}
