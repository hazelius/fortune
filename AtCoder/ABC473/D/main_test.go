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
		{name: "1", args: args{stdin: strings.NewReader(`3 8`)}, wantOut: `0 1 2
0 4 0
1 2 1
2 0 2
2 3 0
3 1 1
4 2 0
5 0 1
6 1 0
8 0 0
`},
		{name: "2", args: args{stdin: strings.NewReader(`1 200000`)}, wantOut: `200000
`},
		{name: "3", args: args{stdin: strings.NewReader(`8 9`)}, wantOut: `0 0 0 1 1 0 0 0
0 0 1 0 0 1 0 0
0 0 3 0 0 0 0 0
0 1 0 0 0 0 1 0
0 1 1 1 0 0 0 0
0 2 0 0 1 0 0 0
0 3 1 0 0 0 0 0
1 0 0 0 0 0 0 1
1 0 0 2 0 0 0 0
1 0 1 0 1 0 0 0
1 1 0 0 0 1 0 0
1 1 2 0 0 0 0 0
1 2 0 1 0 0 0 0
1 4 0 0 0 0 0 0
2 0 0 0 0 0 1 0
2 0 1 1 0 0 0 0
2 1 0 0 1 0 0 0
2 2 1 0 0 0 0 0
3 0 0 0 0 1 0 0
3 0 2 0 0 0 0 0
3 1 0 1 0 0 0 0
3 3 0 0 0 0 0 0
4 0 0 0 1 0 0 0
4 1 1 0 0 0 0 0
5 0 0 1 0 0 0 0
5 2 0 0 0 0 0 0
6 0 1 0 0 0 0 0
7 1 0 0 0 0 0 0
9 0 0 0 0 0 0 0
`},
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
