package slice

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMerge(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7}
	list2 := []int{4, 5, 6, 7}

	result := Merge(list1, list2)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 4, 5, 6, 7}, result)
}

func BenchmarkListMerge(b *testing.B) {
	list1 := make([]int, 0, 500)
	for i := 0; i < 1000; i++ {
		list1 = append(list1, i)
	}
	list2 := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for i := 0; i < b.N; i++ {
		_ = Merge(list1, list2)
	}
}

func TestStrsContains(t *testing.T) {
	cases := []struct {
		name   string
		l      []string
		target string
		want   bool
	}{
		{
			name:   "empty list",
			l:      []string{},
			target: "foo",
			want:   false,
		},
		{
			name:   "target found in first element",
			l:      []string{"foo", "bar", "baz"},
			target: "foo",
			want:   true,
		},
		{
			name:   "target found in last element",
			l:      []string{"foo", "bar", "baz"},
			target: "baz",
			want:   true,
		},
		{
			name:   "target found in middle element",
			l:      []string{"foo", "bar", "baz"},
			target: "bar",
			want:   true,
		},
		{
			name:   "target not found",
			l:      []string{"foo", "bar", "baz"},
			target: "qux",
			want:   false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := StrsContains(c.l, c.target)
			if got != c.want {
				t.Errorf("StrsContains(%q, %q) == %t, want %t", c.l, c.target, got, c.want)
			}
		})
	}
}

func TestListSplit(t *testing.T) {
	// Test splitting a list of integers with a size that evenly divides the list length
	intList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	intSize := 2
	intExpected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	intResult := Split(intList, intSize)
	if !reflect.DeepEqual(intResult, intExpected) {
		t.Errorf("Expected %v but got %v", intExpected, intResult)
	}

	// Test splitting a list of strings with a size that does not evenly divide the list length
	strList := []string{"a", "b", "c", "d", "e"}
	strSize := 3
	strExpected := [][]string{{"a", "b", "c"}, {"d", "e"}}
	strResult := Split(strList, strSize)
	if !reflect.DeepEqual(strResult, strExpected) {
		t.Errorf("Expected %v but got %v", strExpected, strResult)
	}

	// Test splitting an empty list
	emptyList := []int{}
	emptySize := 2
	emptyExpected := [][]int{}
	emptyResult := Split(emptyList, emptySize)
	if !reflect.DeepEqual(emptyResult, emptyExpected) {
		t.Errorf("Expected %v but got %v", emptyExpected, emptyResult)
	}

	// Test splitting a list with a negative size
	negativeList := []int{1, 2, 3, 4}
	negativeSize := -1
	negativeExpected := [][]int{{1}, {2}, {3}, {4}}
	negativeResult := Split(negativeList, negativeSize)
	if !reflect.DeepEqual(negativeResult, negativeExpected) {
		t.Errorf("Expected %v but got %v", negativeExpected, negativeResult)
	}
}

func BenchmarkListSplit(b *testing.B) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	size := 5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split(list, size)
	}
}

func TestMapRW(t *testing.T) {
	m := make(map[int]struct{})
	m1 := make(map[int]struct{})

	fmt.Println(reflect.DeepEqual(m, m1))
}

func TestToSet(t *testing.T) {
	t.Run("Empty input slice", func(t *testing.T) {
		input := []int{}
		result := ToSet(input)
		if len(result) != 0 {
			t.Errorf("Expected an empty map but got a map with length %d", len(result))
		}
	})

	t.Run("Slice with unique elements", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		result := ToSet(input)
		expectedLength := len(input)
		if len(result) != expectedLength {
			t.Errorf("Expected a map with %d elements but got a map with length %d", expectedLength, len(result))
		}
	})

	t.Run("Slice with duplicate elements", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3}
		result := ToSet(input)
		if reflect.DeepEqual(result, map[int]struct{}{1: {}, 2: {}, 3: {}}) {
			t.Errorf("Expected a map with %d elements but got a map with length %d", 3, len(result))
		}
	})
}

