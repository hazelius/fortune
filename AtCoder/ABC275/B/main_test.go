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
		{name: "1", args: args{stdin: strings.NewReader(`2 3 5 1 2 4`)}, wantOut: `22`},
		{name: "2", args: args{stdin: strings.NewReader(`1 1 1000000000 0 0 0`)}, wantOut: `1755647`},
		{name: "3", args: args{stdin: strings.NewReader(`1000000000000000000 1000000000000000000 1000000000000000000 1000000000000000000 1000000000000000000 1000000000000000000`)}, wantOut: `0`},
		{name: "4", args: args{stdin: strings.NewReader(`9999999999999999999 9999999999999999999 9999999999999999999 9999999999999999999 9999999999999999999 9999999999999999999`)}, wantOut: `0`},
		{name: "5", args: args{stdin: strings.NewReader(`1000000000000000000 1000000000000000000 1000000000000000000 2 5 6`)}, wantOut: `5`},
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
