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
		{name: "1", args: args{stdin: strings.NewReader(`3
				3
				1 2
				3`)}, wantOut: `1
12
`},
		{name: "2", args: args{stdin: strings.NewReader(`3
				1 5
				2
				3`)}, wantOut: `5
`},
		{name: "3", args: args{stdin: strings.NewReader(`11
		1 9
		1 9
		1 8
		1 2
		1 4
		1 4
		1 3
		1 5
		1 3
		2
		3`)}, wantOut: `0
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
