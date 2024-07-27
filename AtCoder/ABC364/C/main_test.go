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
		{name: "1", args: args{stdin: strings.NewReader(`4 7 18
		2 3 5 1
		8 8 1 4`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`5 200000000000000 200000000000000
		1 1 1 1 1
		2 2 2 2 2`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`8 30 30
1 2 3 4 5 6 7 8
8 7 6 5 4 3 2 1`)}, wantOut: `6`},
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
