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
		{name: "1", args: args{stdin: strings.NewReader(`3
		2 1 3
		3
		2 3
		1 2
		1 3`)}, wantOut: `2
2
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`7
		3 7 2 1 6 5 4
		13
		2 3
		1 2
		1 3
		3 6
		3 7
		2 4
		3 7
		1 3
		4 7
		1 6
		2 4
		1 3
		1 3`)}, wantOut: `3
2
3
3
3
2
3
3
7
1
2
3
3
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
