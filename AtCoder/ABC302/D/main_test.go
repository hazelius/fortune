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
		{name: "1", args: args{stdin: strings.NewReader(`2 3 2
		3 10
		2 5 15
		`)}, wantOut: `8`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3 0
		1 3 3
		6 2 7`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`1 1 1000000000000000000
		1000000000000000000
		1000000000000000000`)}, wantOut: `2000000000000000000`},
		{name: "4", args: args{stdin: strings.NewReader(`8 6 1
		2 5 6 5 2 1 7 9
		7 2 5 5 2 4`)}, wantOut: `14`},
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
