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
		{name: "1", args: args{stdin: strings.NewReader(`5 3
		1 2
		1 3
		4 5`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`5 0`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`4 6
		1 2
		1 3
		1 4
		2 3
		2 4
		3 4`)}, wantOut: `1`},
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
