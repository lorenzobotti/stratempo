package stratempo

import "testing"

func TestPow(t *testing.T) {
	cases := []struct {
		base, exponent, expected int
	}{
		{2, 3, 8},
		{2, 4, 16},
		{5, 1, 5},
		{3, 0, 1},
		{3, 3, 27},
	}

	for _, tc := range cases {
		got := powInt(tc.base, tc.exponent)
		if tc.expected != got {
			t.Fatalf(
				"%d ^ %d, expected %d, got %d",
				tc.base,
				tc.exponent,
				tc.expected,
				got,
			)
		}
	}
}
