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
		1 1 0 2
		MEEX`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		0 0 0
		XXX`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`15
		1 1 2 0 0 2 0 2 0 0 0 0 0 2 2
		EXMMXXXEMEXEXMM`)}, wantOut: `13`},
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
