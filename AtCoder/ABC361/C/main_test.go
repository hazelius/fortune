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
		{name: "1", args: args{stdin: strings.NewReader(`5 2
3 1 5 4 9`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`6 5
		1 1 1 1 1 1`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`8 3
		31 43 26 6 18 36 22 13`)}, wantOut: `18`},
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
