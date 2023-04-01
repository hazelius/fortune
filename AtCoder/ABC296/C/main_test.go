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
		{name: "1", args: args{stdin: strings.NewReader(`6 5
		3 1 4 1 5 9`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`6 -4
		-2 -7 -1 -8 -2 -8`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`2 0
		141421356 17320508`)}, wantOut: `Yes`},
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
