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
		{name: "1", args: args{stdin: strings.NewReader(`3 5
1 2
2 b
1 1
1 2
2 c`)}, wantOut: `bcc`},
		{name: "2", args: args{stdin: strings.NewReader(`10 15
1 8
2 m
1 3
1 10
2 q
1 6
1 10
2 d
1 1
2 z
1 9
2 f
1 4
1 7
2 k`)}, wantOut: `dkmfkqfazk`},
		{name: "3", args: args{stdin: strings.NewReader(`8 5
1 2
1 1
1 1
1 1
2 c`)}, wantOut: `aacccccc`},
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
