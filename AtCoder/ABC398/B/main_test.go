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
		{name: "1", args: args{stdin: strings.NewReader(`1 4 1 4 2 1 3`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`11 12 13 10 13 12 11`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`7 7 7 7 7 7 7`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`13 13 1 1 7 4 13`)}, wantOut: `Yes`},
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
