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
		4 4
		0 1 0 1
		1 2
		2 3
		1 3
		2 4
		3 3
		0 1 0
		1 2
		2 3
		1 3
		6 6
		0 0 1 1 0 1
		1 2
		2 6
		3 6
		4 6
		4 5
		2 4`)}, wantOut: `3
-1
3
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
