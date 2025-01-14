package slice

import (
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// HasEmpty checks if the given slice contains any empty (zero) value.
func HasEmpty[T comparable](list []T) bool {
	var empty T
	for i := range list {
		if list[i] == empty {
			return true
		}
	}
	return false
}

// FirstEle returns the first element of the slice. If the slice is empty, it returns the zero value of type T.
func FirstEle[T any](list []T) T {
	var first T
	if len(list) > 0 {
		first = list[0]
	}
	return first
}

// Distinct returns a new slice containing only the unique elements from the input slice.
func Distinct[T comparable](list []T) []T {
	if len(list) <= 1 {
		return list
	}

	distinct := make(map[T]struct{}, 3*len(list)/2)
	result := make([]T, 0, len(list))
	for i := range list {
		_, ok := distinct[list[i]]
		if !ok {
			distinct[list[i]] = struct{}{}
			result = append(result, list[i])
		}
	}
	return result
}

// Merge concatenates two slices of the same type and returns the resulting slice.
func Merge[T constraints.Ordered](origin, target []T) []T {
	merged := make([]T, len(origin)+len(target))
	copy(merged, origin)
	copy(merged[len(origin):], target)

	return merged
}

// Contains checks if the given slice contains the specified target element.
func Contains[T comparable](l []T, target T) bool {
	for _, v := range l {
		if v == target {
			return true
		}
	}
	return false
}

// StrsContains checks if any string in the slice contains the target substring.
func StrsContains[K ~string](l []K, target K) bool {
	for i := range l {
		if strings.Contains(string(l[i]), string(target)) {
			return true
		}
	}
	return false
}

// Split divides the input slice into smaller slices of the specified size.
func Split[T any](list []T, size int) [][]T {
	if size <= 0 {
		size = 1
	}

	result := make([][]T, 0, len(list)/size+1)
	for i := 0; i < len(list); i += size {
		end := i + size
		if end > len(list) {
			end = len(list)
		}
		result = append(result, list[i:end])
	}
	return result
}

// IntsToStrings converts a slice of integers to a slice of strings.
func IntsToStrings[K constraints.Integer](l []K) []string {
	result := make([]string, 0, len(l))
	for _, v := range l {
		if v < 0 {
			result = append(result, strconv.FormatInt(int64(v), 10))
		} else {
			result = append(result, strconv.FormatUint(uint64(v), 10))
		}
	}
	return result
}

// StringsToInts converts a slice of strings to a slice of integers.
func StringsToInts[T constraints.Integer](list []string) []T {
	result := make([]T, 0, len(list))
	for _, v := range list {
		n, _ := strconv.ParseInt(v, 10, 64)
		result = append(result, T(n))
	}
	return result
}

// ToSet converts a slice to a set represented as a map with empty struct values.
func ToSet[T comparable](l []T) map[T]struct{} {
	m := make(map[T]struct{}, 2*len(l)/3)
	for i := range l {
		m[l[i]] = struct{}{}
	}
	return m
}

// ToSetFunc converts a slice to a map using a custom function to extract keys and values.
func ToSetFunc[T comparable, V any](l []any, f func(in any) (key T, value V)) map[T]V {
	m := make(map[T]V, len(l))
	for _, v := range l {
		key, value := f(v)
		m[key] = value
	}
	return m
}

// Reverse reverses the elements of the slice in place.
func Reverse[T any](nums []T) {
	if len(nums) == 0 {
		return
	}

	left, right := 0, len(nums)-1
	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// Columns applies a function to each element of the slice and returns a new slice with the results.
func Columns[T any, K any](l []T, f func(T) K) []K {
	list := make([]K, 0, len(l))
	for i := range l {
		list = append(list, f(l[i]))
	}
	return list
}

// MergeSortedAdjacent merges adjacent sorted numbers into groups.
// Example: [1, 2, 3, 5, 7, 8] => [[1, 3], [5], [7, 8]]
func MergeSortedAdjacent[K constraints.Integer](nums []K) [][]K {
	if len(nums) == 0 {
		return nil
	}

	result := make([][]K, 0)
	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= end+1 {
			end = nums[i]
			continue
		}
		cell := []K{start}
		if end != start {
			cell = append(cell, end)
		}
		result = append(result, cell)
		start, end = nums[i], nums[i]
	}
	cell := []K{start}
	if start != end {
		cell = append(cell, end)
	}
	result = append(result, cell)

	return result
}

// GroupSortedAdjacent groups adjacent sorted numbers into sub-slices.
// Example: [1, 2, 3, 5, 7, 8] => [[1, 2, 3], [5], [7, 8]]
func GroupSortedAdjacent[K constraints.Integer](nums []K) [][]K {
	if len(nums) == 0 {
		return nil
	}

	result := [][]K{}
	currentGroup := []K{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1]+1 {
			currentGroup = append(currentGroup, nums[i])
		} else {
			result = append(result, currentGroup)
			currentGroup = []K{nums[i]}
		}
	}

	result = append(result, currentGroup)
	return result
}

// RemoveEmpty removes all empty (zero) values from the slice and returns the result.
func RemoveEmpty[T comparable](l []T) []T {
	var empty T
	list := make([]T, 0, len(l))
	for i := range l {
		if l[i] == empty {
			continue
		}
		list = append(list, l[i])
	}
	return list
}

// ToIndexMap converts a slice to a map where the keys are the indices of the elements.
func ToIndexMap[T any](l []T) map[int]T {
	m := make(map[int]T, len(l))
	for i := range l {
		m[i] = l[i]
	}
	return m
}

// Repeat creates a slice containing the specified value repeated n times.
func Repeat[T any](v T, n int) []T {
	if n <= 0 {
		return nil
	}

	elems := make([]T, 0, n)
	for i := 0; i < n; i++ {
		elems = append(elems, v)
	}
	return elems
}

// Excludes returns a new slice containing elements from the source slice that are not present in the targets slice.
func Excludes[T comparable](source, targets []T) []T {
	if len(source) == 0 || len(targets) == 0 {
		return source
	}

	result := ToSet(targets)
	list := make([]T, 0, len(source))
	for i := range source {
		if _, ok := result[targets[i]]; ok {
			continue
		}
		list = append(list, source[i])
	}
	return list
}
