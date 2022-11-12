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
		{name: "1", args: args{stdin: strings.NewReader(`5 5 2
		1 3 0
		2 3 1
		5 4 1
		2 1 1
		1 4 0
		3 4`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`4 4 2
		4 3 0
		1 2 1
		1 2 0
		2 1 1
		2 4`)}, wantOut: `-1`},
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
