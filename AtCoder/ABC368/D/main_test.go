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
		{name: "1", args: args{stdin: strings.NewReader(`7 3
		1 2
		1 3
		2 4
		2 5
		3 6
		3 7
		1 3 5`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`4 4
				3 1
				1 4
				2 1
				1 2 3 4`)}, wantOut: `4`},
		{name: "3", args: args{stdin: strings.NewReader(`5 1
				1 4
				2 3
				5 2
				1 2
				1`)}, wantOut: `1`},
		{name: "4", args: args{stdin: strings.NewReader(`8 4
		1 2
		1 3
		2 4
		2 5
		3 6
		3 7
		7 8
		2 4 7 8`)}, wantOut: `6`},
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
