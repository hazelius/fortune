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
		3 4 5 4
		1 4`)}, wantOut: `7`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
		1 1 1
		1000000000 1000000000 1000000000`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`7 3
		2 6 8 9 5 1 11
		3 5 7`)}, wantOut: `19`},
		// {name: "4", args: args{stdin: strings.NewReader(`7 3
		// 2 6 11 9 5 1 10
		// 3 5 7`)}, wantOut: `19`},
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
