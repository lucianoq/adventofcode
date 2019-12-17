package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_extract8digits(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		list string
		want string
	}{
		{
			name: "",
			list: "00000001234567890",
			want: "00000001",
		},
		{
			name: "",
			list: "0000001234567890",
			want: "00000123",
		},
		{
			name: "",
			list: "0000012345678901234567890",
			want: "89012345",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			listStr := strings.Split(tt.list, "")
			listInt := make([]int8, len(listStr))
			for i, v := range listStr {
				n, _ := strconv.Atoi(v)
				listInt[i] = int8(n)
			}

			if got := extractDigitsOffset(listInt); got != tt.want {
				t.Errorf("extractDigitsOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		list string
		want string
	}{
		{
			name: "",
			list: "03036732577212944063491565474664",
			want: "84462026",
		},
		{
			name: "",
			list: "02935109699940807407585447034323",
			want: "78725270",
		},
		{
			name: "",
			list: "03081770884921959731165446850517",
			want: "53553731",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			listStr := strings.Split(tt.list, "")
			listInt := make([]int8, len(listStr))
			for i, v := range listStr {
				n, _ := strconv.Atoi(v)
				listInt[i] = int8(n)
			}

			if got := part2(listInt); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
