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
		{name: "1", args: args{stdin: strings.NewReader(`4
		0 0
		2 4
		5 0
		3 4`)}, wantOut: `3
3
1
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`6
		3 2
		1 6
		4 5
		1 3
		5 5
		9 8`)}, wantOut: `6
6
6
6
6
4
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
