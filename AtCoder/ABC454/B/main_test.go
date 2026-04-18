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
1 2 4`)}, wantOut: `Yes
No`},
		{name: "2", args: args{stdin: strings.NewReader(`4 2
1 2 1 2`)}, wantOut: `No
Yes`},
		{name: "3", args: args{stdin: strings.NewReader(`4 4
1 3 2 1`)}, wantOut: `No
No`},
		{name: "4", args: args{stdin: strings.NewReader(`5 5
1 3 4 2 5`)}, wantOut: `Yes
Yes`},
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
