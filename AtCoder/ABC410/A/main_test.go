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
3 1 4 1 5
4`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`1
1
100`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`15
18 89 31 2 15 93 64 78 58 19 79 59 24 50 30
38`)}, wantOut: `8`},
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
