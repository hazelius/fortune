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
		{name: "1", args: args{stdin: strings.NewReader(`5 10
6 11 21 22 30`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`2 100
1 200`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`10 22
47 81 82 95 117 146 165 209 212 215`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`1 10
11`)}, wantOut: `No`},
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
