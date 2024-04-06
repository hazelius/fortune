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
		{name: "1", args: args{stdin: strings.NewReader(`4
		100 1
		20 5
		30 5
		40 1`)}, wantOut: `40`},
		{name: "2", args: args{stdin: strings.NewReader(`10
		68 3
		17 2
		99 2
		92 4
		82 4
		10 3
		100 2
		78 1
		3 1
		35 4`)}, wantOut: `35`},
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
