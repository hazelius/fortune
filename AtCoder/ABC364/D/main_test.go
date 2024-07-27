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
		{name: "1", args: args{stdin: strings.NewReader(`4 3
-3 -1 5 6
-2 3
2 1
10 4`)}, wantOut: `7
3
13
`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
0 0
0 1
0 2`)}, wantOut: `0
0
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 5
-84 -60 -41 -100 8 -8 -52 -62 -61 -76
-52 5
14 4
-2 6
46 2
26 7`)}, wantOut: `11
66
59
54
88
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
