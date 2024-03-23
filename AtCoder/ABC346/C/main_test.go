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
		{name: "1", args: args{stdin: strings.NewReader(`4 5
		1 6 3 1`)}, wantOut: `11`},
		{name: "2", args: args{stdin: strings.NewReader(`1 3
		346`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`10 158260522
		877914575 24979445 623690081 262703497 24979445 1822804784 1430302156 1161735902 923078537 1189330739`)}, wantOut: `12523196466007058`},
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
