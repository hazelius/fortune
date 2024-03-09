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
		2 1 4 3
		4
		2 1
		1 4 5
		2 2
		1 5 1`)}, wantOut: `4 5 1 3 `},
		{name: "2", args: args{stdin: strings.NewReader(`6
		3 1 4 5 9 2
		7
		2 5
		1 3 5
		1 9 7
		2 9
		2 3
		1 2 3
		2 4`)}, wantOut: `5 1 7 2 3 `},
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
