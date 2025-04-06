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
		{name: "1", args: args{stdin: strings.NewReader(`7 3`)}, wantOut: `400`},
		{name: "2", args: args{stdin: strings.NewReader(`1000000 2`)}, wantOut: `inf`},
		{name: "3", args: args{stdin: strings.NewReader(`999999999 1`)}, wantOut: `1000000000`},
		{name: "4", args: args{stdin: strings.NewReader(`998244353 99`)}, wantOut: `inf`},
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
