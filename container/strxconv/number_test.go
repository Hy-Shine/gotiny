package strxconv

import (
	"reflect"
	"testing"
)

func TestIntToStrings(t *testing.T) {
	// Testing for an empty slice
	input1 := []int{}
	expected1 := []string{}
	result1 := IntToStrings(input1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Testing for a slice with positive integers
	input2 := []int{0, 1, 2, 3, 1e5, -100}
	expected2 := []string{"0", "1", "2", "3", "100000", "-100"}
	result2 := IntToStrings(input2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}

	// Testing for a slice with negative integers
	input3 := []int{-1, -2, -3, 0}
	expected3 := []string{"-1", "-2", "-3", "0"}
	result3 := IntToStrings(input3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, result3)
	}
}

func TestFloatToStrings(t *testing.T) {
	cases := []struct {
		input    []float64
		expected []string
	}{
		{input: []float64{1.23, 4.56, 7.89}, expected: []string{"1.23", "4.56", "7.89"}},
		{input: []float64{-1.23, -4.56, -7.89}, expected: []string{"-1.23", "-4.56", "-7.89"}},
		{input: []float64{0.0, 0.0, 0.0}, expected: []string{"0", "0", "0"}},
	}

	for _, c := range cases {
		result := FloatToStrings(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Expected %v, but got %v", c.expected, result)
		}
	}
}
