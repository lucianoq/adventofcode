package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	Size = 3
	code := m.Run()
	os.Exit(code)
}

func Test_crossBorder(t *testing.T) {
	tests := []struct {
		name  string
		p     P
		dir   Dir
		want  P
		want1 Dir
	}{
		{
			name:  "A to H",
			p:     P{-1, 3},
			dir:   N,
			want:  P{9, 0},
			want1: E,
		},
		{
			name:  "H to A",
			p:     P{9, -1},
			dir:   W,
			want:  P{0, 3},
			want1: S,
		},
		{
			name:  "B to F",
			p:     P{-1, 6},
			dir:   N,
			want:  P{11, 0},
			want1: N,
		},
		{
			name:  "F to B",
			p:     P{12, 0},
			dir:   S,
			want:  P{0, 6},
			want1: S,
		},
		{
			name:  "C to L",
			p:     P{3, 8},
			dir:   S,
			want:  P{5, 5},
			want1: W,
		},
		{
			name:  "L to C",
			p:     P{5, 6},
			dir:   E,
			want:  P{2, 8},
			want1: N,
		},
		{
			name:  "D to J",
			p:     P{5, 1},
			dir:   N,
			want:  P{4, 3},
			want1: E,
		},
		{
			name:  "D to J",
			p:     P{5, 2},
			dir:   N,
			want:  P{5, 3},
			want1: E,
		},
		{
			name:  "J to D",
			p:     P{5, 2},
			dir:   W,
			want:  P{6, 2},
			want1: S,
		},
		{
			name:  "E to K",
			p:     P{9, 3},
			dir:   S,
			want:  P{9, 2},
			want1: W,
		},
		{
			name:  "K to E",
			p:     P{9, 3},
			dir:   E,
			want:  P{8, 3},
			want1: N,
		},
		{
			name:  "G to I",
			p:     P{6, -1},
			dir:   W,
			want:  P{2, 3},
			want1: E,
		},
		{
			name:  "I to G",
			p:     P{0, 2},
			dir:   W,
			want:  P{8, 0},
			want1: E,
		},
		{
			name:  "M to N",
			p:     P{6, 6},
			dir:   E,
			want:  P{2, 8},
			want1: W,
		},
		{
			name:  "N to M",
			p:     P{0, 9},
			dir:   E,
			want:  P{8, 5},
			want1: W,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := crossBorder(tt.p, tt.dir)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("crossBorder() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("crossBorder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
