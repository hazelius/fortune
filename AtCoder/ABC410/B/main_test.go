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
		{name: "1", args: args{stdin: strings.NewReader(`4 5
2 0 3 0 0`)}, wantOut: `2 1 3 4 1`},
		{name: "2", args: args{stdin: strings.NewReader(`3 7
1 1 0 0 0 0 0`)}, wantOut: `1 1 2 3 2 3 1`},
		{name: "3", args: args{stdin: strings.NewReader(`6 20
4 6 0 3 4 2 6 5 2 3 0 3 2 5 0 3 5 0 2 0`)}, wantOut: `4 6 1 3 4 2 6 5 2 3 1 3 2 5 1 3 5 4 2 6`},
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
