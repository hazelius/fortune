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
		{name: "1", args: args{stdin: strings.NewReader(`..........
		..........
		..........
		..........
		...######.
		...######.
		...######.
		...######.
		..........
		..........`)}, wantOut: `5 8
4 9`},
		{name: "2", args: args{stdin: strings.NewReader(`..........
		..#.......
		..........
		..........
		..........
		..........
		..........
		..........
		..........
		..........`)}, wantOut: `2 2
3 3`},
		{name: "3", args: args{stdin: strings.NewReader(`##########
		##########
		##########
		##########
		##########
		##########
		##########
		##########
		##########
		##########`)}, wantOut: `1 10
1 10`},
		{name: "4", args: args{stdin: strings.NewReader(`..........
		....######
		....######
		....######
		....######
		....######
		..........
		..........
		..........
		..........`)}, wantOut: `2 6
5 10`},
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
