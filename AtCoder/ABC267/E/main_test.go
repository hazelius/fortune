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
		{name: "1", args: args{stdin: strings.NewReader(`4 3
		3 1 4 2
		1 2
		1 3
		4 1`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`7 13
		464 661 847 514 74 200 188
		5 1
		7 1
		5 7
		4 1
		4 5
		2 4
		5 2
		1 3
		1 6
		3 5
		1 2
		4 6
		2 7`)}, wantOut: `1199`},
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
