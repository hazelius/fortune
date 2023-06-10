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
		{name: "1", args: args{stdin: strings.NewReader(`7
		0 240 720 1320 1440 1800 2160
		3
		480 1920
		720 1200
		0 2160
		`)}, wantOut: `480
0
960
`},
		{name: "2", args: args{stdin: strings.NewReader(`21
		0 20 62 192 284 310 323 324 352 374 409 452 486 512 523 594 677 814 838 946 1000
		10
		77 721
		255 541
		478 970
		369 466
		343 541
		42 165
		16 618
		222 592
		730 983
		338 747`)}, wantOut: `296
150
150
49
89
20
279
183
61
177
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
