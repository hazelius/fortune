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
		{name: "1", args: args{stdin: strings.NewReader(`30 6
		2 9 18 27 18 9`)}, wantOut: `28`},
		{name: "2", args: args{stdin: strings.NewReader(`1 7
		1 1 1 1 1 1 1`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`9 20
		9 5 1 2 2 2 8 9 2 1 6 2 6 5 8 7 8 5 9 8`)}, wantOut: `5`},
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
