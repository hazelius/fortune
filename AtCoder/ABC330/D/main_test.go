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
		{name: "1", args: args{stdin: strings.NewReader(`3
		ooo
		oxx
		xxo`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`4
		oxxx
		xoxx
		xxox
		xxxo`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`15
		xooxxooooxxxoox
		oxxoxoxxxoxoxxo
		oxxoxoxxxoxoxxx
		ooooxooooxxoxxx
		oxxoxoxxxoxoxxx
		oxxoxoxxxoxoxxo
		oxxoxooooxxxoox
		xxxxxxxxxxxxxxx
		xooxxxooxxxooox
		oxxoxoxxoxoxxxo
		xxxoxxxxoxoxxoo
		xooxxxooxxoxoxo
		xxxoxxxxoxooxxo
		oxxoxoxxoxoxxxo
		xooxxxooxxxooox`)}, wantOut: `2960`},
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
