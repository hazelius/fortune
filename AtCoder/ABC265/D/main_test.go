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
		{name: "1", args: args{stdin: strings.NewReader(`10 5 7 5
		1 3 2 2 2 3 1 4 3 2`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`9 100 101 100
		31 41 59 26 53 58 97 93 23`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`7 1 1 1
		1 1 1 1 1 1 1`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`9 10 20 30
		1 5 5 10 10 10 10 9 1`)}, wantOut: `Yes`},
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
