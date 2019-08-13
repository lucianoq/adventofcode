package main

import (
	"reflect"
	"testing"
)

func Test_same(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "",
			want:  true,
		},
		{
			input: "a",
			want:  true,
		},
		{
			input: "aa",
			want:  true,
		},
		{
			input: "aaa",
			want:  true,
		},
		{
			input: "aab",
			want:  false,
		},
		{
			input: "bbb",
			want:  true,
		},
		{
			input: "baaa",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := same(tt.input); got != tt.want {
				t.Errorf("same() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_triples(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]bool
	}{
		{
			input: "abcdefghilmnopq",
			want:  map[string]bool{},
		},
		{
			input: "aa",
			want:  map[string]bool{},
		},
		{
			input: "aabbcc",
			want:  map[string]bool{},
		},
		{
			input: "aaaaabbcc",
			want: map[string]bool{
				"a": true,
			},
		},
		{
			input: "aaaabbbbcc",
			want: map[string]bool{
				"a": true,
				"b": true,
			},
		},
		{
			input: "aaaaabbbbbccccc",
			want: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
		},
		{
			input: "aaaaaaaaaaaaaabb",
			want: map[string]bool{
				"a": true,
			},
		},
		{
			input: "aaaaaaaaaaaaaabbbbb",
			want: map[string]bool{
				"a": true,
				"b": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := triples(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("triples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quintuples(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]bool
	}{
		{
			input: "abcdefghilmnopq",
			want:  map[string]bool{},
		},
		{
			input: "aaaa",
			want:  map[string]bool{},
		},
		{
			input: "aaaabbbbcccc",
			want:  map[string]bool{},
		},
		{
			input: "aaaaabbbbcccc",
			want: map[string]bool{
				"a": true,
			},
		},
		{
			input: "aaaaabbbbbcccc",
			want: map[string]bool{
				"a": true,
				"b": true,
			},
		},
		{
			input: "aaaaabbbbbccccc",
			want: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
		},
		{
			input: "aaaaaaaaaaaaaabbbb",
			want: map[string]bool{
				"a": true,
			},
		},
		{
			input: "aaaaaaaaaaaaaabbbbb",
			want: map[string]bool{
				"a": true,
				"b": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := quintuples(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quintuples() = %v, want %v", got, tt.want)
			}
		})
	}
}