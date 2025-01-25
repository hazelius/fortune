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
3 6 12 24 48`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`3
1 2 3`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`2
10 8`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`3
100 150 225`)}, wantOut: `Yes`},
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
