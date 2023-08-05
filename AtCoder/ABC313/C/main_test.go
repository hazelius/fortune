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
		4 7 3 7`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`1
		313`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10
		999999997 999999999 4 3 2 4 999999990 8 999999991 999999993`)}, wantOut: `2499999974`},
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
