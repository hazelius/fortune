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
		{name: "1", args: args{stdin: strings.NewReader(`5 2
7 13 3 2 5`)}, wantOut: `10`},
		{name: "2", args: args{stdin: strings.NewReader(`2 1
2 5`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`6 18
999999 999999999999 999999999999999 1000000000000 222222222 2`)}, wantOut: `1000`},
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
