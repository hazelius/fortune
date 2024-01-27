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
		{name: "1", args: args{stdin: strings.NewReader(`2
		800 300
		100 100
		200 10`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`2
		800 300
		100 0
		0 10`)}, wantOut: `38`},
		{name: "3", args: args{stdin: strings.NewReader(`2
		800 300
		801 300
		800 301`)}, wantOut: `0`},
		{name: "4", args: args{stdin: strings.NewReader(`10
		1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000
		0 1 2 3 4 5 6 7 8 9
		9 8 7 6 5 4 3 2 1 0`)}, wantOut: `222222`},
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
