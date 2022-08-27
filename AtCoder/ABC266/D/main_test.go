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
		1 0 100
		3 3 10
		5 4 1`)}, wantOut: `101`},
		{name: "2", args: args{stdin: strings.NewReader(`3
		1 4 1
		2 4 1
		3 4 1`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`10
		1 4 602436426
		2 1 623690081
		3 3 262703497
		4 4 628894325
		5 3 450968417
		6 1 161735902
		7 1 707723857
		8 2 802329211
		9 0 317063340
		10 2 125660016`)}, wantOut: `2978279323`},
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
