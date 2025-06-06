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
		{name: "1", args: args{stdin: strings.NewReader(`5
1 4 4 4 2`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`6
2 4 4 2 2 4`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`8
1 4 2 5 7 7 7 2`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`10
1 2 3 4 5 6 7 8 9 10`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`13
1 1 1 1 1 1 1 1 1 1 1 1 1`)}, wantOut: `Yes`},
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
