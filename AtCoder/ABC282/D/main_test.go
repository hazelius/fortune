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
		{name: "1", args: args{stdin: strings.NewReader(`5 4
		4 2
		3 1
		5 2
		3 2`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`4 3
		3 1
		3 2
		1 2`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`9 11
		4 9
		9 1
		8 2
		8 3
		9 2
		8 4
		6 7
		4 6
		7 5
		4 5
		7 8`)}, wantOut: `9`},
		{name: "4", args: args{stdin: strings.NewReader(`5 5
		1 2
		1 3
		3 4
		4 5
		5 1`)}, wantOut: `1`},
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
