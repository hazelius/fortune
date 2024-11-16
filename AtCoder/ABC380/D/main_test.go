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
		{name: "1", args: args{stdin: strings.NewReader(`aB
16
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16`)}, wantOut: `a B A b A b a B A b a B a B A b`},
		{name: "2", args: args{stdin: strings.NewReader(`qWeRtYuIoP
8
1 1 2 3 5 8 13 21`)}, wantOut: `q q W e t I E Q`},
		{name: "3", args: args{stdin: strings.NewReader(`AnUoHrjhgfLMcDIpzxXmEWPwBZvbKqQuiJTtFSlkNGVReOYCdsay
5
1000000000000000000 123456789 1 987654321 999999999999999999`)}, wantOut: `K a A Z L`},
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
