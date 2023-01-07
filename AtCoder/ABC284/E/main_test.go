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
		{name: "1", args: args{stdin: strings.NewReader(`4 2
		1 2
		2 3`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`4 6
		1 2
		1 3
		1 4
		2 3
		2 4
		3 4`)}, wantOut: `16`},
		{name: "3", args: args{stdin: strings.NewReader(`8 21
		2 6
		1 3
		5 6
		3 8
		3 6
		4 7
		4 6
		3 4
		1 5
		2 4
		1 2
		2 7
		1 4
		3 5
		2 5
		2 3
		4 5
		3 7
		6 7
		5 7
		2 8`)}, wantOut: `2023`},
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
