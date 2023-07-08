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
		{name: "1", args: args{stdin: strings.NewReader(`7 3
		1 2 1 3 3 3
		1 1
		1 2
		4 3`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`10 10
		1 1 3 1 2 3 3 5 7
		2 1
		5 1
		4 3
		6 3
		2 1
		7 3
		9 2
		1 2
		6 2
		8 1`)}, wantOut: `10`},
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
