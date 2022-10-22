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
		{name: "1", args: args{stdin: strings.NewReader(`3 -1 1
		2 1 3`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`5 2 0
		2 2 2 2 2`)}, wantOut: `Yes`},
		{name: "3", args: args{stdin: strings.NewReader(`4 5 5
		1 2 3 4`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`3 2 7
		2 7 4`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`10 8 -7
		6 10 4 1 5 9 8 6 5 1`)}, wantOut: `Yes`},
		{name: "6", args: args{stdin: strings.NewReader(`2 2 -1
		2 1`)}, wantOut: `Yes`},
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
