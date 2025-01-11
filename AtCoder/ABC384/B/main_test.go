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
		{name: "1", args: args{stdin: strings.NewReader(`4 1255
2 900
1 521
2 600
1 52`)}, wantOut: `2728`},
		{name: "2", args: args{stdin: strings.NewReader(`2 3031
1 1000
2 -1000`)}, wantOut: `3031`},
		{name: "3", args: args{stdin: strings.NewReader(`15 2352
2 -889
2 420
2 -275
1 957
1 -411
1 -363
1 151
2 -193
2 289
2 -770
2 109
1 345
2 551
1 -702
1 355`)}, wantOut: `1226`},
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
