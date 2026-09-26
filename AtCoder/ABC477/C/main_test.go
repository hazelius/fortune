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
abcdabc
bc
2 6
3 5
3 7`)}, wantOut: `Yes
No
Yes
`},
		{name: "2", args: args{stdin: strings.NewReader(`10
bbbaaabbbaaaaaabbbab
bbb
5 16
5 6
2 16
6 7
17 19
3 7
7 16
7 20
17 19
2 6`)}, wantOut: `Yes
No
Yes
No
No
No
Yes
Yes
No
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
