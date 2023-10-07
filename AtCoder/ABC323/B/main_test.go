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
		-xx
		o-x
		oo-`)}, wantOut: `3 2 1 `},
		{name: "2", args: args{stdin: strings.NewReader(`7
		-oxoxox
		x-xxxox
		oo-xoox
		xoo-ooo
		ooxx-ox
		xxxxx-x
		oooxoo-`)}, wantOut: `4 7 3 1 5 2 6 `},
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
