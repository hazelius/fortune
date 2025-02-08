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
		// mapを使ってるので答えがランダムになる
		// 		{name: "1", args: args{stdin: strings.NewReader(`4 5
		// 1 1
		// 1 2
		// 2 1
		// 3 4
		// 4 4`)}, wantOut: `1
		// 1 1 4
		// `},
		{name: "2", args: args{stdin: strings.NewReader(`4 3
3 4
4 1
1 2`)}, wantOut: `0`},
		// 		{name: "3", args: args{stdin: strings.NewReader(`5 4
		// 3 3
		// 3 3
		// 3 3
		// 3 3`)}, wantOut: `4
		// 1 3 1
		// 2 3 2
		// 3 3 4
		// 4 3 5
		// `},
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
