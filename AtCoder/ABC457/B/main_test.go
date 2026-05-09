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
		{name: "1", args: args{stdin: strings.NewReader(`3
3 10 20 30
1 7
4 5 6 7 8
3 4`)}, wantOut: `8`},
		{name: "2", args: args{stdin: strings.NewReader(`4
2 9 1
3 8 2 6
1 5
2 4 3
2 2`)}, wantOut: `2`},
		{name: "3", args: args{stdin: strings.NewReader(`1
5 100 200 300 400 500
1 5`)}, wantOut: `500`},
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
