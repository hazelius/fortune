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
		{name: "1", args: args{stdin: strings.NewReader(`4 500
		300 900 1300 1700`)}, wantOut: `1300`},
		{name: "2", args: args{stdin: strings.NewReader(`5 99
		100 200 300 400 500`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`4 500
		100 600 1100 1600`)}, wantOut: `600`},
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
