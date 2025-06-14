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
		{name: "1", args: args{stdin: strings.NewReader(`4 10 14
5 8
5 6
7 9
99 99`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3000 3000
3 3
3 3
3 3`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`10 8 8
2 2
2 3
2 2
1 2
2 3
1 2
3 3
3 2
3 1
3 2`)}, wantOut: `9`},
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
