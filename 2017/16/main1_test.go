package main

import (
	"reflect"
	"testing"
)

func TestCmd_Run(t *testing.T) {
	type fields struct {
		Op       byte
		SpinSize int
		Ex1      int
		Ex2      int
		Partner1 byte
		Partner2 byte
	}
	type args struct {
		s []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "s3",
			fields: fields{
				Op:       's',
				SpinSize: 3,
			},
			args: args{s: []byte("abcde")},
			want: []byte("cdeab"),
		},
		{
			name: "s0",
			fields: fields{
				Op:       's',
				SpinSize: 0,
			},
			args: args{s: []byte("abcde")},
			want: []byte("abcde"),
		},
		{
			name: "x3/4",
			fields: fields{
				Op:  'x',
				Ex1: 3,
				Ex2: 4,
			},
			args: args{s: []byte("eabcd")},
			want: []byte("eabdc"),
		},
		{
			name: "pe/b",
			fields: fields{
				Op:       'p',
				Partner1: 'e',
				Partner2: 'b',
			},
			args: args{s: []byte("eabdc")},
			want: []byte("baedc"),
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			c := Cmd{
				Op:       tt.fields.Op,
				SpinSize: tt.fields.SpinSize,
				Ex1:      tt.fields.Ex1,
				Ex2:      tt.fields.Ex2,
				Partner1: tt.fields.Partner1,
				Partner2: tt.fields.Partner2,
			}
			c.Apply(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("Apply() = %s, want %s", string(tt.args.s), string(tt.want))
			}
		})
	}
}
