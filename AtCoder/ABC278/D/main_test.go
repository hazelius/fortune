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
		3 1 4 1 5
		6
		3 2
		2 3 4
		3 3
		1 1
		2 3 4
		3 3`)}, wantOut: `1
8
5
`},
		{name: "2", args: args{stdin: strings.NewReader(`1
		1000000000
		8
		2 1 1000000000
		2 1 1000000000
		2 1 1000000000
		2 1 1000000000
		2 1 1000000000
		2 1 1000000000
		2 1 1000000000
		3 1`)}, wantOut: `8000000000
`},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 8 4 15 7 5 7 5 8 0
20
2 7 0
3 7
3 8
1 7
3 3
2 4 4
2 4 9
2 10 5
1 10
2 4 2
1 10
2 3 1
2 8 11
2 3 14
2 1 9
3 8
3 8
3 1
2 6 5
3 7
`)}, wantOut: `7
5
7
21
21
19
10
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
