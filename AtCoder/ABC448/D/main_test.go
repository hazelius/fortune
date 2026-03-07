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
		{name: "1", args: args{stdin: strings.NewReader(`5
1 3 2 1 2
1 2
1 3
3 4
3 5`)}, wantOut: `No
No
No
Yes
Yes
`},
		{name: "2", args: args{stdin: strings.NewReader(`2
1000000000 1000000000
2 1`)}, wantOut: `No
Yes
`},
		{name: "3", args: args{stdin: strings.NewReader(`10
10 7 3 9 1 3 8 5 7 10
3 6
8 6
6 1
9 7
7 10
5 4
4 2
10 2
1 9`)}, wantOut: `No
Yes
Yes
Yes
Yes
No
No
No
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
