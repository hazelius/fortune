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
		{name: "1", args: args{stdin: strings.NewReader(`4 5
3
3 1 2
3
3 2 1
2
2 3
4
2 5 3 1`)}, wantOut: `3
2
0
5
`},
		{name: "2", args: args{stdin: strings.NewReader(`6 5
1
3
2
3 5
5
5 3 1 4 2
5
5 1 3 4 2
5
3 4 1 5 2
5
5 1 3 2 4`)}, wantOut: `3
5
1
4
2
0
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
