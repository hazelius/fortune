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
		{name: "1", args: args{stdin: strings.NewReader(`5 19 22
17 23
20 23
19 22
0 23
12 20`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3 12 13
0 1
0 1
0 1`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 8 14
5 20
14 21
9 21
5 23
8 10
0 14
3 8
2 6
0 16
5 20`)}, wantOut: `5`},
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
