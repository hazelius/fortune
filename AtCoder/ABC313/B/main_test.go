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
		1 2
		2 3`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
		1 3
		2 3`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`6 6
		1 6
		6 5
		6 2
		2 3
		4 3
		4 2`)}, wantOut: `-1`},
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
