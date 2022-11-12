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
		{name: "1", args: args{stdin: strings.NewReader(`4
		H3
		DA
		D3
		SK`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		H3
		DA
		CK
		H3
		S7`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`4
		3H
		AD
		3D
		KS`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`5
		00
		AA
		XX
		YY
		ZZ`)}, wantOut: `No`},
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
