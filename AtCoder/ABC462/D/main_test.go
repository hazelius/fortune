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
		{name: "1", args: args{stdin: strings.NewReader(`3 2
9 17
10 12
13 20`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`3 5
9 17
10 12
13 20`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`4 1
1 1000000
1 1000000
1 1000000
1 1000000`)}, wantOut: `5999994`},
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
