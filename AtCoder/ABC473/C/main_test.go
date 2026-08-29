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
		{name: "1", args: args{stdin: strings.NewReader(`8 5
3 3 5 5 4 4 3 2`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`6 1
1 1 1 1 1 1`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`14 8
6 1 5 3 8 4 3 4 3 5 1 2 5 1`)}, wantOut: `4`},
		{name: "4", args: args{stdin: strings.NewReader(`8 5
1 1 1 1 2 2 4 2`)}, wantOut: `2`},
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
