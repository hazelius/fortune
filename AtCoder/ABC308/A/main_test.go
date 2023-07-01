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
		{name: "1", args: args{stdin: strings.NewReader(`125 175 250 300 400 525 600 650`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`100 250 300 400 325 575 625 675`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`0 23 24 145 301 413 631 632`)}, wantOut: `No`},
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
