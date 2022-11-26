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
		{name: "1", args: args{stdin: strings.NewReader(`10 1`)}, wantOut: `7.7735026919`},
		{name: "2", args: args{stdin: strings.NewReader(`5 10`)}, wantOut: `5.0000000000`},
		{name: "3", args: args{stdin: strings.NewReader(`1000000000000000000 100`)}, wantOut: `8772053214538.5976562500`},
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
