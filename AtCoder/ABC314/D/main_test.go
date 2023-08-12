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
		{name: "1", args: args{stdin: strings.NewReader(`7
		AtCoder
		5
		1 4 i
		3 0 a
		1 5 b
		2 0 a
		1 4 Y`)}, wantOut: `atcYber`},
		{name: "2", args: args{stdin: strings.NewReader(`35
		TheQuickBrownFoxJumpsOverTheLazyDog
		10
		2 0 a
		1 19 G
		1 13 m
		1 2 E
		1 21 F
		2 0 a
		1 27 b
		3 0 a
		3 0 a
		1 15 i`)}, wantOut: `TEEQUICKBROWMFiXJUGPFOVERTBELAZYDOG`},
		{name: "3", args: args{stdin: strings.NewReader(`7
		AtCoder
		1
		1 5 x`)}, wantOut: `AtCoxer`},
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
