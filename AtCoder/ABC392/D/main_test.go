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
3 1 2 3
4 1 2 2 1
6 1 2 3 4 5 6`)}, wantOut: `0.333333333333333`},
		{name: "2", args: args{stdin: strings.NewReader(`3
5 1 1 1 1 1
4 2 2 2 2
3 1 1 2`)}, wantOut: `0.666666666666667`},
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
