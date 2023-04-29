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
		{name: "1", args: args{stdin: strings.NewReader(`5 9
		#.#.#...#
		.#...#.#.
		#.#...#..
		.....#.#.
		....#...#`)}, wantOut: `1 1 0 0 0`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
		...
		...
		...`)}, wantOut: `0 0 0`},
		{name: "3", args: args{stdin: strings.NewReader(`3 16
		#.#.....#.#..#.#
		.#.......#....#.
		#.#.....#.#..#.#`)}, wantOut: `3 0 0`},
		{name: "4", args: args{stdin: strings.NewReader(`15 20
		#.#..#.............#
		.#....#....#.#....#.
		#.#....#....#....#..
		........#..#.#..#...
		#.....#..#.....#....
		.#...#....#...#..#.#
		..#.#......#.#....#.
		...#........#....#.#
		..#.#......#.#......
		.#...#....#...#.....
		#.....#..#.....#....
		........#.......#...
		#.#....#....#.#..#..
		.#....#......#....#.
		#.#..#......#.#....#`)}, wantOut: `5 0 1 0 0 0 1 0 0 0 0 0 0 0 0`},
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
