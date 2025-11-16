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
		{name: "1", args: args{stdin: strings.NewReader(`3 6 8
11 10 13`)}, wantOut: `18`},
		{name: "2", args: args{stdin: strings.NewReader(`2 3 4
3 5`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`8 4 32
1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000 1000000000`)}, wantOut: `8000000000`},
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
