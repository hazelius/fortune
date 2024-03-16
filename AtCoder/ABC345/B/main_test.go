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
		{name: "1", args: args{stdin: strings.NewReader(`27`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`-13`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`40`)}, wantOut: `4`},
		{name: "4", args: args{stdin: strings.NewReader(`-20`)}, wantOut: `-2`},
		{name: "5", args: args{stdin: strings.NewReader(`123456789123456789`)}, wantOut: `12345678912345679`},
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
