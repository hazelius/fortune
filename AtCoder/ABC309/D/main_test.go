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
		{name: "1", args: args{stdin: strings.NewReader(`3 4 6
		1 2
		2 3
		4 5
		4 6
		1 3
		6 7`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`7 5 20
		10 11
		4 5
		10 12
		1 2
		1 5
		5 6
		2 4
		3 5
		9 10
		2 5
		1 4
		11 12
		9 12
		8 9
		5 7
		3 7
		3 6
		3 4
		8 12
		9 11`)}, wantOut: `4`},
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
