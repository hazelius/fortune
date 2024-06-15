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
oooxx
xooox
xxooo`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
oo
ox
xo`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`8 6
xxoxxo
xxoxxx
xoxxxx
xxxoxx
xxoooo
xxxxox
xoxxox
oxoxxo`)}, wantOut: `3`},
		{name: "4", args: args{stdin: strings.NewReader(`3 6
xxxxxx
xxoxxx
ooxooo`)}, wantOut: `2`},
		{name: "5", args: args{stdin: strings.NewReader(`3 6
xxoxxx
ooxooo
oooooo`)}, wantOut: `1`},
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
