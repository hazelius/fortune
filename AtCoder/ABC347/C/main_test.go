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
		{name: "1", args: args{stdin: strings.NewReader(`3 2 5
		1 2 9`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`2 5 10
		10 15`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`4 347 347
		347 700 705 710`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`3 1 1
		2 4 3`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`3 4 8
		1 2 11`)}, wantOut: `Yes`},
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
