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
		{name: "1", args: args{stdin: strings.NewReader(`5 7
2 3 3 5 1 5 2`)}, wantOut: `1
1
1
2
2
1
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`1 2
1 1`)}, wantOut: `1
0
`},
		{name: "3", args: args{stdin: strings.NewReader(`3 3
1 3 2`)}, wantOut: `1
2
1
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
