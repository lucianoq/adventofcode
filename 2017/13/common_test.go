package main

import "testing"

func TestScanner_IsZero(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		width int
		pico  int
		want  bool
	}{
		{
			name:  "w2-p0",
			width: 2,
			pico:  0,
			want:  true,
		},
		{
			name:  "w2-p1",
			width: 2,
			pico:  1,
			want:  false,
		},
		{
			name:  "w2-p2",
			width: 2,
			pico:  2,
			want:  true,
		},
		{
			name:  "w2-p3",
			width: 2,
			pico:  3,
			want:  false,
		},
		{
			name:  "w2-p4",
			width: 2,
			pico:  4,
			want:  true,
		},
		{
			name:  "w3-p0",
			width: 3,
			pico:  0,
			want:  true,
		},
		{
			name:  "w3-p1",
			width: 3,
			pico:  1,
			want:  false,
		},
		{
			name:  "w3-p2",
			width: 3,
			pico:  2,
			want:  false,
		},
		{
			name:  "w3-p3",
			width: 3,
			pico:  3,
			want:  false,
		},
		{
			name:  "w3-p4",
			width: 3,
			pico:  4,
			want:  true,
		},
		{
			name:  "w4-p0",
			width: 4,
			pico:  0,
			want:  true,
		},
		{
			name:  "w4-p1",
			width: 4,
			pico:  1,
			want:  false,
		},
		{
			name:  "w4-p2",
			width: 4,
			pico:  2,
			want:  false,
		},
		{
			name:  "w4-p3",
			width: 4,
			pico:  3,
			want:  false,
		},
		{
			name:  "w4-p4",
			width: 4,
			pico:  4,
			want:  false,
		},
		{
			name:  "w4-p5",
			width: 4,
			pico:  5,
			want:  false,
		},
		{
			name:  "w4-p6",
			width: 4,
			pico:  6,
			want:  true,
		},
		{
			name:  "w4-p7",
			width: 4,
			pico:  7,
			want:  false,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Scanner(tt.width)
			if got := s.IsZero(tt.pico); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
