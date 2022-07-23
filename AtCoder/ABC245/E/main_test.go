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
		{name: "1", args: args{stdin: strings.NewReader(`2 3
		2 4
		3 2
		8 1 5
		2 10 5`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
		1 1
		2 2
		100 1
		100 1`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`1 1
		10
		100
		100
		10`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`1 1
		10
		100
		10
		100`)}, wantOut: `Yes`},
		{name: "5", args: args{stdin: strings.NewReader(`3 3
		2 4 1
		3 2 10
		8 1 5
		10 10 10`)}, wantOut: `Yes`},
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
