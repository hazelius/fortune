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
		{name: "1", args: args{stdin: strings.NewReader(`5 5
1 2
3 4
1 3
2 3
2 5`)}, wantOut: `1`},
		{name: "2", args: args{stdin: strings.NewReader(`7 8
2 4
1 3
1 7
1 3
1 2
1 6
1 5
1 3`)}, wantOut: `2`},
		{name: "3", args: args{stdin: strings.NewReader(`5 8
1 2
2 4
1 3
1 3
1 2
1 2
1 5
1 2`)}, wantOut: `2`},
		{name: "4", args: args{stdin: strings.NewReader(`5 6
1 2
1 2
1 2
1 2
1 2
1 3`)}, wantOut: `5`},
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
