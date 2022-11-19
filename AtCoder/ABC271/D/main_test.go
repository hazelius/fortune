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
		// 		{name: "1", args: args{stdin: strings.NewReader(`3 11
		// 		1 4
		// 		2 3
		// 		5 7`)}, wantOut: `Yes
		// THH`},
		{name: "2", args: args{stdin: strings.NewReader(`5 25
		2 8
		9 3
		4 11
		5 1
		12 6`)}, wantOut: `No`},
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
