package length

import "testing"

func TestStringLength(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"String", 6},
		{"Len", 3},
		{"", 0},
	}

	for _, tc := range testCases {
		actual := StringLength(tc.input)
		if actual != tc.expected {
			t.Errorf("StringLength(%s) = %d; expected %d", tc.input, actual, tc.expected)
		}
	}
}
