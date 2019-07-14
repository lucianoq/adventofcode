package main

import "testing"

func TestEncodedString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line     string
		expected string
	}{
		{
			line:     `""`,
			expected: `"\"\""`,
		},
		{
			line:     `"abc"`,
			expected: `"\"abc\""`,
		},
		{
			line:     `"aaa\"aaa"`,
			expected: `"\"aaa\\\"aaa\""`,
		},
		{
			line:     `"\x27"`,
			expected: `"\"\\x27\""`,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()

			got := encodedString(tt.line)

			if got != tt.expected {
				t.Errorf("got: %s, expected: %s", got, tt.expected)
			}
		})
	}
}
