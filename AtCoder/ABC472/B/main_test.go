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
5 2 3 8`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`7
31 41 59 26 53 58 97`)}, wantOut: `51`},
		{name: "3", args: args{stdin: strings.NewReader(`10
67011 35764 33042 24098 63738 98760 17199 68579 21812 45408`)}, wantOut: `28105`},
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
