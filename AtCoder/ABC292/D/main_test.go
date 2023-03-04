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
		{name: "1", args: args{stdin: strings.NewReader(`3 3
		2 3
		1 1
		2 3`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`5 5
		1 2
		2 3
		3 4
		3 5
		1 5`)}, wantOut: `Yes`},
		{name: "3", args: args{stdin: strings.NewReader(`13 16
		7 9
		7 11
		3 8
		1 13
		11 11
		6 11
		8 13
		2 11
		3 3
		8 12
		9 11
		1 11
		5 13
		3 12
		6 9
		1 10`)}, wantOut: `No`},
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
