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
1 5 7`)}, wantOut: `1
1`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
1 1 1`)}, wantOut: `0
`},
		{name: "3", args: args{stdin: strings.NewReader(`8 3
4 77 20 26 9 26 22 40`)}, wantOut: `4
1 2 5 8`},
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
