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
		1 3
		4 2
		3 2`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`2 0`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`5 5
		1 2
		2 3
		3 4
		4 5
		5 1`)}, wantOut: `No`},
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
