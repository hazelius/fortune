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
		{name: "1", args: args{stdin: strings.NewReader(`3
5
1 8 2 4 16
5
-16 24 54 81 -36
7
90000 8100 -27000 729 -300000 -2430 1000000`)}, wantOut: `Yes
No
Yes
`},
		{name: "2", args: args{stdin: strings.NewReader(`3
5
8 8 8 8 8
5
8 8 8 8 -8
7
8 8 8 8 -8 -8 -8`)}, wantOut: `Yes
No
Yes
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
