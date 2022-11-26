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
		{name: "1", args: args{stdin: strings.NewReader(`5 4
				1 2 3 2`)}, wantOut: `1
3
2
4
`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
				2 2 2`)}, wantOut: `1
1
1
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
		1 1 1 9 4 4 2 1 3 3`)}, wantOut: `2
2
2
3
3
3
1
3
4
4
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
