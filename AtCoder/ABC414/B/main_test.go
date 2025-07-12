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
		{name: "1", args: args{stdin: strings.NewReader(`8
m 1
i 1
s 2
i 1
s 2
i 1
p 2
i 1`)}, wantOut: `mississippi`},
		{name: "2", args: args{stdin: strings.NewReader(`7
a 1000000000000000000
t 1000000000000000000
c 1000000000000000000
o 1000000000000000000
d 1000000000000000000
e 1000000000000000000
r 1000000000000000000`)}, wantOut: `Too Long`},
		{name: "3", args: args{stdin: strings.NewReader(`1
a 100`)}, wantOut: `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`},
		{name: "4", args: args{stdin: strings.NewReader(`6
g 4
j 1
m 4
e 4
d 3
i 4`)}, wantOut: `ggggjmmmmeeeedddiiii`},
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
