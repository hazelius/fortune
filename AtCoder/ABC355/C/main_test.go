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
		{name: "1", args: args{stdin: strings.NewReader(`3 5
		5 1 8 9 7`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`3 5
		4 2 9 7 5`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`4 12
		13 9 6 5 2 7 16 14 8 3 10 11`)}, wantOut: `9`},
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
