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
		{name: "1", args: args{stdin: strings.NewReader(`4 2
		1 2 3 4
		4 3 1 2`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
		1 2 3
		3 1 2
		1 2 3`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
		4 10 7 2 8 3 9 1 6 5
		3 6 2 9 1 8 10 7 4 5
		9 3 4 5 7 10 1 8 2 6
		7 3 1 8 4 9 5 6 2 10
		5 2 1 4 10 7 9 8 3 6
		5 8 1 6 9 3 2 4 7 10
		8 10 3 4 5 7 2 9 6 1
		3 10 2 7 8 5 1 4 9 6
		10 6 1 5 4 2 3 8 9 7
		4 5 9 1 8 2 7 6 3 10`)}, wantOut: `6`},
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
