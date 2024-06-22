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
3 1 4 1 5`)}, wantOut: `4 5 13 14 26 `},
		{name: "2", args: args{stdin: strings.NewReader(`6
1000000000 1000000000 1000000000 1000000000 1000000000 1000000000`)}, wantOut: `1000000001 2000000001 3000000001 4000000001 5000000001 6000000001 `},
		{name: "3", args: args{stdin: strings.NewReader(`15
748 169 586 329 972 529 432 519 408 587 138 249 656 114 632`)}, wantOut: `749 918 1921 2250 4861 5390 5822 6428 6836 7796 7934 8294 10109 10223 11373 `},
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
