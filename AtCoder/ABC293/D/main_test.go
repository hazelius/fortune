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
		{name: "1", args: args{stdin: strings.NewReader(`5 3
		3 R 5 B
		5 R 3 B
		4 R 2 B`)}, wantOut: `1 2`},
		{name: "2", args: args{stdin: strings.NewReader(`7 0`)}, wantOut: `0 7`},
		{name: "3", args: args{stdin: strings.NewReader(`7 6
		5 R 3 R
		7 R 4 R
		4 B 1 R
		2 R 3 B
		2 B 5 B
		1 B 7 B`)}, wantOut: `2 1`},
		{name: "4", args: args{stdin: strings.NewReader(`4 3
		1 R 1 B
		2 B 3 R
		3 B 4 R`)}, wantOut: `1 1`},
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
