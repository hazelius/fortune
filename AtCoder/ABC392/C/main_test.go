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
4 3 2 1
2 3 1 4`)}, wantOut: `3 4 1 2`},
		{name: "2", args: args{stdin: strings.NewReader(`10
2 6 4 3 7 8 9 10 1 5
1 4 8 2 10 5 7 3 9 6`)}, wantOut: `4 8 6 5 3 10 9 2 1 7`},
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
