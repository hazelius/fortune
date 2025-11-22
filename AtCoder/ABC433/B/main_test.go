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
4 3 2 5`)}, wantOut: `-1
1
2
-1
`},
		{name: "2", args: args{stdin: strings.NewReader(`3
7 7 7`)}, wantOut: `-1
-1
-1
`},
		{name: "3", args: args{stdin: strings.NewReader(`6
31 9 17 10 2 9`)}, wantOut: `-1
1
1
3
4
4
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
