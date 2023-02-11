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
		{name: "1", args: args{stdin: strings.NewReader(`5 3
		1 3 4`)}, wantOut: `2 1 5 4 3 `},
		{name: "2", args: args{stdin: strings.NewReader(`5 0`)}, wantOut: `1 2 3 4 5 `},
		{name: "3", args: args{stdin: strings.NewReader(`10 6
		1 2 3 7 8 9`)}, wantOut: `4 3 2 1 5 6 10 9 8 7 `},
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
