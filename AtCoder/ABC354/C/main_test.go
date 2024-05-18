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
		2 4
		1 1
		3 2`)}, wantOut: `2
2 3`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		1 1
		10 2
		100 3
		1000 4
		10000 5`)}, wantOut: `5
1 2 3 4 5`},
		{name: "3", args: args{stdin: strings.NewReader(`6
		32 101
		65 78
		2 29
		46 55
		103 130
		52 40`)}, wantOut: `4
2 3 5 6`},
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
