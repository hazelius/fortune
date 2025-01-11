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
		{name: "1", args: args{stdin: strings.NewReader(`3 42
3 8 4`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`3 1
3 8 4`)}, wantOut: `No`},
		{name: "3", args: args{stdin: strings.NewReader(`20 83298426
748 169 586 329 972 529 432 519 408 587 138 249 656 114 632 299 984 755 404 772`)}, wantOut: `Yes`},
		{name: "4", args: args{stdin: strings.NewReader(`20 85415869
748 169 586 329 972 529 432 519 408 587 138 249 656 114 632 299 984 755 404 772`)}, wantOut: `No`},
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
