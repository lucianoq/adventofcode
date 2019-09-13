package main

import "testing"

func TestStream_ExtractGroup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		buf  []byte

		want int
	}{
		{
			name: "<>",
			buf:  []byte("<>"),
			want: 0,
		},

		{
			name: "<random characters>",
			buf:  []byte("<random characters>"),
			want: 17,
		},
		{
			name: "<<<<>",
			buf:  []byte("<<<<>"),
			want: 3,
		},
		{
			name: "<{!>}>",
			buf:  []byte("<{!>}>"),
			want: 2,
		},
		{
			name: "<!!>",
			buf:  []byte("<!!>"),
			want: 0,
		},
		{
			name: "<!!!>>",
			buf:  []byte("<!!!>>"),
			want: 0,
		},
		{
			name: `<{o"i!a,<{i<a>`,
			buf:  []byte(`<{o"i!a,<{i<a>`),
			want: 10,
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &Stream{
				buf: tt.buf,
			}
			if got := s.ExtractGroup(); got != tt.want {
				t.Errorf("ExtractGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

//<>, 0 characters.
//<random characters>, 17 characters.
//<<<<>, 3 characters.
//<{!>}>, 2 characters.
//<!!>, 0 characters.
//<!!!>>, 0 characters.
//<{o"i!a,<{i<a>, 10 characters.
