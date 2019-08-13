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
		input  string
		char   string
		exists bool
	}{
		{
			input:  "abcdefghilmnopq",
			char:   "",
			exists: false,
		},
		{
			input:  "aa",
			char:   "",
			exists: false,
		},
		{
			input:  "aabbcc",
			char:   "",
			exists: false,
		},
		{
			input:  "aaaaabbcc",
			char:   "a",
			exists: true,
		},
		{
			input:  "aaaabbbbcc",
			char:   "a",
			exists: true,
		},
		{
			input:  "aaaaabbbbbccccc",
			char:   "a",
			exists: true,
		},
		{
			input:  "aaaaaaaaaaaaaabb",
			char:   "a",
			exists: true,
		},
		{
			input:  "aaaaaaaaaaaaaabbbbb",
			char:   "a",
			exists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotChar, gotExists := triples(tt.input)
			if gotChar != tt.char {
				t.Errorf("triples() = %v, want %v", gotChar, tt.char)
			}
			if gotExists != tt.exists {
				t.Errorf("triples() = %v, want %v", gotExists, tt.exists)
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
