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
		{name: "1", args: args{stdin: strings.NewReader(`3 9
3 1 3 2
1 3
2 4 3
1 3 2`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`3 1
1 7
1 111
1 5
1 100 10000`)}, wantOut: `7`},
		{name: "3", args: args{stdin: strings.NewReader(`3 3163812
5 1 2 3 4 5
4 9 8 7 6
2 10 11
87043 908415 9814`)}, wantOut: `9`},
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
