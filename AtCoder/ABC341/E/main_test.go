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
		10100
		2 1 3
		2 1 5
		1 1 4
		2 1 5
		1 3 3
		2 2 4`)}, wantOut: `Yes
No
Yes
No
`},
		{name: "2", args: args{stdin: strings.NewReader(`1 2
		1
		1 1 1
		2 1 1`)}, wantOut: `Yes
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
