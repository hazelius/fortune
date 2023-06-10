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
		{name: "1", args: args{stdin: strings.NewReader(`53`)}, wantOut: `55`},
		{name: "2", args: args{stdin: strings.NewReader(`21`)}, wantOut: `20`},
		{name: "3", args: args{stdin: strings.NewReader(`100`)}, wantOut: `100`},
		{name: "4", args: args{stdin: strings.NewReader(`19`)}, wantOut: `20`},
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
