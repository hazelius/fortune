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
3
2 3
1 2
8 9`)}, wantOut: `3
2
3
`},
		{name: "2", args: args{stdin: strings.NewReader(`1
4
2 3
4 5
6 7
8 9`)}, wantOut: `2
3
4
5
`},
		{name: "3", args: args{stdin: strings.NewReader(`278117031
7
167642909 517897721
148434323 567739597
319926999 481642530
659199879 252516557
49913403 798318034
89701408 892537201
199166668 742285869`)}, wantOut: `278117031
278117031
319926999
319926999
319926999
319926999
319926999
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
