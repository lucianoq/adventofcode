package main

import (
	"strconv"
	"testing"
)

func TestPresent(t *testing.T) {
	tests := []struct {
		house  int
		wanted int
	}{
		{
			house:  1,
			wanted: 10,
		},
		{
			house:  2,
			wanted: 30,
		},
		{
			house:  3,
			wanted: 40,
		},
		{
			house:  4,
			wanted: 70,
		},
		{
			house:  5,
			wanted: 60,
		},
		{
			house:  6,
			wanted: 120,
		},
		{
			house:  7,
			wanted: 80,
		},
		{
			house:  8,
			wanted: 150,
		},
		{
			house:  9,
			wanted: 130,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(strconv.Itoa(tt.house), func(t *testing.T) {
			t.Parallel()
			got := presents(tt.house)
			if got != tt.wanted {
				t.Errorf("wanted %d, got %d", tt.wanted, got)
			}
		})
	}
}
