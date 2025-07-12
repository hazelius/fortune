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
		{name: "1", args: args{stdin: strings.NewReader(`7 3
5 10 15 20 8 14 15`)}, wantOut: `6`},
		{name: "2", args: args{stdin: strings.NewReader(`7 7
5 10 15 20 8 14 15`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`7 1
5 10 15 20 8 14 15`)}, wantOut: `15`},
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
