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
2 3 5 10 12`)}, wantOut: `17`},
		{name: "2", args: args{stdin: strings.NewReader(`2
1 1`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`6
22 25 26 45 22 31`)}, wantOut: `89`},
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
