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
3 1 4 1`)}, wantOut: `3
1 3 4`},
		{name: "2", args: args{stdin: strings.NewReader(`3
7 7 7`)}, wantOut: `1
7`},
		{name: "3", args: args{stdin: strings.NewReader(`8
19 5 5 19 5 19 4 19`)}, wantOut: `3
4 5 19`},
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
