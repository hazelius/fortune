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
		{name: "1", args: args{stdin: strings.NewReader(`3 5
		xooox
		oooxx
		oooxo`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
		oxo
		oxo
		oxo`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`3 3
		oox
		oxo
		xoo`)}, wantOut: `0`},
		{name: "4", args: args{stdin: strings.NewReader(`1 7
		ooooooo`)}, wantOut: `7`},
		{name: "5", args: args{stdin: strings.NewReader(`5 15
		oxooooooooooooo
		oxooxooooooooox
		oxoooooooooooox
		oxxxooooooxooox
		oxooooooooxooox`)}, wantOut: `5`},
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
