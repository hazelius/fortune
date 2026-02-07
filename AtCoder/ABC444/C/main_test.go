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
		{name: "1", args: args{stdin: strings.NewReader(`4
10 5 5 10`)}, wantOut: `10 15`},
		{name: "2", args: args{stdin: strings.NewReader(`3
4 4 4`)}, wantOut: `4`},
		{name: "3", args: args{stdin: strings.NewReader(`6
10 187 344 100 434 257`)}, wantOut: `444`},
		{name: "4", args: args{stdin: strings.NewReader(`6
10 10 9 2 1 1`)}, wantOut: `11`},
		{name: "5", args: args{stdin: strings.NewReader(`2
10 5`)}, wantOut: `15`},
		{name: "6", args: args{stdin: strings.NewReader(`2
10 10`)}, wantOut: `10 20`},
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
