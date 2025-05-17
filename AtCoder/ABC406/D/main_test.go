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
		{name: "1", args: args{stdin: strings.NewReader(`3 4 5
1 2
1 3
3 4
3 1
2 2
5
1 1
1 2
2 2
2 4
1 2`)}, wantOut: `2
1
0
1
0
`},
		{name: "2", args: args{stdin: strings.NewReader(`1 2 1
1 2
7
2 1
2 1
2 1
2 1
2 1
2 1
2 1
`)}, wantOut: `0
0
0
0
0
0
0
`},
		{name: "3", args: args{stdin: strings.NewReader(`4 4 16
1 1
1 2
1 3
1 4
2 1
2 2
2 3
2 4
3 1
3 2
3 3
3 4
4 1
4 2
4 3
4 4
7
2 1
1 1
2 2
1 2
2 3
1 3
2 4`)}, wantOut: `4
3
3
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
