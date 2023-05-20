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
		{name: "1", args: args{stdin: strings.NewReader(`7 3`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`123456789123456789 987654321`)}, wantOut: `124999999`},
		{name: "3", args: args{stdin: strings.NewReader(`999999999999999998 2`)}, wantOut: `499999999999999999`},
		{name: "4", args: args{stdin: strings.NewReader(`16 3`)}, wantOut: `6`},
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
