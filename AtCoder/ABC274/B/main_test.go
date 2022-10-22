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
		#..#
		.#.#
		.#.#`)}, wantOut: `1 2 0 3 `},
		{name: "2", args: args{stdin: strings.NewReader(`3 7
		.......
		.......
		.......`)}, wantOut: `0 0 0 0 0 0 0 `},
		{name: "3", args: args{stdin: strings.NewReader(`8 3
		.#.
		###
		.#.
		.#.
		.##
		..#
		##.
		.##`)}, wantOut: `2 7 4 `},
		{name: "4", args: args{stdin: strings.NewReader(`5 47
		.#..#..#####..#...#..#####..#...#...###...#####
		.#.#...#.......#.#...#......##..#..#...#..#....
		.##....#####....#....#####..#.#.#..#......#####
		.#.#...#........#....#......#..##..#...#..#....
		.#..#..#####....#....#####..#...#...###...#####`)}, wantOut: `0 5 1 2 2 0 0 5 3 3 3 3 0 0 1 1 3 1 1 0 0 5 3 3 3 3 0 0 5 1 1 1 5 0 0 3 2 2 2 2 0 0 5 3 3 3 3 `},
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
