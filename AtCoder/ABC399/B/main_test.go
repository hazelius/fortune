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
3 12 9 9`)}, wantOut: `4
1
2
2
`},
		{name: "2", args: args{stdin: strings.NewReader(`3
3 9 6`)}, wantOut: `3
1
2
`},
		{name: "3", args: args{stdin: strings.NewReader(`4
100 100 100 100`)}, wantOut: `1
1
1
1
`},
		{name: "4", args: args{stdin: strings.NewReader(`8
87 87 87 88 41 38 41 38`)}, wantOut: `2
2
2
1
5
7
5
7
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
