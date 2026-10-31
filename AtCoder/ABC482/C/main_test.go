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
		{name: "1", args: args{stdin: strings.NewReader(`5 7 5 6
1 1
2 3
2 5
4 3
5 7
1 2 4
2 3 5
1 3 3
2 1 7
1 2 2
2 7 7`)}, wantOut: `3
3
0
5
2
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`17 23 12 12
14 8
2 19
17 3
7 11
2 4
10 19
7 8
1 23
14 11
2 23
17 19
7 2
1 2 7
2 9 20
1 3 6
2 8 8
1 11 17
2 1 23
1 1 1
2 20 22
1 7 14
2 3 11
1 15 17
2 19 23`)}, wantOut: `6
5
0
2
4
12
1
0
6
6
2
5
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
