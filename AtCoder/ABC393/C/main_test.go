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
		{name: "1", args: args{stdin: strings.NewReader(`3 5
1 2
2 3
3 2
3 1
1 1`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`1 0`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`6 10
6 2
4 1
5 1
6 6
5 3
5 1
1 4
6 4
4 2
5 6`)}, wantOut: `3`},
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
