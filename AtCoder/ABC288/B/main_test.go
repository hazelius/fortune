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
		{name: "1", args: args{stdin: strings.NewReader(`5 3
		abc
		aaaaa
		xyz
		a
		def`)}, wantOut: `aaaaa
abc
xyz
`},
		{name: "2", args: args{stdin: strings.NewReader(`4 4
		z
		zyx
		zzz
		rbg`)}, wantOut: `rbg
z
zyx
zzz
`},
		{name: "3", args: args{stdin: strings.NewReader(`3 1
		abc
		arc
		agc`)}, wantOut: `abc
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
