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
		{name: "1", args: args{stdin: strings.NewReader(`4 8
		1 3 2 4`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`3 20
		5 3 2`)}, wantOut: `infinite`},
		{name: "3", args: args{stdin: strings.NewReader(`10 23
		2 5 6 5 2 1 7 9 7 2`)}, wantOut: `2`},
		{name: "4", args: args{stdin: strings.NewReader(`5 5
10 2 3 4 5`)}, wantOut: `1`},
		{name: "5", args: args{stdin: strings.NewReader(`5 44
19 9 9 9 7`)}, wantOut: `10`},
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
