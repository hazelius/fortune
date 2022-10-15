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
		{name: "1", args: args{stdin: strings.NewReader(`6
		2 7 1 8 2 8`)}, wantOut: `2
1
2
1
0
0
`},
		{name: "2", args: args{stdin: strings.NewReader(`1
		1`)}, wantOut: `1
`},
		{name: "3", args: args{stdin: strings.NewReader(`10
979861204 57882493 979861204 447672230 644706927 710511029 763027379 710511029 447672230 136397527`)}, wantOut: `2
1
2
1
2
1
1
0
0
0
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
