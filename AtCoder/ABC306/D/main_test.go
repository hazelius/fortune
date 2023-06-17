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
		{name: "1", args: args{stdin: strings.NewReader(`5
		1 100
		1 300
		0 -200
		1 500
		1 300`)}, wantOut: `600`},
		{name: "2", args: args{stdin: strings.NewReader(`4
		0 -1
		1 -2
		0 -3
		1 -4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`15
		1 900000000
		0 600000000
		1 -300000000
		0 -700000000
		1 200000000
		1 300000000
		0 -600000000
		1 -900000000
		1 600000000
		1 -100000000
		1 -400000000
		0 900000000
		0 200000000
		1 -500000000
		1 900000000`)}, wantOut: `4100000000`},
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
