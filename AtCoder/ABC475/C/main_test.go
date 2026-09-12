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
		{name: "1", args: args{stdin: strings.NewReader(`6 3 10
5 2 4 1 6`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`8 8 17
2 3 4 4 3 5 1`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`2 1 1000000000000000000
10000`)}, wantOut: `2`},
		{name: "4", args: args{stdin: strings.NewReader(`9 6 28
5 4 9 2 3 6 1 4`)}, wantOut: `6`},
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
