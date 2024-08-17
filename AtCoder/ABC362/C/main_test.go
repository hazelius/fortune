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
		3 5
		-4 1
		-2 3`)}, wantOut: `Yes
5 -3 -2`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		1 2
		1 2
		1 2`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`6
		-87 12
		-60 -54
		2 38
		-76 6
		87 96
		-17 38`)}, wantOut: `Yes
12 -54 38 -66 87 -17`},
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
