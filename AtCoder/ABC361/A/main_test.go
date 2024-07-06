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
		{name: "1", args: args{stdin: strings.NewReader(`4 3 7
2 3 5 11`)}, wantOut: `2 3 5 7 11`},
		{name: "2", args: args{stdin: strings.NewReader(`1 1 100
100`)}, wantOut: `100 100`},
		{name: "3", args: args{stdin: strings.NewReader(`8 8 3
9 9 8 2 4 4 3 5`)}, wantOut: `9 9 8 2 4 4 3 5 3`},
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
