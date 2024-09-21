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
		{name: "1", args: args{stdin: strings.NewReader(`6`)}, wantOut: `2
1 1`},
		{name: "2", args: args{stdin: strings.NewReader(`100`)}, wantOut: `4
4 2 2 0`},
		{name: "3", args: args{stdin: strings.NewReader(`59048`)}, wantOut: `20
9 9 8 8 7 7 6 6 5 5 4 4 3 3 2 2 1 1 0 0`},
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
