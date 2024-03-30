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
		{name: "1", args: args{stdin: strings.NewReader(`5 2
		2 5 6 7 10`)}, wantOut: `1 3 5 `},
		{name: "2", args: args{stdin: strings.NewReader(`3 1
		3 4 7`)}, wantOut: `3 4 7 `},
		{name: "3", args: args{stdin: strings.NewReader(`5 10
		50 51 54 60 65`)}, wantOut: `5 6 `},
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
