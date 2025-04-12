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
		{name: "1", args: args{stdin: strings.NewReader(`4 2
o???`)}, wantOut: `o.??`},
		{name: "2", args: args{stdin: strings.NewReader(`5 2
?????`)}, wantOut: `?????`},
		{name: "3", args: args{stdin: strings.NewReader(`7 3
.o???o.`)}, wantOut: `.o.o.o.`},
		{name: "4", args: args{stdin: strings.NewReader(`7 2
.o???o.`)}, wantOut: `.o...o.`},
		{name: "5", args: args{stdin: strings.NewReader(`11 4
.o.?????.o.`)}, wantOut: `.o.?????.o.`},
		{name: "6", args: args{stdin: strings.NewReader(`11 5
.o.?????.o.`)}, wantOut: `.o.o.o.o.o.`},
		{name: "7", args: args{stdin: strings.NewReader(`3 2
???`)}, wantOut: `o.o`},
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
