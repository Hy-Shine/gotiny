package str

import (
	"reflect"
	"testing"
)

func TestToNumber(t *testing.T) {
	// Testing parsing of a valid integer string
	intResult, err := ToNumber[int]("123")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if intResult != 123 {
		t.Errorf("Expected result to be 123, but got %v", intResult)
	}

	// Testing parsing of a valid float string
	floatResult, err := ToNumber[float64]("12.34")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if floatResult != 12.34 {
		t.Errorf("Expected result to be 12.34, but got %v", floatResult)
	}

	// Testing parsing of an invalid string
	_, err = ToNumber[int]("abc")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestToUpper(t *testing.T) {
	s := "test lower"
	ToUpper(&s)
	if s != "TEST LOWER" {
		t.Fatalf("meet %s, but expect %s", s, "TEST LOWER")
	}
}

func BenchmarkToUpper(b *testing.B) {
	s := "Hello World"
	for i := 0; i < b.N; i++ {
		ToUpper(&s)
	}
}

func TestToLower(t *testing.T) {
	s := "TEST lower"
	ToLower(&s)
	if s != "test lower" {
		t.Fatalf("meet %s, but expect %s", s, "test lower")
	}
}

func TestToBytes(t *testing.T) {
	tests := []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"Hello, World!",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			result := ToBytes(test)
			if !reflect.DeepEqual(result, []byte(test)) {
				t.Errorf("Expected %q, but got %q", test, string(result))
			}
		})
	}
}

func BenchmarkStringToBytesUnsafe(b *testing.B) {
	str := "0123456789,abcdefghijklmnopqrstuvwxyz,ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < b.N; i++ {
		_ = ToBytes(str)
	}
}

func _castBytes(str string) []byte {
	return []byte(str)
}

func BenchmarkStringToBytesCasting(b *testing.B) {
	str := "0123456789,abcdefghijklmnopqrstuvwxyz,ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < b.N; i++ {
		_ = _castBytes(str)
	}
}

func TestRepeatJoin(t *testing.T) {
	tests := []struct {
		str, sep string
		times    uint
		exp      string
	}{
		{"", "", 0, ""},
		{"", "", 1, ""},
		{"hello", "", 0, ""},
		{"hello", "", 1, "hello"},
		{"hello", "", 2, "hellohello"},
		{"hello", " ", 0, ""},
		{"hello", " ", 1, "hello"},
		{"hello", " ", 2, "hello hello"},
		{"hello", "world", 0, ""},
		{"hello", "world", 1, "hello"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			result := RepeatJoin(test.str, test.sep, test.times)
			if result != test.exp {
				t.Errorf("Expected %q, but got %q", test.exp, result)
			}
		})
	}
}
