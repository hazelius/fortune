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
		{name: "1", args: args{stdin: strings.NewReader(`5 4
1 3
4 5
1 4
4 2`)}, wantOut: `0 3 1 0 1`},
		{name: "2", args: args{stdin: strings.NewReader(`7 8
3 1
5 4
2 5
5 7
2 3
6 2
3 4
5 1`)}, wantOut: `2 0 0 4 0 0 1`},
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
