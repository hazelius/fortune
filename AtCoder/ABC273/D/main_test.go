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
		{name: "1", args: args{stdin: strings.NewReader(`5 5 4 4
						3
						5 3
						2 2
						1 4
						4
						L 2
						U 3
						L 2
						R 4`)}, wantOut: `4 2
3 2
3 1
3 5
`},
		{name: "2", args: args{stdin: strings.NewReader(`6 6 6 3
		7
		3 1
		4 3
		2 6
		3 4
		5 5
		1 1
		3 2
		10
		D 3
		U 3
		L 2
		D 2
		U 3
		D 3
		U 3
		R 3
		L 3
		D 1`)}, wantOut: `6 3
5 3
5 1
6 1
4 1
6 1
4 1
4 2
4 1
5 1
`},
		{name: "3", args: args{stdin: strings.NewReader(`6 7 5 5
3
6 7
5 2
1 7
1
U 5`)}, wantOut: `1 5
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
