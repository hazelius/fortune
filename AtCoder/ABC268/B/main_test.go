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
		{name: "1", args: args{stdin: strings.NewReader(`atco
		atcoder`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`code
		atcoder`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`abc
		abc`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`aaaa
		aa`)}, wantOut: `No`},
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
