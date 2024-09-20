package container

import (
	"testing"
)

func TestIsAlphabet(t *testing.T) {
	cases := []struct {
		in  byte
		exp bool
	}{
		{in: '0', exp: false},
		{in: '1', exp: false},
		{in: 'A', exp: true},
		{in: 'Z', exp: true},
		{in: 'a', exp: true},
		{in: 'z', exp: true},
		{in: ',', exp: false},
		{in: '?', exp: false},
		{in: ';', exp: false},
	}

	for _, v := range cases {
		if b := IsValidLetter(v.in); b != v.exp {
			t.Errorf("case %s not pass!", string(v.in))
		}
	}
}

func TestIsNumber(t *testing.T) {
	cases := []struct {
		in  byte
		exp bool
	}{
		{in: '1', exp: true},
		{in: '2', exp: true},
		{in: '0', exp: true},
		{in: 'a', exp: false},
		{in: 'Z', exp: false},
		{in: ',', exp: false},
		{in: '?', exp: false},
	}

	for _, v := range cases {
		if b := IsNumber(v.in); b != v.exp {
			t.Errorf("case %s not pass!", string(v.in))
		}
	}
}

func TestSwapCase(t *testing.T) {
	// Test case 1
	input1 := byte('A')
	expected1 := byte('a')
	output1 := SwapCase(input1)
	if output1 != expected1 {
		t.Errorf("SwapCase(%q) = %q, expected %q", input1, output1, expected1)
	}

	// Test case 2
	input2 := byte('a')
	expected2 := byte('A')
	output2 := SwapCase(input2)
	if output2 != expected2 {
		t.Errorf("SwapCase(%q) = %q, expected %q", input2, output2, expected2)
	}

	// Test case 3
	input3 := byte('0')
	expected3 := byte('0')
	output3 := SwapCase(input3)
	if output3 != expected3 {
		t.Errorf("SwapCase(%q) = %q, expected %q", input3, output3, expected3)
	}

	input4 := byte('|')
	expected4 := byte('|')
	output4 := SwapCase(input4)
	if output4 != expected4 {
		t.Errorf("SwapCase(%q) = %q, expected %q", input4, output4, expected4)
	}
}
