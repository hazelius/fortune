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
		{name: "1", args: args{stdin: strings.NewReader(`4
31 4
26 5
3 5
15 9
4
3 4 5 6`)}, wantOut: `31
26
15
15
`},
		{name: "2", args: args{stdin: strings.NewReader(`10
587 138
772 155
755 404
519 408
529 432
169 586
114 632
249 656
329 972
299 984
14
443 801 824 276 399 314 300 510 311 580 498 930 359 5`)}, wantOut: `329
329
329
755
755
755
755
329
755
329
329
329
755
772
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
