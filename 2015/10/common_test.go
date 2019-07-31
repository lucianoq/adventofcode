package main

import (
	"testing"
)

func TestEncodeRecursive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1",
			want:  "11",
		},
		{
			input: "11",
			want:  "21",
		},
		{
			input: "21",
			want:  "1211",
		},
		{
			input: "1211",
			want:  "111221",
		},
		{
			input: "111221",
			want:  "312211",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			if got := EncodeRecursive(tt.input); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeIter(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1",
			want:  "11",
		},
		{
			input: "11",
			want:  "21",
		},
		{
			input: "21",
			want:  "1211",
		},
		{
			input: "1211",
			want:  "111221",
		},
		{
			input: "111221",
			want:  "312211",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			if got := EncodeIter(tt.input); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeIterBuf(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "1",
			want:  "11",
		},
		{
			input: "11",
			want:  "21",
		},
		{
			input: "21",
			want:  "1211",
		},
		{
			input: "1211",
			want:  "111221",
		},
		{
			input: "111221",
			want:  "312211",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			if got := EncodeIterBuf(tt.input); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
