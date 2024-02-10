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
		100 200 3
		50 10 1
		100 200 5
		150 1 2`)}, wantOut: `350`},
		{name: "2", args: args{stdin: strings.NewReader(`10
		1000 10 9
		1000 10 10
		1000 10 2
		1000 10 3
		1000 10 4
		1000 10 5
		1000 10 6
		1000 10 7
		1000 10 8`)}, wantOut: `90`},
		{name: "3", args: args{stdin: strings.NewReader(`6
		1000000000 1000000000 1
		1000000000 1000000000 1
		1000000000 1000000000 1
		1000000000 1000000000 1
		1000000000 1000000000 1`)}, wantOut: `5000000000`},
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
