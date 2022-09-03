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
		{name: "1", args: args{stdin: strings.NewReader(`0101110101`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`0100101001`)}, wantOut: `Yes`},
		{name: "3", args: args{stdin: strings.NewReader(`0000100110`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`1101110101`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`0101100000`)}, wantOut: `No`},
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
