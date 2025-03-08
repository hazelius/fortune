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
		8 5 -1 3
		3 -2 -4`)}, wantOut: `19`},
		{name: "2", args: args{stdin: strings.NewReader(`4 3
5 -10 -2 -5
8 1 4`)}, wantOut: `15`},
		{name: "3", args: args{stdin: strings.NewReader(`3 5
-36 -33 -31
12 12 28 24 27`)}, wantOut: `0`},
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
