package main

import (
	"reflect"
	"testing"
)

const (
	T = true
	F = false
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name    string
		squares []Image
		want    Image
	}{
		{
			name: "empty",
			squares: []Image{
				{},
			},
			want: Image{},
		},
		{
			name: "4 squares 2x2",
			squares: []Image{
				{T, T, F, F},
				{T, T, F, F},
				{T, T, F, F},
				{T, T, F, F},
			},
			want: Image{
				T, T, T, T, F, F, F, F, T, T, T, T, F, F, F, F,
			},
		},
		{
			name: "4 squares 3x3",
			squares: []Image{
				{T, T, T, F, F, F, T, T, T},
				{T, T, T, F, F, F, T, T, T},
				{T, T, T, F, F, F, T, T, T},
				{T, T, T, F, F, F, T, T, T},
			},
			want: Image{
				T, T, T, T, T, T, F, F, F, F, F, F, T, T, T, T, T, T, T, T, T, T, T, T, F, F, F, F, F, F, T, T, T, T, T, T,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.squares); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
