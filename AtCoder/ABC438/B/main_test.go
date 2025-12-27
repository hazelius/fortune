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
		{name: "1", args: args{stdin: strings.NewReader(`4 2
2025
91`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
438
38`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`5 5
00000
11111`)}, wantOut: `45`},
		{name: "4", args: args{stdin: strings.NewReader(`8 3
20251227
438`)}, wantOut: `13`},
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
