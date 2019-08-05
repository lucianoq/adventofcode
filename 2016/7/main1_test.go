package main

import "testing"

func Test_containsABBA(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{
			str:  "xyyx",
			want: true,
		},
		{
			str:  "abba",
			want: true,
		},
		{
			str:  "aaaa",
			want: false,
		},
		{
			str:  "mnop",
			want: false,
		},
		{
			str:  "qrst",
			want: false,
		},
		{
			str:  "ioxxoj",
			want: true,
		},
		{
			str:  "bddb",
			want: true,
		},
	}
	for _, tt := range tests {
		var tt = tt

		t.Run(tt.str, func(t *testing.T) {
			t.Parallel()

			if got := containsABBA(tt.str); got != tt.want {
				t.Errorf("containsABBA = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_supportTLS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ip   string
		want bool
	}{
		{
			ip:   "abba[mnop]qrst",
			want: true,
		},
		{
			ip:   "abcd[bddb]xyyx",
			want: false,
		},
		{
			ip:   "aaaa[qwer]tyui",
			want: false,
		},
		{
			ip:   "ioxxoj[asdfgh]zxcvbn",
			want: true,
		},
	}

	for _, tt := range tests {
		var tt = tt

		t.Run(tt.ip, func(t *testing.T) {
			t.Parallel()

			if got := supportTLS(tt.ip); got != tt.want {
				t.Errorf("supportTLS() = %v, want %v", got, tt.want)
			}
		})
	}
}
