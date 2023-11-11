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
		{name: "1", args: args{stdin: strings.NewReader(`12
		31 29 31 30 31 30 31 31 30 31 30 31`)}, wantOut: `13`},
		{name: "2", args: args{stdin: strings.NewReader(`10
		10 1 2 3 4 5 6 7 8 100`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`30
		73 8 55 26 97 48 37 47 35 55 5 17 62 2 60 23 99 73 34 75 7 46 82 84 29 41 32 31 52 32`)}, wantOut: `15`},
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
