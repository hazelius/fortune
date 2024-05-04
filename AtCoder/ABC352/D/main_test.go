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
		{name: "1", args: args{stdin: strings.NewReader(`4 2
		2 3 1 4`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`4 1
		2 3 1 4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 5
		10 1 6 8 7 2 5 9 3 4`)}, wantOut: `5`},
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
