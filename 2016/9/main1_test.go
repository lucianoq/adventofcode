package main

import (
	"reflect"
	"testing"
)

func Test_decompress(t *testing.T) {
	tests := []struct {
		in  []byte
		out []byte
	}{
		{
			in:  []byte("ADVENT"),
			out: []byte("ADVENT"),
		},
		{
			in:  []byte("A(1x5)BC"),
			out: []byte("ABBBBBC"),
		},
		{
			in:  []byte("(3x3)XYZ"),
			out: []byte("XYZXYZXYZ"),
		},
		{
			in:  []byte("A(2x2)BCD(2x2)EFG"),
			out: []byte("ABCBCDEFEFG"),
		},
		{
			in:  []byte("(6x1)(1x3)A"),
			out: []byte("(1x3)A"),
		},
		{
			in:  []byte("X(8x2)(3x3)ABCY"),
			out: []byte("X(3x3)ABC(3x3)ABCY"),
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.in), func(t *testing.T) {
			if got := decompress(tt.in); !reflect.DeepEqual(got, tt.out) {
				t.Errorf("decompress() = %v, want %v", got, tt.out)
			}
		})
	}
}
