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
		00011
		3 9 2 6 4`)}, wantOut: `7`},
		{name: "2", args: args{stdin: strings.NewReader(`4
		1001
		1 2 3 4`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`11
		11111100111
		512298012 821282085 543342199 868532399 690830957 973970164 928915367 954764623 923012648 540375785 925723427`)}, wantOut: `2286846953`},
		{name: "4", args: args{stdin: strings.NewReader(`5
		10101
		30 9 52 26 54`)}, wantOut: `30`},
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
