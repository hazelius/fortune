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
		{name: "1", args: args{stdin: strings.NewReader(`6 5
1 2
1 4
2 3
5 3
3 1`)}, wantOut: `0 1 0 4 4 10 `},
		{name: "2", args: args{stdin: strings.NewReader(`7 3
1 2
3 4
5 6`)}, wantOut: `10 10 10 10 10 10 20 `},
		{name: "3", args: args{stdin: strings.NewReader(`6 9
3 6
2 5
2 3
4 3
1 5
6 2
3 1
5 3
2 4`)}, wantOut: `1 0 0 1 0 1 `},
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
