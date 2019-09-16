package main

import (
	"reflect"
	"testing"
)

func Test_hash2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty",
			input: "",
			want:  "a2582a3a0e66e6e86e3812dcb672a272",
		},
		{
			name:  "AoC 2017",
			input: "AoC 2017",
			want:  "33efeb34ea91902bb2f59c9920caa6cd",
		},
		{
			name:  "1,2,3",
			input: "1,2,3",
			want:  "3efbe78a8d82f29979031a4aa0b16a9d",
		},
		{
			name:  "1,2,4",
			input: "1,2,4",
			want:  "63960835bcdc130f0b66d7ff4f6a5a8e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash2(tt.input); got != tt.want {
				t.Errorf("hash2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dense(t *testing.T) {
	tests := []struct {
		name  string
		input func() []uint8
		want  func() [block]uint8
	}{
		{
			name: "from assignment",
			input: func() []uint8 {
				long := make([]uint8, 256)
				list := []uint8{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}
				for i := 0; i < block; i++ {
					copy(long[i*block:], list)
				}
				return long
			},
			want: func() [block]uint8 {
				var out [block]uint8
				for i := 0; i < 16; i++ {
					out[i] = 64
				}
				return out
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dense(tt.input())
			wanted := tt.want()
			if !reflect.DeepEqual(got, wanted) {
				t.Errorf("dense() = %v, want %v", got, wanted)
			}
		})
	}
}
