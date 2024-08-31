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
		3 6 9 3`)}, wantOut: `8`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		1 1 1 1 1`)}, wantOut: `15`},
		{name: "3", args: args{stdin: strings.NewReader(`8
87 42 64 86 72 58 44 30`)}, wantOut: `22`},
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
