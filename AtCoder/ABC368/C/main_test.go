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
6 2 2`)}, wantOut: `8`},
		{name: "2", args: args{stdin: strings.NewReader(`9
1 12 123 1234 12345 123456 1234567 12345678 123456789`)}, wantOut: `82304529`},
		{name: "3", args: args{stdin: strings.NewReader(`5
1000000000 1000000000 1000000000 1000000000 1000000000`)}, wantOut: `3000000000`},
		{name: "4", args: args{stdin: strings.NewReader(`3
1 1 2`)}, wantOut: `3`},
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
