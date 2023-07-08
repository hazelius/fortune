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
		0101
		1101
		1111
		0000`)}, wantOut: `1010
1101
0111
0001
`},
		{name: "2", args: args{stdin: strings.NewReader(`2
		11
		11`)}, wantOut: `11
11
`},
		{name: "3", args: args{stdin: strings.NewReader(`5
		01010
		01001
		10110
		00110
		01010`)}, wantOut: `00101
11000
00111
00110
10100
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
