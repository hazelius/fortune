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
		10 2
		10 1
		10 2
		3 2`)}, wantOut: `Takahashi`},
		{name: "2", args: args{stdin: strings.NewReader(`6
		5 4
		4 5
		2 4
		1 6
		7 1
		3 2`)}, wantOut: `Draw`},
		{name: "3", args: args{stdin: strings.NewReader(`4
		0 0
		10 10
		50 50
		0 100`)}, wantOut: `Aoki`},
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
