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
		{name: "1", args: args{stdin: strings.NewReader(`10
		1937458062
		8124690357
		2385760149`)}, wantOut: `6`},
		{name: "2", args: args{stdin: strings.NewReader(`20
		01234567890123456789
		01234567890123456789
		01234567890123456789`)}, wantOut: `20`},
		{name: "3", args: args{stdin: strings.NewReader(`5
		11111
		22222
		33333`)}, wantOut: `-1`},
		{name: "4", args: args{stdin: strings.NewReader(`5
		12345
		12345
		12141`)}, wantOut: `5`},
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
