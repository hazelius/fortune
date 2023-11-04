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
		{name: "2", args: args{stdin: strings.NewReader(`100`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`10000000000`)}, wantOut: `10`},
		{name: "4", args: args{stdin: strings.NewReader(`100000000000001`)}, wantOut: `-1`},
		{name: "5", args: args{stdin: strings.NewReader(`1`)}, wantOut: `1`},
		{name: "6", args: args{stdin: strings.NewReader(`437893890380859375`)}, wantOut: `15`},
		{name: "7", args: args{stdin: strings.NewReader(`437893890380859376`)}, wantOut: `-1`},
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
