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
		{name: "1", args: args{stdin: strings.NewReader(`3
4 2 3`)}, wantOut: `26`},
		{name: "2", args: args{stdin: strings.NewReader(`2
9 45`)}, wantOut: `405`},
		{name: "3", args: args{stdin: strings.NewReader(`10
7781 8803 8630 9065 8831 9182 8593 7660 7548 8617`)}, wantOut: `3227530139`},
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
