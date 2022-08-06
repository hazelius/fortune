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
		{name: "1", args: args{stdin: strings.NewReader(`3
		1 2`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`10
		1 2 3 4 5 6 7 8 9`)}, wantOut: `9`},
		{name: "3", args: args{stdin: strings.NewReader(`10
		1 2 3 4 5 6 7 2 9`)}, wantOut: `3`},
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
