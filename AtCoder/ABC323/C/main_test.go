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
		1000 500 700 2000
		xxxo
		ooxx
		oxox`)}, wantOut: `0
1
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`5 5
		1000 1500 2000 2000 2500
		xxxxx
		oxxxx
		xxxxx
		oxxxx
		oxxxx`)}, wantOut: `1
1
1
1
0
`},
		{name: "3", args: args{stdin: strings.NewReader(`7 8
		500 500 500 500 500 500 500 500
		xxxxxxxx
		oxxxxxxx
		ooxxxxxx
		oooxxxxx
		ooooxxxx
		oooooxxx
		ooooooxx`)}, wantOut: `7
6
5
4
3
2
0
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
