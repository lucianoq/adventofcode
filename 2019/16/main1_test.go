package main

import "testing"

func Test_pattern1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		i    int
		j    int
		want int
	}{
		{
			name: "",
			i:    0,
			j:    0,
			want: 1,
		},
		{
			name: "",
			i:    0,
			j:    1,
			want: 0,
		},
		{
			name: "",
			i:    0,
			j:    2,
			want: -1,
		},
		{
			name: "",
			i:    0,
			j:    3,
			want: 0,
		},
		{
			name: "",
			i:    0,
			j:    4,
			want: 1,
		},

		{
			name: "",
			i:    1,
			j:    0,
			want: 0,
		},
		{
			name: "",
			i:    1,
			j:    1,
			want: 1,
		},
		{
			name: "",
			i:    1,
			j:    2,
			want: 1,
		},
		{
			name: "",
			i:    1,
			j:    3,
			want: 0,
		},
		{
			name: "",
			i:    1,
			j:    4,
			want: 0,
		},

		{
			name: "",
			i:    2,
			j:    0,
			want: 0,
		},
		{
			name: "",
			i:    2,
			j:    1,
			want: 0,
		},
		{
			name: "",
			i:    2,
			j:    2,
			want: 1,
		},
		{
			name: "",
			i:    2,
			j:    3,
			want: 1,
		},
		{
			name: "",
			i:    2,
			j:    4,
			want: 1,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := pattern(tt.i, tt.j); got != tt.want {
				t.Errorf("pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
