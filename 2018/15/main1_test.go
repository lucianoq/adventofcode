package main

import (
	"os"
	"testing"
)

func TestFindOutcome(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		fileName string
		wanted   int
	}{
		{
			name:     "one",
			fileName: "input_test_1",
			wanted:   27730,
		},
		{
			name:     "two",
			fileName: "input_test_2",
			wanted:   36334,
		},
		{
			name:     "three",
			fileName: "input_test_3",
			wanted:   39514,
		},
		{
			name:     "four",
			fileName: "input_test_4",
			wanted:   27755,
		},
		{
			name:     "five",
			fileName: "input_test_5",
			wanted:   28944,
		},
		{
			name:     "six",
			fileName: "input_test_6",
			wanted:   18740,
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(tt.fileName)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			got := FindOutcome(f)
			if got != tt.wanted {
				t.Errorf("wanted %d, got %d", tt.wanted, got)
			}

		})
	}
}
