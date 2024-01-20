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
		{name: "1", args: args{stdin: strings.NewReader(`3 4 3
		xo.x
		..o.
		xx.o`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`4 2 3
		.o
		.o
		.o
		.o`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`3 3 3
		x..
		..x
		.x.`)}, wantOut: `-1`},
		{name: "4", args: args{stdin: strings.NewReader(`10 12 6
		......xo.o..
		x...x.....o.
		x...........
		..o...x.....
		.....oo.....
		o.........x.
		ox.oox.xx..x
		....o...oox.
		..o.....x.x.
		...o........`)}, wantOut: `3`},
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
