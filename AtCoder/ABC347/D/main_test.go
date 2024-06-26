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
		{name: "1", args: args{stdin: strings.NewReader(`3 4 7`)}, wantOut: `26 29`},
		{name: "2", args: args{stdin: strings.NewReader(`34 56 998244353`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`39 47 530423800524412070`)}, wantOut: `174478061544565593 375962048499941375`},
		{name: "4", args: args{stdin: strings.NewReader(`0 0 0`)}, wantOut: `0 0`},
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
