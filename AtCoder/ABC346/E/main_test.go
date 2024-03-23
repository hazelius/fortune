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
		{name: "1", args: args{stdin: strings.NewReader(`3 4 4
				1 2 5
				2 4 0
				1 3 3
				1 3 2`)}, wantOut: `3
0 5
2 4
5 3
`},
		{name: "2", args: args{stdin: strings.NewReader(`1 1 5
		1 1 1
		1 1 10
		2 1 100
		1 1 1000
		2 1 10000`)}, wantOut: `1
10000 1
`},
		{name: "3", args: args{stdin: strings.NewReader(`5 5 10
		1 1 1
		1 2 2
		1 3 3
		1 4 4
		1 5 5
		2 1 6
		2 2 7
		2 3 8
		2 4 9
		2 5 10`)}, wantOut: `5
6 5
7 5
8 5
9 5
10 5
`},
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
