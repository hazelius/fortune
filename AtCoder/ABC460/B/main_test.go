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
		{name: "1", args: args{stdin: strings.NewReader(`7
0 0 2 2 3 2
0 0 2 2 3 1
1 2 5 3 2 1
5 4 2 8 8 3
2 1 5 5 1 2
0 0 1 0 0 1
0 0 500000000 1 1000000000 500000000`)}, wantOut: `Yes
No
No
Yes
Yes
Yes
No
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
