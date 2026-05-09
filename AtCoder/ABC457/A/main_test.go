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
		{name: "1", args: args{stdin: strings.NewReader(`5
1 2 3 4 5
3`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`10
6 6 9 6 10 5 7 2 8 2
4`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`10
4 4 4 3 4 2 1 1 2 1
10`)}, wantOut: `1`},
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
