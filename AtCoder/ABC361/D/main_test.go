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
		BWBWBW
		WWWBBB`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`6
		BBBBBB
		W`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`14
		BBBWBWWWBBWWBW
		WBWWBBWWWBWBBB`)}, wantOut: `7`},
		{name: "4", args: args{stdin: strings.NewReader(`3
		BWB
		BWB`)}, wantOut: `0`},
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
