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
		{name: "1", args: args{stdin: strings.NewReader(`3
		2
		1
		0`)}, wantOut: `0
1
2
3
`},
		{name: "2", args: args{stdin: strings.NewReader(`0`)}, wantOut: `0
`},
		{name: "3", args: args{stdin: strings.NewReader(`123
		456
		789
		987
		654
		321
		0`)}, wantOut: `0
321
654
987
789
456
123
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
