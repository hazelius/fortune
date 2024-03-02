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
		{name: "1", args: args{stdin: strings.NewReader(`3 4
		1 10
		3 20
		2 10
		2 10`)}, wantOut: `2
3
2
2
`},
		{name: "2", args: args{stdin: strings.NewReader(`1 3
		1 3
		1 4
		1 3`)}, wantOut: `1
1
1
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
		7 2620
		9 2620
		8 3375
		1 3375
		6 1395
		5 1395
		6 2923
		10 3375
		9 5929
		5 1225`)}, wantOut: `2
2
3
3
4
4
5
5
6
5
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
