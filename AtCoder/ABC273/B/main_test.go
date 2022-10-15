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
		{name: "1", args: args{stdin: strings.NewReader(`2048 2`)}, wantOut: `2100`},
		{name: "2", args: args{stdin: strings.NewReader(`1 15`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`999 3`)}, wantOut: `1000`},
		{name: "4", args: args{stdin: strings.NewReader(`314159265358979 12`)}, wantOut: `314000000000000`},
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
