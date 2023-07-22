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
		6 7 2 1 3 4 5`)}, wantOut: `3
1 6 4 `},
		{name: "2", args: args{stdin: strings.NewReader(`2
		2 1`)}, wantOut: `2
1 2 `},
		{name: "3", args: args{stdin: strings.NewReader(`8
		3 7 4 7 3 3 8 2`)}, wantOut: `3
7 8 2 `},
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
