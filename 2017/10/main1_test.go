package main

import (
	"reflect"
	"testing"
)

func Test_hash(t *testing.T) {
	type args struct {
		list []uint8
		pos  int
		l    int
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "1",
			args: args{
				list: []uint8{0, 1, 2, 3, 4},
				pos:  0,
				l:    3,
			},
			want: []uint8{2, 1, 0, 3, 4},
		},
		{
			name: "2",
			args: args{
				list: []uint8{2, 1, 0, 3, 4},
				pos:  3,
				l:    4,
			},
			want: []uint8{4, 3, 0, 1, 2},
		},
		{
			name: "3",
			args: args{
				list: []uint8{4, 3, 0, 1, 2},
				pos:  3,
				l:    1,
			},
			want: []uint8{4, 3, 0, 1, 2},
		},
		{
			name: "4",
			args: args{
				list: []uint8{4, 3, 0, 1, 2},
				pos:  1,
				l:    5,
			},
			want: []uint8{3, 4, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knot(tt.args.list, tt.args.pos, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("knot() = %v, want %v", got, tt.want)
			}
		})
	}
}
