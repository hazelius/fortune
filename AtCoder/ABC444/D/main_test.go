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
		{name: "1", args: args{stdin: strings.NewReader(`4
3 3 3 3`)}, wantOut: `444`},
		{name: "2", args: args{stdin: strings.NewReader(`3
30 10 20`)}, wantOut: `111111111122222222223333333333`},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 2 3 4 5 6 7 8 9 10`)}, wantOut: `1234567900`},
		{name: "4", args: args{stdin: strings.NewReader(`10
12 12 12 12 12 12 12 12 12 10`)}, wantOut: `1001111111110`},
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
