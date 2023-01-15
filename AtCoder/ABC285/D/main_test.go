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
		{name: "1", args: args{stdin: strings.NewReader(`2
		b m
		m d`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		a b
		b c
		c a`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`5
		aaa bbb
		yyy zzz
		ccc ddd
		xxx yyy
		bbb ccc`)}, wantOut: `Yes`},
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
