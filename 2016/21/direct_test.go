package main

import (
	"testing"
)

func TestSwapPosition(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			in:  "abcde",
			x:   4,
			y:   0,
			out: "ebcda",
		},
		{
			in:  "abcde",
			x:   2,
			y:   2,
			out: "abcde",
		},
		{
			in:  "abcde",
			x:   1,
			y:   3,
			out: "adcbe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			got := string(SwapPosition([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestSwapLetters(t *testing.T) {
	tests := []struct {
		in  string
		x   byte
		y   byte
		out string
	}{
		{
			in:  "ebcda",
			x:   'd',
			y:   'b',
			out: "edcba",
		},
		{
			in:  "abcde",
			x:   'a',
			y:   'd',
			out: "dbcae",
		},
		{
			in:  "abcde",
			x:   'c',
			y:   'c',
			out: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			got := string(SwapLetters([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			in:  "edcba",
			x:   0,
			y:   4,
			out: "abcde",
		},
		{
			in:  "abcde",
			x:   1,
			y:   3,
			out: "adcbe",
		},
		{
			in:  "abcde",
			x:   2,
			y:   2,
			out: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			got := string(Reverse([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			in:  "abcde",
			out: "edcba",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "a",
			out: "a",
		},
		{
			in:  "ab",
			out: "ba",
		},
		{
			in:  "abba",
			out: "abba",
		},
		{
			in:  "abcba",
			out: "abcba",
		},
		{
			in:  "12345",
			out: "54321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(reverse([]byte(tt.in)))
			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		out string
	}{
		{
			in:  "abcde",
			x:   1,
			out: "bcdea",
		},
		{
			in:  "abcde",
			x:   2,
			out: "cdeab",
		},
		{
			in:  "abcde",
			x:   3,
			out: "deabc",
		},
		{
			in:  "abcde",
			x:   4,
			out: "eabcd",
		},
		{
			in:  "abcde",
			x:   5,
			out: "abcde",
		},
		{
			in:  "abcde",
			x:   -1,
			out: "eabcd",
		},
		{
			in:  "abcde",
			x:   -2,
			out: "deabc",
		},
		{
			in:  "abcde",
			x:   -3,
			out: "cdeab",
		},
		{
			in:  "abcde",
			x:   -4,
			out: "bcdea",
		},
		{
			in:  "abcde",
			x:   -5,
			out: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(RotateLeft([]byte(tt.in), tt.x))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			in:  "bcdea",
			x:   1,
			y:   4,
			out: "bdeac",
		},
		{
			in:  "bdeac",
			x:   3,
			y:   0,
			out: "abdec",
		},
		{
			in:  "abcde",
			x:   0,
			y:   4,
			out: "bcdea",
		},
		{
			in:  "abcde",
			x:   4,
			y:   0,
			out: "eabcd",
		},
		{
			in:  "a",
			x:   0,
			y:   0,
			out: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(Move([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestRotateBasedOnLetter(t *testing.T) {
	tests := []struct {
		in  string
		x   byte
		out string
	}{
		{
			in:  "abdec",
			x:   'b',
			out: "ecabd",
		},
		{
			in:  "ecabd",
			x:   'd',
			out: "decab",
		},

		{
			in:  "abcdefgh",
			x:   'a',
			out: "habcdefg",
		},
		{
			in:  "abcdefgh",
			x:   'b',
			out: "ghabcdef",
		},
		{
			in:  "abcdefgh",
			x:   'c',
			out: "fghabcde",
		},
		{
			in:  "abcdefgh",
			x:   'd',
			out: "efghabcd",
		},
		{
			in:  "abcdefgh",
			x:   'e',
			out: "cdefghab",
		},
		{
			in:  "abcdefgh",
			x:   'f',
			out: "bcdefgha",
		},
		{
			in:  "abcdefgh",
			x:   'g',
			out: "abcdefgh",
		},
		{
			in:  "abcdefgh",
			x:   'h',
			out: "habcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(RotateBasedOnLetter([]byte(tt.in), tt.x))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}
