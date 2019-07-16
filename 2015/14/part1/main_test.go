package main

import (
	"strconv"
	"testing"
)

var (
	comet = Reindeer{
		name:     "Comet",
		speed:    14,
		runTime:  10,
		restTime: 127,
	}

	dancer = Reindeer{
		name:     "Dancer",
		speed:    16,
		runTime:  11,
		restTime: 162,
	}
)

func Test(t *testing.T) {
	t.Parallel()
	tests := []struct {
		r      Reindeer
		time   int
		wanted int
	}{
		{
			r:      comet,
			time:   1,
			wanted: 14,
		},
		{
			r:      dancer,
			time:   1,
			wanted: 16,
		},
		{
			r:      comet,
			time:   10,
			wanted: 140,
		},
		{
			r:      dancer,
			time:   10,
			wanted: 160,
		},
		{
			r:      comet,
			time:   11,
			wanted: 140,
		},
		{
			r:      dancer,
			time:   11,
			wanted: 176,
		},
		{
			r:      comet,
			time:   12,
			wanted: 140,
		},
		{
			r:      dancer,
			time:   12,
			wanted: 176,
		},
		{
			r:      comet,
			time:   1000,
			wanted: 1120,
		},
		{
			r:      dancer,
			time:   1000,
			wanted: 1056,
		},
	}
	for _, tt := range tests {
		var tt = tt

		name := tt.r.name + "at time " + strconv.Itoa(tt.time)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.r.Travelled(tt.time)

			if got != tt.wanted {
				t.Errorf("%s: expected %d, got %d", name, tt.wanted, got)
			}
		})
	}
}
