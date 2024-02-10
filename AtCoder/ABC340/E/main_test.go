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
		1 2 3 4 5
		2 4 0`)}, wantOut: `0 4 2 7 2 `},
		{name: "2", args: args{stdin: strings.NewReader(`3 10
		1000000000 1000000000 1000000000
		0 1 0 1 0 1 0 1 0 1`)}, wantOut: `104320141 45436840 2850243019 `},
		{name: "3", args: args{stdin: strings.NewReader(`1 4
		1
		0 0 0 0`)}, wantOut: `1 `},
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
