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
		{name: "1", args: args{stdin: strings.NewReader(`7 6
		5
		6 1
		3 1
		4 2
		1 5
		6 2
		2
		2 5
		2
		3 4`)}, wantOut: `0 2`},
		{name: "2", args: args{stdin: strings.NewReader(`4 4
		4
		1 1
		3 1
		3 3
		1 3
		1
		2
		1
		2`)}, wantOut: `1 1`},
		{name: "3", args: args{stdin: strings.NewReader(`20 20
		5
		3 3
		3 4
		4 3
		3 5
		3 6
		3
		2 6 15
		3
		2 8 19`)}, wantOut: `0 5`},
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
