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
		{name: "1", args: args{stdin: strings.NewReader(`4 10
30 300 3
60 45 4
90 45 5
120 45 2`)}, wantOut: `30
60
105
120
`},
		{name: "2", args: args{stdin: strings.NewReader(`4 10
30 300 10
60 45 2
90 45 3
120 45 4`)}, wantOut: `30
330
330
330
`},
		{name: "3", args: args{stdin: strings.NewReader(`10 24
279290 9485601 1
1094410 8022270 4
1314176 7214745 5
1897674 5924694 10
1921802 5769841 4
2506394 2765234 2
2558629 2727489 9
2681289 4061363 5
3022540 2291905 3
4407692 1313036 8`)}, wantOut: `279290
1094410
1314176
1897674
1921802
7691643
7822368
8528921
8528921
10549857
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
