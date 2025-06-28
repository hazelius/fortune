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
		{name: "1", args: args{stdin: strings.NewReader(`4
2 8
5 5
5 4
6 7`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`5
1 1
1 1
1 1
1 1
1 1`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`6
1 6
2 5
3 4
4 3
5 2
6 1`)}, wantOut: `3`},
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
