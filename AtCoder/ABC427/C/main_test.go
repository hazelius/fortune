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
		{name: "1", args: args{stdin: strings.NewReader(`5 8
1 2
1 3
1 4
2 3
2 5
3 4
3 5
4 5`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`2 1
1 2`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10 20
5 9
1 4
3 8
1 6
4 10
5 7
5 6
3 7
3 6
5 10
1 3
3 4
6 7
1 2
4 7
1 5
1 9
9 10
4 5
8 9`)}, wantOut: `5`},
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
