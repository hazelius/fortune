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
		{name: "1", args: args{stdin: strings.NewReader(`4 5
4 2 1 8
14 9 3 2 9`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
5 5 3
11 1000 1000`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`8 7
2 3 4 4 4 3 2 3
8 5 5 9 9 7 1`)}, wantOut: `5`},
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