func TestReverse(t *testing.T) {
	// empty string slice
	empty := []string{}
	Reverse(empty)
	if !reflect.DeepEqual(empty, []string{}) {
		t.Errorf("Expected [] but got %v", empty)
	}

	l0 := []int{1, 2, 3, 4}
	Reverse(l0)
	if !reflect.DeepEqual(l0, []int{4, 3, 2, 1}) {
		t.Errorf("Expected [4, 3, 2, 1] but got %v", l0)
	}

	l1 := []string{"a", "b", "c", "d"}
	Reverse(l1)
	if !reflect.DeepEqual(l1, []string{"d", "c", "b", "a"}) {
		t.Errorf("Expected [d, c, b, a] but got %v", l1)
	}

	l2 := []bool{true, false, true, false}
	Reverse(l2)
	if !reflect.DeepEqual(l2, []bool{false, true, false, true}) {
		t.Errorf("Expected [false, true, false, true] but got %v", l2)
	}

	l3 := []byte{'a', 'b', 'c', 'd', 'e'}
	Reverse(l3)
	if !reflect.DeepEqual(l3, []byte{'e', 'd', 'c', 'b', 'a'}) {
		t.Errorf("Expected [d, c, b, a] but got %v", l3)
	}
}

func TestColumns(t *testing.T) {
	// Test case 1: Empty slice
	input1 := []int{}
	expected1 := []string{}
	result1 := Columns(input1, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, but got %v", expected1, result1)
	}

	// Test case 2: Slice with positive integers
	input2 := []int{1, 2, 3, 4, 5}
	expected2 := []string{"1", "2", "3", "4", "5"}
	result2 := Columns(input2, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, but got %v", expected2, result2)
	}

	// Test case 3: Slice with negative integers
	input3 := []int{-5, -4, -3, -2, -1}
	expected3 := []string{"-5", "-4", "-3", "-2", "-1"}
	result3 := Columns(input3, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, but got %v", expected3, result3)
	}

	// Test case 4: Slice with mixed integers
	input4 := []int{-5, 0, 3, -2, 8}
	expected4 := []string{"-5", "0", "3", "-2", "8"}
	result4 := Columns(input4, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Test case 4 failed. Expected %v, but got %v", expected4, result4)
	}

	// Test case 5: Slice with custom transformation
	input5 := []string{"apple", "banana", "cherry", "date"}
	expected5 := []string{"A", "B", "C", "D"}
	result5 := Columns(input5, func(s string) string {
		return strings.ToUpper(string(s[0]))
	})
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf("Test case 5 failed. Expected %v, but got %v", expected5, result5)
	}
}

func TestGroupSortedAdjacent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Empty slice",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "Single element",
			args: args{nums: []int{1}},
			want: [][]int{{1}},
		},
		{
			name: "Sorted consecutive numbers",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "Non-consecutive numbers",
			args: args{nums: []int{1, 3, 5, 7}},
			want: [][]int{{1}, {3}, {5}, {7}},
		},
		{
			name: "Multiple groups",
			args: args{nums: []int{1, 2, 2, 3, 5, 6, 7, 8, 10, 10}},
			want: [][]int{{1, 2, 2, 3}, {5, 6, 7, 8}, {10, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupSortedAdjacent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupSortedAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortedAdjacent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Empty slice",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "Single element",
			args: args{nums: []int{1}},
			want: [][]int{{1}},
		},
		{
			name: "Sorted consecutive numbers",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: [][]int{{1, 5}},
		},
		{
			name: "Non-consecutive numbers",
			args: args{nums: []int{1, 3, 5, 7}},
			want: [][]int{{1}, {3}, {5}, {7}},
		},
		{
			name: "Multiple groups",
			args: args{nums: []int{1, 2, 2, 3, 5, 6, 7, 8, 10, 10}},
			want: [][]int{{1, 3}, {5, 8}, {10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedAdjacent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortedAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
