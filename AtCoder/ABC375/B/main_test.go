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
		{name: "1", args: args{stdin: strings.NewReader(`2
1 2
-1 0`)}, wantOut: `6.064495102246`},
		{name: "2", args: args{stdin: strings.NewReader(`7
-14142 13562
-17320 50807
-22360 67977
24494 89742
-26457 51311
28284 27124
31622 77660`)}, wantOut: `384694.575879320735`},
		{name: "3", args: args{stdin: strings.NewReader(`5
-100000 100000
100000 -100000
-100000 100000
100000 -100000
-100000 100000`)}, wantOut: `1414213.562373095192`},
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
