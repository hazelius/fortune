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
1 3
3 1
4 4
7 1`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3
1 8
10 11
21 5`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`10
2 1
22 10
26 17
29 2
45 20
47 32
72 12
75 1
81 31
97 7`)}, wantOut: `57`},
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
