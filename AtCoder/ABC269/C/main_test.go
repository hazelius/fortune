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
		{name: "1", args: args{stdin: strings.NewReader(`11`)}, wantOut: `0
1
2
3
8
9
10
11
`},
		{name: "2", args: args{stdin: strings.NewReader(`0`)}, wantOut: `0
`},
		{name: "3", args: args{stdin: strings.NewReader(`576461302059761664`)}, wantOut: `0
524288
549755813888
549756338176
576460752303423488
576460752303947776
576461302059237376
576461302059761664
`},
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
