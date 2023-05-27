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
		{name: "1", args: args{stdin: strings.NewReader(`1 3 3
		AAaA`)}, wantOut: `9`},
		{name: "2", args: args{stdin: strings.NewReader(`1 1 100
		aAaAaA`)}, wantOut: `6`},
		{name: "3", args: args{stdin: strings.NewReader(`1 2 4
		aaAaAaaAAAAaAaaAaAAaaaAAAAA`)}, wantOut: `40`},
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
