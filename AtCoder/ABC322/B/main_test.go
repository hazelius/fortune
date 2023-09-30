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
		{name: "1", args: args{stdin: strings.NewReader(`3 7
		abc
		abcdefg`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`3 4
		abc
		aabc`)}, wantOut: `2`},
		{name: "3", args: args{stdin: strings.NewReader(`3 3
		abc
		xyz`)}, wantOut: `3`},
		{name: "4", args: args{stdin: strings.NewReader(`3 3
		aaa
		aaa`)}, wantOut: `0`},
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
