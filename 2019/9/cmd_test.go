package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCmd_Modes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		cmd   Cmd
		arity int
		want  []Mode
	}{
		{
			cmd:   2,
			arity: 0,
			want:  []Mode{},
		},
		{
			cmd:   2,
			arity: 1,
			want:  []Mode{0},
		},
		{
			cmd:   2,
			arity: 2,
			want:  []Mode{0, 0},
		},
		{
			cmd:   2,
			arity: 3,
			want:  []Mode{0, 0, 0},
		},
		{
			cmd:   102,
			arity: 1,
			want:  []Mode{1},
		},
		{
			cmd:   202,
			arity: 1,
			want:  []Mode{2},
		},
		{
			cmd:   102,
			arity: 3,
			want:  []Mode{1, 0, 0},
		},
		{
			cmd:   1102,
			arity: 3,
			want:  []Mode{1, 1, 0},
		},
		{
			cmd:   11102,
			arity: 3,
			want:  []Mode{1, 1, 1},
		},
		{
			cmd:   12102,
			arity: 3,
			want:  []Mode{1, 2, 1},
		},
		{
			cmd:   12202,
			arity: 3,
			want:  []Mode{2, 2, 1},
		},
		{
			cmd:   22202,
			arity: 3,
			want:  []Mode{2, 2, 2},
		},
		{
			cmd:   2202,
			arity: 3,
			want:  []Mode{2, 2, 0},
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(fmt.Sprintf("%d/%d", tt.cmd, tt.arity), func(t *testing.T) {
			t.Parallel()

			if got := tt.cmd.Modes(tt.arity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Modes() = %v, want %v", got, tt.want)
			}
		})
	}
}
