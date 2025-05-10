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
3 1 4 1 5`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`4 3
1 6 1 8`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 3
1 6 2 10 2 3 2 10 6 4`)}, wantOut: `2`},
		{name: "4", args: args{stdin: strings.NewReader(`10 0
7 10 7 1 1 10 1 10 7 4`)}, wantOut: `6`},
		{name: "5", args: args{stdin: strings.NewReader(`12 2
1 1 1 1 3 3 5 5 15 15 17 17`)}, wantOut: `4`},
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
