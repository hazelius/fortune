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
		{name: "1", args: args{stdin: strings.NewReader(`3 2 10 20`)}, wantOut: `20.0000000`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2 20 20`)}, wantOut: `32.0000000`},
		{name: "3", args: args{stdin: strings.NewReader(`314159265358979323 4 223606797 173205080`)}, wantOut: `6418410657.7408381`},
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
