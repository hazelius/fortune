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
		{name: "1", args: args{stdin: strings.NewReader(`6
2 3 4 4 7 10
`)}, wantOut: `3`},
		{name: "2", args: args{stdin: strings.NewReader(`3
387 388 389`)}, wantOut: `0`},
		{name: "3", args: args{stdin: strings.NewReader(`24
307 321 330 339 349 392 422 430 477 481 488 537 541 571 575 602 614 660 669 678 712 723 785 792`)}, wantOut: `6`},
		{name: "4", args: args{stdin: strings.NewReader(`6
1 2 3 4 5 6`)}, wantOut: `3`},
		{name: "5", args: args{stdin: strings.NewReader(`6
2 9 10 16 20 32`)}, wantOut: `3`},
		{name: "6", args: args{stdin: strings.NewReader(`6
2 2 4 4 5 8`)}, wantOut: `3`},
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
