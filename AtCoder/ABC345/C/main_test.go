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
		{name: "1", args: args{stdin: strings.NewReader(`abc`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`aaaaa`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`aab`)}, wantOut: `3`},
		{name: "4", args: args{stdin: strings.NewReader(`aaba`)}, wantOut: `4`},
		{name: "4", args: args{stdin: strings.NewReader(`aabc`)}, wantOut: `6`},
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
