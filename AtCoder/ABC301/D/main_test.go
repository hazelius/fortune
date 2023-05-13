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
		{name: "1", args: args{stdin: strings.NewReader(`?0?
		2`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`101
		4`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`?0?
		1000000000000000000`)}, wantOut: `5`},
		{name: "4", args: args{stdin: strings.NewReader(`?11?
		11`)}, wantOut: `7`},
		{name: "5", args: args{stdin: strings.NewReader(`1?111?11
		250`)}, wantOut: `191`},
		{name: "6", args: args{stdin: strings.NewReader(`?
		1`)}, wantOut: `1`},
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
