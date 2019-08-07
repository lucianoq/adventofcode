package main

import "testing"

func Test_computeSize(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "(3x3)XYZ",
			out: 9,
		},
		{
			in:  "X(8x2)(3x3)ABCY",
			out: 20,
		},
		{
			in:  "(27x12)(20x12)(13x14)(7x10)(1x12)A",
			out: 241920,
		},
		{
			in:  "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			out: 445,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := computeSize([]byte(tt.in)); got != tt.out {
				t.Errorf("computeSize() = %v, want %v", got, tt.out)
			}
		})
	}
}
