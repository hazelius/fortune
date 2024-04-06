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
		{name: "1", args: args{stdin: strings.NewReader(`4 4
		S...
		#..#
		#...
		..#T
		4
		1 1 3
		1 3 5
		3 2 1
		2 3 1`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
		S.
		T.
		1
		1 2 4`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`4 5
		..#..
		.S##.
		.##T.
		.....
		3
		3 1 5
		1 2 3
		2 2 1`)}, wantOut: `Yes`},
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
