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
		{name: "1", args: args{stdin: strings.NewReader(`2 6
2 1 at
3 1
2 2 on
1 2
2 2 coder
3 2`)}, wantOut: `atcoder`},
		{name: "2", args: args{stdin: strings.NewReader(`100000 3
1 100
2 300 abc
3 200`)}, wantOut: ``},
		{name: "3", args: args{stdin: strings.NewReader(`10 10
2 7 ladxf
2 7 zz
2 7 kfm
3 7
1 5
2 5 irur
3 5
1 6
2 6 ptilun
3 6`)}, wantOut: `ladxfzzkfmirurptilun`},
		{name: "4", args: args{stdin: strings.NewReader(`100000 7
2 1 a
2 2 b 
3 2
1 100
2 2 c
2 2 d
3 100`)}, wantOut: `b`},
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
