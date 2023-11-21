package php_test

import (
	"testing"

	"github.com/reiver/go-php"
)

func TestIsNumeric_uint8(t *testing.T) {

	tests := []struct{
		Value uint8
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_uint16(t *testing.T) {

	tests := []struct{
		Value uint16
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_uint32(t *testing.T) {

	tests := []struct{
		Value uint32
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_uint64(t *testing.T) {

	tests := []struct{
		Value uint64
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_int8(t *testing.T) {

	tests := []struct{
		Value int8
	}{
		{
			Value: -127,
		},
		{
			Value: -123,
		},
		{
			Value: -5,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_int16(t *testing.T) {

	tests := []struct{
		Value int16
	}{
		{
			Value: -255,
		},
		{
			Value: -127,
		},
		{
			Value: -123,
		},
		{
			Value: -5,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_int32(t *testing.T) {

	tests := []struct{
		Value int32
	}{
		{
			Value: -255,
		},
		{
			Value: -127,
		},
		{
			Value: -123,
		},
		{
			Value: -5,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}

}

func TestIsNumeric_int64(t *testing.T) {

	tests := []struct{
		Value int64
	}{
		{
			Value: -255,
		},
		{
			Value: -127,
		},
		{
			Value: -123,
		},
		{
			Value: -5,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %d", test.Value)
			continue
		}
	}
}

func TestIsNumeric_float32(t *testing.T) {

	tests := []struct{
		Value float32
	}{
		{
			Value: -255.0,
		},
		{
			Value: -127.0,
		},
		{
			Value: -123.0,
		},
		{
			Value: -12.34,
		},
		{
			Value: -5.0,
		},
		{
			Value: -1.0,
		},
		{
			Value: 0.0,
		},
		{
			Value: 1.0,
		},
		{
			Value: 5.0,
		},
		{
			Value: 12.34,
		},
		{
			Value: 123.0,
		},
		{
			Value: 255.0,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %v", test.Value)
			continue
		}
	}
}

func TestIsNumeric_float64(t *testing.T) {

	tests := []struct{
		Value float64
	}{
		{
			Value: -255.0,
		},
		{
			Value: -127.0,
		},
		{
			Value: -123.0,
		},
		{
			Value: -12.34,
		},
		{
			Value: -5.0,
		},
		{
			Value: -1.0,
		},
		{
			Value: 0.0,
		},
		{
			Value: 1.0,
		},
		{
			Value: 5.0,
		},
		{
			Value: 12.34,
		},
		{
			Value: 123.0,
		},
		{
			Value: 255.0,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %v", test.Value)
			continue
		}
	}
}

func TestIsNumeric_complex64(t *testing.T) {

	tests := []struct{
		Value complex64
	}{
		{
			Value: -255,
		},
		{
			Value: -127,
		},
		{
			Value: -123,
		},
		{
			Value: -12.34i,
		},
		{
			Value: -5,
		},
		{
			Value: -1i,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 1i,
		},
		{
			Value: 5,
		},
		{
			Value: 12.34i,
		},
		{
			Value: 123,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %v", test.Value)
			continue
		}
	}
}

func TestIsNumeric_complex128(t *testing.T) {

	tests := []struct{
		Value complex128
	}{
		{
			Value: -255,
		},
		{
			Value: -127,
		},
		{
			Value: -123,
		},
		{
			Value: -12.34i,
		},
		{
			Value: -5,
		},
		{
			Value: -1i,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 1i,
		},
		{
			Value: 5,
		},
		{
			Value: 12.34i,
		},
		{
			Value: 123,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %v", test.Value)
			continue
		}
	}
}

func TestIsNumeric_uint(t *testing.T) {

	tests := []struct{
		Value uint
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %v", test.Value)
			continue
		}
	}
}

func TestIsNumeric_int(t *testing.T) {

	tests := []struct{
		Value int
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: 5,
		},
		{
			Value: 123,
		},
		{
			Value: 127,
		},
		{
			Value: 255,
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %v", test.Value)
			continue
		}
	}

}

func TestIsNumeric_string(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "-123.45",
		},
		{
			Value: "-123",
		},
		{
			Value: "-12.34",
		},
		{
			Value: "-5",
		},
		{
			Value: "-1",
		},
		{
			Value: "0",
		},
		{
			Value: "1",
		},
		{
			Value: "5",
		},
		{
			Value: "12.34",
		},
		{
			Value: "123",
		},
		{
			Value: "123.45",
		},
		{
			Value: "127",
		},
		{
			Value: "255",
		},
	}

	for testNumber, test := range tests {
		expected := true
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}

func TestIsNumeric_string_fail(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "apple",
		},
		{
			Value: "1-23",
		},
		{
			Value: "123-",
		},
		{
			Value: ".123",
		},
		{
			Value: "123.",
		},
		{
			Value: "1.2.3",
		},
	}

	for testNumber, test := range tests {
		expected := false
		actual := php.IsNumeric(test.Value)

		if expected != actual {
			t.Errorf("For test #%d, the actual return value for php.IsNumeric() on a %T is not what was expected.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
