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
		{name: "1", args: args{stdin: strings.NewReader(`3 1 9
		2 5 6
		2 7 1`)}, wantOut: `0.666666666667`},
		{name: "2", args: args{stdin: strings.NewReader(`7 7 6
		8 6 8
		7 7 6`)}, wantOut: `0.004982363316`},
		{name: "3", args: args{stdin: strings.NewReader(`3 6 7
		1 9 7
		5 7 5`)}, wantOut: `0.400000000000`},
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
