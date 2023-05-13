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
		{name: "1", args: args{stdin: strings.NewReader(`5
		TTAAT`)}, wantOut: `T`},
		{name: "2", args: args{stdin: strings.NewReader(`6
		ATTATA`)}, wantOut: `T`},
		{name: "3", args: args{stdin: strings.NewReader(`1
		A`)}, wantOut: `A`},
		{name: "3", args: args{stdin: strings.NewReader(`6
		ATTAAT`)}, wantOut: `A`},
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
