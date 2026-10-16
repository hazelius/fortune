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
		{name: "1", args: args{stdin: strings.NewReader(`4 13
4 5
3 3
5 9
8 6`)}, wantOut: `12`},
		{name: "2", args: args{stdin: strings.NewReader(`3 99
1 100
10 100
100 100`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`5 9
6 4
9 1
2 1
4 3
5 4`)}, wantOut: `21`},
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
