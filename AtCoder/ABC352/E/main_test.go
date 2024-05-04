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
		{name: "1", args: args{stdin: strings.NewReader(`4 3
		3 3
		1 2 3
		2 2
		1 2
		3 4
		1 3 4`)}, wantOut: `9`},
		{name: "2", args: args{stdin: strings.NewReader(`3 2
		2 1
		1 2
		2 1
		1 2`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`10 5
		6 158260522
		1 3 6 8 9 10
		10 877914575
		1 2 3 4 5 6 7 8 9 10
		4 602436426
		2 6 7 9
		6 24979445
		2 3 4 5 8 10
		4 861648772
		2 4 8 9`)}, wantOut: `1202115217`},
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
