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
3 L
6 R
9 L
1 R`)}, wantOut: `11`},
		{name: "2", args: args{stdin: strings.NewReader(`3
2 L
2 L
100 L`)}, wantOut: `98`},
		{name: "3", args: args{stdin: strings.NewReader(`8
22 L
75 L
26 R
45 R
72 R
81 R
47 L
29 R`)}, wantOut: `188`},
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
