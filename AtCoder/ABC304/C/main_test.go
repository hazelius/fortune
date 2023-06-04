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
		{name: "1", args: args{stdin: strings.NewReader(`4 5
		2 -1
		3 1
		8 8
		0 5`)}, wantOut: `Yes
Yes
No
Yes
`},
		{name: "2", args: args{stdin: strings.NewReader(`3 1
		0 0
		-1000 -1000
		1000 1000`)}, wantOut: `Yes
No
No
`},
		{name: "3", args: args{stdin: strings.NewReader(`9 4
		3 2
		6 -1
		1 6
		6 5
		-2 -3
		5 3
		2 -3
		2 1
		2 6`)}, wantOut: `Yes
No
No
Yes
Yes
Yes
Yes
Yes
No
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
