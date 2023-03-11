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
		3 1 4 5 4`)}, wantOut: `2
2 4`},
		{name: "2", args: args{stdin: strings.NewReader(`20
		9 7 19 7 10 4 13 9 4 8 10 15 16 3 18 19 12 13 2 12`)}, wantOut: `10
1 2 5 6 8 11 14 17 18 20`},
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
