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
		{name: "1", args: args{stdin: strings.NewReader(`4 6 6
2 2
1 4
3 1
1 2
3 1
4 2`)}, wantOut: `2 1
4 1
1 4
1 1
1 3
1 5
`},
		{name: "2", args: args{stdin: strings.NewReader(`100 100 20
72 7
4 74
17 67
20 67
96 3
1 95
5 81
87 3
3 87
16 67
95 8
100 2
92 1
3 98
92 5
87 3
1 98
19 67
11 74
87 1`)}, wantOut: `7 89
90 22
62 22
7 22
4 96
99 1
94 15
7 19
4 9
27 22
4 1
1 99
7 14
1 1
7 9
7 16
100 1
43 22
79 22
7 15
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
