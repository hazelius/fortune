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
		1 5
		4 5
		2 3
		1 4
		3 5
		2 5`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`3 1
		1 2`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`7 10
		1 7
		5 7
		2 5
		3 6
		4 7
		1 5
		2 4
		1 3
		1 6
		2 7`)}, wantOut: `4`},
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
