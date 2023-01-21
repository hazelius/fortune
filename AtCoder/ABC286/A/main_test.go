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
		{name: "1", args: args{stdin: strings.NewReader(`8 1 3 5 7
		1 2 3 4 5 6 7 8`)}, wantOut: `5 6 7 4 1 2 3 8`},
		{name: "2", args: args{stdin: strings.NewReader(`5 2 3 4 5
		2 2 1 1 1`)}, wantOut: `2 1 1 2 1`},
		{name: "3", args: args{stdin: strings.NewReader(`2 1 1 2 2
		50 100`)}, wantOut: `100 50`},
		{name: "4", args: args{stdin: strings.NewReader(`10 2 4 7 9
		22 75 26 45 72 81 47 29 97 2`)}, wantOut: `22 47 29 97 72 81 75 26 45 2`},
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
