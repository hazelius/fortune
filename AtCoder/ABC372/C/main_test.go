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
		{name: "1", args: args{stdin: strings.NewReader(`3 3
				ABC
				1 A
				2 B
				3 C`)}, wantOut: `1
1
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`7 4
ABCDABC
4 B
3 A
5 C
4 G`)}, wantOut: `2
1
1
0
`},
		{name: "3", args: args{stdin: strings.NewReader(`15 10
BBCCBCACCBACACA
9 C
11 B
5 B
11 B
4 A
8 C
8 B
5 B
7 B
14 B`)}, wantOut: `0
0
0
0
1
1
2
2
1
1
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
