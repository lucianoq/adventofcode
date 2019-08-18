package main

import "testing"

func Test_next(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "..^^.",
			output: ".^^^^",
		},
		{
			input:  ".^^^^",
			output: "^^..^",
		},
		{
			input:  ".^^.^.^^^^",
			output: "^^^...^..^",
		},
		{
			input:  "^^^...^..^",
			output: "^.^^.^.^^.",
		},
		{
			input:  "^.^^.^.^^.",
			output: "..^^...^^^",
		},
		{
			input:  "..^^...^^^",
			output: ".^^^^.^^.^",
		},
		{
			input:  ".^^^^.^^.^",
			output: "^^..^.^^..",
		},
		{
			input:  "^^..^.^^..",
			output: "^^^^..^^^.",
		},
		{
			input:  "^^^^..^^^.",
			output: "^..^^^^.^^",
		},
		{
			input:  "^..^^^^.^^",
			output: ".^^^..^.^^",
		},
		{
			input:  ".^^^..^.^^",
			output: "^^.^^^..^^",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := next(tt.input); got != tt.output {
				t.Errorf("next() = %v, want %v", got, tt.output)
			}
		})
	}
}
