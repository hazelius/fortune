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
		{name: "1", args: args{stdin: strings.NewReader(`6 6
		######
		#....#
		#.#..#
		#..#.#
		#....#
		######`)}, wantOut: `12`},
		{name: "2", args: args{stdin: strings.NewReader(`21 25
		#########################
		#..............###...####
		#..............#..#...###
		#........###...#...#...##
		#........#..#..#........#
		#...##...#..#..#...#....#
		#..#..#..###...#..#.....#
		#..#..#..#..#..###......#
		#..####..#..#...........#
		#..#..#..###............#
		#..#..#.................#
		#........##.............#
		#.......#..#............#
		#..........#....#.......#
		#........###...##....#..#
		#..........#..#.#...##..#
		#.......#..#....#..#.#..#
		##.......##.....#....#..#
		###.............#....#..#
		####.................#..#
		#########################`)}, wantOut: `215`},
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
