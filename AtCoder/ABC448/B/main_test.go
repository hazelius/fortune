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
		{name: "1", args: args{stdin: strings.NewReader(`7 5
4 4 8 3 7
1 2
2 3
5 2
4 10
2 3
5 4
2 3`)}, wantOut: `15`},
		{name: "2", args: args{stdin: strings.NewReader(`1 1
1
1 1`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`15 10
7 94 100 82 63 81 75 2 76 73
10 44
5 77
10 47
7 32
2 82
5 90
3 37
6 70
6 28
3 25
2 26
10 56
1 69
5 46
7 26`)}, wantOut: `438`},
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
