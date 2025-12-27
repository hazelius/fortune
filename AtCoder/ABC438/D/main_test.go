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
1 4 2 4 3
2 3 4 2 2
3 2 4 4 3`)}, wantOut: `16`},
		{name: "2", args: args{stdin: strings.NewReader(`3
1 1 1
1 1 1
1 1 1`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`6
2 10 7 7 7 11
5 7 9 10 9 12
6 6 7 10 12 7`)}, wantOut: `50`},
		{name: "4", args: args{stdin: strings.NewReader(`6
1 4 9 8 6 6
2 1 2 1 1 2
3 6 9 8 3 9`)}, wantOut: `32`},
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
