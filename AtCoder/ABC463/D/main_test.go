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
		{name: "1", args: args{stdin: strings.NewReader(`6 3
1 12
2 7
5 9
9 13
10 18
15 20`)}, wantOut: `2`},
		{name: "2", args: args{stdin: strings.NewReader(`2 2
1 5
5 9`)}, wantOut: `-1`},
		{name: "3", args: args{stdin: strings.NewReader(`20 5
169 748
329 586
529 972
432 520
408 587
138 250
114 656
299 632
755 984
404 772
155 506
832 854
353 465
374 387
384 567
555 631
428 951
104 705
405 530
102 258`)}, wantOut: `35`},
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
