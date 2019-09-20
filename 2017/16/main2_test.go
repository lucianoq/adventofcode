package main

import (
	"reflect"
	"testing"
)

func Test_applyPermutation(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		perm  []int
		want  []byte
	}{
		{
			name:  "empty",
			input: []byte{},
			perm:  []int{},
			want:  []byte{},
		},
		{
			name:  "1 elem",
			input: []byte{'a'},
			perm:  []int{0},
			want:  []byte{'a'},
		},
		{
			name:  "2 elems still",
			input: []byte{'a', 'b'},
			perm:  []int{0, 1},
			want:  []byte{'a', 'b'},
		},
		{
			name:  "2 elems changing",
			input: []byte{'a', 'b'},
			perm:  []int{1, 0},
			want:  []byte{'b', 'a'},
		},
		{
			name:  "many elems changing",
			input: []byte("abcde"),
			perm:  []int{4, 0, 1, 2, 3},
			want:  []byte("eabcd"),
		},
		// {
		// 	name:  "many elems changing",
		// 	input: []byte("abcde"),
		// 	perm:  []int{4, 0, 1, 2, 3},
		// 	want:  []byte("bcdea"),
		// },
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			applyPermutation(tt.input, tt.perm)

			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("got: %s, want: %s", tt.input, tt.want)
			}
		})
	}
}

func Test_Naive(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		perm  []int
		want  []byte
	}{
		{
			name:  "empty",
			input: []byte{},
			perm:  []int{},
			want:  []byte{},
		},
		{
			name:  "1 elem",
			input: []byte{'a'},
			perm:  []int{0},
			want:  []byte{'a'},
		},
		{
			name:  "2 elems still",
			input: []byte{'a', 'b'},
			perm:  []int{0, 1},
			want:  []byte{'a', 'b'},
		},
		{
			name:  "2 elems changing",
			input: []byte{'a', 'b'},
			perm:  []int{1, 0},
			want:  []byte{'b', 'a'},
		},
		{
			name:  "many elems changing",
			input: []byte("abcde"),
			perm:  []int{4, 0, 1, 2, 3},
			want:  []byte("eabcd"),
		},
		// {
		// 	name:  "many elems changing",
		// 	input: []byte("abcde"),
		// 	perm:  []int{4, 0, 1, 2, 3},
		// 	want:  []byte("bcdea"),
		// },
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			permute(tt.input, tt.perm)

			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("got: %s, want: %s", tt.input, tt.want)
			}
		})
	}
}
