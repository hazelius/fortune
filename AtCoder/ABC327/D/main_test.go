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
		{name: "1", args: args{stdin: strings.NewReader(`3 2
		1 2
		2 3`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
		1 2 3
		2 3 1`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`10 1
		1
		1`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`7 8
		1 6 2 7 5 4 2 2
		3 2 7 2 1 2 3 3`)}, wantOut: `Yes`},
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
