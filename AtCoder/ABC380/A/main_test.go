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
		{name: "1", args: args{stdin: strings.NewReader(`123233`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`123234`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`323132`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`500000`)}, wantOut: `No`},
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
