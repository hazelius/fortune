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
		{name: "1", args: args{stdin: strings.NewReader(`7
2 4 7 5 5 6 7`)}, wantOut: `5 5 7 5 5 6 7 `},
		{name: "2", args: args{stdin: strings.NewReader(`5
1 2 3 4 5`)}, wantOut: `1 2 3 4 5 `},
		{name: "3", args: args{stdin: strings.NewReader(`15
11 3 10 7 15 10 10 11 11 13 11 12 14 14 15`)}, wantOut: `11 14 14 14 15 14 14 11 11 14 11 12 14 14 15 `},
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
