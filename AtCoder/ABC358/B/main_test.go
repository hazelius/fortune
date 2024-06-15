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
		{name: "1", args: args{stdin: strings.NewReader(`3 4
0 2 10`)}, wantOut: `4
8
14
`},
		{name: "2", args: args{stdin: strings.NewReader(`3 3
1 4 7`)}, wantOut: `4
7
10
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 50000
120190 165111 196897 456895 540000 552614 561627 743796 757613 991216`)}, wantOut: `170190
220190
270190
506895
590000
640000
690000
793796
843796
1041216
`},
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
