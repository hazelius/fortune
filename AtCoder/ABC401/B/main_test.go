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
		{name: "1", args: args{stdin: strings.NewReader(`6
login
private
public
logout
private
public`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`4
private
private
private
logout`)}, wantOut: `3`},
		{name: "3", args: args{stdin: strings.NewReader(`20
private
login
private
logout
public
logout
logout
logout
logout
private
login
login
private
login
private
login
public
private
logout
private`)}, wantOut: `3`},
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
