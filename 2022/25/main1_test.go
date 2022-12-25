package main

import (
	"strconv"
	"testing"
)

var tests = []struct {
	snafu string
	dec   int
}{
	{"1", 1},
	{"2", 2},
	{"1=", 3},
	{"1-", 4},
	{"10", 5},
	{"11", 6},
	{"12", 7},
	{"2=", 8},
	{"2-", 9},
	{"20", 10},
	{"1=0", 15},
	{"1-0", 20},
	{"1=11-2", 2022},
	{"1-0---0", 12345},
	{"1121-1110-1=0", 314159265},
	{"1=-0-2", 1747},
	{"12111", 906},
	{"2=0=", 198},
	{"21", 11},
	{"2=01", 201},
	{"111", 31},
	{"20012", 1257},
	{"112", 32},
	{"1=-1=", 353},
	{"1-12", 107},
	{"12", 7},
	{"1=", 3},
	{"122", 37},
}

func Test_toDecimal(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.snafu+" to decimal", func(t *testing.T) {
			t.Parallel()
			if got := toDecimal(tt.snafu); got != tt.dec {
				t.Errorf("toDecimal() = %v, want %v", got, tt.dec)
			}
		})
	}
}

func Test_toSNAFU(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		var tt = tt
		t.Run(strconv.Itoa(tt.dec)+" to SNAFU", func(t *testing.T) {
			t.Parallel()
			if got := toSNAFU(tt.dec); got != tt.snafu {
				t.Errorf("toSNAFU() = %v, want %v", got, tt.snafu)
			}
		})
	}
}
