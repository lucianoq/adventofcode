package main

import (
	"reflect"
	"testing"
)

func Test_findABAs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str  string
		want []string
	}{
		{
			str:  "aba",
			want: []string{"aba"},
		},
		{
			str:  "bab",
			want: []string{"bab"},
		},
		{
			str:  "xyz",
			want: []string{},
		},
		{
			str:  "aaa",
			want: []string{},
		},
		{
			str:  "kek",
			want: []string{"kek"},
		},
		{
			str:  "eke",
			want: []string{"eke"},
		},
		{
			str:  "zazbz",
			want: []string{"zaz", "zbz"},
		},
		{
			str:  "bzb",
			want: []string{"bzb"},
		},
		{
			str:  "cdb",
			want: []string{},
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.str, func(t *testing.T) {
			t.Parallel()

			if got := findABAs(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findABAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsBAB(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str  string
		aba  string
		want bool
	}{
		{
			str:  "bab",
			aba:  "aba",
			want: true,
		},
		{
			str:  "xyz",
			aba:  "xyx",
			want: false,
		},
		{
			str:  "kek",
			aba:  "eke",
			want: true,
		},
		{
			str:  "bzb",
			aba:  "zaz",
			want: false,
		},
		{
			str:  "bzb",
			aba:  "zbz",
			want: true,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.str, func(t *testing.T) {
			t.Parallel()

			if got := containsBAB(tt.str, tt.aba); got != tt.want {
				t.Errorf("containsBAB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_supportSSL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ip   string
		want bool
	}{
		{
			ip:   "aba[bab]xyz",
			want: true,
		},
		{
			ip:   "xyx[xyx]xyx",
			want: false,
		},
		{
			ip:   "aaa[kek]eke",
			want: true,
		},
		{
			ip:   "zazbz[bzb]cdb",
			want: true,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.ip, func(t *testing.T) {
			t.Parallel()

			if got := supportSSL(tt.ip); got != tt.want {
				t.Errorf("supportSSL() = %v, want %v", got, tt.want)
			}
		})
	}
}
