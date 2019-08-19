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
