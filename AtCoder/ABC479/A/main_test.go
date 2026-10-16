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
3 1 4 1 5`)}, wantOut: `12`},
		{name: "2", args: args{stdin: strings.NewReader(`6
2 7 1 8 2 8`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`9
100 40 30 50 20 90 100 80 60`)}, wantOut: `310`},
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
