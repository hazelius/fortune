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
		alice 31
		bob 41
		carol 5
		dave 92
		ellen 65`)}, wantOut: `carol
dave
ellen
alice
bob
`},
		{name: "2", args: args{stdin: strings.NewReader(`2
		takahashi 1000000000
		aoki 999999999`)}, wantOut: `aoki
takahashi
`},
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
