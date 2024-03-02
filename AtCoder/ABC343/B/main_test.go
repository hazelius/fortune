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
		{name: "1", args: args{stdin: strings.NewReader(`4
		0 1 1 0
		1 0 0 1
		1 0 0 0
		0 1 0 0`)}, wantOut: `2 3
1 4
1
2
`},
		{name: "2", args: args{stdin: strings.NewReader(`2
		0 0
		0 0`)}, wantOut: ``},
		{name: "3", args: args{stdin: strings.NewReader(`5
		0 1 0 1 1
		1 0 0 1 0
		0 0 0 0 1
		1 1 0 0 1
		1 0 1 1 0`)}, wantOut: `2 4 5
1 4
5
1 2 5
1 3 4
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
