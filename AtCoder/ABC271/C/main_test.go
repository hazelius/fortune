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
		{name: "1", args: args{stdin: strings.NewReader(`6
		1 2 4 6 7 271`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`10
		1 1 1 1 1 1 1 1 1 1`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`1 5`)}, wantOut: `0`},
		{name: "4", args: args{stdin: strings.NewReader(`6 1 1 2 4 5 7000`)}, wantOut: `5`},
		{name: "5", args: args{stdin: strings.NewReader(`2 6 40`)}, wantOut: `1`},
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
