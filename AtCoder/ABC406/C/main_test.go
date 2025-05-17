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
		{name: "1", args: args{stdin: strings.NewReader(`6
1 3 6 4 2 5`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`6
1 2 3 4 5 6`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`12
11 3 8 9 5 2 10 4 1 6 12 7`)}, wantOut: `4`},
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
