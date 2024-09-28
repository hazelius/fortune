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
		{name: "1", args: args{stdin: strings.NewReader(`3 3
		1 2 2
		3 2 3
		1 3 -1`)}, wantOut: `0 2 -1`},
		{name: "2", args: args{stdin: strings.NewReader(`4 2
2 1 5
3 4 -3`)}, wantOut: `0 -5 0 -3`},
		{name: "3", args: args{stdin: strings.NewReader(`5 7
2 1 18169343
3 1 307110901
4 1 130955934
2 3 -288941558
2 5 96267410
5 3 -385208968
4 3 -176154967`)}, wantOut: `0 -18169343 -307110901 -130955934 78098067`},
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
