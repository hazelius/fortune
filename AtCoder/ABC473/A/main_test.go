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
		{name: "1", args: args{stdin: strings.NewReader(`8
1 3 7 8 4 2 5 6`)}, wantOut: `17`},
		{name: "2", args: args{stdin: strings.NewReader(`2
1 100`)}, wantOut: `100`},
		{name: "3", args: args{stdin: strings.NewReader(`10
31 41 59 26 53 58 97 93 23 84`)}, wantOut: `355`},
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
