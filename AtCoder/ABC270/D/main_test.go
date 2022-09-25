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
		{name: "1", args: args{stdin: strings.NewReader(`10 2
		1 4`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`11 4
		1 2 3 6`)}, wantOut: `8`},
		{name: "3", args: args{stdin: strings.NewReader(`10000 10
		1 2 4 8 16 32 64 128 256 512`)}, wantOut: `5136`},
		{name: "4", args: args{stdin: strings.NewReader(`100 4
		1 31 34 35`)}, wantOut: `65`},
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
