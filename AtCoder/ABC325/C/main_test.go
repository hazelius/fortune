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
		{name: "1", args: args{stdin: strings.NewReader(`5 6
		.##...
		...#..
		....##
		#.#...
		..#...`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
		#.#
		.#.
		#.#`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`4 2
		..
		..
		..
		..`)}, wantOut: `0`},
		{name: "4", args: args{stdin: strings.NewReader(`5 47
		.#..#..#####..#...#..#####..#...#...###...#####
		.#.#...#.......#.#...#......##..#..#...#..#....
		.##....#####....#....#####..#.#.#..#......#####
		.#.#...#........#....#......#..##..#...#..#....
		.#..#..#####....#....#####..#...#...###...#####`)}, wantOut: `7`},
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
