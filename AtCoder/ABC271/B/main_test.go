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
		{name: "1", args: args{stdin: strings.NewReader(`2 2
		3 1 4 7
		2 5 9
		1 3
		2 1`)}, wantOut: `7
5
`},
		{name: "2", args: args{stdin: strings.NewReader(`3 4
		4 128 741 239 901
		2 1 1
		3 314 159 26535
		1 1
		2 2
		3 3
		1 4`)}, wantOut: `128
1
26535
901
`},
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
