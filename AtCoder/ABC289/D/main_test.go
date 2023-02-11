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
		3 4 5
		4
		4 5 6 8
		15`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`4
		2 3 4 5
		4
		3 4 5 6
		8`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`4
		2 5 7 8
		5
		2 9 10 11 19
		20`)}, wantOut: `Yes`},
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
