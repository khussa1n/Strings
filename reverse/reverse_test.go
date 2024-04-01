package reverse

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"string", "gnirts"},
		{"aaa  a", "a  aaa"},
		{"ABC", "CBA"},
	}

	for _, tc := range testCases {
		actual := ReverseString(tc.input)
		if actual != tc.expected {
			t.Errorf("ReverseString(%s) = %s; expected %s", tc.input, actual, tc.expected)
		}
	}
}
