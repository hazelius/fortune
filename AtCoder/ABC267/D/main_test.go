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
		// {name: "1", args: args{stdin: strings.NewReader(`4 2
		// 5 4 -1 8`)}, wantOut: `21`},
		// {name: "2", args: args{stdin: strings.NewReader(`10 4
		// -3 1 -4 1 -5 9 -2 6 -5 3`)}, wantOut: `54`},
		{name: "3", args: args{stdin: strings.NewReader(`5 3
		-10 4 -1 -8 1`)}, wantOut: `5`},
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
