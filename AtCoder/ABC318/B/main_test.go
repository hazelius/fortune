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
		0 5 1 3
		1 4 0 5
		2 5 2 4`)}, wantOut: `20`},
		{name: "2", args: args{stdin: strings.NewReader(`2
		0 100 0 100
		0 100 0 100`)}, wantOut: `10000`},
		{name: "3", args: args{stdin: strings.NewReader(`3
		0 1 0 1
		0 3 0 5
		5 10 0 10`)}, wantOut: `65`},
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
