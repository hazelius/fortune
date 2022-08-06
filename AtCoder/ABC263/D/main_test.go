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
		{name: "1", args: args{stdin: strings.NewReader(`5 4 3
		5 5 0 6 3`)}, wantOut: `14`},
		{name: "2", args: args{stdin: strings.NewReader(`4 10 10
		1 2 3 4`)}, wantOut: `10`},
		{name: "3", args: args{stdin: strings.NewReader(`10 -5 -3
		9 -6 10 -1 2 10 -1 7 -15 5`)}, wantOut: `-58`},
		{name: "4", args: args{stdin: strings.NewReader(`3 2 3
		5 6 7`)}, wantOut: `6`},
		{name: "5", args: args{stdin: strings.NewReader(`5 3 3
		8 2 0 0 5`)}, wantOut: `8`},
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
