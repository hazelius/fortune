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
		{name: "1", args: args{stdin: strings.NewReader(`5 2
		1 4
		3 2`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`10 3
		1 4 8
		4 2 4`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`3 2
		2 3
		1 2`)}, wantOut: `-1`},
		{name: "4", args: args{stdin: strings.NewReader(`3 2
1 3
2 1`)}, wantOut: `1`},
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
