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
5 0 9 3`)}, wantOut: `2 0 10 5`},
		{name: "2", args: args{stdin: strings.NewReader(`5
4 6 7 2 5`)}, wantOut: `0 4 7 4 9`},
		{name: "3", args: args{stdin: strings.NewReader(`10
2 9 1 2 0 4 6 7 1 5`)}, wantOut: `0 2 0 0 0 4 7 10 4 10`},
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
