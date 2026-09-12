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
		{name: "1", args: args{stdin: strings.NewReader(`motor`)}, wantOut: `10607`},
		{name: "2", args: args{stdin: strings.NewReader(`byebye`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`coconut`)}, wantOut: `1010237`},
		{name: "4", args: args{stdin: strings.NewReader(`coeonuu`)}, wantOut: `1`},
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
