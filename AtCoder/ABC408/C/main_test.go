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
		{name: "1", args: args{stdin: strings.NewReader(`10 4
1 6
4 5
5 10
7 10`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`5 2
1 2
3 4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`5 10
2 5
1 5
1 2
2 4
2 2
5 5
2 4
1 2
2 2
2 3`)}, wantOut: `3`},
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
