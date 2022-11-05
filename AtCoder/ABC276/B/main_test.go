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
		{name: "1", args: args{stdin: strings.NewReader(`6 6
		3 6
		1 3
		5 6
		2 5
		1 2
		1 6`)}, wantOut: `3 2 3 6
2 1 5
2 1 6
0
2 2 6
3 1 3 5
`},
		{name: "2", args: args{stdin: strings.NewReader(`5 10
		1 2
		1 3
		1 4
		1 5
		2 3
		2 4
		2 5
		3 4
		3 5
		4 5`)}, wantOut: `4 2 3 4 5
4 1 3 4 5
4 1 2 4 5
4 1 2 3 5
4 1 2 3 4
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
