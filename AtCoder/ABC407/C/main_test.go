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
		{name: "1", args: args{stdin: strings.NewReader(`21`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`407`)}, wantOut: `17`},
		{name: "3", args: args{stdin: strings.NewReader(`2025524202552420255242025524`)}, wantOut: `150`},
		{name: "4", args: args{stdin: strings.NewReader(`5`)}, wantOut: `6`},
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
