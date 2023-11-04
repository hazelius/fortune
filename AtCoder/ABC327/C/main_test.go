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
		{name: "1", args: args{stdin: strings.NewReader(`1 2 3 4 5 6 7 8 9
		4 5 6 7 8 9 1 2 3
		7 8 9 1 2 3 4 5 6
		2 3 4 5 6 7 8 9 1
		5 6 7 8 9 1 2 3 4
		8 9 1 2 3 4 5 6 7
		3 4 5 6 7 8 9 1 2
		6 7 8 9 1 2 3 4 5
		9 1 2 3 4 5 6 7 8`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`1 2 3 4 5 6 7 8 9
		2 3 4 5 6 7 8 9 1
		3 4 5 6 7 8 9 1 2
		4 5 6 7 8 9 1 2 3
		5 6 7 8 9 1 2 3 4
		6 7 8 9 1 2 3 4 5
		7 8 9 1 2 3 4 5 6
		8 9 1 2 3 4 5 6 7
		9 1 2 3 4 5 6 7 8`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`1 2 3 4 5 6 7 8 9
		4 5 6 7 8 9 1 2 3
		7 8 9 1 2 3 4 5 6
		1 2 3 4 5 6 7 8 9
		4 5 6 7 8 9 1 2 3
		7 8 9 1 2 3 4 5 6
		1 2 3 4 5 6 7 8 9
		4 5 6 7 8 9 1 2 3
		7 8 9 1 2 3 4 5 6`)}, wantOut: `No`},
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
