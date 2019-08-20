package main

import (
	"testing"
)

func TestInverseSwapPosition(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			out: "abcde",
			x:   4,
			y:   0,
			in:  "ebcda",
		},
		{
			out: "abcde",
			x:   2,
			y:   2,
			in:  "abcde",
		},
		{
			out: "abcde",
			x:   1,
			y:   3,
			in:  "adcbe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			got := string(InverseSwapPosition([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestInverseSwapLetters(t *testing.T) {
	tests := []struct {
		in  string
		x   byte
		y   byte
		out string
	}{
		{
			out: "ebcda",
			x:   'd',
			y:   'b',
			in:  "edcba",
		},
		{
			out: "abcde",
			x:   'a',
			y:   'd',
			in:  "dbcae",
		},
		{
			out: "abcde",
			x:   'c',
			y:   'c',
			in:  "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			got := string(InverseSwapLetters([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestInverseReverse(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			out: "edcba",
			x:   0,
			y:   4,
			in:  "abcde",
		},
		{
			out: "abcde",
			x:   1,
			y:   3,
			in:  "adcbe",
		},
		{
			out: "abcde",
			x:   2,
			y:   2,
			in:  "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			got := string(InverseReverse([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestInverseRotateLeft(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		out string
	}{
		{
			out: "abcde",
			x:   1,
			in:  "bcdea",
		},
		{
			out: "abcde",
			x:   2,
			in:  "cdeab",
		},
		{
			out: "abcde",
			x:   3,
			in:  "deabc",
		},
		{
			out: "abcde",
			x:   4,
			in:  "eabcd",
		},
		{
			out: "abcde",
			x:   5,
			in:  "abcde",
		},
		{
			out: "abcde",
			x:   -1,
			in:  "eabcd",
		},
		{
			out: "abcde",
			x:   -2,
			in:  "deabc",
		},
		{
			out: "abcde",
			x:   -3,
			in:  "cdeab",
		},
		{
			out: "abcde",
			x:   -4,
			in:  "bcdea",
		},
		{
			out: "abcde",
			x:   -5,
			in:  "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(InverseRotateLeft([]byte(tt.in), tt.x))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestInverseMove(t *testing.T) {
	tests := []struct {
		in  string
		x   int
		y   int
		out string
	}{
		{
			out: "bcdea",
			x:   1,
			y:   4,
			in:  "bdeac",
		},
		{
			out: "bdeac",
			x:   3,
			y:   0,
			in:  "abdec",
		},
		{
			out: "abcde",
			x:   0,
			y:   4,
			in:  "bcdea",
		},
		{
			out: "abcde",
			x:   4,
			y:   0,
			in:  "eabcd",
		},
		{
			out: "a",
			x:   0,
			y:   0,
			in:  "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(InverseMove([]byte(tt.in), tt.x, tt.y))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}

func TestInverseRotateBasedOnLetter(t *testing.T) {
	tests := []struct {
		in  string
		x   byte
		out string
	}{
		{
			out: "abdec",
			x:   'b',
			in:  "ecabd",
		},
		{
			out: "ecabd",
			x:   'd',
			in:  "decab",
		},

		{
			in:  "habcdefg",
			x:   'a',
			out: "abcdefgh",
		},
		{
			in:  "ghabcdef",
			x:   'b',
			out: "abcdefgh",
		},
		{
			in:  "fghabcde",
			x:   'c',
			out: "abcdefgh",
		},
		{
			in:  "efghabcd",
			x:   'd',
			out: "abcdefgh",
		},
		{
			in:  "cdefghab",
			x:   'e',
			out: "abcdefgh",
		},
		{
			in:  "bcdefgha",
			x:   'f',
			out: "abcdefgh",
		},
		{
			in:  "abcdefgh",
			x:   'g',
			out: "abcdefgh",
		},
		{
			in:  "habcdefg",
			x:   'h',
			out: "abcdefgh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := string(InverseRotateBasedOnLetter([]byte(tt.in), tt.x))

			if got != tt.out {
				t.Errorf("got: %s, wanted: %s", got, tt.out)
			}
		})
	}
}
