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
		{name: "1", args: args{stdin: strings.NewReader(`3 2
		2 7 8`)}, wantOut: `8 0 0`},
		{name: "2", args: args{stdin: strings.NewReader(`3 4
		9 9 9`)}, wantOut: `0 0 0`},
		{name: "3", args: args{stdin: strings.NewReader(`9 5
		1 2 3 4 5 6 7 8 9`)}, wantOut: `6 7 8 9 0 0 0 0 0`},
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
