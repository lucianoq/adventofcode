package main

import "testing"

func Test_valid(t *testing.T) {
	tests := []struct {
		name  string
		floor uint16
		want  bool
	}{
		{
			name:  "floor full. Valid",
			floor: 0xFFFF,
			want:  true,
		},
		{
			name:  "floor empty. Valid",
			floor: 0x0000,
			want:  true,
		},
		{
			name:  "one chip missing -> one generator alone. Valid",
			floor: 0xFFFE,
			want:  true,
		},
		{
			name:  "one generator missing -> one chip alone. Not Valid",
			floor: 0xFEFF,
			want:  false,
		},
		{
			name:  "one chip with someone else generator. Not Valid",
			floor: 0x0201,
			want:  false,
		},
		{
			name:  "one chip with its and someone else generator. Valid",
			floor: 0x0301,
			want:  true,
		},
		{
			name:  "one chip paired with its generator. Valid",
			floor: 0x0101,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid(tt.floor); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
