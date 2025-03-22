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
		{name: "1", args: args{stdin: strings.NewReader(`6 -2 1
NNEEWS`)}, wantOut: `001010`},
		{name: "2", args: args{stdin: strings.NewReader(`10 1 2
NEESESWEES`)}, wantOut: `0001101011`},
		{name: "3", args: args{stdin: strings.NewReader(`20 -1 -2
WWNNWSWEWNSWWENSNWWN`)}, wantOut: `00100111111000101111`},
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
