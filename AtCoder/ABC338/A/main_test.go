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
		{name: "1", args: args{stdin: strings.NewReader(`Capitalized`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`AtCoder`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`yes`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`A`)}, wantOut: `Yes`},
		{name: "5", args: args{stdin: strings.NewReader(`a`)}, wantOut: `No`},
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
