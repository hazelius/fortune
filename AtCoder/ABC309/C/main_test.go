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
		{name: "1", args: args{stdin: strings.NewReader(`4 8
		6 3
		2 5
		1 9
		4 2
		`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`4 100
		6 3
		2 5
		1 9
		4 2`)}, wantOut: `1`},
		{name: "3", args: args{stdin: strings.NewReader(`15 158260522
		877914575 2436426
		24979445 61648772
		623690081 33933447
		476190629 62703497
		211047202 71407775
		628894325 31963982
		822804784 50968417
		430302156 82631932
		161735902 80895728
		923078537 7723857
		189330739 10286918
		802329211 4539679
		303238506 17063340
		492686568 73361868
		125660016 50287940`)}, wantOut: `492686569`},
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
