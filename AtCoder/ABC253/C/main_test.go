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
		{name: "1", args: args{stdin: strings.NewReader(`8
		1 3
		1 2
		3
		1 2
		1 7
		3
		2 2 3
		3`)}, wantOut: `1
5
4
`},
		{name: "2", args: args{stdin: strings.NewReader(`4
		1 10000
		1 1000
		2 100 3
		1 10`)}, wantOut: ``},
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
