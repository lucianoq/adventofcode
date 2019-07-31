package main

import "testing"

func TestNext(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "a",
			want:  "b",
		},
		{
			input: "aa",
			want:  "ab",
		},
		{
			input: "aaa",
			want:  "aab",
		},
		{
			input: "z",
			want:  "aa",
		},
		{
			input: "az",
			want:  "ba",
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := Next(tt.input)
			if got != tt.want {
				t.Errorf("got %s, expected %s", got, tt.want)
			}
		})
	}
}

func TestContains3Straight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "aaa",
			want:  false,
		},
		{
			input: "abc",
			want:  true,
		},
		{
			input: "aaaaaaa",
			want:  false,
		},
		{
			input: "aaabacc",
			want:  false,
		},
		{
			input: "aaaaaaaaaabc",
			want:  true,
		},
		{
			input: "hhhhxyzhhhh",
			want:  true,
		},
		{
			input: "aaaaabdeffff",
			want:  true,
		},
		{
			input: "qwegssgsagasdgf",
			want:  false,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := Contains3Straight(tt.input)
			if got != tt.want {
				t.Errorf("got %t, expected %t", got, tt.want)
			}
		})
	}
}

func TestNotContainsIOL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "aaa",
			want:  true,
		},
		{
			input: "abc",
			want:  true,
		},
		{
			input: "aaalaaaa",
			want:  false,
		},
		{
			input: "aaabiacc",
			want:  false,
		},
		{
			input: "aaaaaoaaaaabc",
			want:  false,
		},
		{
			input: "hhhhxyzhhhh",
			want:  true,
		},
		{
			input: "aaaaabdeffff",
			want:  true,
		},
		{
			input: "qwegssgiolsagasdgf",
			want:  false,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := NotContainsIOL(tt.input)
			if got != tt.want {
				t.Errorf("got %t, expected %t", got, tt.want)
			}
		})
	}
}

func TestContains2Pairs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "a",
			want:  false,
		},
		{
			input: "aa",
			want:  false,
		},
		{
			input: "aaa",
			want:  false,
		},
		{
			input: "aaaa",
			want:  true,
		},
		{
			input: "aabb",
			want:  true,
		},
		{
			input: "aabba",
			want:  true,
		},
		{
			input: "aaxbb",
			want:  true,
		},
		{
			input: "aaxxxxbb",
			want:  true,
		},
		{
			input: "aabbccddee",
			want:  true,
		},
		{
			input: "aaaqqqqq",
			want:  true,
		},
		{
			input: "aaaqxqxqxqxqx",
			want:  false,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := Contains2Pairs(tt.input)
			if got != tt.want {
				t.Errorf("got %t, expected %t", got, tt.want)
			}
		})
	}
}

func TestValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "hijklmmn",
			want:  false,
		},
		{
			input: "abbceffg",
			want:  false,
		},
		{
			input: "abbcegjk",
			want:  false,
		},
		{
			input: "abcdffaa",
			want:  true,
		},
		{
			input: "ghjaabcc",
			want:  true,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := Valid(tt.input)
			if got != tt.want {
				t.Errorf("got %t, expected %t", got, tt.want)
			}
		})
	}
}

func TestNextValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "abcdefgh",
			want:  "abcdffaa",
		},
		{
			input: "ghijklmn",
			want:  "ghjaabcc",
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := NextValid(tt.input)
			if got != tt.want {
				t.Errorf("got %s, expected %s", got, tt.want)
			}
		})
	}
}
