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
3 4 3 5 7 6 2`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`5
5 4 3 2 1`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 2 3 4 5 6 7 8 9 10`)}, wantOut: `10`},
		{name: "4", args: args{stdin: strings.NewReader(`10
1 2 3 2 3 4 5 6 7 8`)}, wantOut: `8`},
		{name: "5", args: args{stdin: strings.NewReader(`10
4 5 1 2 3 4 5 6 7 8`)}, wantOut: `8`},
		{name: "6", args: args{stdin: strings.NewReader(`5
1 2 3 3 4`)}, wantOut: `4`},
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
