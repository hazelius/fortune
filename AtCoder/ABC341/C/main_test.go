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
		{name: "1", args: args{stdin: strings.NewReader(`6 7 5
		LULDR
		#######
		#...#.#
		##...##
		#.#...#
		#...#.#
		#######`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`13 16 9
		ULURDLURD
		################
		##..##.#..####.#
		###.#..#.....#.#
		#..##..#####.###
		#...#..#......##
		###.##.#..#....#
		##.#####....##.#
		###.###.#.#.#..#
		######.....##..#
		#...#.#.######.#
		##..###..#..#.##
		#...#.#.#...#..#
		################`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`6 7 5
		DDDLL
		#######
		#...#.#
		##...##
		#.#...#
		#...#.#
		#######`)}, wantOut: `1`},
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
