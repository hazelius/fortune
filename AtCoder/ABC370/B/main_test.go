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
3
2 4
3 1 2
2 1 2 4`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`5
5
5 5
5 5 5
5 5 5 5
5 5 5 5 5`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`6
2
1 5
1 6 3
2 6 1 4
2 1 1 1 6
5 6 1 2 2 5`)}, wantOut: `5`},
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
