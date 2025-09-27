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
-1 -1 2 -1`)}, wantOut: `Yes
1 3 2 4 `},
		{name: "2", args: args{stdin: strings.NewReader(`5
-1 -1 1 -1 1`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`7
3 -1 4 -1 5 -1 2`)}, wantOut: `Yes
3 1 4 6 5 7 2 `},
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
