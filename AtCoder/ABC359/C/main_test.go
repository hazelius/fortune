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
		{name: "1", args: args{stdin: strings.NewReader(`5 0
		2 5`)}, wantOut: `5`},
		{name: "2", args: args{stdin: strings.NewReader(`3 1
		4 1`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`2552608206527595 5411232866732612
		771856005518028 7206210729152763`)}, wantOut: `1794977862420151`},
		{name: "4", args: args{stdin: strings.NewReader(`0 0 -1 0`)}, wantOut: `1`},
		{name: "5", args: args{stdin: strings.NewReader(`0 0 -2 0`)}, wantOut: `1`},
		{name: "6", args: args{stdin: strings.NewReader(`0 0 1 0`)}, wantOut: `0`},
		{name: "6", args: args{stdin: strings.NewReader(`0 0 2 0`)}, wantOut: `1`},
		{name: "6", args: args{stdin: strings.NewReader(`0 0 2 1`)}, wantOut: `1`},
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
