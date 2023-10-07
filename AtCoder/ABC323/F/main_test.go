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
		{name: "1", args: args{stdin: strings.NewReader(`1 2 3 3 0 5`)}, wantOut: `9`},
		{name: "2", args: args{stdin: strings.NewReader(`0 0 1 0 -1 0`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`-100000000000000000 -100000000000000000 100000000000000000 100000000000000000 -100000000000000000 -100000000000000000`)}, wantOut: `800000000000000003`},
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
