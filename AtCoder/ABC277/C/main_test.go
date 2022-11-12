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
		1 4
		4 3
		4 10
		8 3`)}, wantOut: `10`},
		{name: "2", args: args{stdin: strings.NewReader(`6
		1 3
		1 5
		1 12
		3 5
		3 12
		5 12`)}, wantOut: `12`},
		{name: "3", args: args{stdin: strings.NewReader(`3
		500000000 600000000
		600000000 700000000
		700000000 800000000`)}, wantOut: `1`},
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
