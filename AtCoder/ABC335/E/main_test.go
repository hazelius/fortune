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
		{name: "1", args: args{stdin: strings.NewReader(`5 6
		10 20 30 40 50
		1 2
		1 3
		2 5
		3 4
		3 5
		4 5`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`4 5
		1 10 11 4
		1 2
		1 3
		2 3
		2 4
		3 4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 12
		1 2 3 3 4 4 4 6 5 7
		1 3
		2 9
		3 4
		5 6
		1 2
		8 9
		4 5
		8 10
		7 10
		4 6
		2 8
		6 7`)}, wantOut: `5`},
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
