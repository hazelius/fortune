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
1 at
2 watcoder
2 atcoder
1 wa`)}, wantOut: `0
1
1
0
`},
		{name: "2", args: args{stdin: strings.NewReader(`10
1 w
1 avko
2 atcoder
1 bzginn
2 beginner
1 atco
2 contest
1 ntxcdg
1 atc
1 contest`)}, wantOut: `0
0
1
1
2
1
2
2
2
1
`},
		{name: "3", args: args{stdin: strings.NewReader(`4
2 wise
2 wise
1 wi
2 drop`)}, wantOut: `1
2
0
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
