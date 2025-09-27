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
		{name: "1", args: args{stdin: strings.NewReader(`4 3
3 1 4 5
2 1 3
1 1
2 2 3`)}, wantOut: `8
9
`},
		{name: "2", args: args{stdin: strings.NewReader(`5 7
1 2 4 8 16
2 1 5
1 4
1 5
2 1 5
2 2 4
1 1
2 3 3`)}, wantOut: `31
31
7
4
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
