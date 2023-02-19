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
		{name: "1", args: args{stdin: strings.NewReader(`3 2
		10 20 30
		1 3`)}, wantOut: `40`},
		{name: "2", args: args{stdin: strings.NewReader(`4 1
		1 1 1 100
		4`)}, wantOut: `100`},
		{name: "3", args: args{stdin: strings.NewReader(`8 4
		22 75 26 45 72 81 47 29
		4 6 7 8`)}, wantOut: `202`},
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
