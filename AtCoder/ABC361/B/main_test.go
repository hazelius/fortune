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
		{name: "1", args: args{stdin: strings.NewReader(`0 0 0 4 5 6
				2 3 4 5 6 7`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`0 0 0 2 2 2
		0 0 2 2 2 4`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`0 0 0 1000 1000 1000
		10 10 10 100 100 100`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`1 0 0 1000 1000 1000
		0 0 0 1000 1000 1000`)}, wantOut: `Yes`},
		{name: "5", args: args{stdin: strings.NewReader(`0 0 0 1000 1000 1000
		0 0 0 1000 1000 1000`)}, wantOut: `Yes`},
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
