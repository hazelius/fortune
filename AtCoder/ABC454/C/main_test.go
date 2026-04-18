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
		{name: "1", args: args{stdin: strings.NewReader(`5 5
1 2
2 3
3 4
2 4
5 2`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
2 1
3 2`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`7 8
2 6
2 5
3 6
1 6
1 2
5 6
2 3
3 7`)}, wantOut: `6`},
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
