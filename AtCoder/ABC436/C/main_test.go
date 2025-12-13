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
		{name: "1", args: args{stdin: strings.NewReader(`4 3
1 1
2 2
2 3`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`1000000000 4
1 1
1 101
101 1
101 101`)}, wantOut: `4`},
		{name: "3", args: args{stdin: strings.NewReader(`8 10
6 5
7 3
6 7
3 4
4 2
3 7
1 3
7 4
6 1
6 1`)}, wantOut: `8`},
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
