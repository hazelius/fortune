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
		{name: "1", args: args{stdin: strings.NewReader(`9 9
.........
.........
.........
.........
....#....
.........
.........
.........
.........`)}, wantOut: `57`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
..
..`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
..........
....#.....
#.......#.
......#...
.......#..
.....#....
..........
..........
..#...#...
.......#..`)}, wantOut: `64`},
		{name: "4", args: args{stdin: strings.NewReader(`1 4
.###`)}, wantOut: `4`},
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
