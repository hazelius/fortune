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
		{name: "1", args: args{stdin: strings.NewReader(`3
		1 3
		3 1
		2 2`)}, wantOut: `2 3 1 `},
		{name: "2", args: args{stdin: strings.NewReader(`2
		1 3
		2 6`)}, wantOut: `1 2 `},
		{name: "3", args: args{stdin: strings.NewReader(`4
		999999999 1000000000
		333333333 999999999
		1000000000 999999997
		999999998 1000000000`)}, wantOut: `3 1 4 2 `},
		{name: "4", args: args{stdin: strings.NewReader(`3
		0 1
		1 0
		1 1`)}, wantOut: `2 3 1 `},
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
