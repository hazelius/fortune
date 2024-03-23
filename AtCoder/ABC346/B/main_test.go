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
		{name: "1", args: args{stdin: strings.NewReader(`3 2`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`3 0`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`92 66`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`0 0`)}, wantOut: `Yes`},
		{name: "5", args: args{stdin: strings.NewReader(`8 5`)}, wantOut: `Yes`},
		{name: "6", args: args{stdin: strings.NewReader(`9 5`)}, wantOut: `Yes`},
		{name: "7", args: args{stdin: strings.NewReader(`3 1`)}, wantOut: `Yes`},
		{name: "8", args: args{stdin: strings.NewReader(`3 0`)}, wantOut: `No`},
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
