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
		4 10
		5 8
		2 9`)}, wantOut: `18`},
		{name: "2", args: args{stdin: strings.NewReader(`5
		1 1
		1 1
		1 1
		1 1
		1 1`)}, wantOut: `5`},
		{name: "3", args: args{stdin: strings.NewReader(`10
		690830957 868532399
		741145463 930111470
		612846445 948344128
		540375785 925723427
		723092548 925021315
		928915367 973970164
		563314352 832796216
		562681294 868338948
		923012648 954764623
		691107436 891127278`)}, wantOut: `7362669937`},
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
