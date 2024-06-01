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
		{name: "1", args: args{stdin: strings.NewReader(`3 2 2
		3 1 2 3 o
		2 2 3 x`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`4 5 3
		3 1 2 3 o
		3 2 3 4 o
		3 3 4 1 o
		3 4 1 2 o
		4 1 2 3 4 x`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`11 4 9
		10 1 2 3 4 5 6 7 8 9 10 o
		11 1 2 3 4 5 6 7 8 9 10 11 o
		10 11 10 9 8 7 6 5 4 3 2 x
		10 11 9 1 4 3 7 5 6 2 10 x`)}, wantOut: `8`},
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
