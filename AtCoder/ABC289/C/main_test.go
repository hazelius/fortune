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
		{name: "1", args: args{stdin: strings.NewReader(`3 3
		2
		1 2
		2
		1 3
		1
		2`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`4 2
		2
		1 2
		2
		1 3`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`6 6
		3
		2 3 6
		3
		2 4 6
		2
		3 6
		3
		1 5 6
		3
		1 3 6
		2
		1 4`)}, wantOut: `18`},
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
