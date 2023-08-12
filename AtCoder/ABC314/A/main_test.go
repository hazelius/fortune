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
		{name: "1", args: args{stdin: strings.NewReader(`2`)}, wantOut: `3.14`},
		{name: "2", args: args{stdin: strings.NewReader(`32`)}, wantOut: `3.14159265358979323846264338327950`},
		{name: "3", args: args{stdin: strings.NewReader(`100`)}, wantOut: `3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679`},
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
