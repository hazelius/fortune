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
		{name: "1", args: args{stdin: strings.NewReader(`5
2 1 4 3 5`)}, wantOut: `3 2 2 1 0`},
		{name: "2", args: args{stdin: strings.NewReader(`4
1 2 3 4`)}, wantOut: `3 2 1 0`},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 9 6 5 2 7 10 4 8 3`)}, wantOut: `2 3 3 3 2 1 2 1 1 0`},
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
