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
		1 2
		2 3
		1 4`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3 0`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 8
		1 2
		2 3
		3 4
		4 5
		6 7
		7 8
		8 9
		9 10`)}, wantOut: `12`},
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
