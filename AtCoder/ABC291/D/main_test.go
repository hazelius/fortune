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
		1 2
		4 2
		3 4`)}, wantOut: `4`},
		{name: "2", args: args{stdin: strings.NewReader(`4
		1 5
		2 6
		3 7
		4 8`)}, wantOut: `16`},
		{name: "3", args: args{stdin: strings.NewReader(`8
		877914575 602436426
		861648772 623690081
		476190629 262703497
		971407775 628894325
		822804784 450968417
		161735902 822804784
		161735902 822804784
		822804784 161735902`)}, wantOut: `48`},
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
