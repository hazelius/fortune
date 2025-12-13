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
		{name: "1", args: args{stdin: strings.NewReader(`3 4
..a.
####
ba#b`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`3 4
..a.
####
b.#b`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`4 4
xxxx
xxxx
xxxx
xxxx`)}, wantOut: `1`},
		{name: "4", args: args{stdin: strings.NewReader(`7 11
u..#y..#...
k..#.z.#.k.
iju#...#x..
###########
..x#.t.#..n
abc#y..#...
..z#..t#.y.`)}, wantOut: `12`},
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
