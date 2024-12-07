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
		{name: "1", args: args{stdin: strings.NewReader(`2 5 1
.###.
.#.##`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`5 5 2
.#.#.
.....
.#.#.
#.#.#
.....`)}, wantOut: `15`},
		{name: "3", args: args{stdin: strings.NewReader(`4 4 2
....
.##.
.##.
....`)}, wantOut: `10`},
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
