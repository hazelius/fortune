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
		{name: "1", args: args{stdin: strings.NewReader(`.o..o
..oo.
.o..o
o.o..
o..oo`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`o..oo
...o.
o.ooo
oo.o.
.o.oo`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`o..o.
.o...
..o.o
..o..
o...o`)}, wantOut: `1`},
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
