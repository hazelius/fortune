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
		{name: "1", args: args{stdin: strings.NewReader(`4 4
1 2
1 3
2 4
3 4`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`5 0`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
7 9
4 6
6 10
2 5
5 6
5 9
6 8
4 8
1 5
1 4`)}, wantOut: `2`},
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
