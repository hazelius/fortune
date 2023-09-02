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
		{name: "1", args: args{stdin: strings.NewReader(`5 2 10
		7 1 6 3 6`)}, wantOut: `20`},
		{name: "2", args: args{stdin: strings.NewReader(`3 1 10
		1 2 3`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`8 3 1000000000
		1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000`)}, wantOut: `3000000000`},
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
