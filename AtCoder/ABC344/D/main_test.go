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
		{name: "1", args: args{stdin: strings.NewReader(`abcde
		3
		3 ab abc abcd
		4 f c cd bcde
		2 e de`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`abcde
		3
		2 ab abc
		3 f c bcde
		1 e`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`aaabbbbcccc
		6
		2 aa aaa
		2 dd ddd
		2 ab aabb
		4 bbaa bbbc bbb bbcc
		2 cc bcc
		3 ccc cccc ccccc`)}, wantOut: `4`},
		{name: "4", args: args{stdin: strings.NewReader(`abcde
		3
		3 a bc de
		4 f s cd baacde
		2 ec ddde`)}, wantOut: `-1`},
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
