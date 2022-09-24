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
		{name: "1", args: args{stdin: strings.NewReader(`10 -10 1`)}, wantOut: `10`},
		{name: "2", args: args{stdin: strings.NewReader(`20 10 -10`)}, wantOut: `40`},
		{name: "3", args: args{stdin: strings.NewReader(`100 1 1000`)}, wantOut: `-1`},
		{name: "4", args: args{stdin: strings.NewReader(`-10 -5 5`)}, wantOut: `20`},
		{name: "5", args: args{stdin: strings.NewReader(`10 5 -5`)}, wantOut: `20`},
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
