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
		{name: "1", args: args{stdin: strings.NewReader(`4 3 5
		5 3 0 2
		3 1 2 3
		3 2 4 0
		1 0 1 4`)}, wantOut: `9`},
		{name: "2", args: args{stdin: strings.NewReader(`7 3 5
		85 1 0 1
		37 1 1 0
		38 2 0 0
		45 0 2 2
		67 1 1 0
		12 2 2 0
		94 2 2 1`)}, wantOut: `-1`},
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
