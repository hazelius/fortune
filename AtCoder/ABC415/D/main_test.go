package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		stdin io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{name: "1", args: args{stdin: strings.NewReader(`5 3
5 1
4 3
3 1`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
5 1
5 1
4 2`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`415 8
327 299
413 396
99 67
108 51
195 98
262 180
250 175
234 187`)}, wantOut: `11`},
		{name: "4", args: args{stdin: strings.NewReader(`8 2
8 3
3 2`)}, wantOut: `6`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			run(tt.args.stdin, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
