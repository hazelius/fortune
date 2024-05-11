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
		{name: "1", args: args{stdin: strings.NewReader(`7 6
		2 5 1 4 1 2 3`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`7 10
		1 10 1 10 1 10 1`)}, wantOut: `7`},
		{name: "3", args: args{stdin: strings.NewReader(`15 100
		73 8 55 26 97 48 37 47 35 55 5 17 62 2 60`)}, wantOut: `8`},
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
