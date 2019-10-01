package main

import (
	"os"
	"testing"
)

func TestFindOutcomeLowestAttack(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		fileName string
		wanted   int
	}{
		{
			name:     "one",
			fileName: "input_test_1",
			wanted:   4988,
		},
		{
			name:     "three",
			fileName: "input_test_3",
			wanted:   31284,
		},
		{
			name:     "four",
			fileName: "input_test_4",
			wanted:   3478,
		},
		{
			name:     "five",
			fileName: "input_test_5",
			wanted:   6474,
		},
		{
			name:     "six",
			fileName: "input_test_6",
			wanted:   1140,
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

			got := FindOutcomeLowestAttack(f)
			if got != tt.wanted {
				t.Errorf("wanted %d, got %d", tt.wanted, got)
			}
		})
	}
}
