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
		142857
		004159
		071028
		159
		287
		857`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`5 4
		235983
		109467
		823476
		592801
		000333
		333
		108
		467
		983`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`4 4
		000000
		123456
		987111
		000000
		000
		111
		999
		111`)}, wantOut: `3`},
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
