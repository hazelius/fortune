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
		{name: "1", args: args{stdin: strings.NewReader(`5
3 1 4 2 5`)}, wantOut: `6`},
		{name: "2", args: args{stdin: strings.NewReader(`2
2 1`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`20
15 5 13 17 9 11 20 4 14 16 6 3 8 19 12 7 10 18 2 1`)}, wantOut: `77`},
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
