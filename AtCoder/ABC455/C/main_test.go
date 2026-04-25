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
		{name: "1", args: args{stdin: strings.NewReader(`6 2
7 2 7 2 2 9`)}, wantOut: `6`},
		{name: "2", args: args{stdin: strings.NewReader(`8 6
1 2 3 4 1 2 3 4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 2
3 3 4 1 1 3 3 1 5 1`)}, wantOut: `8`},
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
