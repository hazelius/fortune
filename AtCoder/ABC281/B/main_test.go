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
		{name: "1", args: args{stdin: strings.NewReader(`Q142857Z`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`AB912278C`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`X900000`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`K012345K`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`K0123450I`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`K000000K`)}, wantOut: `No`},
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
