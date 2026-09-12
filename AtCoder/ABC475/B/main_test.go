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
1296 110 1`)}, wantOut: `13 18 24`},
		{name: "2", args: args{stdin: strings.NewReader(`12
3141 592 65358 9 79 323 84 6264 3 38327 950 28`)}, wantOut: `52 59 82`},
		{name: "3", args: args{stdin: strings.NewReader(`3
1000 1000 1000`)}, wantOut: `0 0 0`},
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
