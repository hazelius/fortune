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
		{name: "1", args: args{stdin: strings.NewReader(`9 7
		3 0 2 5 5 3 0 6 3`)}, wantOut: `11`},
		{name: "2", args: args{stdin: strings.NewReader(`1 10
		4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`20 20
		18 16 15 9 8 8 17 1 3 17 11 9 12 11 7 3 2 14 3 12`)}, wantOut: `99`},
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
