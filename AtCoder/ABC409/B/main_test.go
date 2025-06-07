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
1 2 1`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`7
1 6 2 10 2 3 2`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 2 7 8 9 10 11 12 13 999999999`)}, wantOut: `7`},
		{name: "4", args: args{stdin: strings.NewReader(`3
4 8 8`)}, wantOut: `3`},
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
