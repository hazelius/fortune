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
		1 5 4
		7 8
		6`)}, wantOut: `13`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		1 2
		3`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`16
		5 6 5 2 1 7 9 7 2 5 5 2 4 7 6
		8 7 7 9 8 1 9 6 10 8 8 6 10 3
		10 5 8 1 10 7 8 4 8 6 5 1 10
		7 4 1 4 5 4 5 10 1 5 1 2
		2 9 9 7 6 2 2 8 3 5 2
		9 10 3 1 1 2 10 7 7 5
		10 6 1 8 9 3 2 4 2
		10 10 8 9 2 10 7 9
		5 8 8 7 5 8 2
		4 2 2 6 8 3
		2 7 3 10 3
		5 7 10 3
		8 5 7
		9 1
		4`)}, wantOut: `75`},
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
