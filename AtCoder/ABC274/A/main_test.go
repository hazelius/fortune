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
		{name: "1", args: args{stdin: strings.NewReader(`7 4`)}, wantOut: `0.571`},
		{name: "2", args: args{stdin: strings.NewReader(`7 3`)}, wantOut: `0.429`},
		{name: "3", args: args{stdin: strings.NewReader(`2 1`)}, wantOut: `0.500`},
		{name: "4", args: args{stdin: strings.NewReader(`10 10`)}, wantOut: `1.000`},
		{name: "5", args: args{stdin: strings.NewReader(`1 0`)}, wantOut: `0.000`},
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
