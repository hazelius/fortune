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
		{name: "1", args: args{stdin: strings.NewReader(`a?c
				b?`)}, wantOut: `Yes
No
No
`},
		{name: "2", args: args{stdin: strings.NewReader(`atcoder
		?????`)}, wantOut: `Yes
Yes
Yes
Yes
Yes
Yes
`},
		{name: "3", args: args{stdin: strings.NewReader(`beginner
		contest`)}, wantOut: `No
No
No
No
No
No
No
No
`},
		{name: "4", args: args{stdin: strings.NewReader(`?bcoefg?i?
				d?o`)}, wantOut: `Yes
Yes
Yes
No
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
