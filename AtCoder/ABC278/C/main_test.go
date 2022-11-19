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
		1 1 2
		3 1 2
		1 2 1
		3 1 2
		1 2 3
		1 3 2
		3 1 3
		2 1 2
		3 1 2`)}, wantOut: `No
Yes
No
No
`},
		{name: "2", args: args{stdin: strings.NewReader(`2 8
		1 1 2
		1 2 1
		3 1 2
		1 1 2
		1 1 2
		1 1 2
		2 1 2
		3 1 2`)}, wantOut: `Yes
No
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 30
		3 1 6
		3 5 4
		1 6 1
		3 1 7
		3 8 4
		1 1 6
		2 4 3
		1 6 5
		1 5 6
		1 1 8
		1 8 1
		2 3 10
		1 7 6
		3 5 6
		1 6 7
		3 6 7
		1 9 5
		3 8 6
		3 3 8
		2 6 9
		1 7 1
		3 10 8
		2 9 2
		1 10 9
		2 6 10
		2 6 8
		3 1 6
		3 1 8
		2 8 5
		1 9 10`)}, wantOut: `No
No
No
No
Yes
Yes
No
No
No
Yes
Yes
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
