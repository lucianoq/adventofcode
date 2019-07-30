package main

import (
	"reflect"
	"testing"
)

func Test_complementary(t *testing.T) {
	tests := []struct {
		name  string
		list  []int
		total []int
		want  []int
	}{
		{
			name:  "U - {} = U",
			total: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			list:  []int{},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "{1} - {1} = {}",
			total: []int{1},
			list:  []int{1},
			want:  []int{},
		},
		{
			name:  "{1, 2} - {1} = {2}",
			total: []int{1, 2},
			list:  []int{1},
			want:  []int{2},
		},
		{
			name:  "{1...9} - {2,5,8} = {1,3,4,6,7,9}",
			total: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			list:  []int{2, 5, 8},
			want:  []int{1, 3, 4, 6, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complementary(tt.list, tt.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("complementary() = %v, want %v", got, tt.want)
			}
		})
	}
}
