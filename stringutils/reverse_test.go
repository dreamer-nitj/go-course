package stringutils

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"Hello world", "dlrow olleH"},
		{"potato", "otatop"},
	}
	for _, c := range cases {
		got := Reverse(c.input)
		if got != c.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", c.input, got, c.expected)
		}
	}
}
