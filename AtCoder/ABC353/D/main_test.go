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
		{name: "1", args: args{stdin: strings.NewReader(`3
		3 14 15`)}, wantOut: `2044`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		1001 5 1000000 1000000000 100000`)}, wantOut: `625549048`},
		// {name: "3", args: args{stdin: strings.NewReader(`3
		// 1000000000 1000000000 1000000000`)}, wantOut: `22`},
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
