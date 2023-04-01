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
		{name: "1", args: args{stdin: strings.NewReader(`5 7`)}, wantOut: `8`},
		{name: "2", args: args{stdin: strings.NewReader(`2 5`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`100000 10000000000`)}, wantOut: `10000000000`},
		{name: "4", args: args{stdin: strings.NewReader(`8 19`)}, wantOut: `20`},
		{name: "5", args: args{stdin: strings.NewReader(`8 15`)}, wantOut: `15`},
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
