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
		3 3
		5 1
		6 1`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		1 1
		2 1
		3 1`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`1
		1000000000 1000000000`)}, wantOut: `13`},
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
