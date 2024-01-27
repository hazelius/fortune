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
		{name: "1", args: args{stdin: strings.NewReader(`3 3
		1 3 2`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`4 5
		2 4 2 4 2`)}, wantOut: `8`},
		{name: "3", args: args{stdin: strings.NewReader(`163054 10
		62874 19143 77750 111403 29327 56303 6659 18896 64175 26369`)}, wantOut: `390009`},
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
