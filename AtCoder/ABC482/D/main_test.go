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
2 4 10 14 18`)}, wantOut: `8
7
6
6
8
`},
		{name: "2", args: args{stdin: strings.NewReader(`9
140322306 171833796 356303608 595703832 701965684 798563428 857902324 933683272 971602066`)}, wantOut: `415639880
399884135
307649229
280821689
280821689
329120561
358790009
396680483
415639880
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
