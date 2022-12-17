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
		{name: "1", args: args{stdin: strings.NewReader(`5 5
		ooooo
		oooxx
		xxooo
		oxoxo
		xxxxx`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
		ox
		xo
		xx`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`2 4
		xxxx
		oxox`)}, wantOut: `0`},
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
