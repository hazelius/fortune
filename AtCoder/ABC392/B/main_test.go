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
		{name: "1", args: args{stdin: strings.NewReader(`10 3
3 9 2`)}, wantOut: `7
1 4 5 6 7 8 10`},
		{name: "2", args: args{stdin: strings.NewReader(`6 6
1 3 5 2 4 6`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`9 1
9`)}, wantOut: `8
1 2 3 4 5 6 7 8`},
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
