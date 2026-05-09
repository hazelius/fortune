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
		{name: "1", args: args{stdin: strings.NewReader(`3 3
1 2 3`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`4 5
10 1 10 1`)}, wantOut: `7`},
		{name: "3", args: args{stdin: strings.NewReader(`20 457
8 9 10 9 8 8 4 6 8 1 5 10 2 8 2 6 8 1 6 6`)}, wantOut: `132`},
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
