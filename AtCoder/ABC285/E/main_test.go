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
		{name: "1", args: args{stdin: strings.NewReader(`7
		10 10 1 1 1 1 1`)}, wantOut: `50`},
		{name: "2", args: args{stdin: strings.NewReader(`10
		200000000 500000000 1000000000 800000000 100000000 80000000 600000 900000000 1 20`)}, wantOut: `5100000000`},
		{name: "3", args: args{stdin: strings.NewReader(`20
		38 7719 21238 2437 8855 11797 8365 32285 10450 30612 5853 28100 1142 281 20537 15921 8945 26285 2997 14680`)}, wantOut: `236980`},
		{name: "4", args: args{stdin: strings.NewReader(`5
		1 2 3 4 5`)}, wantOut: `6`},
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
