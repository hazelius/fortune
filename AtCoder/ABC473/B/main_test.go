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
		{name: "1", args: args{stdin: strings.NewReader(`8
2 7 1 8 2 8 1 8`)}, wantOut: `15`},
		{name: "2", args: args{stdin: strings.NewReader(`5
1 2 3 4 5`)}, wantOut: `15`},
		{name: "3", args: args{stdin: strings.NewReader(`15
58 97 74 16 97 74 97 16 51 52 58 52 74 32 43`)}, wantOut: `297`},
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
