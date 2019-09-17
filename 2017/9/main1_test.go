package main

import "testing"

func TestStream_ExtractGroup(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string

		buf []byte
		pos int

		level int

		want int
	}{
		{
			name: "{}",
			buf:  []byte("{}"),
			want: 1,
		},
		{
			name: "{{{}}}",
			buf:  []byte("{{{}}}"),
			want: 6,
		},
		{
			name: "{{},{}}",
			buf:  []byte("{{},{}}"),
			want: 5,
		},
		{
			name: "{{{},{},{{}}}}",
			buf:  []byte("{{{},{},{{}}}}"),
			want: 16,
		},
		{
			name: "{<a>,<a>,<a>,<a>}",
			buf:  []byte("{<a>,<a>,<a>,<a>}"),
			want: 1,
		},
		{
			name: "{{<ab>},{<ab>},{<ab>},{<ab>}}",
			buf:  []byte("{{<ab>},{<ab>},{<ab>},{<ab>}}"),
			want: 9,
		},
		{
			name: "{{<!!>},{<!!>},{<!!>},{<!!>}}",
			buf:  []byte("{{<!!>},{<!!>},{<!!>},{<!!>}}"),
			want: 9,
		},
		{
			name: "{{<a!>},{<a!>},{<a!>},{<ab>}}",
			buf:  []byte("{{<a!>},{<a!>},{<a!>},{<ab>}}"),
			want: 3,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Stream{
				buf: tt.buf,
				pos: tt.pos,
			}
			if got := s.ExtractGroup(tt.level); got != tt.want {
				t.Errorf("ExtractGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
