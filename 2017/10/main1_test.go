package main

import (
	"reflect"
	"testing"
)

func Test_hash(t *testing.T) {
	type args struct {
		list []int
		pos  int
		l    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				list: []int{0, 1, 2, 3, 4},
				pos:  0,
				l:    3,
			},
			want: []int{2, 1, 0, 3, 4},
		},
		{
			name: "2",
			args: args{
				list: []int{2, 1, 0, 3, 4},
				pos:  3,
				l:    4,
			},
			want: []int{4, 3, 0, 1, 2},
		},
		{
			name: "3",
			args: args{
				list: []int{4, 3, 0, 1, 2},
				pos:  3,
				l:    1,
			},
			want: []int{4, 3, 0, 1, 2},
		},
		{
			name: "4",
			args: args{
				list: []int{4, 3, 0, 1, 2},
				pos:  1,
				l:    5,
			},
			want: []int{3, 4, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.list, tt.args.pos, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}